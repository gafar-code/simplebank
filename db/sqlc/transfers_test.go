package db

import (
	"context"
	"testing"

	"github.com/gafar-code/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	arg := CreateTransferParams{
		FromAccountID: 1,
		ToAccountID:   3,
		Amount:        utils.RandomMoney(),
	}

	account, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Amount, account.Amount)
	require.Equal(t, arg.FromAccountID, account.FromAccountID)
	require.Equal(t, arg.ToAccountID, account.ToAccountID)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestListTransfers(t *testing.T) {
	arg := ListTransfersParams{
		Limit:  1,
		Offset: 1,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfers)
}

func TestUpdateTransaction(t *testing.T) {
	arg := UpdateTransferParams{
		ID:     2,
		Amount: utils.RandomMoney(),
	}

	transfer, err := testQueries.UpdateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
}

func TestDeleteTransfer(t *testing.T) {
	err := testQueries.DeleteTransfer(context.Background(), 5)

	require.NoError(t, err)
}
