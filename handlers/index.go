package handlers

import (
	"encoding/json"
	"fmt"
	"go-basic-rest-api/models"
	"go-basic-rest-api/utils"
	"io/ioutil"

	"mime/multipart"
	"net/http"

	"github.com/jinzhu/gorm"
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

func saveFile(file multipart.File, handle *multipart.FileHeader) error {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("err %s", err.Error())
		return err
	}

	err = ioutil.WriteFile("./files/"+handle.Filename, data, 0666)
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
