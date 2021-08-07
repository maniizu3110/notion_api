package middlewares

import (
	"net/http"
	"server/database"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func CORS(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))
}

func SetDB(db *gorm.DB) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
			tx := db.Begin()
			//TODO：API叩かれたときではなく、サーバー起動時に呼ばれるべき処理
			database.Migrate(tx)
			c.Set("Tx", tx)
			if err := next(c); err != nil {
				tx.Rollback()
				logrus.Debugln("Transaction Rollback")
				return err
			}
			tx.Commit()
			return nil
        }
    }
}