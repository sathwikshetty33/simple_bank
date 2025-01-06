package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sathwikshetty33/golang_bank/db/sqlc"
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

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *server) getAccount(c *gin.Context) {
	var req getAccountRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
}
	account, err := s.store.GetAccount(c, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
		} else {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return
	} 
	c.JSON(http.StatusOK, account)

}
type listAccountRequest struct {
	page_id int32 `form:"page_id" binding:"required,min=1"`
	page_size int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (s *server) listAccount(c *gin.Context) {
	var req listAccountRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
}
	arg := db.ListAccountsParams{
		Limit: int32(req.page_size),
		Offset: int32((req.page_id - 1) * req.page_size),
	}
	accounts, err:= s.store.ListAccounts(c, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
		} else {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	return
	} 
	c.JSON(http.StatusOK, accounts)

}
