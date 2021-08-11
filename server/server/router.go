package server

import (
	"server/controller/handler"
	"server/middlewares"

	"github.com/labstack/echo/v4"
)


func(server *Server) SetUpRouter(){
	e := echo.New()
	middlewares.CORS(e)

	//認証なしで利用できるエンドポイント
	{
		g := e.Group("/api/v1")
		handler.AssignUserHandlers(g.Group("/user"))
	}
	
	//認証が必要なエンドポイント
	{
		g := e.Group("/api/v1",middlewares.AuthMiddleware(server.tokenMaker))
		handler.AssignNotionDatabaseHandlers(g.Group("/databases"))                                                  // auth ok
		handler.AssignBrockHandlers(g.Group("/block"))
	}
	server.router = e

}