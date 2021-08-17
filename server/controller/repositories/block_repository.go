package repositories

import (
	"errors"
	"server/controller/services"
	"server/models"
	"server/util"

	"github.com/jinzhu/gorm"
)

type blockRepositoryImpl struct {
	config util.Config
	db     *gorm.DB
}

func NewBlockRepository(config util.Config, db *gorm.DB) services.BlockRepository {
	res := &blockRepositoryImpl{
		config: config,
		db:     db,
	}
	return res
}

func (u *blockRepositoryImpl) AddChild(data *models.MyBlock) (*models.MyBlock, error) {
	if err := u.db.Create(data).Error; err != nil {
		return nil, errors.New("ブロックの追加に失敗しました")
	}
	return data, nil
}
