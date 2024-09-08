package repo

import (
	"context"
	"delivery-stream-service/internal/domain"

	"go.uber.org/zap"
)

type DeliveryRepo interface {
	Create(ctx context.Context, delivery *domain.Delivery, orderUID string, lg *zap.Logger) error
	GetById(ctx context.Context, orderUID string, lg *zap.Logger) (*domain.Delivery, error)
	// Update(ctx context.Context, newDelivery *domain.Delivery) error
	// Delete(ctx context.Context, orderUID string) error
}