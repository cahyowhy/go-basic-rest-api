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

func RenderHome(_ *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	mapClaims, ok := utils.GetTokenParsed(r)

	if ok {
		username, okUsername := mapClaims["username"].(string)

		if okUsername && len(username) != 0 {

		}
	}

	// if err := db.Find(&users).Error; err != nil {
	// 	respondError(w, http.StatusInternalServerError, "FAILED RENDER PAGE", utils.DATA_NOT_FOUND)

	// 	return
	// }

	p := &templates.HomePage{
		User: user,
	}

	templates.WritePageTemplate(w, p)
}
