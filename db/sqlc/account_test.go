package db

import (
	"context"
	"testing"
	"time"

	"github.com/sathwikshetty33/golang_bank/db/util"
	"github.com/stretchr/testify/require"
)
func createCreateAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner: util.RandomOwner(),
		Balance: util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	acc1 := createCreateAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), acc1.ID)
	require.NoError(t, err)
	require.Equal(t, account2.Owner, acc1.Owner)
	require.Equal(t, account2.ID, acc1.ID)
	require.NotEmpty(t, account2)
	require.Equal(t, account2.Balance, acc1.Balance)
	require.WithinDuration(t, account2.CreatedAt, acc1.CreatedAt, time.Second)
}