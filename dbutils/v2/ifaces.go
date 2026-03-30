package dbutils

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Querier is an interface for database queries.
type Querier interface {
	Query(context.Context, string, ...any) (pgx.Rows, error)
	QueryRow(context.Context, string, ...any) pgx.Row
}

// Execer is an interface for database statements.
type Execer interface {
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
}

// Txer is an interface for database transaction.
type Txer interface {
	BeginTx(context.Context, pgx.TxOptions) (pgx.Tx, error)
}

var (
	_ Querier = (*pgxpool.Pool)(nil)
	_ Execer  = (*pgxpool.Pool)(nil)
	_ Txer    = (*pgxpool.Pool)(nil)

	_ Querier = (*pgxpool.Pool)(nil)
	_ Execer  = (*pgxpool.Pool)(nil)
)
