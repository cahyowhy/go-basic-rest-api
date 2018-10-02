package handlers

import (
	"go-basic-rest-api/models"
	"go-basic-rest-api/templates"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
)

func RenderIndex(_ *gorm.DB, w http.ResponseWriter, r *http.Request) {
	//user := models.User{}
	p := &templates.RootPage{}
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")

	// if err := db.Find(&users).Error; err != nil {
	// 	templates.(w, strconv.Itoa(http.StatusInternalServerError)+" INTERNAL SERVER ERR")

	// 	return
	// }

	templates.WritePageTemplate(w, p);
}
