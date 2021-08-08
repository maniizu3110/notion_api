package handler

import (
	"fmt"
	"net/http"
	"server/models"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func AssignUserHandlers(g *echo.Group) {
	//ここでDI実装する
	
	g.POST("/", CereateUserHandler)
}

func CereateUserHandler(c echo.Context)error{
	user := new(models.User)
	err := c.Bind(user)
	fmt.Printf("%+v",user)
	if err != nil {
		return err
	}
	fmt.Printf("%v",user)
	return c.JSON(http.StatusOK,user)
}

func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", fmt.Errorf("failed to hash password: %w", err)
    }
    return string(hashedPassword), nil
}