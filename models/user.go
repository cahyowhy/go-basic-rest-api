package models

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/icrowley/fake"
	"github.com/tidwall/sjson"
)

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Name      string     `json:"name"`
	Todos     []Todo     `gorm:"ForeignKey:TodoID" json:"todos,omitempty"`
}

func (u User) Serialize() ([]byte, error) {
	jsonVal, err := json.Marshal(u)
	clonedJson := jsonVal
	emits := []string{}

	for index, _ := range u.Todos {
		emits = append(emits, "todos."+strconv.Itoa(index)+".user")
	}

	for _, element := range emits {
		clonedJson, err = sjson.DeleteBytes(clonedJson, element)
	}

	return clonedJson, err
}

func (u *User) FakeIt() {
	u.Name = fake.FullName()
}

type Users []User

func SerializeUsers(users []User) ([]byte, error) {
	jsonVal, err := json.Marshal([]User{})
	clonedJson := jsonVal

	for index, user := range users {
		userJson, _ := user.Serialize()
		clonedJson, err = sjson.SetRawBytes(clonedJson, strconv.Itoa(index), userJson)
	}

	return clonedJson, err
}
