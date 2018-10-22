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
	"github.com/tidwall/sjson"
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

func (u User) SerializeUploadImageProfile() ([]byte, error) {
	var payload map[string]string;
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
	jsonVal, err := json.Marshal([]User{})
	clonedJson := jsonVal

	for index, user := range users {
		userJson, _ := user.Serialize()
		clonedJson, err = sjson.SetRawBytes(clonedJson, strconv.Itoa(index), userJson)
	}

	return clonedJson, err
}

func (u User) MergeToken(token string) ([]byte, error) {
	jsonInitial, err := u.Serialize()

	clonedJson := jsonInitial
	clonedJson, err = sjson.SetBytes(clonedJson, "token", token)

	return clonedJson, err
}
