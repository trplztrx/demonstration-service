package repo

import (
	"context"
	"delivery-stream-service/internal/domain"

	"go.uber.org/zap"
)

type PaymentRepo interface {
	Create(ctx context.Context, payment *domain.Payment, orderUID string, lg *zap.Logger) error
	GetById(ctx context.Context, orderUID string, lg *zap.Logger) (*domain.Payment, error)
	// Update(ctx context.Context, newPayment *domain.Payment) error
	// Delete(ctx context.Context, orderUID string) error
}