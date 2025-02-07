package service

import (
	"cats/internal/domain/entity"
	"context"
	"github.com/google/uuid"
)

//go:generate go run github.com/vektra/mockery/v2@v2.52.1 --name=MissionsService
type MissionsService interface {
	List(ctx context.Context) ([]*entity.Mission, error)
	Get(ctx context.Context, id uuid.UUID) (*entity.Mission, error)
	Persist(ctx context.Context, cat *entity.Mission) error
	Delete(ctx context.Context, id uuid.UUID) error
}
