package db

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	acc1 := createCreateAccount(t)
	acc2 := createCreateAccount(t)

	n := int64(5)
	amount := int64(10)

	errs := make(chan error, n)
	results := make(chan TransferTxResult, n)

	for i := int64(0); i < n; i++ {
		arg := TransferTxParams{
			FromAccountID: acc1.ID,
			ToAccountID:   acc2.ID,
			Amount:        amount,
		}

		go func() {
			var result TransferTxResult
			var err error

			// Retry logic for deadlock
			for j := 0; j < 3; j++ {
				result, err = store.TransferTx(context.Background(), arg)
				if err == nil || !isDeadlockError(err) {
					break
				}
			}

			errs <- err
			results <- result
		}()
	}

	// Validate the results
	for i := int64(0); i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)
	}
}

// Helper function to check for deadlock errors
func isDeadlockError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "deadlock detected")
}
