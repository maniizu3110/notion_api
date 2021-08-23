package models

import (
	"github.com/jinzhu/gorm"
	"server/api/notion"
)

type MyRichText struct {
	gorm.Model
	notion.RichText
	MyRichTextBlockID uint
}

func ChangeToMyRichText(richText notion.RichText, blockID uint) *MyRichText {
	newMyRichText := new(MyRichText)
	newMyRichText.RichText = richText
	newMyRichText.MyRichTextBlockID = blockID
	return newMyRichText
}
