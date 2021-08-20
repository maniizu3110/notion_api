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

func AssignMyRichTextBlockHandlers(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			config := c.Get("Ck").(util.Config)
			db := c.Get(config.DatabaseKey).(*gorm.DB)
			user := c.Get("user").(*models.MyUser)
			r := repositories.NewMyRichTextBlockRepository(config, db,user.ID)
			s := services.NewMyRichTextBlockService(r, user)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.GET("/:id", GetMyRichTextBlockByBlockIDHandler)
}

func GetMyRichTextBlockByBlockIDHandler(c echo.Context) error{
	service := c.Get("Service").(services.MyRichTextBlockService)
	blockID := c.Param("id")
	blocks, err := service.GetRichTextBlockByBlockID(blockID)
	if err != nil {
		return errors.New("RichTextBlockの取得に失敗しました")
	}
	return c.JSON(http.StatusOK, blocks)

}

