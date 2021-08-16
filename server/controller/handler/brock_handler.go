package handler

import (
	"context"
	"net/http"
	"server/models"
	"time"

	"github.com/dstotijn/go-notion"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func AssignBrockHandlers(g *echo.Group) {
	//ここでDI実装する
	g.GET("/", GetBrockByIDHandler)
	g.POST("/", CreateBrockChildrenHandler)
}

func GetBrockByIDHandler(c echo.Context) error {
	//この辺りはユーザー情報に紐づけるデータベースを作成しておいて、そこからどれを使うか選択できるようにしておく
	testKey := "secret_sJNm41y3NfWLd59bUvSeyYXlxKL4VmwvpYgHikItzhB"
	client := notion.NewClient(testKey)
	query := new(notion.PaginationQuery)
	block, err := client.FindBlockChildrenByID(context.Background(), "b8643af68f9340438e99b30f4586a332", query)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, block)

}

func CreateBrockChildrenHandler(c echo.Context) error {
	testKey := "secret_sJNm41y3NfWLd59bUvSeyYXlxKL4VmwvpYgHikItzhB"
	// 	//データベースに紐づく値を入れる(状態管理でログイン時に持っておきたい情報)
	client := notion.NewClient(testKey)

	childText := new(notion.Text)
	childText.Content = "test"

	parentText := new(notion.RichText)
	parentText.Type = "text"
	parentText.Text = childText

	textBlock := new(notion.RichTextBlock)
	textBlock.Text = []notion.RichText{*parentText}

	block := new(notion.Block)
	block.Object = "block"
	block.Type = "paragraph"
	block.Paragraph = textBlock
	res, err := client.AppendBlockChildren(context.Background(), "349f28e31be94105b461ccde34cd6496", []notion.Block{*block})
	if err != nil {
		//一回成功メッセージ出した方が良さげ(本当はもしミスがあったら直したい)
		return err
	}
	// 自分のデータベースに保存する
	myBlock := new(models.MyBlock)
	myBlock.Block = res
	myBlock.DisplayTime = time.Now()

	db := c.Get("Tx").(*gorm.DB)
	result := db.Create(myBlock)

	return c.JSON(http.StatusOK, result)

}
