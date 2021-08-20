package database

import (
	"server/models"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.MyBlock{})
	db.AutoMigrate(&models.MyUser{})
	db.AutoMigrate(&models.MyRichTextBlock{})
	db.AutoMigrate(&models.MyRichText{})
}
