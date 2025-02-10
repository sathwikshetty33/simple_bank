package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)
	type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" `
	Password string `json:"password" binding:"required,min=6"`
}
 type renewAccessTokenResponse struct{
        AccessToken string `json:"access_token"`
		AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`


 }

 func (server *server) renewAccessToken(c *gin.Context) {
    var req renewAccessTokenRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, errorResponse(err))
        return
    }
	refreshPayload, err :=server.tokenMaker.VerifyToken(req.RefreshToken)
    if err != nil {
        c.JSON(http.StatusUnauthorized, errorResponse(err)) // Unauthorized for invalid password
        return
    }
	session, err := server.store.GetSession(c, refreshPayload.ID)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, errorResponse(err))
        } else {
            c.JSON(http.StatusInternalServerError, errorResponse(err))
        }
        return
    }

    if session.IsBlocked{
		err := fmt.Errorf("blocked session")
		c.JSON(http.StatusForbidden, errorResponse(err))
	}
	if session.Username != refreshPayload.Username{
		err := fmt.Errorf("username mismatch")
		c.JSON(http.StatusUnauthorized, errorResponse(err))
	}
	if session.RefreshToken != req.RefreshToken{
		err := fmt.Errorf("access token mismatch")
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		}
		if time.Now().After(session.ExpiresAt){
			err := fmt.Errorf("access token expired")
			c.JSON(http.StatusUnauthorized, errorResponse(err))
		}
    accessToken,accessPayload, err := server.tokenMaker.CreateToken(session.Username, server.config.AccessTokenDuration)
    if err != nil {
        c.JSON(http.StatusInternalServerError, errorResponse(err))
        return
    }
		

    rsp := renewAccessTokenResponse{
        AccessToken: accessToken,
		AccessTokenExpiresAt : accessPayload.ExpiresAt,
    }
    c.JSON(http.StatusOK, rsp)
}
