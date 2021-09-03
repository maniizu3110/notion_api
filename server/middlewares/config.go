package middlewares

import (
	"server/database"
	"server/util"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// SetConfMiddleware set a config and db in context
func SetConf(config util.Config, db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			//TODO:txいらなそう？？migrate部分と合わせて確認する
			tx := db.Begin()
			//TODO：API叩かれたときではなく、サーバー起動時に呼ばれるべき処理?
			database.Migrate(tx)
			c.Set(config.DatabaseKey, tx)
			c.Set("Ck", config)
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
func SetConfWithTokenMaker(config util.Config, db *gorm.DB, tokenMaker util.Maker) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tx := db.Begin()
			database.Migrate(tx)
			c.Set(config.DatabaseKey, tx)
			c.Set("Ck", config)
			c.Set("Tk", tokenMaker)
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
