package db

import (
	"RoaminRoninXBank/db/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(test *testing.T, account Account) Entry {
	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)
	require.NoError(test, err)
	require.NotEmpty(test, entry)

	require.Equal(test, args.AccountID, entry.AccountID)
	require.Equal(test, args.Amount, entry.Amount)

	require.NotZero(test, entry.ID)
	require.NotZero(test, entry.CreatedAt)
	return entry
}

func TestCreateEntry(test *testing.T) {
	account := createRandomAccount(test)
	createRandomEntry(test, account)
}

func TestGetEntry(test *testing.T) {
	account := createRandomAccount(test)
	entry1 := createRandomEntry(test, account)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(test, err)
	require.NotEmpty(test, entry2)

	require.Equal(test, entry1.ID, entry2.ID)
	require.Equal(test, entry1.AccountID, entry2.AccountID)
	require.Equal(test, entry1.Amount, entry2.Amount)
	require.WithinDuration(test, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(test *testing.T) {
	account := createRandomAccount(test)
	for i := 0; i < 10; i++ {
		createRandomEntry(test, account)
	}

	args := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), args)
	require.NoError(test, err)
	require.Len(test, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(test, entry)
		require.Equal(test, args.AccountID, entry.AccountID)
	}

}
