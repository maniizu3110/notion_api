package models

import "github.com/jinzhu/gorm"


type User struct {
	gorm.Model
	User string `json:"user" gorm:"uniqueIndex" validate:"required" `
	HashedPassword string `json:"password" validate:"required,min=6"`
}