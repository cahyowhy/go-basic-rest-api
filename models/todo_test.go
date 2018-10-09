package models

import (
	"encoding/json"
	"testing"
)

func TestSerializeTodo(t *testing.T) {
	testSerializeTodoOpt(t, true)
	testSerializeTodoOpt(t, false)
}

func testSerializeTodoOpt(t *testing.T, hasUserId bool) {
	message := "Check Todo on serialize has user"
	if !hasUserId {
		message = "Check Todo on serialize has no user"
	}

	t.Log(message)

	todo := Todo{}
	todo.FakeIt()

	if hasUserId {
		user := User{}
		user.FakeIt()
		user.ID = 2

		todo.User = user
		todoJson, _ := todo.Serialize()

		var todoChecked map[string]interface{}
		json.Unmarshal(todoJson, &todoChecked)
		_, todoOk := todoChecked["user"]

		if !todoOk {
			t.Error("Expected has user but hasn't")
		}
	} else {
		var todoChecked map[string]interface{}
		todoJson, _ := todo.Serialize()

		json.Unmarshal(todoJson, &todoChecked)
		_, todoOk := todoChecked["user"]

		if todoOk {
			t.Error("Expected has'nt user but has")
		}
	}
}
