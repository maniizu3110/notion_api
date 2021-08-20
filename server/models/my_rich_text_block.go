package models

import (
	"github.com/jinzhu/gorm"
	"server/api/notion"
)

type MyRichTextBlock struct {
	gorm.Model
	BlockID string
	notion.RichTextBlock
	MyRichText MyRichText
}

func ChangeToMyRichTextBlock(richTextBlock *notion.RichTextBlock, blockID string) *MyRichTextBlock {
	newMyRichTextBlock := new(MyRichTextBlock)
	newMyRichTextBlock.RichTextBlock = *richTextBlock
	newMyRichTextBlock.BlockID = blockID
	return newMyRichTextBlock
}
