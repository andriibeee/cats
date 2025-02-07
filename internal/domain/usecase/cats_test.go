package usecase_test

import (
	"cats/internal/domain/dto"
	"cats/internal/domain/entity"
	"cats/internal/domain/service/mocks"
	"cats/internal/domain/usecase"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateCat(t *testing.T) {
	ctx := context.TODO()
	mockBS := new(mocks.BreedService)
	mockCS := new(mocks.CatsService)
	useCase := usecase.NewCatsUseCase(mockBS, mockCS)

	dto := dto.CreateCatDTO{Name: "Keks", YearsOfExperience: 3, Breed: "Siamese", Salary: 50000}
	mockBS.On("CheckBreed", ctx, dto.Breed).Return(true, nil)
	mockCS.On("Persist", ctx, mock.Anything).Return(nil)

	cat, err := useCase.Create(ctx, dto)
	require.NoError(t, err)
	assert.Equal(t, dto.Name, cat.Name)
}

func TestFindByID(t *testing.T) {
	ctx := context.TODO()
	mockCS := new(mocks.CatsService)
	useCase := usecase.NewCatsUseCase(nil, mockCS)
	id := uuid.New()
	cat := &entity.Cat{ID: id, Name: "Keks"}
	mockCS.On("Get", ctx, id).Return(cat, nil)

	result, err := useCase.FindByID(ctx, id)
	require.NoError(t, err)
	assert.Equal(t, cat, result)
}

func TestUpdateCat(t *testing.T) {
	ctx := context.TODO()
	mockCS := new(mocks.CatsService)
	useCase := usecase.NewCatsUseCase(nil, mockCS)
	id := uuid.New()
	cat := &entity.Cat{ID: id, Name: "Keks", Salary: 50000}
	dto := dto.UpdateCatDTO{Salary: 60000}
	mockCS.On("Get", ctx, id).Return(cat, nil)
	mockCS.On("Persist", ctx, cat).Return(nil)

	updatedCat, err := useCase.Update(ctx, id, dto)
	require.NoError(t, err)
	assert.Equal(t, dto.Salary, updatedCat.Salary)
}

func TestDeleteCat(t *testing.T) {
	ctx := context.TODO()
	mockCS := new(mocks.CatsService)
	useCase := usecase.NewCatsUseCase(nil, mockCS)
	id := uuid.New()
	mockCS.On("Delete", ctx, id).Return(nil)

	err := useCase.Delete(ctx, id)
	require.NoError(t, err)
}
