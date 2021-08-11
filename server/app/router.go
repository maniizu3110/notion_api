package app

import (
	"server/controller/handler"
	"server/middlewares"

	"github.com/labstack/echo/v4"
)

func InitRouter(serverConf *Server) {
	e := echo.New()
	middlewares.CORS(e)

	//認証なしで利用できるエンドポイント
	{
		g := e.Group("/api/v1", middlewares.SetConf(serverConf.config, serverConf.db))
		handler.AssignUserHandlers(g.Group("/user"))
	}

	//認証が必要なエンドポイント
	{
		g := e.Group("/api/v1", middlewares.SetConf(serverConf.config, serverConf.db))
		handler.AssignNotionDatabaseHandlers(g.Group("/databases")) // auth ok
		handler.AssignBrockHandlers(g.Group("/block"))
	}
	e.Debug = true
	e.Logger.Fatal(e.Start(serverConf.config.ServerAddress))

}
