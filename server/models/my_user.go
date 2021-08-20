package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type MyUser struct {
	gorm.Model
	User           string `json:"user" gorm:"unique not null"`
	HashedPassword string `json:"-" validate:"min=6"`
	MyBlocks []MyBlock
}

type LoginUserResponse struct {
	AccessToken string       `json:"access_token" gorm:"not null"`
	User        UserResponse `json:"user"`
}

type Payload struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type UserResponse struct {
	User      string    `json:"user"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *MyUser) NewUserResponse() UserResponse {
	return UserResponse{
		User:      u.User,
		CreatedAt: u.CreatedAt,
	}
}
