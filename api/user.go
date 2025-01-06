package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/sathwikshetty33/golang_bank/db/sqlc"
	"github.com/sathwikshetty33/golang_bank/db/util"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	FullName string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (s *server) createUser(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPs,err := util.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return	
	}
	arg := db.CreateUserParams{
		Username:    req.Username,
		FullName:  req.FullName,
		Email: req.Email,
		Pass : hashedPs,
	}
	account, err := s.store.CreateUser(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	c.JSON(http.StatusOK, account)
}
