package util

import "server/models"

func SetUser(payload *Payload) *models.MyUser {
	user := new(models.MyUser)
	user.ID = payload.ID
	user.User = payload.User
	return user
}
