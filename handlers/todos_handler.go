package handlers

import (
	"encoding/json"
	"go-basic-rest-api/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllTodos(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	todos := []models.Todo{}

	if query.Get("offset") == "" || query.Get("limit") == "" {
		respondError(w, http.StatusBadRequest, "Required offset limit but not present")

		return
	}

	if err := db.Preload("User").Offset(query.Get("offset")).Limit(query.Get("limit")).Find(&todos).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())

		return
	}

	todoJsons, err := models.SerializeTodos(todos)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
	}

	ProcessJSON(w, http.StatusOK, todoJsons)
}

func CreateTodo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	todo := models.Todo{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())

		return
	}

	defer r.Body.Close()

	if err := db.Save(&todo).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())

		return
	}

	respondJSON(w, http.StatusOK, todo)
}

func GetTodo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	todo := getTodoOr404(db, id, w, r)

	if todo == nil {
		return
	}

	respondJSON(w, http.StatusOK, todo)
}

func UpdateTodo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	todo := getTodoOr404(db, id, w, r)

	if todo == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	if err := db.Save(&todo).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, todo)
}

func DeleteTodo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	todo := getTodoOr404(db, id, w, r)

	if todo == nil {
		return
	}

	if err := db.Delete(&todo).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ProcessJSON(w, http.StatusOK, []byte(`{message: "Delete Suceed"}`))
}

func getTodoOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *models.Todo {
	todo := models.Todo{}
	if err := db.Preload("User").First(&todo, id).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())

		return nil
	}

	return &todo
}
