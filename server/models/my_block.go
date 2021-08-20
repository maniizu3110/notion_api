package models

import (
	"server/api/notion"
	"time"
)

type MyBlock struct {
	notion.Block
	DisplayTime time.Time
	UserID      uint
	MyRichTextBlock MyRichTextBlock
}

type AddChildBlockResponse struct {
	ParentNotionBlock        notion.Block
	RegisteredNotionBlocks   []notion.Block
	RegisteredMyNotionBlocks []MyBlock
}

func ChangeToMyBlock(block notion.Block, user *MyUser) *MyBlock {
	myblock := new(MyBlock)
	myblock.Block = block
	myblock.DisplayTime = time.Now()
	myblock.UserID = user.ID
	return myblock
}
