package main

import (
	"bytes"
	"encoding/json"
	"go-basic-rest-api/config"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a App
var token string

func TestMain(m *testing.M) {
	configApp := config.GetConfig()
	a = App{}

	a.Initialize(configApp)
	code := m.Run()

	os.Exit(code)
}

func TestGetAllTodos(t *testing.T) {
	m1 := "test get all todos handler, expect 200 as status code"
	m2 := "test get all todos handler, expect 400 as status code"

	executeBasicHandlerGetTest(t, m1, "/api/todos?offset=0&limit=9", http.StatusOK)
	executeBasicHandlerGetTest(t, m2, "/api/todos", http.StatusBadRequest)
}
func TestGetTodos(t *testing.T) {
	m1 := "test get todo handler, expect 200 as status code"
	m2 := "test get todo handler, expect 404 as status code"

	executeBasicHandlerGetTest(t, m1, "/api/todos/1", http.StatusOK)
	executeBasicHandlerGetTest(t, m2, "/api/todos/1001", http.StatusNotFound)
}

func TestCreateTodos(t *testing.T) {
	TestLogin(t)

	response := executePost(t, "POST", "/api/todos", []byte(`{"name": "kill someone", "completed": false, "user_id": 3, "due": "2017-02-20T17:00:00.000Z"}`))
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestLogin(t *testing.T) {
	var mapper map[string]interface{}
	response := executePost(t, "POST", "/api/login", []byte(`{"username": "quepasacontigo","password":"123456"}`))
	err := json.Unmarshal(response.Body.Bytes(), &mapper)

	if err != nil {
		t.Fatal(err)
	}

	tokenDecoded, _ := mapper["token"].(string)
	token = tokenDecoded

	checkResponseCode(t, http.StatusOK, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func executeBasicHandlerGetTest(t *testing.T, message string, url string, status int) {
	t.Log(message)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		t.Fatal(err)
	}

	response := executeRequest(req)
	checkResponseCode(t, status, response.Code)
}

func executePost(t *testing.T, method string, url string, payload interface{}) *httptest.ResponseRecorder {
	bytePayload, ok := payload.([]byte)
	body := bytes.NewReader(bytePayload)

	if !ok {
		body = nil
	}

	req, err := http.NewRequest(method, url, body)

	if len(token) != 0 {
		req.Header.Set("Authorization", token)
	}

	if err != nil {
		t.Fatal(err)
	}

	return executeRequest(req)
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
