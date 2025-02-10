package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/sathwikshetty33/golang_bank/db/sqlc"
	"github.com/sathwikshetty33/golang_bank/db/util"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	FullName string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
type userResponse struct {
	Username string `json:"username"`
FullName string `json:"fullname"`
Email   string `json:"email" binding:"required,email"`
PasswordChangedAt time.Time `json:"password_changed_at"`
CreatedAt time.Time `json:"created_at"`
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
	resp := newUserResponse(account)
	c.JSON(http.StatusOK, resp)
}
func newUserResponse(account db.User) userResponse {
	return userResponse{
		Username : account.Username,
	FullName : account.FullName,
	Email    : account.Email,
	PasswordChangedAt : account.PassChanged,
	CreatedAt : account.CreatedAt,
	}
}
type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}
 type loginUserResponse struct{
		SessionID uuid.UUID `json:"session_id"`
        AccessToken string `json:"access_token"`
		AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
        RefreshToken string `json:"refresh_token"`
		RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
		User userResponse `json:"user"`

 }

 func (server *server) login(c *gin.Context) {
    var req loginUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, errorResponse(err))
        return
    }

    user, err := server.store.GetUser(c, req.Username)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, errorResponse(err))
        } else {
            c.JSON(http.StatusInternalServerError, errorResponse(err))
        }
        return
    }

    err = util.CheckPasswordHash(req.Password, user.Pass)
    if err != nil {
        c.JSON(http.StatusUnauthorized, errorResponse(err)) // Unauthorized for invalid password
        return
    }

    accessToken,accessPayload, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
    if err != nil {
        c.JSON(http.StatusInternalServerError, errorResponse(err))
        return
    }
	Token ,refreshPayload ,err := server.tokenMaker.CreateToken(user.Username, server.config.RefreshTokenDuration)
	if err != nil {
        c.JSON(http.StatusInternalServerError, errorResponse(err))
        return
    }
	session, err := server.store.CreateSession(c, db.CreateSessionParams{
		ID :     refreshPayload.ID,
	Username     : user.Username,
	RefreshToken  :Token ,
	UserAgent    : "",
	ClientIp     : "",
	IsBlocked    : false,
	ExpiresAt  : refreshPayload.ExpiresAt,
	})
	if err != nil {
        c.JSON(http.StatusInternalServerError, errorResponse(err))
        return
    }
    rsp := loginUserResponse{
		SessionID : session.ID,
        AccessToken: accessToken,
		AccessTokenExpiresAt : accessPayload.ExpiresAt,
        RefreshToken : Token,
		RefreshTokenExpiresAt : refreshPayload.ExpiresAt,
		User:        newUserResponse(user),
    }
    c.JSON(http.StatusOK, rsp)
}
