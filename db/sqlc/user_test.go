package db

import (
	"context"
	"testing"
	"time"

	"github.com/sathwikshetty33/golang_bank/db/util"
	"github.com/stretchr/testify/require"
)
func createCreateUser(t *testing.T) User {
		var name string = util.RandomOwner()
		hashedPassword, err := util.HashPassword(util.RandomOwner())
		require.NoError(t, err)
		arg := CreateUserParams{
			Username: name, 
			Pass: hashedPassword,
			FullName:	name + " " + util.RandomOwner(),
			Email: name + string('@')+util.RandomOwner()+".com",
		}
		account, err := testQueries.CreateUser(context.Background(), arg)
		if err != nil {
			require.NoError(t, err)
		}
		return account
}
func TestCreateUser(t *testing.T) {
	var name string = util.RandomOwner()
	hashedPassword, err := util.HashPassword(util.RandomOwner())
	require.NoError(t, err)
	arg := CreateUserParams{
		Username: name, 
		Pass: hashedPassword,
		FullName:	name + " " + util.RandomOwner(),
		Email: name + string('@')+util.RandomOwner()+".com",
	}
	account, err := testQueries.CreateUser(context.Background(), arg)
	if err != nil {
		require.NoError(t, err)
	}
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Username, account.Username)
	require.Equal(t, arg.Pass, account.Pass)
	require.Equal(t, arg.FullName, account.FullName)
	require.NotZero(t, account.CreatedAt)
}

func TestGetUser(t *testing.T) {
	acc1 := createCreateUser(t)
	account2, err := testQueries.GetUser(context.Background(), acc1.Username)
	require.NoError(t, err)
	require.Equal(t, account2.Username, acc1.Username)
	require.Equal(t, account2.Pass, acc1.Pass)
	require.NotEmpty(t, account2)
	require.Equal(t, account2.FullName, acc1.FullName)
	require.WithinDuration(t, account2.CreatedAt, acc1.CreatedAt, time.Second)
}
