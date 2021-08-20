package services

import "server/models"

type MyRichTextRepository interface {
	Create(*models.MyRichText) (*models.MyRichText, error)
}
