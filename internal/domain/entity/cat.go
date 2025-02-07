package entity

import "github.com/google/uuid"

type Cat struct {
	ID                uuid.UUID `json:"id"`
	Name              string    `json:"name"`
	YearsOfExperience int       `json:"experience"`
	Breed             string    `json:"breed"`
	Salary            int       `json:"salary"`
}

func NewCat(id uuid.UUID, name string, yoe int, breed string, salary int) *Cat {
	return &Cat{
		ID:                id,
		Name:              name,
		YearsOfExperience: yoe,
		Breed:             breed,
		Salary:            salary,
	}
}

func (cat *Cat) UpdateSalary(s int) {
	cat.Salary = s
}
