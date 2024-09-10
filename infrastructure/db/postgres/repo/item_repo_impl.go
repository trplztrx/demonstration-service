package postgres

import (
	"context"
	"delivery-stream-service/infrastructure/db/adapter"
	"delivery-stream-service/internal/domain"
	"delivery-stream-service/internal/transaction"

	"go.uber.org/zap"
)

// type PostgresItemRepo struct {
// 	dbAdapter adapter.DBAdapter
// }

// func NewPostgresItemRepo(dbAdapter adapter.DBAdapter) *PostgresItemRepo {
// 	return &PostgresItemRepo{
// 		dbAdapter: dbAdapter,
// 	}
// }

func (p *PostgresItemRepo) Create(ctx context.Context, executor transaction.SQLExecutor, item *domain.Item, orderUID string, lg *zap.Logger) error {
	lg.Info("Create item", zap.String("order_uid", orderUID))

	query := `
        INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err := p.executor.Exec(ctx, query,
		orderUID, item.ChrtID, item.TrackNumber, item.Price, item.RID, item.Name,
		item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status)
	if err != nil {
		lg.Warn("Postgres create item error", zap.Error(err))
		return err
	}

	return nil
}
