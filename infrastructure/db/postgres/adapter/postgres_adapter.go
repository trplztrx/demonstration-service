package pgadap

import (
    "context"
    "delivery-stream-service/internal/transaction"

    "github.com/jackc/pgconn"
    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"
)

// основной адаптер для выполнения запросов через пул подключений
type PostgresExecutor struct {
    db *pgxpool.Pool
}

func NewPostgresExecutor(db *pgxpool.Pool) *PostgresExecutor {
    return &PostgresExecutor{db: db}
}

func (p *PostgresExecutor) Exec(ctx context.Context, query string, args ...interface{}) (transaction.Result, error) {
    cmdTag, err := p.db.Exec(ctx, query, args...)
    return PostgresResult{cmdTag}, err
}

func (p *PostgresExecutor) QueryRow(ctx context.Context, query string, args ...interface{}) transaction.Row {
    return PostgresRow{p.db.QueryRow(ctx, query, args...)}
}

func (p *PostgresExecutor) Query(ctx context.Context, query string, args ...interface{}) (transaction.Rows, error) {
    rows, err := p.db.Query(ctx, query, args...)
    if err != nil {
        if rows != nil {
            rows.Close()
        }
        return nil, err
    }
    return PostgresRows{rows}, nil
}
