package models

import (
)

type User struct {
	Model
	User string `json:"user"`
	HashedPassword string `json:"hashed_password"`
}
