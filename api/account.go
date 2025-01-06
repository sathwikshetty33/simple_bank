package api

import (
	"net/http"
	"github.com/sathwikshetty33/golang_bank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
    Owner    string `json:"owner" binding:"required"`
    Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (s *server) createAccount(c *gin.Context) {
	var req createAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
}
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  0,
		Currency : req.Currency,
	}
	account, err := s.store.CreateAccount(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	} 
	c.JSON(http.StatusOK, account)
}