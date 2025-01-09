package token

import (
	"testing"
	"time"

	"github.com/sathwikshetty33/golang_bank/db/util"
	"github.com/stretchr/testify/require"
)

func TestNewPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)
	require.NotNil(t, maker)
	username := util.RandomOwner()
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)
	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, expiredAt, payload.ExpiresAt, time.Second)

}

func TestExpiredPaestoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)
	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	//require.EqualError(t, err, jwt.NewValidationError("", jwt.ValidationErrorExpired).Error())
	require.Nil(t, payload)

}

// func TestInvalidToken(t *testing.T){
// 	payload, err := NewPayload(util.RandomOwner(), time.Minute)
// 	require.NoError(t, err)
// 	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
// 	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
// 	require.NoError(t, err)
// 	maker, err := NewJWTMaker(util.RandomString(32))
// 	require.NoError(t, err)
// 	payload , err = maker.VerifyToken(token)
// }