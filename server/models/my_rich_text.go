
package models

import "server/api/notion"


type MyRichText struct {
	Model
	MyRichTextBlockID uint
	notion.RichText
}

func ChangeToMyRichText(richText notion.RichText,blockID uint)(*MyRichText){
	newMyRichText := new(MyRichText)
	newMyRichText.RichText = richText
	newMyRichText.MyRichTextBlockID = blockID
	return newMyRichText
}