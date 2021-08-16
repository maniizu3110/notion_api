package services

import (
	"errors"
	"fmt"
	"server/models"
	"server/util"
)

type UserRepository interface {
	Create(data *models.User) (*models.User, error)
	SearchByUserName(username string) ([]models.User, error)
}

type UserService interface {
	Create(data *models.User) (*models.User, error)
	Login(data *models.User) (*models.LoginUserResponse, error)
	SearchByUserName(username string, password string) ([]models.User, error)
}

type userServiceImpl struct {
	config     util.Config
	repo       UserRepository
	tokenMaker util.Maker
}

func NewUserService(repo UserRepository, config util.Config, tokenMaker util.Maker) UserService {
	res := &userServiceImpl{
		config:     config,
		repo:       repo,
		tokenMaker: tokenMaker,
	}
	return res
}

func (u *userServiceImpl) Create(data *models.User) (*models.User, error) {
	//naked passowrd in this point
	password := data.HashedPassword
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, errors.New("ハッシュパスワードの生成に失敗したためユーザーを登録できませんでした")
	}
	data.HashedPassword = hashedPassword
	//TODO:ハッシュパスワードはリターンしない
	newData, err := u.repo.Create(data)
	if err != nil {
		return nil, errors.New("ユーザーの作成に失敗しました")
	}
	return newData, nil
}
func (u *userServiceImpl) SearchByUserName(username string, password string) ([]models.User, error) {
	users, err := u.repo.SearchByUserName(username)
	if err != nil {
		return nil, errors.New("ユーザー名での検索に失敗しました")
	}
	return users, nil
}

//TODO:実装大きすぎるので区切る
//TODO:同じ名前で登録できないようにバリデーションを修正する
func (u *userServiceImpl) Login(inputData *models.User) (*models.LoginUserResponse, error) {
	users, err := u.repo.SearchByUserName(inputData.User)
	fmt.Printf("%v", users)
	if err != nil {
		return nil, errors.New("ユーザー名での検索に失敗しました")
	}
	if len(users) > 1 {
		return nil, errors.New("複数のユーザーが該当しました。システム管理者に連絡してください。")
	}
	user := users[0]
	err = util.CheckPassword(inputData.HashedPassword, user.HashedPassword)
	if err != nil {
		return nil, errors.New("パスワードのチェックに失敗しました")
	}

	duration := u.config.AccessTokenDuration

	accessToken, err := u.tokenMaker.CreateToken(user.ID, user.User, duration)
	if err != nil {
		return nil, errors.New("トークンの作成に失敗しました")
	}
	userRes := user.NewUserResponse()
	res := &models.LoginUserResponse{
		AccessToken: accessToken,
		User:        userRes,
	}

	return res, nil
}
