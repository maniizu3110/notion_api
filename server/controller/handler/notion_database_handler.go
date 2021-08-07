package handler

import (
	"context"
	"net/http"

	"github.com/dstotijn/go-notion"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func AssignNotionDatabaseHandlers(g *echo.Group) {
	//ここでDI実装する
	g.GET("/", GetDatabaseByID)
}

func GetDatabaseByID(c echo.Context)error{
	testKey := viper.GetString(`notion_test.key`)
	testDB := viper.GetString(`notion_test.db_id`)
	client := notion.NewClient(testKey)
			page, err := client.FindDatabaseByID(context.Background(), testDB)
		if err != nil {
			return err
		}
	return c.JSON(http.StatusOK, page)

	
}