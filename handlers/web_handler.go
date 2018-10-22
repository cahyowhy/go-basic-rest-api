package handlers

import (
	"go-basic-rest-api/models"
	"go-basic-rest-api/templates"
	"go-basic-rest-api/utils"

	"net/http"

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

func RenderHome(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	var userJson []byte
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	mapClaims, ok := utils.GetTokenParsed(r)

	if ok {
		var err error
		username, okUsername := mapClaims["username"].(string)

		if err = db.First(&user, "username = ?", username).Error; err != nil && okUsername {
			respondError(w, http.StatusInternalServerError, `"HOLD REDIRECT ERR PAGE"`, utils.DATA_NOT_FOUND)
			return
		}

		if userJson, err = user.Serialize(); err != nil {
			respondError(w, http.StatusInternalServerError, `"HOLD REDIRECT ERR PAGE"`, utils.DATA_NOT_FOUND)
			return
		}

		p := &templates.HomePage{
			User:     user,
			UserJSON: userJson,
		}

		templates.WritePageTemplate(w, p)
	} else {
		respondError(w, http.StatusInternalServerError, `"HOLD REDIRECT ERR PAGE"`, utils.DATA_NOT_FOUND)
		return
	}
}
