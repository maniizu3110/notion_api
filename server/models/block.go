package models

import (
	"server/api/notion"
	"time"
)

type MyBlock struct {
	Model
	notion.Block
	//カスタムで表示する時間
	DisplayTime time.Time
}


type AddChildBlockResponse struct{
	ParentNotionBlock notion.Block
	RegisteredNotionBlocks []notion.Block
	RegisteredMyNotionBlocks []MyBlock 
}