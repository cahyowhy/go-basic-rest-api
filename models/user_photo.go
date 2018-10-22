package models

import (
	"encoding/json"
	"time"
)

type UserPhoto struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	UserID    uint       `json:"user_id"`
	User      User       `gorm:"association_autoupdate:false;association_autocreate:false" json:"user,omitempty"`
	Path      string     `json:"path"`
}

func (u UserPhoto) ValidValue() bool {
	return u.UserID != 0
}

func (u UserPhoto) Serialize() ([]byte, error) {
	return json.Marshal(u)
}
