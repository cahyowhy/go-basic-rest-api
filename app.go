package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"go-basic-rest-api/models"
	"go-basic-rest-api/routes"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (app *App) Initialize(config *Config) {
	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username, config.DB.Password, config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	app.DB = models.DBMigrate(db)
	app.Router = mux.NewRouter()

	app.setRouters()
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) seedsDb() {
	users := []models.User{}
	a.DB.Find(&users)

	countUser := len(users)

	if countUser < 200 {
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
			log.Printf(
				"%s\t%s\t%s\t%s",
				r.Method,
				r.RequestURI,
				newRoute.Name,
				time.Since(start),
			)

			newRoute.RouteHandle(app.DB, w, r)
		}

		app.Router.HandleFunc(newRoute.Pattern, handler).
			Name(newRoute.Name).
			Methods(newRoute.Method)
	}
}
