package services

import (
	"errors"
	"server/models"
	"server/util"
)

type MyUserRepository interface {
	Create(data *models.MyUser) (*models.MyUser, error)
	SearchByUserName(username string) ([]models.MyUser, error)
}

type MyUserService interface {
	Create(data *models.MyUser) (*models.MyUser, error)
	Login(data *models.MyUser) (*models.LoginUserResponse, error)
	SearchByUserName(username string, password string) ([]models.MyUser, error)
}

type myUserServiceImpl struct {
	config     util.Config
	repo       MyUserRepository
	tokenMaker util.Maker
}

func NewMyUserService(repo MyUserRepository, config util.Config, tokenMaker util.Maker) MyUserService {
	res := &myUserServiceImpl{
		config:     config,
		repo:       repo,
		tokenMaker: tokenMaker,
	}
	return res
}

func (u *myUserServiceImpl) Create(data *models.MyUser) (*models.MyUser, error) {
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
func (u *myUserServiceImpl) SearchByUserName(username string, password string) ([]models.MyUser, error) {
	users, err := u.repo.SearchByUserName(username)
	if err != nil {
		return nil, errors.New("ユーザー名での検索に失敗しました")
	}
	return users, nil
}

//TODO:実装大きすぎるので区切る
//TODO:同じ名前で登録できないようにバリデーションを修正する
func (u *myUserServiceImpl) Login(inputData *models.MyUser) (*models.LoginUserResponse, error) {
	users, err := u.repo.SearchByUserName(inputData.User)
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
