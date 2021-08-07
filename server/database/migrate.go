package database

import (
	"server/models"
	"github.com/jinzhu/gorm"
)




func Migrate(db *gorm.DB){
	db.AutoMigrate(&models.MyBlock{})
}