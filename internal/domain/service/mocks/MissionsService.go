// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	entity "cats/internal/domain/entity"
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MissionsService is an autogenerated mock type for the MissionsService type
type MissionsService struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MissionsService) Delete(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *MissionsService) Get(ctx context.Context, id uuid.UUID) (*entity.Mission, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *entity.Mission
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*entity.Mission, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *entity.Mission); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Mission)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx
func (_m *MissionsService) List(ctx context.Context) ([]*entity.Mission, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*entity.Mission
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*entity.Mission, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*entity.Mission); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Mission)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Persist provides a mock function with given fields: ctx, cat
func (_m *MissionsService) Persist(ctx context.Context, cat *entity.Mission) error {
	ret := _m.Called(ctx, cat)

	if len(ret) == 0 {
		panic("no return value specified for Persist")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Mission) error); ok {
		r0 = rf(ctx, cat)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMissionsService creates a new instance of MissionsService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMissionsService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MissionsService {
	mock := &MissionsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
