package handlers

import (
	"encoding/json"
	"fmt"
	"go-basic-rest-api/models"
	"go-basic-rest-api/utils"
	"net/http"
	"strconv"

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

func UploadPhotoProfile(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	userPhoto, err := uploadPhotoUserMixin(db, w, r)
	if err != nil {
		return
	}

	userId := fmt.Sprint(userPhoto.UserID)
	user := getUserOr404(db, userId, w, r)
	if user == nil {
		return
	}

	user.ImageProfile = userPhoto.Path
	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Sprintf(`"%s"`, err.Error()), utils.UPDATE_FAILED)
		return
	}

	respondJSON(w, http.StatusOK, user, utils.UPDATE_SUCCESS)
}

func UpdateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	intId, err := strconv.ParseUint(id, 10, 32)

	if id == "" || err != nil {
		respondError(w, http.StatusNotFound, fmt.Sprintf(`"%s"`, err.Error()), utils.DATA_NOT_FOUND)
		return
	}

	user := models.User{ID: uint(intId)}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil || (len(user.Name) == 0 && len(user.Username) == 0) {
		respondError(w, http.StatusBadRequest, fmt.Sprintf(`"%s"`, err.Error()), utils.UPDATE_FAILED)
		return
	}

	defer r.Body.Close()

	if err := db.Model(&user).Omit("password").Updates(user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Sprintf(`"%s"`, err.Error()), utils.UPDATE_FAILED)
		return
	}

	userData := getUserOr404(db, id, w, r)
	if userData == nil {
		return
	}

	respondJSON(w, http.StatusOK, userData, utils.UPDATE_SUCCESS)
}

func UpdateUserPassword(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var user map[string]interface{} = make(map[string]interface{})

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, fmt.Sprintf(`"%s"`, err.Error()), utils.UPDATE_FAILED)
		return
	}

	defer r.Body.Close()

	id, _ := user["id"].(string)
	passwordOld, _ := user["passwordOld"].(string)
	password, _ := user["password"].(string)

	if len(id) == 0 || len(passwordOld) == 0 || len(password) == 0 {
		respondError(w, http.StatusBadRequest, `"Required passwordOld, password & id but not present"`, utils.INPUT_NOT_VALID)
		return
	}

	userStored := getUserOr404(db, id, w, r)
	if userStored == nil {
		return
	}

	isMatch := utils.CompareHashPassword(passwordOld, userStored.Password)
	if !isMatch {
		respondError(w, http.StatusUnauthorized, `"Password not valid!"`, utils.PASSWORD_NOT_VALID)
		return
	}

	newPassword, errPass := utils.GeneratePassword(password)
	if errPass != nil {
		respondError(w, http.StatusInternalServerError, fmt.Sprintf(`"%s"`, errPass.Error()), utils.UPDATE_FAILED)
		return
	}

	if err := db.Model(userStored).Update("password", newPassword).Error; err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Sprintf(`"%s"`, err.Error()), utils.UPDATE_FAILED)
		return
	}

	respondJSON(w, http.StatusOK, userStored, utils.UPDATE_SUCCESS)
}
