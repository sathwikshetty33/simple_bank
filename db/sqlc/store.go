package db

import (
	"context"
	"database/sql"
	"fmt"
	
)

type Store struct {
	*Queries
	db *sql.DB
}

// Constructor for Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

// Executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted, // or sql.LevelReadUncommitted
	})
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}


// Input parameters for a transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// Output result of a transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// Performs a money transfer transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
    var result TransferTxResult

    err := store.execTx(ctx, func(q *Queries) error {
        var err error

        // Step 1: Get and lock both accounts first, in a consistent order
        if arg.FromAccountID < arg.ToAccountID {
            result.FromAccount, err = q.GetAccountForUpdate(ctx, arg.FromAccountID)
            if err != nil {
                return err
            }
            result.ToAccount, err = q.GetAccountForUpdate(ctx, arg.ToAccountID)
            if err != nil {
                return err
            }
        } else {
            result.ToAccount, err = q.GetAccountForUpdate(ctx, arg.ToAccountID)
            if err != nil {
                return err
            }
            result.FromAccount, err = q.GetAccountForUpdate(ctx, arg.FromAccountID)
            if err != nil {
                return err
            }
        }
		
        // Step 2: Create transfer record
        result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
            FromAccID: arg.FromAccountID,
            ToAccID:   arg.ToAccountID,
            Amount:    arg.Amount,
        })
        if err != nil {
            return err
        }

        // Step 3: Create account entries
        result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
            AccID:  arg.FromAccountID,
            Amount: -arg.Amount,
        })
        if err != nil {
            return err
        }

        result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
            AccID:  arg.ToAccountID,
            Amount: arg.Amount,
        })
        if err != nil {
            return err
        }

        // Step 4: Update account balances in the same order as locks
        if arg.FromAccountID < arg.ToAccountID {
            // Update sender's balance
            result.FromAccount, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
                ID:      arg.FromAccountID,
                Balance: -arg.Amount,  // SQLC will handle adding this to current balance
            })
            if err != nil {
                return err
            }

            // Update receiver's balance
            result.ToAccount, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
                ID:      arg.ToAccountID,
                Balance: arg.Amount,  // SQLC will handle adding this to current balance
            })
            if err != nil {
                return err
            }
        } else {
            // Update receiver's balance
            result.ToAccount, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
                ID:      arg.ToAccountID,
                Balance: arg.Amount,
            })
            if err != nil {
                return err
            }

            // Update sender's balance
            result.FromAccount, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
                ID:      arg.FromAccountID,
                Balance: -arg.Amount,
            })
            if err != nil {
                return err
            }
        }

        return nil
    })

    return result, err
}