package app

import (
	"server/controller/handler"
	"server/middlewares"

	"github.com/labstack/echo/v4"
)

func InitRouter(serverConf *Server) {
	e := echo.New()
	middlewares.CORS(e)

	//認証不要
	{
		g := e.Group("/api/v1", middlewares.SetConfWithTokenMaker(serverConf.config, serverConf.db, serverConf.tokenMaker))
		handler.AssignUserHandlers(g.Group("/user"))
	}

	//認証が必要
	{
		g := e.Group("/api/v1", middlewares.SetConf(serverConf.config, serverConf.db),middlewares.AuthMiddleware(serverConf.tokenMaker))
		handler.AssignNotionDatabaseHandlers(g.Group("/databases")) // auth ok
		handler.AssignMyBlockHandlers(g.Group("/myblock"))
		handler.AssignBlockHandlers(g.Group("/block"))
	}
	e.Debug = true
	e.Logger.Fatal(e.Start(serverConf.config.ServerAddress))

}
