package handlers

import (
	"encoding/json"
	"fmt"
	"go-basic-rest-api/models"
	"go-basic-rest-api/utils"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllTodos(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	todos := []models.Todo{}
	var count string

	if query.Get("offset") == "" || query.Get("limit") == "" {
		respondError(w, http.StatusBadRequest, `"Required offset limit but not present"`, utils.INPUT_NOT_VALID)

		return
	}

	tx := db.Preload("User").Order("id DESC").Offset(query.Get("offset")).Limit(query.Get("limit"))
	txCount := db.Model(&models.Todo{})

	if query.Get("user_id") != "" {
		tx = tx.Where("user_id = ?", query.Get("user_id"))
		txCount = txCount.Where("user_id = ?", query.Get("user_id"))
	}

	if err := tx.Find(&todos).Error; err != nil {
		respondError(w, http.StatusNotFound, fmt.Sprintf(`"%s"`, err.Error()), utils.DATA_NOT_FOUND)

		return
	}

	if err := txCount.Count(&count).Error; err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Sprintf(`"%s"`, err.Error()), utils.DB_EXCEPTION)

		return
	}

	todoJsons, err := models.SerializeTodos(todos)
	if err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Sprintf(`"%s"`, err.Error()), utils.FAILED_SERIALIZE)
	}

	ProcessJSON(w, http.StatusOK, todoJsons, utils.STATUS_OK, count)
}

func CreateTodo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	todo := models.Todo{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		respondError(w, http.StatusBadRequest, fmt.Sprintf(`"%s"`, err.Error()), utils.INPUT_NOT_VALID)

		return
	}

	defer r.Body.Close()

	if err := db.Save(&todo).Error; err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Sprintf(`"%s"`, err.Error()), utils.UPDATE_FAILED)

		return
	}

	respondJSON(w, http.StatusCreated, todo, utils.SAVE_SUCCESS)
}

func GetTodo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	todo := getTodoOr404(db, id, w, r)

	if todo == nil {
		return
	}

	respondJSON(w, http.StatusOK, todo, utils.STATUS_OK)
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
		respondError(w, http.StatusBadRequest, fmt.Sprintf(`"%s"`, err.Error()), utils.UPDATE_FAILED)
		return
	}

	defer r.Body.Close()

	if err := db.Save(&todo).Error; err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Sprintf(`"%s"`, err.Error()), utils.UPDATE_FAILED)
		return
	}

	respondJSON(w, http.StatusOK, todo, utils.UPDATE_SUCCESS)
}

func DeleteTodo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	todo := getTodoOr404(db, id, w, r)

	if todo == nil {
		return
	}

	if err := db.Delete(&todo).Error; err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Sprintf(`"%s"`, err.Error()), utils.DELETE_FAILED)
		return
	}

	ProcessJSON(w, http.StatusOK, []byte(`"Delete Suceed"`), utils.DELETE_SUCCESS, "")
}

func getTodoOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *models.Todo {
	todo := models.Todo{}
	if err := db.Preload("User").First(&todo, id).Error; err != nil {
		respondError(w, http.StatusNotFound, fmt.Sprintf(`"%s"`, err.Error()), utils.DATA_NOT_FOUND)

		return nil
	}

	return &todo
}
