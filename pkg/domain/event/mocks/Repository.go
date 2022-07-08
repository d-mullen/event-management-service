// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	event "github.com/zenoss/event-management-service/pkg/domain/event"

	testing "testing"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *Repository) Create(_a0 context.Context, _a1 *event.Event) (*event.Event, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *event.Event
	if rf, ok := ret.Get(0).(func(context.Context, *event.Event) *event.Event); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*event.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *event.Event) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: _a0, _a1
func (_m *Repository) Find(_a0 context.Context, _a1 *event.Query) (*event.Page, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *event.Page
	if rf, ok := ret.Get(0).(func(context.Context, *event.Query) *event.Page); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*event.Page)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *event.Query) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *Repository) Get(_a0 context.Context, _a1 *event.GetRequest) ([]*event.Event, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*event.Event
	if rf, ok := ret.Get(0).(func(context.Context, *event.GetRequest) []*event.Event); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*event.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *event.GetRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *Repository) Update(_a0 context.Context, _a1 *event.Event) (*event.Event, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *event.Event
	if rf, ok := ret.Get(0).(func(context.Context, *event.Event) *event.Event); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*event.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *event.Event) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t testing.TB) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}