package services

import (
	"context"
	"fmt"
	"server/models"

	"github.com/dstotijn/go-notion"
)

type BlockRepository interface {
	AddChild(data *models.MyBlock) (*models.MyBlock, error)
}

type BlockService interface {
	AddChild(key string, blockID string, myBlock []models.MyBlock) (notion.Block, error)
	GetChildren(key string, blockID string) (notion.BlockChildrenResponse, error)
}

type blockServiceImpl struct {
	user *models.User
	repo BlockRepository
}

func NewBlockService(repo BlockRepository, user *models.User) BlockService {
	res := &blockServiceImpl{
		user: user,
		repo: repo,
	}
	return res
}

func (u *blockServiceImpl) AddChild(key string, blockID string, myBlock []models.MyBlock) (notion.Block, error) {
	client := notion.NewClient(key)
	//自分のデータベースに保存する処理の追加(最初に追加してその後notionAPIに合わせた方がロールバックしやすい)

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
	res, err := client.AppendBlockChildren(context.Background(), blockID, []notion.Block{*block})
	if err != nil {
		return notion.Block{}, err
	}
	return res, nil
	// 自分のデータベースに保存する
	// myBlock := new(models.MyBlock)
	// myBlock.Block = res
	// myBlock.DisplayTime = time.Now()

	// db := c.Get("Tx").(*gorm.DB)
	// result := db.Create(myBlock)

	// return c.JSON(http.StatusOK, result)

	// return new(models.MyBlock), nil
}
func (u *blockServiceImpl) GetChildren(key string, blockID string) (notion.BlockChildrenResponse, error) {
	client := notion.NewClient(key)
	query := new(notion.PaginationQuery)
	block, err := client.FindBlockChildrenByID(context.Background(), blockID, query)
	if err != nil {
		fmt.Println(err)
		return notion.BlockChildrenResponse{}, err
	}
	if len(block.Results) == 0 {
		fmt.Println(err)
		return notion.BlockChildrenResponse{}, err
	}

	return block, nil
}

// {
// 	"children": [
// 		{
// 			"object": "block",
// 			"type": "heading_2",
// 			"heading_2": {
// 				"text": [{ "type": "text", "text": { "content": "Lacinato kale" } }]
// 			}
// 		},
// 		{
// 			"object": "block",
// 			"type": "paragraph",
// 			"paragraph": {
// 				"text": [
// 					{
// 						"type": "text",
// 						"text": {
// 							"content": "Lacinato kale is a variety of kale with a long tradition in Italian cuisine, especially that of Tuscany. It is also known as Tuscan kale, Italian kale, dinosaur kale, kale, flat back kale, palm tree kale, or black Tuscan palm.",
// 							"link": { "url": "https://en.wikipedia.org/wiki/Lacinato_kale" }
// 						}
// 					}
// 				]
// 			}
// 		}
// 	]
// }
