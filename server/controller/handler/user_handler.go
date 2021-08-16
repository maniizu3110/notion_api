package handler

import (
	"errors"
	"net/http"
	"server/controller/repositories"
	"server/controller/services"
	"server/models"
	"server/util"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func AssignUserHandlers(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			config := c.Get("Ck").(util.Config)
			tokenMaker := c.Get("Tk").(util.Maker)
			db := c.Get(config.DatabaseKey).(*gorm.DB)
			r := repositories.NewUserRepository(config, db)
			s := services.NewUserService(r, config, tokenMaker)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("/", CreateUserHandler)
	g.POST("/login", LoginUserHandler)
}

//TODO:同じ名前のユーザーを登録することはできない
//TODO:Createした後にもログイン処理をする
func CreateUserHandler(c echo.Context) error {
	service := c.Get("Service").(services.UserService)
	user := new(models.User)
	err := c.Bind(user)
	data, err := service.Create(user)
	if err != nil {
		return errors.New("ユーザー情報の取得に失敗したため登録できませんでした")
	}
	return c.JSON(http.StatusOK, data)
}

func LoginUserHandler(c echo.Context) error {
	service := c.Get("Service").(services.UserService)
	inputedUser := new(models.User)
	err := c.Bind(inputedUser)
	if err != nil {
		return errors.New("ログイン情報の取得に失敗しました")
	}
	res, err := service.Login(inputedUser)
	if err != nil {
		return errors.New("ログインできませんでした")
	}

	return c.JSON(http.StatusOK, res)
}
