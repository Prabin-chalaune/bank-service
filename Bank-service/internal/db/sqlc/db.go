// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	// "github.com/jackc/pgx/v5/pgxpool"
)

type DBTX interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

// type Store struct {
//     Queries *Queries
//     Pool    *pgxpool.Pool
// }
// type Store interface {
// 	Querier      // Embed the SQLC-generated Querier interface
// 	ExecTx(ctx context.Context, fn func(*Queries) error) error
// }

// type SQLStore struct {
//     *Queries
// 	pool *pgxpool.Pool // Database connection pool
// }

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

type Queries struct {
	db DBTX
}

func (q *Queries) WithTx(tx pgx.Tx) *Queries {
	return &Queries{
		db: tx,
	}
}

// func NewStore(pool *pgxpool.Pool) Store {
//     return &SQLStore{
//         Queries: New(pool),
//         pool:    pool,
//     }
// }

// // ExecTx executes a function within a database transaction.
// func (store *SQLStore) ExecTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := store.pool.BeginTx(ctx, pgx.TxOptions{})
// 	if err != nil {
// 		return err
// 	}

// 	q := New(tx) // Create a new Queries instance for the transaction
// 	err = fn(q)  // Execute the transaction logic
// 	if err != nil {
// 		if rbErr := tx.Rollback(ctx); rbErr != nil {
// 			return rbErr
// 		}
// 		return err
// 	}

// 	return tx.Commit(ctx)
// }

