package api

import (
	"database/sql"
	"errors"
	"net/http"
	"github.com/sathwikshetty33/golang_bank/token"
	"github.com/gin-gonic/gin"
	"github.com/sathwikshetty33/golang_bank/db/sqlc"
)

type transferRequest struct {
    FromAccID    int64 `json:"from_account_id" binding:"required,min=1"`
    ToAccID int64 `json:"to_account_id" binding:"required,min=1"`
	Ammount int64 `json:"amount" binding:"required,min=1"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (s *server) createTransfer(c *gin.Context) {
	var req transferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
}
	fromAccount, valid := s.validAcc(c,req.FromAccID,req.Currency)
	if !valid {
		return
	}
	authPayload := c.MustGet(authorizationPayloadKey).(*token.Payload)
	if fromAccount.Owner != authPayload.Username{
		err := errors.New("account does not belong to you")
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	_, valid = s.validAcc(c,req.ToAccID,req.Currency)
	if !valid {
		return
	}
	arg := db.TransferTxParams{
		FromAccountID: req.FromAccID,
		ToAccountID: req.ToAccID,
		Amount: req.Ammount,
	}
	result, err := s.store.TransferTx(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	} 
	c.JSON(http.StatusOK, result)
}

func(server *server) validAcc(c *gin.Context, accID int64, currency string)(db.Account, bool) {
	acc, err := server.store.GetAccount(c, accID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
		} else {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		
	}
	return acc, false
}
if acc.Currency != currency {
	err := errors.New("account currency does not match")
	c.JSON(http.StatusBadRequest, errorResponse(err))
	return acc, false
}
return acc, true
}