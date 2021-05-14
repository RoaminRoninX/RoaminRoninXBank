package db

import (
	"RoaminRoninXBank/db/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(test *testing.T, account1, account2 Account) Transfer {
	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(test, err)
	require.NotEmpty(test, transfer)

	require.Equal(test, args.FromAccountID, transfer.FromAccountID)
	require.Equal(test, args.ToAccountID, transfer.ToAccountID)
	require.Equal(test, args.Amount, transfer.Amount)

	require.NotZero(test, transfer.ID)
	require.NotZero(test, transfer.CreatedAt)

	return transfer
}

func TestCreateTranser(test *testing.T) {
	account1 := createRandomAccount(test)
	account2 := createRandomAccount(test)
	createRandomTransfer(test, account1, account2)
}

func TestGetTransfer(test *testing.T) {
	account1 := createRandomAccount(test)
	account2 := createRandomAccount(test)
	transfer1 := createRandomTransfer(test, account1, account2)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(test, err)
	require.NotEmpty(test, transfer2)
	require.Equal(test, transfer1.ID, transfer2.ID)
	require.Equal(test, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(test, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(test, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(test, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestLisTransfer(test *testing.T) {
	account1 := createRandomAccount(test)
	account2 := createRandomAccount(test)

	for i := 0; i < 5; i++ {
		createRandomTransfer(test, account1, account2)
		createRandomTransfer(test, account2, account1)
	}

	args := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account1.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), args)
	require.NoError(test, err)
	require.Len(test, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(test, transfer)
		require.True(test, transfer.FromAccountID == account1.ID || transfer.ToAccountID == account1.ID)
	}

}
