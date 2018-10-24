package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"go-basic-rest-api/config"
	"go-basic-rest-api/handlers"
	"go-basic-rest-api/models"
	"go-basic-rest-api/routes"
	"go-basic-rest-api/utils"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (app *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Name,
		config.DB.Charset)
	fmt.Println(dbURI)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	app.DB = models.DBMigrate(db)
	app.Router = mux.NewRouter()

	app.DB.LogMode(true)
	app.setRouters()
	app.Router.NotFoundHandler = http.HandlerFunc(routes.NotFoundRoute)
	app.Router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	app.Router.PathPrefix("/user-files/").Handler(http.StripPrefix("/user-files/", http.FileServer(http.Dir("./files/"))))
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) seedsDb() {
	users := []models.User{}
	a.DB.Find(&users)

	countUser := len(users)

	if countUser < 50 {
		for index := 0; index < 50; index++ {
			user := models.User{}
			user.FakeIt()

			if err := a.DB.Create(&user).Error; err != nil {
				fmt.Println(err.Error())
			}

			todo := models.Todo{}
			todo.FakeIt()
			todo.User = user

			if err := a.DB.Create(&todo).Error; err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

func (app *App) setRouters() {
	for _, route := range routes.DefinedRoutes {
		newRoute := route

		var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			var fullUrl string = r.Host + r.RequestURI
			var isIndexPath bool = false

			if !strings.HasPrefix(fullUrl, "http") || !strings.HasPrefix(fullUrl, "https") {
				fullUrl = "http://" + fullUrl
			}

			urlParse, err := url.Parse(fullUrl)

			log.Printf(
				"%s\t%s\t%s\t%s",
				r.Method,
				fullUrl,
				newRoute.Name,
				time.Since(start),
			)

			if err != nil {
				log.Printf("err parse url %s %s", fullUrl, err.Error())
			} else {
				isIndexPath = urlParse.Path == "/"
			}

			redirectUnauthorize := func(err string, w http.ResponseWriter, r *http.Request, isIndex bool) {
				if !isIndex {
					log.Printf("err : %s", err)
					http.Redirect(w, r, "/?login-first=true", http.StatusTemporaryRedirect)

					return
				}
			}

			if newRoute.AuthFirst {
				if validHeaderAuth := utils.DecodedToken(r); validHeaderAuth != nil {
					handlers.ProcessJSON(w, http.StatusUnauthorized, validHeaderAuth, utils.TOKEN_NOT_VALID)

					return
				}
			}

			if newRoute.AuthCookie {
				cookie, err := r.Cookie("token")

				if err == nil {
					if validTokenCookie := utils.ValidToken(cookie.Value); !validTokenCookie {
						// remove cookie if has cookie but failed auth
						c := &http.Cookie{
							Name:    "token",
							Value:   "",
							Path:    "/",
							Expires: time.Unix(0, 0),

							HttpOnly: true,
						}
						http.SetCookie(w, c)

						redirectUnauthorize(utils.TOKEN_NOT_VALID, w, r, isIndexPath)
					} else {
						if isIndexPath {
							http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
						}
					}
				} else {
					redirectUnauthorize(err.Error(), w, r, isIndexPath)
				}

			}

			newRoute.RouteHandle(app.DB, w, r)
		}

		app.Router.HandleFunc(newRoute.Pattern, handler).
			Name(newRoute.Name).
			Methods(newRoute.Method)
	}
}
