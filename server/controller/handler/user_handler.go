package handler

import (
	"errors"
	"fmt"
	"net/http"
	"server/controller/repositories"
	"server/controller/services"
	"server/models"
	"server/util"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/o1egl/paseto"
	"github.com/spf13/viper"
)

func AssignUserHandlers(g *echo.Group) {
	g = g.Group("",func(handler echo.HandlerFunc)echo.HandlerFunc{
		return func(c echo.Context) error {
			config := c.Get("Ck").(util.Config)
			db := c.Get(config.DatabaseKey).(*gorm.DB)
			r := repositories.NewUserRepository(config,db)
			s := services.NewUserService(r)
			c.Set("Service",s)
			return handler(c)
		}
	})
	g.POST("/", CreateUserHandler)
	g.POST("/login", LoginUserHandler)
}
//TODO:同じ名前のユーザーを登録することはできない
//TODO:Createした後にもログイン処理をする
func CreateUserHandler(c echo.Context)error{
	service := c.Get("Service").(services.UserService)
	user := new(models.User)
	err := c.Bind(user)
	data,err := service.Create(user)
	if err != nil {
		return errors.New("ユーザー情報の取得に失敗したため登録できませんでした")
	}
	return c.JSON(http.StatusOK,data)
}

func LoginUserHandler(c echo.Context)error{
	inputedUser := new(models.User)
	err := c.Bind(inputedUser)
	if err != nil {
		return errors.New("ログイン情報の取得に失敗しました")
	}
	db := c.Get("heyhey").(*gorm.DB)
	data := new(models.User)
	if err := db.Where("user = ?", inputedUser.User).First(data).Error; err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = util.CheckPassword(inputedUser.HashedPassword,data.HashedPassword)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	duration,_ := time.ParseDuration("10m")
	//DIから取ってくるように変更(冗長なのでまとめる)
	key := []byte(viper.GetString(`database.token_symmetric_key`))
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  data.User,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	//ここでcreate処理しているので抽象化する
	accessToken,err := paseto.NewV2().Encrypt(key,payload,nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	type loginUserResponse struct {
		AccessToken string       `json:"access_token"`
		User        userResponse `json:"user"`
	}
	resUser := newUserResponse(data)
	res := loginUserResponse{
		AccessToken: accessToken,
		User:        resUser,
	}
	return c.JSON(http.StatusOK,res)
}
type userResponse struct {
	User             string    `json:"user"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user *models.User) userResponse {
	return userResponse{
		User:             user.User,
		CreatedAt:         user.CreatedAt,
	}
}