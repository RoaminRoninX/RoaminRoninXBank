package db

import (
	"RoaminRoninXBank/db/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(test *testing.T) {
	createRandomAccount(test)
}

func createRandomAccount(test *testing.T) Account {
	args := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(test, err)
	require.NotEmpty(test, account)
	require.Equal(test, args.Owner, account.Owner)
	require.Equal(test, args.Balance, account.Balance)
	require.Equal(test, args.Currency, account.Currency)
	require.NotZero(test, account.ID)
	require.NotZero(test, account.CreatedAt)

	return account
}

func TestGetAcount(test *testing.T) {
	account1 := createRandomAccount(test)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(test, err)
	require.NotEmpty(test, account2)
	require.Equal(test, account1.ID, account2.ID)
	require.Equal(test, account1.Owner, account2.Owner)
	require.Equal(test, account1.Balance, account2.Balance)
	require.Equal(test, account1.Currency, account2.Currency)
	require.WithinDuration(test, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(test *testing.T) {
	account1 := createRandomAccount(test)
	args := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(test, err)
	require.NotEmpty(test, account2)
	require.Equal(test, account1.ID, account2.ID)
	require.Equal(test, account1.Owner, account2.Owner)
	require.Equal(test, args.Balance, account2.Balance)
	require.Equal(test, account1.Currency, account2.Currency)
	require.WithinDuration(test, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(test *testing.T) {
	account1 := createRandomAccount(test)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(test, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(test, err)
	require.EqualError(test, err, sql.ErrNoRows.Error())
	require.Empty(test, account2)
}

func TestListAccount(test *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(test)
	}

	args := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), args)
	require.NoError(test, err)
	require.Len(test, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(test, account)
	}

}
