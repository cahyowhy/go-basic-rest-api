package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/icrowley/fake"
)

type Todo struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Name      string     `json:"name"`
	Completed bool       `json:"completed"`
	Due       *time.Time `gorm:"default:null" json:"due"`
	UserID    uint       `json:"user_id"`
	User      User       `gorm:"association_autoupdate:false;association_autocreate:false" json:"user,omitempty"`
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
	t.Completed = false
	t.Due = &duration
}

func (t Todo) Serialize() []byte {
	jsonVal, err := json.Marshal(t)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return jsonVal
}

type Todos []Todo
