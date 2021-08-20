package repositories

import (
	"errors"
	"server/controller/services"
	"server/models"
	"server/util"

	"github.com/jinzhu/gorm"
)

type myRichTextBlockRepositoryImpl struct {
	config util.Config
	db     *gorm.DB
}

func NewMyRichTextBlockRepository(config util.Config, db *gorm.DB, userID uint) services.MyRichTextBlockRepository {
	res := &myRichTextBlockRepositoryImpl{
		config: config,
		db:     db.Where("user_id = ?", userID),
	}
	return res
}

func (u *myRichTextBlockRepositoryImpl) Create(data *models.MyRichTextBlock) (*models.MyRichTextBlock, error) {
	if err := u.db.Create(data).Error; err != nil {
		return nil, errors.New("ブロックの追加に失敗しました")
	}
	return data, nil
}
