package models

import "server/api/notion"


type MyRichTextBlock struct {
	BlockID string
	notion.RichTextBlock
}

func ChangeToMyRichTextBlock(richTextBlock *notion.RichTextBlock,blockID string)(*MyRichTextBlock){
	newMyRichTextBlock := new(MyRichTextBlock)
	newMyRichTextBlock.RichTextBlock = *richTextBlock
	newMyRichTextBlock.BlockID = blockID
	return newMyRichTextBlock
}