
package infrastructure

import (
	"context"
	"delivery-stream-service/infrastructure/db/adapter"
	"delivery-stream-service/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PostgresItemRepo struct {
	db *pgxpool.Pool
	retryAdapter adapter.IPostgresRetryAdapter
}

func NewPostgresItemRepo(db *pgxpool.Pool, retryAdapter adapter.IPostgresRetryAdapter) *PostgresItemRepo {
	return &PostgresItemRepo{
		db: db,
		retryAdapter: retryAdapter,
	}
}

func (p *PostgresItemRepo) Create(ctx context.Context, item *domain.Item, orderUID string, lg *zap.Logger) error {
	lg.Info("Create item", zap.String("order_uid", orderUID))

	query := `
        INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err := p.retryAdapter.Exec(ctx, query,
		orderUID, item.ChrtID, item.TrackNumber, item.Price, item.RID, item.Name,
		item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status)
	if err != nil {
		lg.Warn("Postgres create item error", zap.Error(err))
		return err
	}

	return nil
}

