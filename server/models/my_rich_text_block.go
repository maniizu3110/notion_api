package models

import "server/api/notion"


type MyRichTextBlock struct {
	Model
	BlockID string
	notion.RichTextBlock
	MyRichText MyRichText
}

func ChangeToMyRichTextBlock(richTextBlock *notion.RichTextBlock,blockID string)(*MyRichTextBlock){
	newMyRichTextBlock := new(MyRichTextBlock)
	newMyRichTextBlock.RichTextBlock = *richTextBlock
	newMyRichTextBlock.BlockID = blockID
	return newMyRichTextBlock
}