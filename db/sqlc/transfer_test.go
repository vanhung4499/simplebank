package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/vanhung4499/simplebank/util"
	"testing"
	"time"
)

func createRandomTransaction(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transaction, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, arg.FromAccountID, transaction.FromAccountID)
	require.Equal(t, arg.ToAccountID, transaction.ToAccountID)
	require.Equal(t, arg.Amount, transaction.Amount)

	require.NotZero(t, transaction.ID)
	require.NotZero(t, transaction.CreatedAt)

	return transaction
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransaction(t, account1, account2)
}

func TestGetTransaction(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transaction1 := createRandomTransaction(t, account1, account2)

	transaction2, err := testQueries.GetTransfer(context.Background(), transaction1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transaction2)

	require.Equal(t, transaction1.ID, transaction2.ID)
	require.Equal(t, transaction1.FromAccountID, transaction2.FromAccountID)
	require.Equal(t, transaction1.ToAccountID, transaction2.ToAccountID)
	require.Equal(t, transaction1.Amount, transaction2.Amount)
	require.WithinDuration(t, transaction1.CreatedAt, transaction2.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		createRandomTransaction(t, account1, account2)
		createRandomTransaction(t, account2, account1)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account1.ID,
		Limit:         5,
		Offset:        5,
	}

	transactions, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transactions, 5)

	for _, transaction := range transactions {
		require.NotEmpty(t, transaction)
		require.True(t, transaction.FromAccountID == account1.ID || transaction.ToAccountID == account1.ID)
	}
}
