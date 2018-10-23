package models

import (
	"encoding/json"
	"errors"
	"go-basic-rest-api/utils"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/icrowley/fake"
)

type User struct {
	ID           uint       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"created_at,omitempty"`
	UpdatedAt    time.Time  `json:"updated_at,omitempty"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
	Name         string     `json:"name"`
	ImageProfile string     `json:"image_profile"`
	Username     string     `gorm:"not null;unique" json:"username"`
	Password     string     `gorm:"not null" json:"password"`
	Todos        []Todo     `gorm:"ForeignKey:TodoID" json:"todos,omitempty"`
}

func (u *User) ToMap() map[string]interface{} {
	var userMap map[string]interface{} = make(map[string]interface{})

	userMap["id"] = u.ID
	userMap["created_at"] = u.CreatedAt
	userMap["name"] = u.Name
	userMap["image_profile"] = u.ImageProfile
	userMap["username"] = u.Username

	if len(u.Todos) != 0 {
		var todoMaps []map[string]interface{} = make([]map[string]interface{}, len(u.Todos))

		for i, todo := range u.Todos {
			todoMap := todo.ToMap()
			delete(todoMap, "user")
			delete(todoMap, "subcontent")
			delete(todoMap, "content")

			todoMaps[i] = todoMap
		}

		userMap["todos"] = todoMaps
	}

	return userMap
}

func (u User) Serialize() ([]byte, error) {
	return json.Marshal(u.ToMap())
}

func (u User) SerializeUploadImageProfile() ([]byte, error) {
	var payload map[string]string = make(map[string]string)
	payload["image_profile"] = u.ImageProfile

	return json.Marshal(payload)
}

func (u User) ValidValue(checkLogin bool) bool {
	validCreate := len(u.Name) > 4

	if checkLogin {
		validCreate = true
	}

	return len(u.Username) > 6 && len(u.Password) > 3 && validCreate
}

func (u User) UserMapToken() (jwt.MapClaims, error) {
	if !u.ValidValue(true) {
		return nil, errors.New("User not valid, Must has username & password")
	}

	expired := strconv.FormatInt(time.Now().Add(time.Hour*3).Unix(), 10)
	return jwt.MapClaims{
		"username": u.Username,
		"password": u.Password,
		"expired":  expired,
	}, nil
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
	var userMaps []map[string]interface{} = make([]map[string]interface{}, len(users))
	for i, user := range users {
		userMaps[i] = user.ToMap()
	}

	return json.Marshal(userMaps)
}

func (u User) MergeToken(token string) ([]byte, error) {
	userMap := u.ToMap()
	userMap["token"] = token

	return json.Marshal(userMap)
}
