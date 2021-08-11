package database

import (
	"github.com/jinzhu/gorm"
	"server/models"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.MyBlock{})
	db.AutoMigrate(&models.User{})
}
