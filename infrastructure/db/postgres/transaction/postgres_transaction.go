package pgtrans

import (
	"context"
	"database/sql"
	"delivery-stream-service/infrastructure/db/postgres"
	"delivery-stream-service/internal/transaction"

	"github.com/jackc/pgx/v5"
)

type PostgresTransaction struct {
	tx pgx.Tx
}

func NewPostgresTransaction(tx pgx.Tx) *PostgresTransaction {
	return &PostgresTransaction{tx: tx}
}

func (p *PostgresTransaction) Exec(ctx context.Context, query string, args ...any) (transaction.Result, error) {
	cmdTag, err := p.tx.Exec(ctx, query, args...)
	return PostgresResult{cmdTag}, err
}

func (p *PostgresTransaction) QueryRow(ctx context.Context, query string, args ...any) transaction.Row {
	return PostgresRow{p.tx.QueryRow(ctx, query, args...)}
}

func (p *PostgresTransaction) Query(ctx context.Context, query string, args ...any) (transaction.Rows, error) {
	rows, err := p.tx.Query(ctx, query, args...)
	if err != nil {
		if rows != nil {
			rows.Close()
		}
		return nil, err
	}
	return PostgresRows{rows}, nil
}

func (p *PostgresTransaction) Commit(ctx context.Context) error {
	return p.tx.Commit(ctx)
}

func (p *PostgresTransaction) Rollback(ctx context.Context) error {
	return p.tx.Rollback(ctx)
}
