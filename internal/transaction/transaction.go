// уровень repo
package transaction

import "context"

type Transaction interface {
    SQLExecutor
    Commit(ctx context.Context) error
    Rollback(ctx context.Context) error
}

// абстракция для выполнения SQL запросов
type SQLExecutor interface {
    Exec(ctx context.Context, query string, args ...any) (Result, error)
    QueryRow(ctx context.Context, query string, args ...any) Row
    Query(ctx context.Context, query string, args ...any) (Rows, error)
}

// абстракция результата выполнения SQL запроса
type Result interface {
    RowsAffected() (int64, error)
}

// абстракция строки результата запроса
type Row interface {
    Scan(dest ...any) error
}

// абстракция для множества строк результата запроса
type Rows interface {
    Next() bool
    Scan(dest ...any) error
    Close() error
    Err() error
}

