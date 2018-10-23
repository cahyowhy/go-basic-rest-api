package models

import (
	"encoding/json"
	"time"

	"github.com/icrowley/fake"
)

type Todo struct {
	ID         uint       `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time  `json:"created_at,omitempty"`
	UpdatedAt  time.Time  `json:"updated_at,omitempty"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
	Name       string     `json:"name"`
	Subcontent string     `json:"subcontent"`
	Content    string     `gorm:"type:text" json:"content"`
	Completed  bool       `json:"completed"`
	Due        *time.Time `gorm:"default:null" json:"due"`
	UserID     uint       `json:"user_id"`
	User       User       `gorm:"association_autoupdate:false;association_autocreate:false" json:"user,omitempty"`
}

func (todo *Todo) ToMap() map[string]interface{} {
	var todoMap map[string]interface{} = make(map[string]interface{})

	todoMap["id"] = todo.ID
	todoMap["created_at"] = todo.CreatedAt
	todoMap["name"] = todo.Name
	todoMap["subcontent"] = todo.Subcontent
	todoMap["content"] = todo.Content
	todoMap["completed"] = todo.Completed
	todoMap["user_id"] = todo.UserID
	todoMap["due"] = todo.Due

	if todo.User.ID != 0 {
		todoMap["user_id"] = todo.UserID
		todoMap["user"] = todo.User.ToMap()
	}

	return todoMap
}

func (todo *Todo) completed() {
	todo.Completed = true
}

func (todo *Todo) unCompleted() {
	todo.Completed = false
}

func (t *Todo) FakeIt() {
	duration := time.Now().AddDate(1, 2, fake.Day())

	t.Name = fake.Company()
	t.Content = "<p>" + fake.WordsN(128) + "</p>"
	t.Subcontent = fake.WordsN(100)
	t.Completed = false
	t.Due = &duration
}

func (t Todo) Serialize() ([]byte, error) {
	return json.Marshal(t.ToMap())
}

type Todos []Todo

func SerializeTodos(todos []Todo) ([]byte, error) {
	var todoMaps []map[string]interface{} = make([]map[string]interface{}, len(todos))

	for i, todo := range todos {
		todoMap := todo.ToMap()
		delete(todoMap, "content")

		todoMaps[i] = todoMap
	}

	return json.Marshal(todoMaps)
}
