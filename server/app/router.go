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
		g := e.Group("/api/v1", middlewares.SetConfWithTokenMaker(serverConf.config, serverConf.db, serverConf.tokenMaker))
		handler.AssignUserHandlers(g.Group("/user"))
	}

	//認証が必要なエンドポイント
	{
		//TODO:複数middlewareある時のnext(c)の挙動に注意
		//開発環境だったら基本的に認証なしでテストできるようにしておきたい
		// g := e.Group("/api/v1", middlewares.SetConf(serverConf.config, serverConf.db))
		g := e.Group("/api/v1", middlewares.SetConf(serverConf.config, serverConf.db),middlewares.AuthMiddleware(serverConf.tokenMaker))
		handler.AssignNotionDatabaseHandlers(g.Group("/databases")) // auth ok
		handler.AssignBrockHandlers(g.Group("/block"))
	}
	e.Debug = true
	e.Logger.Fatal(e.Start(serverConf.config.ServerAddress))

}
