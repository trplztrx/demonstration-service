package repo

import (
	"context"
	"delivery-stream-service/internal/domain"
)

type ItemRepo interface {
	Create(ctx context.Context, item *domain.Item) error
	GetById(ctx context.Context, orderUID string) (*domain.Item, error)
	Update(ctx context.Context, newItem *domain.Item) error
	Delete(ctx context.Context, orderUID string) error
}