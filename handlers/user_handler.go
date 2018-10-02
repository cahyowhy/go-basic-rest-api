package handlers

import (
	"encoding/json"
	"go-basic-rest-api/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())

		return
	}

	defer r.Body.Close()

	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())

		return
	}

	respondJSON(w, http.StatusOK, user)
}

func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	users := []models.User{}

	if query.Get("offset") == "" || query.Get("limit") == "" {
		respondError(w, http.StatusBadRequest, "Required offset limit but not present")

		return
	}

	if err := db.Offset(query.Get("offset")).Limit(query.Get("limit")).Find(&users).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())

		return
	}

	userJsons, err := models.SerializeUsers(users)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}

	processJSON(w, http.StatusOK, userJsons)
}

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user := getUserOr404(db, id, w, r)

	if user == nil {
		return
	}

	respondJSON(w, http.StatusOK, user)
}

func getUserOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *models.User {
	user := models.User{}
	todos := []models.Todo{}

	if err := db.First(&user, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())

		return nil
	}

	db.Where("user_id=?", user.ID).Find(&todos)

	user.Todos = todos

	return &user
}
