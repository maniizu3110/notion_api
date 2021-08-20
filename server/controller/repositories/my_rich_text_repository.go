package repositories

import (
	"errors"
	"server/controller/services"
	"server/models"
	"server/util"

	"github.com/jinzhu/gorm"
)

type myRichTextRepositoryImpl struct {
	config util.Config
	db     *gorm.DB
}

func NewMyRichTextRepository(config util.Config, db *gorm.DB, userID uint) services.MyRichTextRepository {
	res := &myRichTextRepositoryImpl{
		config: config,
		db:     db,
	}
	return res
}

func (u *myRichTextRepositoryImpl) Create(data *models.MyRichText) (*models.MyRichText, error) {
	if err := u.db.Create(data).Error; err != nil {
		return nil, errors.New("ブロックの追加に失敗しました")
	}
	return data, nil
}
