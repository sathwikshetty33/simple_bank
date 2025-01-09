package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/sathwikshetty33/golang_bank/db/sqlc"
	"github.com/sathwikshetty33/golang_bank/db/util"
	"github.com/sathwikshetty33/golang_bank/token"
)

type server struct {
	config util.Config
	store *db.Store
	tokenMaker token.Maker
	router *gin.Engine
}

func NewServer(config util.Config, store *db.Store)( *server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.SymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w ",err)
	}
	server := &server{
		config: config,
		store: store,
		tokenMaker: tokenMaker,
		}
	router := gin.Default()
	router.POST("/createuser", server.createUser)
	router.POST("/users/login", server.login)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.POST("/maketransfer", server.createTransfer)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)
	server.router = router
	return server, nil
}

func (s *server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}