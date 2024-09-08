package repo

import (
	"context"
	"delivery-stream-service/internal/domain"

	"go.uber.org/zap"
)
type OrderRepo interface {
	Create(ctx context.Context, order *domain.Order, lg *zap.Logger) error
	GetById(ctx context.Context, orderUID string, lg *zap.Logger) (*domain.Order, error)
	// Update(ctx context.Context, newOrder *domain.Order) error
	// Delete(ctx context.Context, orderUID string, lg *zap.Logger) error
}