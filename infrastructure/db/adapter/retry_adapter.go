package adapter

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IPostgresRetryAdapter interface {
	Exec(ctx context.Context, sql string, args ...any) (cmdTag pgconn.CommandTag, err error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Rows
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

type PostgresRetryAdapter struct {
	db            *pgxpool.Pool
	retriesNumber int
	sleepTimeMs   time.Duration
}

func NewPostgresRetryAdapter(db *pgxpool.Pool, retriesNumber int, sleepTimeMs time.Duration) *PostgresRetryAdapter {
	return &PostgresRetryAdapter{
		db:            db,
		retriesNumber: retriesNumber,
		sleepTimeMs:   sleepTimeMs,
	}
}

func (p *PostgresRetryAdapter) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	var (
		cmdTag pgconn.CommandTag
		err error
	)
	for i := 0; i < p.retriesNumber; i++ {
		cmdTag, err = p.db.Exec(ctx, sql, args...)
		if err == nil {
			return cmdTag, nil
		}
		time.Sleep(p.sleepTimeMs)
	}
	return pgconn.CommandTag{}, err
}

func (p *PostgresRetryAdapter) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	for i := 0; i < p.retriesNumber; i++ {
		row := p.db.QueryRow(ctx, sql, args...)
		if row != nil {
			return row
		}
		time.Sleep(p.sleepTimeMs)
	}
	return nil
}

func (p *PostgresRetryAdapter) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	var (
		rows pgx.Rows
		err error
	)
	for i := 0; i < p.retriesNumber; i++ {
		rows, err := p.db.Query(ctx, sql, args...)
		if err == nil {
			return rows, nil
		}
		if rows != nil {
			rows.Close()
		}
		time.Sleep(p.sleepTimeMs)
	}
	return rows, err
}