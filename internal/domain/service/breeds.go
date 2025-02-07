package service

import "context"

//go:generate go run github.com/vektra/mockery/v2@v2.52.1 --name=BreedService
type BreedService interface {
	CheckBreed(ctx context.Context, name string) (bool, error)
}
