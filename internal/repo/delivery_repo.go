package repo

import (
	"context"
	"delivery-stream-service/internal/domain"
)

type DeliveryRepo interface {
	Create(ctx context.Context, delivery *domain.Delivery) error
	GetById(ctx context.Context, orderUID string) (*domain.Delivery, error)
	Update(ctx context.Context, newDelivery *domain.Delivery) error
	Delete(ctx context.Context, orderUID string) error
}