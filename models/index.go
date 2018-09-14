package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Todo{}, &User{})
	db.Model(&Todo{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	return db
}
