package infrastructure

import (
	"context"
	"delivery-stream-service/infrastructure/db/adapter"
	"delivery-stream-service/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PostgresPaymentRepo struct {
	db *pgxpool.Pool
	retryAdapter adapter.IPostgresRetryAdapter
}

func NewPostgresPaymentRepo(db *pgxpool.Pool, retryAdapter adapter.IPostgresRetryAdapter) *PostgresPaymentRepo {
	return &PostgresPaymentRepo{
		db: db,
		retryAdapter: retryAdapter,
	}
}

func (p *PostgresPaymentRepo)Create(ctx context.Context, payment *domain.Payment, orderUID string, lg *zap.Logger) error {
	lg.Info("Create payment", zap.String("order_uid", orderUID))

	query := `
        INSERT INTO payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := p.retryAdapter.Exec(ctx, query,
		orderUID, payment.Transaction, payment.RequestID, payment.Currency, payment.Provider, payment.Amount,
		payment.PaymentDt, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee)
	if err != nil {
		lg.Warn("Postgres create payment error", zap.Error(err))
		return err
	}

	return nil
}