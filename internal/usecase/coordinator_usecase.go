package usecase

import (
	"context"
	"delivery-stream-service/internal/domain"
	"delivery-stream-service/internal/repo"

	"go.uber.org/zap"
)

type CoordinatorUsecase struct {
	coordinatorRepo repo.CoordinatorRepo
}

func NewCoordinatorUsecase(coordinatorRepo repo.CoordinatorRepo) *CoordinatorUsecase {
	return &CoordinatorUsecase{
		coordinatorRepo: coordinatorRepo,
	}
}

func (u *CoordinatorUsecase) Create(ctx context.Context, coordinator *domain.Coordinator, lg *zap.Logger) error {
	lg.Info("Start processing new order", zap.String("order_uid", coordinator.OrderUID))

	err := u.coordinatorRepo.Create(ctx, coordinator, lg)
	if err != nil {
		lg.Warn("Failed processing coordinator", zap.Error(err))
		return err
	}
	lg.Info("Succesful processing new order", zap.String("order_uid", coordinator.OrderUID))
	return nil
}