package usecase

import (
	"cats/internal/domain/dto"
	"cats/internal/domain/entity"
	"cats/internal/domain/service"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type MissionsUseCase struct {
	ms service.MissionsService
	cs service.CatsService
}

func NewMissionsUseCase(
	ms service.MissionsService,
	cs service.CatsService,
) *MissionsUseCase {
	return &MissionsUseCase{
		ms: ms,
		cs: cs,
	}
}

func (uc *MissionsUseCase) List(ctx context.Context) ([]*entity.Mission, error) {
	return uc.ms.List(ctx)
}

func (uc *MissionsUseCase) Get(ctx context.Context, id uuid.UUID) (*entity.Mission, error) {
	return uc.ms.Get(ctx, id)
}

func (uc *MissionsUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	miss, err := uc.ms.Get(ctx, id)
	if err != nil {
		return err
	}

	if miss.HasAssignee() {
		return errors.New("cannot delete mission because it has assignee")
	}

	return uc.ms.Delete(ctx, id)
}

func (uc *MissionsUseCase) Update(ctx context.Context, id uuid.UUID, dto *dto.UpdateMissionDTO) (*entity.Mission, error) {
	miss, err := uc.ms.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get: %w", err)
	}

	if dto.Complete {
		miss.Finish()
	}

	if dto.Assignee != uuid.Nil {
		ass, err := uc.cs.Get(ctx, dto.Assignee)
		if err != nil {
			return nil, fmt.Errorf("get assignee: %w", err)
		}

		if err := miss.AddAssignee(ass); err != nil {
			return nil, fmt.Errorf("add assignee: %w", err)
		}
	}

	if err := uc.ms.Persist(ctx, miss); err != nil {
		return nil, fmt.Errorf("persist: %w", err)
	}

	return miss, nil
}

func (uc *MissionsUseCase) Create(ctx context.Context, dto *dto.CreateMissionDTO) (*entity.Mission, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	t := make([]*entity.Target, len(dto.Targets))
	for i, v := range dto.Targets {
		t[i] = entity.NewTarget(
			v.Name,
			v.Country,
			v.Notes,
			false,
		)
	}

	var as *entity.Cat
	if dto.AssigneeID != uuid.Nil {
		c, err := uc.cs.Get(ctx, dto.AssigneeID)
		if err != nil {
			return nil, err
		}
		as = c
	}
	miss := entity.NewMission(
		id,
		t,
		as,
		dto.Complete,
	)
	if err := uc.ms.Persist(ctx, miss); err != nil {
		return nil, err
	}
	return miss, nil
}

func (uc *MissionsUseCase) AddTarget(ctx context.Context, id uuid.UUID, p *dto.CreateTargetDTO) (*entity.Mission, error) {
	miss, err := uc.ms.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := miss.AddTarget(entity.NewTarget(
		p.Name,
		p.Country,
		p.Notes,
		false,
	)); err != nil {
		return nil, err
	}

	if err := uc.ms.Persist(ctx, miss); err != nil {
		return nil, err
	}

	return miss, nil
}

func (uc *MissionsUseCase) UpdateTarget(ctx context.Context, id uuid.UUID, targetName string, p *dto.UpdateTargetDTO) (*entity.Mission, error) {
	miss, err := uc.ms.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	target, err := miss.GetTarget(targetName)
	if err != nil {

		return nil, err
	}

	if p.Name != target.Name {
		target.Name = p.Name
	}

	if p.Complete {
		target.SetComplete()
	}

	if p.Notes != target.Notes {
		if err := target.UpdateNotes(p.Notes); err != nil {
			return nil, err
		}
	}

	if err := uc.ms.Persist(ctx, miss); err != nil {
		return nil, err
	}

	return miss, nil
}

func (uc *MissionsUseCase) DeleteTarget(ctx context.Context, id uuid.UUID, target string) (*entity.Mission, error) {
	miss, err := uc.ms.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := miss.RemoveTargetByName(target); err != nil {
		return nil, err
	}

	if err := uc.ms.Persist(ctx, miss); err != nil {
		return nil, err
	}

	return miss, nil
}
