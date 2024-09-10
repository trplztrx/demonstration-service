package repo

import (
	"context"
	"delivery-stream-service/internal/domain"
	"delivery-stream-service/internal/transaction"

	"go.uber.org/zap"
)

type ItemRepo interface {
	Create(ctx context.Context, executor transaction.SQLExecutor, item *domain.Item, orderUID string, lg *zap.Logger) error
	// GetById(ctx context.Context, orderUID string, lg *zap.Logger) (*domain.Item, error)
	// Update(ctx context.Context, newItem *domain.Item) error
	// Delete(ctx context.Context, orderUID string) error
}