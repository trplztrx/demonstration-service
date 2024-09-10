package repo

import (
	"context"
	"delivery-stream-service/internal/domain"

	"go.uber.org/zap"
)

type OrderCoordinatorRepo interface {
	Create(ctx context.Context, orderCoordinator *domain.OrderCoordinator, lg *zap.Logger) error
}