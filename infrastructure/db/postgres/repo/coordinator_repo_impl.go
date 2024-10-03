package postgres

import (
	"context"
	"delivery-stream-service/infrastructure/db/adapter"
	pgtrans "delivery-stream-service/infrastructure/db/postgres/transaction"
	"delivery-stream-service/internal/domain"
	"delivery-stream-service/internal/repo"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PostgresCoordinatorRepo struct {
	dbPool       *pgxpool.Pool
	deliveryRepo repo.DeliveryRepo
	itemRepo     repo.ItemRepo
	paymentRepo  repo.PaymentRepo
	orderRepo    repo.OrderRepo
}

func NewPostgresCoordinatorRepo(dbPool *pgxpool.Pool, orderRepo repo.OrderRepo, deliveryRepo repo.DeliveryRepo, itemRepo repo.ItemRepo, paymentRepo repo.PaymentRepo) *PostgresCoordinatorRepo {
	return &PostgresCoordinatorRepo{
		dbPool:       dbPool,
		orderRepo:    orderRepo,
		deliveryRepo: deliveryRepo,
		itemRepo:     itemRepo,
		paymentRepo:  paymentRepo,
	}
}

func (p *PostgresCoordinatorRepo) Create(ctx context.Context, coordinator *domain.Coordinator, lg *zap.Logger) error {
	lg.Info("Starting transaction for order", zap.String("order_uid", coordinator.OrderUID))

	tx, err := p.dbPool.Begin(ctx)
	if err != nil {
		lg.Warn("Failed to begin transaction", zap.Error(err))
		return err
	}

	postgresTransaction := pgtrans.NewPostgresTransaction(tx)
	dbAdapter := adapter.NewDBAdapter(postgresTransaction)

	defer func() {
		if err != nil {
			lg.Warn("Transaction rollback due to error", zap.Error(err))
			tx.Rollback(ctx)
		}
	}()

	err = p.orderRepo.Create(ctx, dbAdapter, &coordinator.Order, lg)
	if err != nil {
		lg.Warn("Failed to create order", zap.Error(err))
		return err
	}

	err = p.deliveryRepo.Create(ctx, dbAdapter, &coordinator.Delivery, lg)
	if err != nil {
		lg.Warn("Failed to create delivery", zap.Error(err))
		return err
	}

	err = p.paymentRepo.Create(ctx, dbAdapter, &coordinator.Payment, lg)
	if err != nil {
		lg.Warn("Failed to create payment", zap.Error(err))
		return err
	}

	for _, item := range coordinator.Items {
		err = p.itemRepo.Create(ctx, dbAdapter, &item, lg)
		if err != nil {
			lg.Warn("Error occurred while creating item", zap.Error(err))
			return err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		lg.Warn("Failed to commit transaction", zap.Error(err))
		return err
	}

	lg.Info("Transaction committed successfully for order", zap.String("order_uid", coordinator.OrderUID))

	return nil
}
