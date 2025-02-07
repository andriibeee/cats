package dto

import (
	"github.com/google/uuid"
)

type CreateTargetDTO struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Notes   string `json:"notes"`
}

type CreateMissionDTO struct {
	AssigneeID uuid.UUID         `json:"assignee"`
	Targets    []CreateTargetDTO `json:"targets"`
	Complete   bool              `json:"complete"`
}

type UpdateMissionDTO struct {
	Assignee uuid.UUID `json:"assignee,omitempty"`
	Complete bool      `json:"complete"`
}

type UpdateTargetDTO struct {
	Name     string `json:"name"`
	Notes    string `json:"notes"`
	Complete bool   `json:"complete"`
}
