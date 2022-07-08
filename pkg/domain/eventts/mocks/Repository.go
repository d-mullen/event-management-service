// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	eventts "github.com/zenoss/event-management-service/pkg/domain/eventts"

	testing "testing"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Frequency provides a mock function with given fields: _a0, _a1
func (_m *Repository) Frequency(_a0 context.Context, _a1 *eventts.FrequencyRequest) (*eventts.FrequencyResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *eventts.FrequencyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *eventts.FrequencyRequest) *eventts.FrequencyResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*eventts.FrequencyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *eventts.FrequencyRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FrequencyStream provides a mock function with given fields: _a0, _a1
func (_m *Repository) FrequencyStream(_a0 context.Context, _a1 *eventts.FrequencyRequest) chan *eventts.FrequencyOptional {
	ret := _m.Called(_a0, _a1)

	var r0 chan *eventts.FrequencyOptional
	if rf, ok := ret.Get(0).(func(context.Context, *eventts.FrequencyRequest) chan *eventts.FrequencyOptional); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *eventts.FrequencyOptional)
		}
	}

	return r0
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *Repository) Get(_a0 context.Context, _a1 *eventts.GetRequest) ([]*eventts.Occurrence, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*eventts.Occurrence
	if rf, ok := ret.Get(0).(func(context.Context, *eventts.GetRequest) []*eventts.Occurrence); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*eventts.Occurrence)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *eventts.GetRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStream provides a mock function with given fields: _a0, _a1
func (_m *Repository) GetStream(_a0 context.Context, _a1 *eventts.GetRequest) <-chan *eventts.OccurrenceOptional {
	ret := _m.Called(_a0, _a1)

	var r0 <-chan *eventts.OccurrenceOptional
	if rf, ok := ret.Get(0).(func(context.Context, *eventts.GetRequest) <-chan *eventts.OccurrenceOptional); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *eventts.OccurrenceOptional)
		}
	}

	return r0
}

// NewRepository creates a new instance of Repository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t testing.TB) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}