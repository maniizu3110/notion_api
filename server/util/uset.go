package util

import "server/models"

func SetUser(payload *Payload) *models.User {
	user := new(models.User)
	user.ID = payload.ID
	user.User = payload.User
	return user
}
