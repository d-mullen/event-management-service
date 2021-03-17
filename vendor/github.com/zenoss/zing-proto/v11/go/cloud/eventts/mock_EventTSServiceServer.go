// Code generated by mockery v1.0.0. DO NOT EDIT.

package eventts

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockEventTSServiceServer is an autogenerated mock type for the EventTSServiceServer type
type MockEventTSServiceServer struct {
	mock.Mock
}

// GetEventCounts provides a mock function with given fields: _a0, _a1
func (_m *MockEventTSServiceServer) GetEventCounts(_a0 context.Context, _a1 *EventTSCountsRequest) (*EventTSCountsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *EventTSCountsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *EventTSCountsRequest) *EventTSCountsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*EventTSCountsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *EventTSCountsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEventCountsStream provides a mock function with given fields: _a0, _a1
func (_m *MockEventTSServiceServer) GetEventCountsStream(_a0 *EventTSCountsRequest, _a1 EventTSService_GetEventCountsStreamServer) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*EventTSCountsRequest, EventTSService_GetEventCountsStreamServer) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetEventFrequency provides a mock function with given fields: _a0, _a1
func (_m *MockEventTSServiceServer) GetEventFrequency(_a0 context.Context, _a1 *EventTSFrequencyRequest) (*EventTSFrequencyResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *EventTSFrequencyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *EventTSFrequencyRequest) *EventTSFrequencyResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*EventTSFrequencyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *EventTSFrequencyRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEvents provides a mock function with given fields: _a0, _a1
func (_m *MockEventTSServiceServer) GetEvents(_a0 context.Context, _a1 *EventTSRequest) (*EventTSResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *EventTSResponse
	if rf, ok := ret.Get(0).(func(context.Context, *EventTSRequest) *EventTSResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*EventTSResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *EventTSRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEventsStream provides a mock function with given fields: _a0, _a1
func (_m *MockEventTSServiceServer) GetEventsStream(_a0 *EventTSRequest, _a1 EventTSService_GetEventsStreamServer) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*EventTSRequest, EventTSService_GetEventsStreamServer) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}