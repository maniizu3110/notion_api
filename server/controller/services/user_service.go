package services

import (
	"errors"
	"server/models"
	"server/util"
)



type UserRepository interface {
	Create(data *models.User)(*models.User,error)
	SearchByUserName(username string)([]models.User,error)
}

type UserService interface{
	Create(data *models.User)(*models.User,error)
	SearchByUserName(username string,password string)([]models.User,error)

}

type userServiceImpl struct{
	repo UserRepository
}

func NewUserService(repo UserRepository)UserService{
	res := &userServiceImpl{
		repo:repo,
	}
	return res
}

func(u *userServiceImpl) Create(data *models.User)(*models.User,error){
	//naked passowrd in this point
	password := data.HashedPassword
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return nil,errors.New("ハッシュパスワードの生成に失敗したためユーザーを登録できませんでした")
	}
	data.HashedPassword = hashedPassword
	//TODO:ハッシュパスワードはリターンしない
	newData,err := u.repo.Create(data)
	if err != nil {
		return nil,errors.New("ユーザーの作成に失敗しました")
	}
	return newData,nil
}
func(u *userServiceImpl) SearchByUserName(username string,password string)([]models.User,error){
	users,err := u.repo.SearchByUserName(username)
	if err != nil {
		return nil,errors.New("ユーザー名での検索に失敗しました")
	} 
	return users,nil
}

