package models

import (
	"server/api/notion"
	"time"
)

type MyBlock struct {
	//notion.Blockの内容
	Object         string           `json:"object"`
	ID             string           `json:"id,omitempty"`
	Type           notion.BlockType `json:"type"`
	CreatedTime    *time.Time       `json:"created_time,omitempty"`
	LastEditedTime *time.Time       `json:"last_edited_time,omitempty"`
	HasChildren    bool             `json:"has_children,omitempty"`

	Paragraph        *notion.RichTextBlock `json:"paragraph,omitempty"`
	Heading1         *notion.Heading       `json:"heading_1,omitempty"`
	Heading2         *notion.Heading       `json:"heading_2,omitempty"`
	Heading3         *notion.Heading       `json:"heading_3,omitempty"`
	BulletedListItem *notion.RichTextBlock `json:"bulleted_list_item,omitempty"`
	NumberedListItem *notion.RichTextBlock `json:"numbered_list_item,omitempty"`
	ToDo             *notion.ToDo          `json:"to_do,omitempty"`
	Toggle           *notion.RichTextBlock `json:"toggle,omitempty"`
	ChildPage        *notion.ChildPage     `json:"child_page,omitempty"`
	//

	DisplayTime time.Time `json:"display_time"`
	MyBlockID   string    `json:"my_block_id"`
	UserID      uint      `json:"user_id"`
	MyBlocks    []MyBlock
	//TODO:notion.Blockに紐づいているのでそこからとってきたい
	MyRichTextBlock MyRichTextBlock `json:"my_rich_text_block" gorm:"foreignKey:MyBlockID"`
}

type AddChildBlockResponse struct {
	ParentNotionBlock        notion.Block
	RegisteredNotionBlocks   []notion.Block
	RegisteredMyNotionBlocks []MyBlock
}

//TODO：receive指定してやる
func ChangeToMyBlock(block notion.Block, user *MyUser, parentBlockID string) *MyBlock {
	myblock := new(MyBlock)
	myblock.Object = block.Object
	myblock.ID = block.ID
	myblock.Type = block.Type
	myblock.CreatedTime = block.CreatedTime
	myblock.LastEditedTime = block.LastEditedTime
	myblock.HasChildren = block.HasChildren

	myblock.Paragraph = block.Paragraph
	myblock.Heading1 = block.Heading1
	myblock.Heading2 = block.Heading2
	myblock.Heading3 = block.Heading3
	myblock.BulletedListItem = block.BulletedListItem
	myblock.NumberedListItem = block.NumberedListItem
	myblock.ToDo = block.ToDo
	myblock.Toggle = block.Toggle
	myblock.ChildPage = block.ChildPage

	myblock.DisplayTime = time.Now()
	myblock.UserID = user.ID
	myblock.MyBlockID = parentBlockID
	return myblock
}
