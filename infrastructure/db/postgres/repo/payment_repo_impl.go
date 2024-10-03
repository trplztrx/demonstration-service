package postgres

import (
	"context"
	"delivery-stream-service/internal/domain"
	"delivery-stream-service/internal/transaction"

	"go.uber.org/zap"
)

type PostgresPaymentRepo struct {
	executor transaction.SQLExecutor
}

func NewPostgresPaymentRepo(executor transaction.SQLExecutor) *PostgresPaymentRepo {
	return &PostgresPaymentRepo{
		executor: executor,
	}
}

func (p *PostgresPaymentRepo) Create(ctx context.Context, executor transaction.SQLExecutor, payment *domain.Payment, orderUID string, lg *zap.Logger) error {
	lg.Info("Create payment", zap.String("order_uid", orderUID))

	query := `
		INSERT INTO payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := p.executor.Exec(ctx, query,
		orderUID, payment.Transaction, payment.RequestID, payment.Currency, payment.Provider, payment.Amount,
		payment.PaymentDt, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee)

	if err != nil {
		lg.Warn("Postgres create payment error", zap.Error(err))
		return err
	}

	return nil
}
