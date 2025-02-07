package entity_test

import (
	"cats/internal/domain/entity"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewCat(t *testing.T) {
	id := uuid.New()
	cat := entity.NewCat(id, "Keks", 5, "Siamese", 50000)
	assert.Equal(t, id, cat.ID)
	assert.Equal(t, "Keks", cat.Name)
	assert.Equal(t, 5, cat.YearsOfExperience)
	assert.Equal(t, "Siamese", cat.Breed)
	assert.Equal(t, 50000, cat.Salary)
}

func TestUpdateSalary(t *testing.T) {
	id := uuid.New()
	cat := entity.NewCat(id, "Keks", 5, "Siamese", 50000)

	cat.UpdateSalary(60000)
	assert.Equal(t, 60000, cat.Salary)
}
