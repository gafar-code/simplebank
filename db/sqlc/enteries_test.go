package db

import (
	"context"
	"testing"

	"github.com/gafar-code/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateEntries(t *testing.T) {
	arg := CreateEntryParams{
		AccountID: 1,
		Amount:    utils.RandomMoney(),
	}

	account, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Amount, account.Amount)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestListEntries(t *testing.T) {
	arg := ListentriesParams{
		Limit:  1,
		Offset: 1,
	}

	entries, err := testQueries.Listentries(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entries)
}

func TestUpdateEntries(t *testing.T) {
	arg := UpdateEntryParams{
		ID:     1,
		Amount: utils.RandomMoney(),
	}

	entry, err := testQueries.UpdateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

}

func TestDeleteEntries(t *testing.T) {
	err := testQueries.DeleteEntry(context.Background(), 5)

	require.NoError(t, err)
}
