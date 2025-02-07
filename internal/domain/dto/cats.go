package dto

type CreateCatDTO struct {
	Name              string `json:"name"`
	YearsOfExperience int    `json:"experience"`
	Breed             string `json:"breed"`
	Salary            int    `json:"salary"`
}

type UpdateCatDTO struct {
	Salary int `json:"salary"`
}
