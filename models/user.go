package models

import (
	"encoding/json"
	"strconv"
	"time"
	"log"
	"go-basic-rest-api/utils"
	"github.com/icrowley/fake"
	"github.com/tidwall/sjson"
)

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Name      string     `json:"name"`
	Username  string     `gorm:"not null;unique" json:"username"`
	Password  string     `gorm:"not null" json:"password"`
	Todos     []Todo     `gorm:"ForeignKey:TodoID" json:"todos,omitempty"`
}

func (u User) Serialize() ([]byte, error) {
	jsonVal, err := json.Marshal(u)
	clonedJson := jsonVal
	emits := []string{}

	emits = append(emits, "password")
	for index, _ := range u.Todos {
		emits = append(emits, "todos."+strconv.Itoa(index)+".user")
	}

	for _, element := range emits {
		clonedJson, err = sjson.DeleteBytes(clonedJson, element)
	}

	return clonedJson, err
}

func (u User) ValidValue(checkLogin bool) bool {
	validCreate := len(u.Name) > 4
	
	if checkLogin {
		validCreate = true
	}

	return len(u.Username) > 6 && len(u.Password) > 6 && validCreate
}

func (u *User) FakeIt() {
	u.Name = fake.FullName()
	u.Username = fake.UserName()
	password, err := utils.GeneratePassword("1234")
	
	if err != nil {
		log.Fatal(err)
		password = "123456"
	}

	u.Password = password
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
