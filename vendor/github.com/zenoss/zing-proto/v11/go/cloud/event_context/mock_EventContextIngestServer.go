// Code generated by mockery v1.0.0. DO NOT EDIT.

package event_context

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockEventContextIngestServer is an autogenerated mock type for the EventContextIngestServer type
type MockEventContextIngestServer struct {
	mock.Mock
}

// PutEvent provides a mock function with given fields: _a0, _a1
func (_m *MockEventContextIngestServer) PutEvent(_a0 context.Context, _a1 *PutEventRequest) (*PutEventResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *PutEventResponse
	if rf, ok := ret.Get(0).(func(context.Context, *PutEventRequest) *PutEventResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PutEventResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *PutEventRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutEventBulk provides a mock function with given fields: _a0, _a1
func (_m *MockEventContextIngestServer) PutEventBulk(_a0 context.Context, _a1 *PutEventBulkRequest) (*PutEventBulkResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *PutEventBulkResponse
	if rf, ok := ret.Get(0).(func(context.Context, *PutEventBulkRequest) *PutEventBulkResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PutEventBulkResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *PutEventBulkRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateEvent provides a mock function with given fields: _a0, _a1
func (_m *MockEventContextIngestServer) UpdateEvent(_a0 context.Context, _a1 *UpdateEventRequest) (*UpdateEventResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *UpdateEventResponse
	if rf, ok := ret.Get(0).(func(context.Context, *UpdateEventRequest) *UpdateEventResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UpdateEventResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *UpdateEventRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
