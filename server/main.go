package main

import (
	"context"
	"net/http"

	"github.com/dstotijn/go-notion"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	//notionのAPIに接続してデータをとってくる処理に変更
	client := notion.NewClient("secret_sJNm41y3NfWLd59bUvSeyYXlxKL4VmwvpYgHikItzhB")
	e.GET("/", func(c echo.Context) error {
		// TODO:環境変数をgitignoreから読み出す
		page, err := client.FindDatabaseByID(context.Background(), "6be60152910541968272879f95ca6594")
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, page)
	})
	e.GET("/brock", func(c echo.Context) error {
		// TODO:環境変数をgitignoreから読み出す
		query := new(notion.PaginationQuery)
		block, err := client.FindBlockChildrenByID(context.Background(), "5edbea1047824fa38a42f815489f0206",query)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, block)
	})
	e.Logger.Fatal(e.Start(":8080"))

}

//データをとってくることができるcurlコマンド
