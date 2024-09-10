package postgres

import (
	"context"
	"delivery-stream-service/internal/domain"
	"delivery-stream-service/internal/transaction"

	"go.uber.org/zap"
)

// type PostgresDeliveryRepo struct {
// 	dbAdapter adapter.DBAdapter
// }

// func NewPostgresDeliveryRepo(dbAdapter adapter.DBAdapter) *PostgresDeliveryRepo {
// 	return &PostgresDeliveryRepo{
// 		dbAdapter: dbAdapter,
// 	}
// }

func (p *PostgresDeliveryRepo) Create(ctx context.Context, executor transaction.SQLExecutor, delivery *domain.Delivery, orderUID string, lg *zap.Logger) error {
	lg.Info("Create delivery", zap.String("order_uid", orderUID))

	query := `
        INSERT INTO delivery (order_uid, name, phone, zip, city, address, region, email)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := p.executor.Exec(ctx, query,
		orderUID, delivery.Name, delivery.Phone, delivery.Zip,
		delivery.City, delivery.Address, delivery.Region, delivery.Email)
	if err != nil {
		lg.Warn("Postgres create delivery error", zap.Error(err))
		return err
	}

	return nil
}

// func (p *PostgresDeliveryRepo) GetById(ctx context.Context, orderUID string, lg *zap.Logger) (*domain.Delivery, error) {
// 	lg.Info("Get delivery", zap.String("order_uid", orderUID))

// 	var delivery domain.Delivery
// 	query := `
// 		SELECT name, phone, zip, city, address, region, email
// 		FROM delivery
// 		WHERE order_uid = $1`

// 	rows := p.retryAdapter.QueryRow(ctx, query, orderUID)
// 	defer rows.Close()

// 	err := rows.Scan(
// 		&deelivery.Delivery.Name,
// 		&order.Delivery.Phone,
// 		&order.Delivery.Zip,
// 		&order.Delivery.City,
// 		&order.Delivery.Address,
// 		&order.Delivery.Region,
// 		&order.Delivery.Email,
// 	)
// 	if err != nil {
// 		lg.Warn("Postgres get delivery by id failed", zap.String("order_uid", orderUID))
// 		return &domain.Order{}, err
// 	}

// 	return &order, nil
// }
