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
	server.router = router
	return server
}

func (s *server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}