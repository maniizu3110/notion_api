package handler

import (
	"fmt"
	"net/http"
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
	//create DI
	
	g.POST("/", CereateUserHandler)
	g.POST("/login", LoginUserHandler)
}
//TODO:同じ名前のユーザーを登録することはできない
func CereateUserHandler(c echo.Context)error{
	user := new(models.User)
	err := c.Bind(user)
	//naked password in this point
	password := user.HashedPassword
	if err != nil {
		return err
	}
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,err)
	}
	user.HashedPassword = hashedPassword
	db := c.Get("Tx").(*gorm.DB)
	//TODO:ハッシュパスワードはリターンしない
	result := db.Create(user)
	
	return c.JSON(http.StatusOK,result)
}

func LoginUserHandler(c echo.Context)error{
	inputedUser := new(models.User)
	err := c.Bind(inputedUser)
	db := c.Get("Tx").(*gorm.DB)
	if err != nil {
		return err
	}
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
	duration,_ := time.ParseDuration("2m")
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