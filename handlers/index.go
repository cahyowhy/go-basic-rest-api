package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-basic-rest-api/models"
	"go-basic-rest-api/utils"
	"io/ioutil"
	"strings"

	"mime/multipart"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/thanhpk/randstr"
	"github.com/tidwall/sjson"
)

type HandlerRoute func(*gorm.DB, http.ResponseWriter, *http.Request)

func respondJSON(w http.ResponseWriter, status int, payload interface{ models.Serialize }, statuscode string) {
	response, err := payload.Serialize()
	response, err = handleJson(response, statuscode)

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error(), utils.FAILED_SERIALIZE)

		return
	}

	ProcessJSON(w, status, response, "")
}

func ProcessJSON(w http.ResponseWriter, status int, response []byte, statuscode string) {
	responseJson := response
	if len(statuscode) != 0 {
		responseJson, _ = handleJson(response, statuscode)
	}

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(status)
	w.Write([]byte(responseJson))

	return
}

func respondError(w http.ResponseWriter, code int, message string, statuscode string) {
	response, err := handleJson([]byte(message), statuscode)
	if err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	ProcessJSON(w, code, response, "")
}

func filterFileFormat(fileFormats []string, fileFormat string) bool {
	fileFormat = strings.ToLower(fileFormat)
	for _, file := range fileFormats {
		if file == fileFormat {
			return true
		}
	}

	return false
}

func saveFile(file multipart.File, handle *multipart.FileHeader, filename string, allowedFormatFiles []string) error {
	filepath := "./files/" + filename
	splitString := strings.Split(filename, ".")
	formatFile := splitString[len(splitString)-1]

	if len(allowedFormatFiles) != 0 && !filterFileFormat(allowedFormatFiles, formatFile) {
		return errors.New("Format file is not permitted")
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("err %s", err.Error())
		return err
	}

	err = ioutil.WriteFile(filepath, data, 0666)
	if err != nil {
		fmt.Printf("err %s", err.Error())
		return err
	}

	return nil
}

func handleJson(payload []byte, status string) ([]byte, error) {
	jsonval, err := json.Marshal("{}")

	clonedJson := jsonval
	clonedJson, err = sjson.SetRawBytes(clonedJson, "data", payload)
	clonedJson, err = sjson.SetRawBytes(clonedJson, "status", []byte(fmt.Sprintf(`"%s"`, status)))

	return clonedJson, err
}

func uploadPhotoUserMixin(db *gorm.DB, w http.ResponseWriter, r *http.Request) (models.UserPhoto, error) {
	file, handle, err := r.FormFile("file")
	userphotoForm := r.FormValue("userPhoto")
	userPhoto := models.UserPhoto{}

	if len(userphotoForm) == 0 {
		respondError(w, http.StatusOK, `"Form value userPhoto has no value"`, utils.INPUT_NOT_VALID)
		return models.UserPhoto{}, errors.New("Form value userPhoto has no value")
	}

	if json.Unmarshal([]byte(userphotoForm), &userPhoto); !userPhoto.ValidValue() {
		respondError(w, http.StatusBadRequest, `"Form value userPhoto has no user_id"`, utils.INPUT_NOT_VALID)
		return models.UserPhoto{}, errors.New("Form value userPhoto has no user_id")
	}

	defer file.Close()

	filename := randstr.Hex(16) + handle.Filename
	mimeType := r.Header.Get("Content-Type")
	allowedFormatFiles := []string{"jpg", "png", "gif", "jpeg"}

	if !strings.Contains(mimeType, "multipart/form-data") {
		ProcessJSON(w, http.StatusBadRequest, []byte(`"invalid file header"`), utils.UPLOAD_FAILED)
		return models.UserPhoto{}, errors.New("invalid file header")
	}

	if err = saveFile(file, handle, filename, allowedFormatFiles); err != nil {
		ProcessJSON(w, http.StatusInternalServerError, []byte(fmt.Sprintf(`"%s"`, err.Error())), utils.UPLOAD_FAILED)
		return models.UserPhoto{}, err
	}

	userPhoto.Path = filename
	if err := db.Save(&userPhoto).Error; err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Sprintf(`"%s"`, err.Error()), utils.UPLOAD_FAILED)

		return models.UserPhoto{}, err
	}

	return userPhoto, nil
}
