package handlers

import (
	"go-basic-rest-api/models"
	"go-basic-rest-api/templates"
	"net/http"

	"github.com/jinzhu/gorm"
)

func RenderIndex(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	users := []models.User{}
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	if err := db.Find(&users).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "FAILED RENDER PAGE")

		return
	}

	p := &templates.IndexPage{
		Users: users,
	}

	templates.WritePageTemplate(w, p)
}
