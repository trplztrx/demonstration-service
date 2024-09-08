package infrastructure

import (
	"context"
	"delivery-stream-service/infrastructure/db/adapter"
	"delivery-stream-service/internal/domain"
	"delivery-stream-service/internal/repo"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PostgresOrderRepo struct {
	db           *pgxpool.Pool
	retryAdapter adapter.IPostgresRetryAdapter
	deliveryRepo repo.DeliveryRepo
	itemRepo     repo.ItemRepo
	paymentRepo  repo.PaymentRepo
}

func NewPostgresOrderRepo(db *pgxpool.Pool, retryAdapter adapter.IPostgresRetryAdapter, deliveryRepo repo.DeliveryRepo, itemRepo repo.ItemRepo, paymentRepo repo.PaymentRepo) *PostgresOrderRepo {
	return &PostgresOrderRepo{
		db:           db,
		retryAdapter: retryAdapter,
		deliveryRepo: deliveryRepo,
		itemRepo:     itemRepo,
		paymentRepo:  paymentRepo,
	}
}

func (p *PostgresOrderRepo) Create(ctx context.Context, order *domain.Order, lg *zap.Logger) error {
	lg.Info("Create order", zap.String("order_uid", order.OrderUID))

	tx, err := p.db.Begin(ctx)
	if err != nil {
		lg.Warn("Begin transaction failed", zap.Error(err))
		return err
	}
	defer tx.Rollback(ctx)

	query := `
	INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err = p.retryAdapter.Exec(ctx, query,
		order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature,
		order.CustomerID, order.DeliveryService, order.ShardKey, order.SmID, order.CreatedAt, order.OofShard)
	if err != nil {
		lg.Warn("Postgres create order error", zap.Error(err))
		return err
	}
	
	err = p.deliveryRepo.Create(ctx, &order.Delivery, order.OrderUID, lg)
	if err != nil {
		// lg.Warn("Postgres create order error", zap.Error(err))
		return err
	}

	err = p.paymentRepo.Create(ctx, &order.Payment, order.OrderUID, lg)
	if err != nil {
		// lg.Warn("Postgres create order error", zap.Error(err))
		return err
	}

	for _, item := range order.Items {
		err = p.itemRepo.Create(ctx, &item, order.OrderUID, lg)
		if err != nil {
			lg.Warn("Error occurred while creating item", zap.Error(err))
			return err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		lg.Warn("Commit transaction failed", zap.Error(err))
		return err
	}

	return nil
}

func (p *PostgresOrderRepo) GetById(ctx context.Context, orderUID string, lg *zap.Logger) (*domain.Order, error) {
	lg.Info("Get order", zap.String("order_uid", orderUID))
	
	var order domain.Order
	query := ` 
		SELECT order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard 
		FROM orders 
		WHERE order_uid = $1`

	rows := p.retryAdapter.QueryRow(ctx, query, orderUID)
	defer rows.Close()

	err := rows.Scan(
		&order.OrderUID,
		&order.TrackNumber,
		&order.Entry,
		&order.Locale,
		&order.InternalSignature,
		&order.CustomerID,
		&order.DeliveryService,
		&order.ShardKey,
		&order.SmID,
		&order.CreatedAt,
		&order.OofShard,
	)
	if err != nil {
		lg.Warn("Postgres get by id failed", zap.String("order_uid", orderUID))
		return &domain.Order{}, err
	}

	
	return &order, nil
}
