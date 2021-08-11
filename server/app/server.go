package app

import (
	"fmt"
	"server/util"

	"github.com/jinzhu/gorm"
)

type Server struct {
	config     util.Config
	db         *gorm.DB
	tokenMaker util.Maker
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, database *gorm.DB) (*Server, error) {
	tokenMaker, err := util.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		db:         database,
		tokenMaker: tokenMaker,
	}
	return server, nil
}
