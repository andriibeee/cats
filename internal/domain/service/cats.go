package service

import (
	"cats/internal/domain/entity"
	"context"
	"github.com/google/uuid"
)

//go:generate go run github.com/vektra/mockery/v2@v2.52.1 --name=CatsService
type CatsService interface {
	List(ctx context.Context) ([]*entity.Cat, error)
	Get(ctx context.Context, id uuid.UUID) (*entity.Cat, error)
	Persist(ctx context.Context, cat *entity.Cat) error
	Delete(ctx context.Context, id uuid.UUID) error
}
