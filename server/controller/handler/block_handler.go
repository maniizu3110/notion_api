package handler

import (
	"errors"
	"net/http"
	"server/controller/repositories"
	"server/controller/services"
	"server/models"
	"server/util"

	"github.com/dstotijn/go-notion"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func AssignBlockHandlers(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			config := c.Get("Ck").(util.Config)
			db := c.Get(config.DatabaseKey).(*gorm.DB)
			user := c.Get("user").(*models.User)
			r := repositories.NewBlockRepository(config, db)
			s := services.NewBlockService(r, user)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.GET("/", GetBlockByIDHandler)
	g.POST("/", CreateBlockChildrenHandler)
}

func GetBlockByIDHandler(c echo.Context) error {
	service := c.Get("Service").(services.BlockService)
	var params struct {
		SecretKey     string
		ParentBlockID string
	}
	if err := c.Bind(&params); err != nil {
		return errors.New("送られた情報を取得できませんでした")
	}
	data, err := service.GetChildren(params.SecretKey, params.ParentBlockID)
	if err != nil {
		return errors.New("ブロックの取得に失敗しました")
	}
	return c.JSON(http.StatusOK, data)

}

func CreateBlockChildrenHandler(c echo.Context) error {
	service := c.Get("Service").(services.BlockService)
	//大文字と小文字を統一しておきたい
	var params struct {
		SecretKey     string
		ParentBlockID string
		Blocks      []notion.Block
	}
	if err := c.Bind(&params); err != nil {
		return errors.New("送られた情報を取得できませんでした")
	}
	block, err := service.AddChild(params.SecretKey, params.ParentBlockID, params.Blocks)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, block)
}
