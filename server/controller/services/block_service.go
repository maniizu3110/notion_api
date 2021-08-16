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
	AddChild(key string,blockID string,block models.MyBlock) (*models.MyBlock, error)
	GetChildren(key string,blockID string) (notion.BlockChildrenResponse, error)
}

type blockServiceImpl struct {
	user	  *models.User
	repo       BlockRepository
}

func NewBlockService(repo BlockRepository, user *models.User) BlockService {
	res := &blockServiceImpl{
		user:	user,
		repo:       repo,
	}
	return res
}

func (u *blockServiceImpl) AddChild(key string,blockID string,block models.MyBlock) (*models.MyBlock, error) {
	client := notion.NewClient(key)

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

	return new(models.MyBlock), nil
}
func (u *blockServiceImpl) GetChildren(key string,blockID string) (notion.BlockChildrenResponse, error) {
	client := notion.NewClient(key)
	query := new(notion.PaginationQuery)
	block, err := client.FindBlockChildrenByID(context.Background(), blockID, query)
	if err != nil {
		fmt.Println(err)
		return notion.BlockChildrenResponse{},err
	}
	if len(block.Results)==0{
		fmt.Println(err)
		return notion.BlockChildrenResponse{},err
	}

	return block, nil
}
