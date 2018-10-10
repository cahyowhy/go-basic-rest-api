package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-basic-rest-api/config"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/icrowley/fake"
	"github.com/jinzhu/gorm"

	"testing"
)

var a App
var token string

func TestMain(m *testing.M) {
	configApp := config.GetConfig()
	a = App{}

	a.Initialize(configApp)
	code := m.Run()

	prepareDb(a.DB)
	os.Exit(code)
}

func TestGetAllTodos(t *testing.T) {
	m1 := "test get all todos handler, expect 200 as status code"
	m2 := "test get all todos handler, expect 400 as status code"

	executeBasicHandlerGetTest(t, m1, "/api/todos?offset=0&limit=9", http.StatusOK)
	executeBasicHandlerGetTest(t, m2, "/api/todos", http.StatusBadRequest)
}

func TestGetAllUsers(t *testing.T) {
	m1 := "test get all users handler, expect 200 as status code"
	m2 := "test get all users handler, expect 400 as status code"

	executeBasicHandlerGetTest(t, m1, "/api/users?offset=0&limit=9", http.StatusOK)
	executeBasicHandlerGetTest(t, m2, "/api/users", http.StatusBadRequest)
}

func TestGetTodos(t *testing.T) {
	m1 := "test get todo handler, expect 200 as status code"
	m2 := "test get todo handler, expect 404 as status code"

	executeBasicHandlerGetTest(t, m1, "/api/todos/1", http.StatusOK)
	executeBasicHandlerGetTest(t, m2, "/api/todos/1001", http.StatusNotFound)
}

func TestGetUsers(t *testing.T) {
	m1 := "test get users handler, expect 200 as status code"
	m2 := "test get users handler, expect 404 as status code"

	executeBasicHandlerGetTest(t, m1, "/api/users/1", http.StatusOK)
	executeBasicHandlerGetTest(t, m2, "/api/users/1001", http.StatusNotFound)
}

func TestCreateTodos(t *testing.T) {
	TestLogin(t)

	response := executePostPut(t, "POST", "/api/todos", []byte(`{"name": "kill someone", "completed": false, "user_id": 3, "due": "2017-02-20T17:00:00.000Z"}`))
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestCreateUsers(t *testing.T) {
	payload := fmt.Sprintf(`{"name": "Sanata dharma", "username":"%s","password":"123456"}`, fake.UserName())
	response := executePostPut(t, "POST", "/api/users", []byte(payload))
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateTodo(t *testing.T) {
	TestLogin(t)

	response := executePostPut(t, "PUT", "/api/todos/1", []byte(`{"name": "kill someone uno", "completed": true, "due": "2017-02-20T17:00:00.000Z"}`))
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestDeleteTodo(t *testing.T) {
	TestLogin(t)

	req := getReq(t, "DELETE", "/api/todos/11", nil)

	if req == nil {
		return
	}

	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestLogin(t *testing.T) {
	var mapper map[string]interface{}
	response := executePostPut(t, "POST", "/api/login", []byte(`{"username": "quepasacontigo","password":"123456"}`))
	err := json.Unmarshal(response.Body.Bytes(), &mapper)

	if err != nil {
		t.Fatal(err)
	}

	tokenDecoded, _ := mapper["token"].(string)
	token = tokenDecoded

	checkResponseCode(t, http.StatusOK, response.Code)
}

func prepareDb(db *gorm.DB) {
	db.Exec("update todos set deleted_at = null")
	db.Exec("update users set deleted_at = null")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func executeBasicHandlerGetTest(t *testing.T, message string, url string, status int) {
	t.Log(message)
	req := getReq(t, "GET", url, nil)

	if req == nil {
		return
	}

	response := executeRequest(req)
	checkResponseCode(t, status, response.Code)
}

func executePostPut(t *testing.T, method string, url string, payload interface{}) *httptest.ResponseRecorder {
	bytePayload, _ := payload.([]byte)
	body := bytes.NewReader(bytePayload)

	if payload == nil {
		body = nil
	}

	req := getReq(t, method, url, body)

	if req == nil {
		return nil
	}

	return executeRequest(req)
}

func getReq(t *testing.T, method string, url string, payload *bytes.Reader) *http.Request {
	if payload == nil {
		req2, err2 := http.NewRequest(method, url, nil)

		if err2 != nil {
			t.Fatal(err2)

			return nil
		}

		if len(token) != 0 {
			req2.Header.Set("Authorization", token)
		}

		return req2
	}

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		t.Fatal(err)

		return nil
	}

	if len(token) != 0 {
		req.Header.Set("Authorization", token)
	}

	return req
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
