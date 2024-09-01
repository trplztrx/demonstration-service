package repo

import (
	"context"
	"delivery-stream-service/internal/domain"
)
type OrderRepo interface {
	Create(ctx context.Context, order *domain.Order) error
	GetById(ctx context.Context, orderUID string) (*domain.Order, error)
	Update(ctx context.Context, newOrder *domain.Order) error
	Delete(ctx context.Context, orderUID string) error
}