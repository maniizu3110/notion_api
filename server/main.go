package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/dstotijn/go-notion"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	server.Run()
	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	//TODO:テスト用
	integrations_api_key := os.Getenv("INTEGRATIONS_APIKEY")
	database_id := os.Getenv("DATABASE_ID")
	
	client := notion.NewClient(integrations_api_key)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		page, err := client.FindDatabaseByID(context.Background(), database_id)
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
