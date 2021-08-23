package services

import "server/models"

type MyRichTextBlockRepository interface {
	Create(*models.MyRichTextBlock) (*models.MyRichTextBlock, error)
	GetRichTextBlockByBlockID(string) (*models.MyRichTextBlock, error)
}

type MyRichTextBlockService interface {
	GetRichTextBlockByBlockID(blockID string) (*models.MyRichTextBlock, error)
}

type myRichTextBlockServiceImpl struct {
	user *models.MyUser
	repo MyRichTextBlockRepository
}

func NewMyRichTextBlockService(repo MyRichTextBlockRepository, user *models.MyUser) MyRichTextBlockService {
	res := &myRichTextBlockServiceImpl{
		user: user,
		repo: repo,
	}
	return res
}

func (m *myRichTextBlockServiceImpl) GetRichTextBlockByBlockID(blockID string) (*models.MyRichTextBlock, error) {
	data, err := m.repo.GetRichTextBlockByBlockID(blockID)
	if err != nil {
		return nil, err
	}
	return data, nil
}
