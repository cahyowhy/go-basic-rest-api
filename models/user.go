package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Todos []Todo `gorm:"ForeignKey:TodoID" json:"todos,omitempty"`
}

type Users []User