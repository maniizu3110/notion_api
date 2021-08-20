package services

import "server/models"


type MyRichTextBlockRepository interface {
	Create(*models.MyRichTextBlock) (*models.MyRichTextBlock, error)
}