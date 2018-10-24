package handlers

import (
	"errors"
	"go-basic-rest-api/models"
	"go-basic-rest-api/templates"
	"go-basic-rest-api/utils"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
)

func RenderIndex(_ *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	p := &templates.IndexPage{}

	templates.WritePageTemplate(w, p)
}

func RenderNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	p := &templates.NotFoundPage{}

	templates.WritePageTemplate(w, p)
}

func RenderTodo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	var err error
	var userJson []byte

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	if len(name) == 0 {
		renderErrPage(w, r)
		return
	}

	todo := models.Todo{}
	if err = db.Preload("User").First(&todo, "name = ?", name).Error; err != nil {
		renderErrPage(w, r)
		return
	}

	if userJson, err = todo.User.Serialize(); err != nil {
		renderErrPage(w, r)
		return
	}

	p := &templates.TodoPage{
		Todo:     todo,
		UserJSON: userJson,
	}

	templates.WritePageTemplate(w, p)
}

func RenderHome(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user, userJson, err := getUserAuth(db, w, r)
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	if err != nil {
		return
	}

	p := &templates.HomePage{
		User:     *user,
		UserJSON: userJson,
	}

	templates.WritePageTemplate(w, p)
}

func RenderSetting(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	user, userJson, err := getUserAuth(db, w, r)

	if err != nil {
		return
	}

	p := &templates.SettingPage{
		User:     *user,
		UserJSON: userJson,
	}

	templates.WritePageTemplate(w, p)
}

func RenderAdmin(_ *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	p := &templates.AdminPage{}
	templates.WritePageTemplate(w, p)
}

func renderErrPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	p := &templates.ErrorPage{}

	templates.WritePageTemplate(w, p)
}

func getUserAuth(db *gorm.DB, w http.ResponseWriter, r *http.Request) (*models.User, []byte, error) {
	user := models.User{}
	var userJson []byte
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	mapClaims, ok := utils.GetTokenParsed(r)

	if ok {
		var err error
		username, okUsername := mapClaims["username"].(string)

		if err = db.First(&user, "username = ?", username).Error; err != nil && okUsername {
			renderErrPage(w, r)

			return nil, nil, err
		}

		if userJson, err = user.Serialize(); err != nil {
			renderErrPage(w, r)

			return nil, nil, err
		}

		return &user, userJson, nil
	}

	renderErrPage(w, r)

	return nil, nil, errors.New(`"Failed get user auth"`)
}
