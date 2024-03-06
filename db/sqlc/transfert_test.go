package db

import (
	"context"
	util "simplebank/Util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)


func createRandomTransfer(t *testing.T, accountTo Account, accountFrom Account) Transfer {
	arg := CreateTransfertParams{
		Amount: util.RandomMoney(),
		FromAccountID: accountFrom.ID,
		ToAccountID: accountTo.ID,
	}

	transfert, err := testQueries.CreateTransfert(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfert)

	require.Equal(t, transfert.Amount, arg.Amount)
	require.Equal(t, transfert.ToAccountID, arg.ToAccountID)
	require.Equal(t, transfert.FromAccountID, arg.FromAccountID)

	require.NotZero(t, transfert.ID)
	require.NotZero(t, transfert.CreatedAt)

	return transfert
}

func TestCreateTransfert(t *testing.T) {
	accountFrom := createRandomAccount(t)
	accountTo := createRandomAccount(t)
	createRandomTransfer(t, accountTo, accountFrom)
}

func TestGetTransfert(t *testing.T) {
	accountFrom := createRandomAccount(t)
	accountTo := createRandomAccount(t)

	transfert1 := createRandomTransfer(t, accountTo, accountFrom)

	transfert2, err := testQueries.GetTransfert(context.Background(), transfert1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfert2)

	require.Equal(t, transfert1.ID, transfert2.ID)
	require.Equal(t, transfert1.FromAccountID, transfert2.FromAccountID)
	require.Equal(t, transfert1.ToAccountID, transfert2.ToAccountID)
	require.Equal(t, transfert1.Amount, transfert2.Amount)
	require.WithinDuration(t, transfert1.CreatedAt, transfert2.CreatedAt, time.Second)
}

func TestListTransfert(t *testing.T) {
	accountFrom := createRandomAccount(t)
	accountTo := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, accountTo, accountFrom)
	}

	arg := ListTransfertParams{
		FromAccountID: accountFrom.ID,
		Limit: 5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfert(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)

	require.Len(t, transfers, 5)
}
