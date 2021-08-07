package models

import (
)

type User struct {
	Model
	User string 
	HashedPassword string
}
