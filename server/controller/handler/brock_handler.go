package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dstotijn/go-notion"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)



func AssignBrockHandlers(g *echo.Group) {
	//ここでDI実装する
	g.GET("/", GetBrockByIDHandler)
	g.POST("/", CreateBrockChildrenHandler)
}

func GetBrockByIDHandler(c echo.Context)error{
	//この辺はフロントから指定したい
	testKey := viper.GetString(`notion_test.key`)
	client := notion.NewClient(testKey)
	query := new(notion.PaginationQuery)
	block, err := client.FindBlockChildrenByID(context.Background(), "349f28e31be94105b461ccde34cd6496", query)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, block)
}

func CreateBrockChildrenHandler(c echo.Context)error{
	testKey := viper.GetString(`notion_test.key`)
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
	res,err := client.AppendBlockChildren(context.Background(),"349f28e31be94105b461ccde34cd6496",[]notion.Block{*block})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return c.JSON(http.StatusOK, res)


}

//www.notion.so/api-test-349f28e31be94105b461ccde34cd6496#32584c68d2a24e48b7ab93e7c4f15a85
	// {
    //         "object": "block",
    //         "id": "f66b6724-872b-4e74-85d1-90612a45e047",
    //         "type": "paragraph",
    //         "created_time": "2021-08-01T04:14:00Z",
    //         "last_edited_time": "2021-08-01T04:14:00Z",
    //         "paragraph": {
    //             "text": [
    //                 {
    //                     "type": "text",
    //                     "annotations": {
    //                         "color": "default"
    //                     },
    //                     "plain_text": "hey",
    //                     "text": {
    //                         "content": "hey"
    //                     }
    //                 }
    //             ]
    //         }
    //     },


