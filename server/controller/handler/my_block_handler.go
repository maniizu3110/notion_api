package handler

import (
	"errors"
	"fmt"
	"net/http"
	"server/controller/repositories"
	"server/controller/services"
	"server/models"
	"server/util"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func AssignMyBlockHandlers(g *echo.Group) {
	g = g.Group("", func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			config := c.Get("Ck").(util.Config)
			db := c.Get(config.DatabaseKey).(*gorm.DB)
			user := c.Get("user").(*models.MyUser)
			r := repositories.NewMyBlockRepository(config, db, user.ID)
			bs := services.NewBlockService(r, user)
			rr := repositories.NewMyRichTextBlockRepository(config, db, user.ID)
			rtr := repositories.NewMyRichTextRepository(config, db, user.ID)
			s := services.NewMyBlockService(r, user, bs, rr, rtr)
			c.Set("Service", s)
			return handler(c)
		}
	})
	g.POST("/", GetAndCreateMyBlockChildrenHandler)
	g.GET("/:id", GetMyBlockByIDHandler)
	g.GET("/:id/children", GetMyBlockChildenByIDHandler)
	g.GET("/:id/childrenInfo", GetMyBlockChildenInfoByIDHandler)
	g.GET("/", GetAllMyBlockHandler)
}

func GetAndCreateMyBlockChildrenHandler(c echo.Context) error {
	service := c.Get("Service").(services.MyBlockService)
	var params struct {
		SecretKey     string
		ParentBlockID string
	}
	if err := c.Bind(&params); err != nil {
		return errors.New("送られた情報を取得できませんでした")
	}
	data, err := service.GetAndCreateChildren(params.SecretKey, params.ParentBlockID)
	if err != nil {
		return errors.New("ブロックの取得に失敗しました")
	}
	return c.JSON(http.StatusOK, data)
}

func GetAllMyBlockHandler(c echo.Context) error {
	service := c.Get("Service").(services.MyBlockService)
	blocks, err := service.GetAllBlocks()
	if err != nil {
		return errors.New("ブロックの取得に失敗しました")
	}
	return c.JSON(http.StatusOK, blocks)
}

func GetMyBlockByIDHandler(c echo.Context) error {
	service := c.Get("Service").(services.MyBlockService)
	blockID := c.Param("id")
	blocks, err := service.GetChildrenByID(blockID)
	if err != nil {
		return errors.New("ブロックの取得に失敗しました")
	}
	return c.JSON(http.StatusOK, blocks)
}

//GetMyBlockChildenByIDHandlerは親IDに紐づくMyBlockの配列を返す(my_rich_text_blockの情報は返さない)
func GetMyBlockChildenByIDHandler(c echo.Context) error {
	service := c.Get("Service").(services.MyBlockService)
	blockID := c.Param("id")
	blocks, err := service.GetMyBlockChildrenByParentID(blockID)
	if err != nil {
		fmt.Println(err)
		return errors.New("ブロックの取得に失敗しました")
	}
	return c.JSON(http.StatusOK, blocks)
}

//GetMyBlockChildenInfoByIDHandlerは親IDに紐づくMyBlockの配列を返す(my_rich_text_blockの情報を返す)
func GetMyBlockChildenInfoByIDHandler(c echo.Context) error {
	service := c.Get("Service").(services.MyBlockService)
	blockID := c.Param("id")
	blocks, err := service.GetMyBlockChildrenInfoByParentID(blockID)
	var result []models.MyBlock
	for i := range blocks {
		result = append(result, *blocks[i])
	}
	if err != nil {
		fmt.Println(err)
		return errors.New("ブロックの取得に失敗しました")
	}
	return c.JSON(http.StatusOK, result)
}

// func GetAllMyChildrenHandler(c echo.Context) error {
// 	service := c.Get("Service").(services.MyBlockService)
// 	blocks, err := service.GetAllBlocks()
// 	if err != nil {
// 		return errors.New("ブロックの取得に失敗しました")
// 	}
// 	return c.JSON(http.StatusOK, blocks)
// }
