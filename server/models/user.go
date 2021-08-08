package models

import (
)

type User struct {
	Model
	User string `json:"user" validate:"required"`
	HashedPassword string `json:"password" validate:"required,min=6"`
}
