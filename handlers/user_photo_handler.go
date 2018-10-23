package handlers

import (
	"encoding/json"
	"fmt"
	"go-basic-rest-api/models"
	"go-basic-rest-api/utils"
	"net/http"

	"github.com/jinzhu/gorm"
)

func GetUserPhoto(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userPhotos := []models.UserPhoto{}

	if query.Get("offset") == "" || query.Get("limit") == "" || query.Get("user_id") == "" {
		respondError(w, http.StatusBadRequest, `"Required offset, limit and userid. but not present"`, utils.INPUT_NOT_VALID)

		return
	}

	if err := db.Where("user_id = ?", query.Get("user_id")).Order("id DESC").Offset(query.Get("offset")).Limit(query.Get("limit")).Find(&userPhotos).Error; err != nil {
		respondError(w, http.StatusNotFound, fmt.Sprintf(`"%s"`, err.Error()), utils.DATA_NOT_FOUND)

		return
	}

	userPhotoJsons, err := json.Marshal(userPhotos)
	if err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Sprintf(`"%s"`, err.Error()), utils.FAILED_SERIALIZE)
	}

	ProcessJSON(w, http.StatusOK, userPhotoJsons, utils.STATUS_OK)
}

func UploadUserPhoto(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	userPhoto, err := uploadPhotoUserMixin(db, w, r)

	if err != nil {
		return
	}

	respondJSON(w, http.StatusCreated, userPhoto, utils.SAVE_SUCCESS)
}
