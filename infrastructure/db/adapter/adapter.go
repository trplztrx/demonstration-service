package adapter

import (
	"context"
	"delivery-stream-service/internal/transaction"
)

type DBAdapter struct {
    executor transaction.SQLExecutor
}

func NewDBAdapter(executor transaction.SQLExecutor) *DBAdapter {
    return &DBAdapter{
        executor: executor,
    }
}

func (p *DBAdapter) Exec(ctx context.Context, query string, args ...any) (transaction.Result, error) {
    return p.executor.Exec(ctx, query, args...)
}

func (p *DBAdapter) QueryRow(ctx context.Context, query string, args ...any) transaction.Row {
    return p.executor.QueryRow(ctx, query, args...)
}

func (p *DBAdapter) Query(ctx context.Context, query string, args ...any) (transaction.Rows, error) {
    rows, err := p.executor.Query(ctx, query, args...)
    if err != nil {
        if rows != nil {
            rows.Close()
        }
        return nil, err
    }
    return rows, nil
}
