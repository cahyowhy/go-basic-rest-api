package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

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

func (u User) serialize() []byte {
	jsonVal, err := json.Marshal(u)
	clonedJson := jsonVal
	emits := []string{}

	if err != nil {
		fmt.Println(err)
		return nil
	}

	for index, element := range u.Todos {
		if element.ID == 0 {
			emits = append(emits, "todos."+strconv.Itoa(index)+".user")
		}
	}

	for _, element := range emits {
		clonedJson, err = sjson.DeleteBytes(clonedJson, element)

		if err != nil {
			fmt.Println(err)
			return nil
		}
	}

	return clonedJson
}

type Users []User
