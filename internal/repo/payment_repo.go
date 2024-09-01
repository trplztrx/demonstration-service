package repo

import (
	"context"
	"delivery-stream-service/internal/domain"
)

type PaymentRepo interface {
	Create(ctx context.Context, payment *domain.Payment) error
	GetById(ctx context.Context, orderUID string) (*domain.Payment, error)
	Update(ctx context.Context, newPayment *domain.Payment) error
	Delete(ctx context.Context, orderUID string) error
}