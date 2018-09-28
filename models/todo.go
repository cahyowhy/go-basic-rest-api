package models

import (
	"time"
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

type Todos []Todo
