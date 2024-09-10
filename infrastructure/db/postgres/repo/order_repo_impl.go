package postgres

import (
	"context"
	"delivery-stream-service/infrastructure/db/adapter"
	"delivery-stream-service/internal/domain"

	"go.uber.org/zap"
)

// type PostgresOrderRepo struct {
// 	dbAdapter adapter.DBAdapter
// }

// func NewPostgresOrderRepo(dbAdapter adapter.DBAdapter) *PostgresOrderRepo {
// 	return &PostgresOrderRepo{
// 		dbAdapter: dbAdapter,
// 	}
// }

func (p *PostgresOrderRepo) Create(ctx context.Context, dbAdapter transaction.SQLExecutor, order *domain.Order, lg *zap.Logger) error {
	lg.Info("Create order", zap.String("order_uid", order.OrderUID))

	query := `
	INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := p.dbAdapter.Exec(ctx, query,
		order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature,
		order.CustomerID, order.DeliveryService, order.ShardKey, order.SmID, order.CreatedAt, order.OofShard)
	if err != nil {
		lg.Warn("Postgres create order error", zap.Error(err))
		return err
	}

	return nil
}

// func (p *PostgresOrderRepo) GetById(ctx context.Context, orderUID string, lg *zap.Logger) (*domain.Order, error) {
// 	lg.Info("Get order", zap.String("order_uid", orderUID))

// 	var order domain.Order
// 	query := `
// 		SELECT order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard
// 		FROM orders
// 		WHERE order_uid = $1`

// 	rows := p.retryAdapter.QueryRow(ctx, query, orderUID)
// 	defer rows.Close()

// 	err := rows.Scan(
// 		&order.OrderUID,
// 		&order.TrackNumber,
// 		&order.Entry,
// 		&order.Locale,
// 		&order.InternalSignature,
// 		&order.CustomerID,
// 		&order.DeliveryService,
// 		&order.ShardKey,
// 		&order.SmID,
// 		&order.CreatedAt,
// 		&order.OofShard,
// 	)
// 	if err != nil {
// 		lg.Warn("Postgres get by id failed", zap.String("order_uid", orderUID))
// 		return &domain.Order{}, err
// 	}

// 	return &order, nil
// }
