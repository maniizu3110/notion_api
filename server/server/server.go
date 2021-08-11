package server

import (
	"fmt"
	"server/util"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type Server struct {
	config     util.Config
	db     		*gorm.DB
	tokenMaker 	util.Maker
	router     	*echo.Echo
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, database *gorm.DB) (*Server, error) {
	tokenMaker, err := util.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		db:      database,
		tokenMaker: tokenMaker,
	}

	server.SetUpRouter()
	return server, nil
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string){
	server.router.Logger.Fatal(server.router.Start(":"+ address));
}