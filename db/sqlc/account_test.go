package db

import (
	"context"
	"testing"

	"github.com/gafar-code/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
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

func TestListAccounts(t *testing.T) {
	arg := ListAccountsParams{
		Limit:  5,
		Offset: 6,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, accounts)
}

func TestUpdateAccount(t *testing.T) {
	arg := UpdateAccountParams{
		ID:      1,
		Balance: utils.RandomMoney(),
	}

	account, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Balance, account.Balance)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	err := testQueries.DeleteAccount(context.Background(), 8)

	require.NoError(t, err)
}
