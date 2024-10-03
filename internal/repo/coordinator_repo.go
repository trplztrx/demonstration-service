package repo

import (
	"context"
	"delivery-stream-service/internal/domain"

	"go.uber.org/zap"
)

type CoordinatorRepo interface {
	Create(ctx context.Context, Coordinator *domain.Coordinator, lg *zap.Logger) error
}