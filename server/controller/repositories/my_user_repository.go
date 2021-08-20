package repositories

import (
	"errors"
	"server/controller/services"
	"server/models"
	"server/util"

	"github.com/jinzhu/gorm"
)

type myUserRepositoryImpl struct {
	config util.Config
	db     *gorm.DB
}

func NewMyUserRepository(config util.Config, db *gorm.DB) services.MyUserRepository {
	res := &myUserRepositoryImpl{
		config: config,
		db:     db,
	}
	return res
}

func (u *myUserRepositoryImpl) Create(data *models.MyUser) (*models.MyUser, error) {
	if err := u.db.Create(data).Error; err != nil {
		return nil, errors.New("ユーザーの作成に失敗しました")
	}
	return data, nil
}

func (u *myUserRepositoryImpl) SearchByUserName(username string) ([]models.MyUser, error) {
	var users []models.MyUser
	if err := u.db.Where("user = ?", username).Find(&users).Error; err != nil {
		return nil, errors.New("ユーザーネームでの検索に失敗しました")
	}
	if len(users) < 1 {
		return nil, errors.New("該当するユーザーはいませんでした")
	}
	return users, nil
}
