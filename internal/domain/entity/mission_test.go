package entity_test

import (
	"cats/internal/domain/entity"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTarget(t *testing.T) {
	target := entity.NewTarget("Donald Trump", "USA", "KILL", false)
	assert.Equal(t, "Donald Trump", target.Name)
	assert.Equal(t, "USA", target.Country)
	assert.Equal(t, "KILL", target.Notes)
	assert.False(t, target.Complete)
}

func TestUpdateNotes(t *testing.T) {
	target := entity.NewTarget("Elon Musk", "USA", "The 3rd worst techbro ever", false)
	err := target.UpdateNotes("The worst techbro ever")
	assert.NoError(t, err)
	assert.Equal(t, "The worst techbro ever", target.Notes)

	target.SetComplete()
	err = target.UpdateNotes("ELIMINATED WINNER WINNER CHICKEN DINNER")
	assert.Error(t, err)
	assert.Equal(t, "cannot update notes because the target is already complete", err.Error())
}

func TestSetComplete(t *testing.T) {
	target := entity.NewTarget("Shaman", "RF", "Writes worst music ever", false)
	target.SetComplete()
	assert.True(t, target.Complete)
}

func TestNewMission(t *testing.T) {
	id := uuid.New()
	mission := entity.NewMission(id, nil, nil, false)
	assert.Equal(t, id, mission.ID)
	assert.Nil(t, mission.Targets)
	assert.Nil(t, mission.Assignee)
	assert.False(t, mission.Complete)
}

func TestAddTarget(t *testing.T) {
	id := uuid.New()
	mission := entity.NewMission(id, nil, nil, false)
	target := entity.NewTarget("Operation Monaco", "USA", "top secret", false)

	err := mission.AddTarget(target)
	require.NoError(t, err)
	assert.Contains(t, mission.Targets, target)

	err = mission.AddTarget(target)
	assert.Error(t, err)
	assert.Equal(t, "the mission already exists", err.Error())
}

func TestGetTarget(t *testing.T) {
	id := uuid.New()
	mission := entity.NewMission(id, nil, nil, false)
	target := entity.NewTarget("Donal Duck", "USA", "Not sure", false)
	mission.AddTarget(target)

	foundTarget, err := mission.GetTarget("Donal Duck")
	assert.NoError(t, err)
	assert.Equal(t, target, foundTarget)

	_, err = mission.GetTarget("missingno")
	assert.Error(t, err)
	assert.Equal(t, "target not found", err.Error())
}

func TestRemoveTarget(t *testing.T) {
	id := uuid.New()
	mission := entity.NewMission(id, nil, nil, false)
	target := entity.NewTarget("Shaman", "RF", "writes bad music", false)
	err := mission.AddTarget(target)
	assert.NoError(t, err)

	err = mission.RemoveTargetByName(target.Name)
	assert.NoError(t, err)

	_, err = mission.GetTarget("Shaman")
	assert.Error(t, err)
}

func TestHasAssignee(t *testing.T) {
	id := uuid.New()
	mission := entity.NewMission(id, nil, nil, false)
	assert.False(t, mission.HasAssignee())
	mission.Assignee = &entity.Cat{Name: "Keks"}
	assert.True(t, mission.HasAssignee())
}

func TestAddAssignee(t *testing.T) {
	id := uuid.New()
	mission := entity.NewMission(id, nil, nil, false)
	assignee := &entity.Cat{Name: "Keks"}

	err := mission.AddAssignee(assignee)
	assert.NoError(t, err)
	assert.Equal(t, assignee, mission.Assignee)

	err = mission.AddAssignee(&entity.Cat{Name: "Keks"})
	assert.Error(t, err)
	assert.Equal(t, "this mission already have an assignee", err.Error())
}

func TestFinish(t *testing.T) {
	id := uuid.New()
	target1 := entity.NewTarget("Shaman", "RF", "writes bad music", false)
	target2 := entity.NewTarget("Clonnex", "PL", "shaman's stunt double", false)
	mission := entity.NewMission(id, []*entity.Target{target1, target2}, nil, false)

	mission.Finish()
	assert.True(t, mission.Complete)
	assert.True(t, target1.Complete)
	assert.True(t, target2.Complete)
}
