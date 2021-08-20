
package models

import "server/api/notion"


type MyRichText struct {
	Model
	RichTextBlockID string
	notion.RichText
}

func ChangeToMyRichText(richText *notion.RichText,blockID string)(*MyRichText){
	newMyRichText := new(MyRichText)
	newMyRichText.RichText = *richText
	newMyRichText.RichTextBlockID = blockID
	return newMyRichText
}