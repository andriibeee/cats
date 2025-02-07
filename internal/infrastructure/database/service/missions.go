package service

import (
	"cats/internal/domain/entity"
	"cats/internal/domain/types"
	"cats/internal/infrastructure/database/dbgen"
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MissionsService struct {
	q *dbgen.Queries
}

func NewMissionsService(db *pgxpool.Pool) *MissionsService {
	return &MissionsService{
		q: dbgen.New(db),
	}
}

func (s *MissionsService) List(ctx context.Context) ([]*entity.Mission, error) {
	res, err := s.q.GetMissions(ctx)
	if err != nil {
		return nil, err
	}
	rows := make([]*entity.Mission, len(res))
	for i, row := range res {
		uid, err := uuid.Parse(row.ID.String())
		if err != nil {
			return nil, err
		}

		var t []*entity.Target
		if err := json.Unmarshal(row.Targets, &t); err != nil {
			return nil, err
		}

		var ass *entity.Cat

		if row.AssigneeID.Valid {

			catID, err := uuid.Parse(row.AssigneeID.String())
			if err != nil {
				return nil, err
			}
			ass = entity.NewCat(
				catID,
				row.AssigneeName.String,
				int(row.AssigneeExperience.Int32),
				row.AssigneeBreed.String,
				int(row.AssigneeSalary.Int32),
			)
		}

		rows[i] = entity.NewMission(
			uid,
			t,
			ass,
			row.Complete,
		)
	}

	return rows, nil
}

func (s *MissionsService) Get(ctx context.Context, id uuid.UUID) (*entity.Mission, error) {
	var u pgtype.UUID
	if err := u.Scan(id.String()); err != nil {
		return nil, err
	}

	row, err := s.q.GetMission(ctx, u)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, types.ErrMissionNotFound
		}
		return nil, err
	}

	uid, err := uuid.Parse(row.ID.String())
	if err != nil {
		return nil, err
	}

	var t []*entity.Target
	if err := json.Unmarshal(row.Targets, &t); err != nil {
		return nil, err
	}

	var ass *entity.Cat

	if row.AssigneeID.Valid {

		catID, err := uuid.Parse(row.AssigneeID.String())
		if err != nil {
			return nil, err
		}
		ass = entity.NewCat(
			catID,
			row.AssigneeName.String,
			int(row.AssigneeExperience.Int32),
			row.AssigneeBreed.String,
			int(row.AssigneeSalary.Int32),
		)
	}

	return entity.NewMission(
		uid,
		t,
		ass,
		row.Complete,
	), nil
}

func (s *MissionsService) Persist(ctx context.Context, cat *entity.Mission) error {
	p := dbgen.CreateMissionParams{}

	tar, err := json.Marshal(cat.Targets)
	if err != nil {
		return err
	}

	p.Targets = tar
	if cat.Assignee != nil {
		if err := p.AssigneeID.Scan(cat.Assignee.ID.String()); err != nil {
			return err
		}
	}

	if err := p.ID.Scan(cat.ID.String()); err != nil {
		return err
	}

	p.Complete = cat.Complete

	return s.q.CreateMission(ctx, p)
}

func (s *MissionsService) Delete(ctx context.Context, id uuid.UUID) error {
	var u pgtype.UUID
	if err := u.Scan(id.String()); err != nil {
		return err
	}
	return s.q.DeleteMission(ctx, u)
}
