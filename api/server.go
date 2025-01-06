package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/sathwikshetty33/golang_bank/db/sqlc"
)

type server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *server {
	server := &server{store: store}
	router := gin.Default()
	router.POST("/accounts", server.createAccount)
	router.POST("/createuser", server.createUser)
	router.POST("/maketransfer", server.createTransfer)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	server.router = router
	return server
}

func (s *server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}