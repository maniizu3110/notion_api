package handler

import (
	"net/http"
	"server/models"
	"server/util"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func AssignUserHandlers(g *echo.Group) {
	//create DI
	
	g.POST("/", CereateUserHandler)
}

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
	result := db.Create(user)

	return c.JSON(http.StatusOK,result)
}