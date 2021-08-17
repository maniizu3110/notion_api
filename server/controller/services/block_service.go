package services

import (
	"context"
	"server/models"

	"github.com/dstotijn/go-notion"
)

type BlockRepository interface {
	AddChild(data *models.MyBlock) (*models.MyBlock, error)
}

type BlockService interface {
	AddChild(key string, blockID string, myBlocks []notion.Block) (notion.Block, error)
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

func (u *blockServiceImpl) AddChild(key string, blockID string, blocks []notion.Block) (notion.Block, error) {
	client := notion.NewClient(key)
	res, err := client.AppendBlockChildren(context.Background(), blockID, blocks)
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
		return notion.BlockChildrenResponse{}, err
	}
	if len(block.Results) == 0 {
		return notion.BlockChildrenResponse{}, err
	}

	return block, nil
}

