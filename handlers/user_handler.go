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

func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	password, errPass := utils.GeneratePassword(user.Password)

	if err != nil || !user.ValidValue(false) || errPass != nil {
		respondError(w, http.StatusBadRequest, `"Bad JSON request"`, utils.DATA_NOT_FOUND)

		return
	}

	user.Password = password

	defer r.Body.Close()

	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Sprintf(`"%s"`, err.Error()), utils.UPDATE_FAILED)

		return
	}

	respondJSON(w, http.StatusOK, user, utils.SAVE_SUCCESS)
}

func AuthUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	userStored := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil || !user.ValidValue(true) {
		respondError(w, http.StatusBadRequest, `"Bad JSON body"`, utils.INPUT_NOT_VALID)

		return
	}

	if errUser := db.Where("username = ?", user.Username).First(&userStored).Error; errUser != nil || userStored.ID == 0 {
		respondError(w, http.StatusNotFound, fmt.Sprintf(`"%s"`, errUser.Error()), utils.DATA_NOT_FOUND)

		return
	}

	isMatch := utils.CompareHashPassword(user.Password, userStored.Password)
	userTokenMapper, errToken := userStored.UserMapToken()
	tokenString, errGenToken := utils.GenerateToken(userTokenMapper)
	responseJson, errMergToken := userStored.MergeToken(tokenString)

	if !isMatch || errToken != nil || errGenToken != nil || errMergToken != nil {
		respondError(w, http.StatusUnauthorized, `"Password not valid!"`, utils.PASSWORD_NOT_VALID)

		return
	}

	ProcessJSON(w, http.StatusOK, responseJson, utils.LOGIN_SUCCESS)
}

func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	users := []models.User{}

	if query.Get("offset") == "" || query.Get("limit") == "" {
		respondError(w, http.StatusBadRequest, `"Required offset limit but not present"`, utils.INPUT_NOT_VALID)

		return
	}

	if err := db.Offset(query.Get("offset")).Limit(query.Get("limit")).Find(&users).Error; err != nil {
		respondError(w, http.StatusNotFound, fmt.Sprintf(`"%s"`, err.Error()), utils.DATA_NOT_FOUND)

		return
	}

	userJsons, err := models.SerializeUsers(users)
	if err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Sprintf(`"%s"`, err.Error()), utils.FAILED_SERIALIZE)
	}

	ProcessJSON(w, http.StatusOK, userJsons, utils.STATUS_OK)
}

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	user := getUserOr404(db, id, w, r)

	if user == nil {
		return
	}

	respondJSON(w, http.StatusOK, user, utils.STATUS_OK)
}

func getUserOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *models.User {
	user := models.User{}
	todos := []models.Todo{}

	if err := db.First(&user, id).Error; err != nil {
		respondError(w, http.StatusNotFound, fmt.Sprintf(`"%s"`, err.Error()), utils.DATA_NOT_FOUND)

		return nil
	}

	db.Where("user_id=?", user.ID).Find(&todos)

	user.Todos = todos

	return &user
}
