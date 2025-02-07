package service

import (
	"cats/internal/domain/entity"
	"cats/internal/domain/types"
	"cats/internal/infrastructure/database/dbgen"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CatsService struct {
	q *dbgen.Queries
}

func NewCatsService(db *pgxpool.Pool) *CatsService {
	return &CatsService{
		q: dbgen.New(db),
	}
}

func (cs *CatsService) List(ctx context.Context) ([]*entity.Cat, error) {
	rows, err := cs.q.GetCats(ctx)
	if err != nil {
		return nil, err
	}
	cats := make([]*entity.Cat, len(rows))
	for i, row := range rows {
		e := new(entity.Cat)
		if err := e.ID.Scan(row.ID.String()); err != nil {
			return nil, err
		}
		e.Breed = row.Breed
		e.Name = row.Name
		e.Salary = int(row.Salary)
		e.YearsOfExperience = int(row.Experience)

		cats[i] = e
	}

	return cats, nil
}

func (cs *CatsService) Get(ctx context.Context, id uuid.UUID) (*entity.Cat, error) {
	var u pgtype.UUID
	if err := u.Scan(id.String()); err != nil {
		return nil, err
	}
	row, err := cs.q.GetCat(ctx, u)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, types.ErrCatNotFound
		}
		return nil, err
	}
	e := new(entity.Cat)
	_ = e.ID.Scan(id.String())
	e.Breed = row.Breed
	e.Name = row.Name
	e.Salary = int(row.Salary)
	e.YearsOfExperience = int(row.Experience)
	return e, nil
}

func (cs *CatsService) Persist(ctx context.Context, cat *entity.Cat) error {
	params := dbgen.CreateCatParams{}
	if err := params.ID.Scan(cat.ID.String()); err != nil {
		return err
	}
	params.Name = cat.Name
	params.Breed = cat.Breed
	params.Experience = int32(cat.YearsOfExperience)
	params.Salary = int32(cat.Salary)

	return cs.q.CreateCat(ctx, params)
}

func (cs *CatsService) Delete(ctx context.Context, id uuid.UUID) error {
	var u pgtype.UUID
	if err := u.Scan(id.String()); err != nil {
		return err
	}
	return cs.q.DeleteCat(ctx, u)
}
