package models

import (
	"encoding/json"
	"testing"
)

func TestSerializeUser(t *testing.T) {
	t.Log("Check user on serialize, must has'nt password ans todos.user")

	user := User{}
	user.FakeIt()

	todo := Todo{}
	todo.FakeIt()

	todos := []Todo{todo}
	user.Todos = todos

	userJson, _ := user.Serialize()
	var userChecked map[string]interface{}
	json.Unmarshal(userJson, &userChecked)

	_, passwordOk := userChecked["password"]

	if passwordOk {
		t.Error("Expected has'nt password but has")
	}

	if todos, ok := userChecked["todos"].([]interface{}); ok {
		for _, todo := range todos {
			todoCast, todoOk := todo.(map[string]interface{})
			if todoOk {
				_, userOk := todoCast["user"]

				if userOk {
					t.Error("Expected has'nt todo.user but has")
				}
			} else {
				t.Error("Expected todo as map[string]interface{} but it has'nt")
			}
		}
	} else {
		t.Error("Expected user todos as slice, but has'nt")
	}
}
