package entity

import (
	"cats/internal/domain/types"
	"errors"
	"github.com/google/uuid"
)

type Target struct {
	Name     string `json:"name"`
	Country  string `json:"country"`
	Notes    string `json:"notes"`
	Complete bool   `json:"complete"`
}

func NewTarget(name string, country string, notes string, complete bool) *Target {
	return &Target{
		Name:     name,
		Country:  country,
		Notes:    notes,
		Complete: complete,
	}
}

func (t *Target) UpdateNotes(notes string) error {
	if t.Complete {
		return errors.New("cannot update notes because the target is already complete")
	}
	t.Notes = notes
	return nil
}

func (t *Target) SetComplete() {
	t.Complete = true
}

type Mission struct {
	ID       uuid.UUID          `json:"id"`
	Targets  []*Target          `json:"targets"`
	Index    map[string]*Target `json:"-"`
	Assignee *Cat               `json:"assignee,omitempty"`
	Complete bool               `json:"complete"`
}

func indexTargets(targets []*Target) map[string]*Target {
	index := map[string]*Target{}
	for _, target := range targets {
		index[target.Name] = target
	}
	return index
}

func NewMission(id uuid.UUID, targets []*Target, assignee *Cat, complete bool) *Mission {
	return &Mission{
		ID:       id,
		Index:    indexTargets(targets),
		Targets:  targets,
		Assignee: assignee,
		Complete: complete,
	}
}

func (m *Mission) AddTarget(t *Target) error {
	if m.Complete {
		return errors.New("the mission is already complete")
	}

	if _, ok := m.Index[t.Name]; ok {
		return errors.New("the mission already exists")
	}

	m.Targets = append(m.Targets, t)
	m.Index[t.Name] = t
	return nil
}

func (m *Mission) GetTarget(name string) (*Target, error) {

	if mi, ok := m.Index[name]; ok {
		return mi, nil
	}

	return nil, types.ErrTargetNotFound
}

func (m *Mission) RemoveTargetByName(name string) error {
	if m.Complete {
		return errors.New("the mission is already complete")
	}
	_, ok := m.Index[name]
	if !ok {
		return errors.New("target not found")
	}

	for i, target := range m.Targets {
		if target.Name == name {
			m.Targets = append(m.Targets[:i], m.Targets[i+1:]...)
			delete(m.Index, name)
			break
		}
	}

	return nil
}

func (m *Mission) HasAssignee() bool {
	return m.Assignee != nil
}

func (m *Mission) AddAssignee(a *Cat) error {
	if m.HasAssignee() {
		return errors.New("this mission already have an assignee")
	}
	m.Assignee = a
	return nil
}

func (m *Mission) Finish() {
	m.Complete = true
	for _, t := range m.Targets {
		t.Complete = true
	}
}
