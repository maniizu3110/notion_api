package services

import (
	"context"
	"server/api/notion"
	"server/models"
)

type BlockRepository interface {
	AddChild(data *models.MyBlock) (*models.MyBlock, error)
}

type BlockService interface {
	AddChild(key string, blockID string, blocks []notion.Block) (notion.Block, error)
	GetChildren(key string, blockID string) (notion.BlockChildrenResponse, error)
}

type blockServiceImpl struct {
	user *models.MyUser
	repo BlockRepository
}

func NewBlockService(repo BlockRepository, user *models.MyUser) BlockService {
	res := &blockServiceImpl{
		user: user,
		repo: repo,
	}
	return res
}

func (u *blockServiceImpl) AddChild(key string, blockID string, blocks []notion.Block) (notion.Block, error) {
	client := notion.NewClient(key)
	parentBlock, err := client.AppendBlockChildren(context.Background(), blockID, blocks)

	if err != nil {
		return notion.Block{}, err
	}
	return parentBlock, nil
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

