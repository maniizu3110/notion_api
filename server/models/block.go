package models

import (
	"time"

	"github.com/dstotijn/go-notion"
)

type MyBlock struct {
	notion.Block
	//カスタムで表示する時間
	DisplayTime time.Time
}
