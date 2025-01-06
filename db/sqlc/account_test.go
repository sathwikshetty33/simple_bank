package db

import (
	"context"
	"database/sql"
	"testing"
	"time"
	"github.com/sathwikshetty33/golang_bank/db/util"
	"github.com/stretchr/testify/require"
)
func createCreateAccount(t *testing.T) Account {
	acc := createCreateUser(t)
	arg := CreateAccountParams{
		Owner: acc.Username,
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
	acc := createCreateUser(t)
	arg := CreateAccountParams{
		Owner: acc.Username,
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

func TestUpdateAccount(t *testing.T) {
	acc1 := createCreateAccount(t)
	arg := UpdateAccountParams{
		ID: acc1.ID,
		Balance: util.RandomMoney(),
	}
	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account2.Owner, acc1.Owner)
	require.Equal(t, account2.ID, acc1.ID)
	require.NotEqual(t, account2.Balance, acc1.Balance)
	require.WithinDuration(t, account2.CreatedAt, acc1.CreatedAt,time.Second)
}      
func TestDeleteAccount(t *testing.T) {
	acc1 := createCreateAccount(t)
	err := testQueries.DeleteAccount(context.Background(), acc1.ID)
	require.NoError(t, err)
	account2, err := testQueries.GetAccount(context.Background(), acc1.ID)
	require.Error(t, err)
	require.Empty(t, account2)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestListAccounts(t *testing.T) {
	for i:= 0;i < 10;i++ {
		createCreateAccount(t)
	}
	arg := ListAccountsParams{
	Limit: 5,
	Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}