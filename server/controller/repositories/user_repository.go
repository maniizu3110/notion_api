package repositories

import (
	"errors"
	"server/controller/services"
	"server/models"
	"server/util"

	"github.com/jinzhu/gorm"
)

type userRepositoryImpl struct {
	config util.Config
	db *gorm.DB
}

func NewUserRepository(config util.Config,db *gorm.DB) services.UserRepository{
	res := &userRepositoryImpl{
		config:config,
		db:db,
	}
	return res
}


func (u *userRepositoryImpl)Create(data *models.User)(*models.User,error){
	if err := u.db.Create(data).Error;err!=nil{
		return nil,errors.New("ユーザーの作成に失敗しました")
	}
	return data,nil
}

func (u *userRepositoryImpl)SearchByUserName(username string)([]models.User,error){
	var users []models.User
	if err := u.db.Where("user = ?",username).Find(&users).Error; err!= nil {
		return nil, errors.New("ユーザーネームでの検索に失敗しました")
	}
	if len(users) < 1 {
		return nil,errors.New("該当するユーザーはいませんでした")
	}
	return users,nil
}