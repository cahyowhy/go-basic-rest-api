package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Name      string     `json:"name"`
	Completed bool       `json:"completed"`
	Due       *time.Time `gorm:"default:null" json:"due"`
	UserID    uint       `json:"user_id"`
	User      User       `gorm:"association_autoupdate:false;association_autocreate:false"`
}

func (todo *Todo) completed() {
	todo.Completed = true
}

func (todo *Todo) unCompleted() {
	todo.Completed = false
}

type Todos []Todo
