package usecase

import (
	"cats/internal/domain/dto"
	"cats/internal/domain/entity"
	"cats/internal/domain/service"
	"cats/internal/domain/types"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type CatsUseCase struct {
	bs service.BreedService
	cs service.CatsService
}

func NewCatsUseCase(
	bs service.BreedService,
	cs service.CatsService,
) *CatsUseCase {
	return &CatsUseCase{
		bs: bs,
		cs: cs,
	}
}

func (uc *CatsUseCase) Create(ctx context.Context, dto *dto.CreateCatDTO) (*entity.Cat, error) {
	validity, err := uc.bs.CheckBreed(ctx, dto.Breed)
	if err != nil {
		return nil, fmt.Errorf("failed to check breed: %w", err)
	}
	if !validity {
		return nil, types.ErrBreedNotValid
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	cat := entity.NewCat(
		id,
		dto.Name,
		dto.YearsOfExperience,
		dto.Breed,
		dto.Salary,
	)

	if err := uc.cs.Persist(ctx, cat); err != nil {
		return nil, err
	}

	return cat, nil
}

func (uc *CatsUseCase) FindByID(ctx context.Context, id uuid.UUID) (*entity.Cat, error) {
	return uc.cs.Get(ctx, id)
}

func (uc *CatsUseCase) List(ctx context.Context) ([]*entity.Cat, error) {
	return uc.cs.List(ctx)
}

func (uc *CatsUseCase) Update(ctx context.Context, id uuid.UUID, dto *dto.UpdateCatDTO) (*entity.Cat, error) {
	ent, err := uc.cs.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	ent.UpdateSalary(dto.Salary)
	if err := uc.cs.Persist(ctx, ent); err != nil {
		return nil, err
	}
	return ent, nil
}

func (uc *CatsUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.cs.Delete(ctx, id)
}
