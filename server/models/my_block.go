package models

import (
	"server/api/notion"
	"time"
)

type MyBlock struct {
	notion.Block
	DisplayTime time.Time
	//親ブロックのID
	MyBlockID string 
	UserID      uint
	MyBlocks []MyBlock
	MyRichTextBlock MyRichTextBlock `gorm:"foreignKey:MyBlockID"`
}

type AddChildBlockResponse struct {
	ParentNotionBlock        notion.Block
	RegisteredNotionBlocks   []notion.Block
	RegisteredMyNotionBlocks []MyBlock
}

func ChangeToMyBlock(block notion.Block, user *MyUser,parentBlockID string) *MyBlock {
	myblock := new(MyBlock)
	myblock.Block = block
	myblock.DisplayTime = time.Now()
	myblock.UserID = user.ID
	myblock.MyBlockID = parentBlockID
	return myblock
}
