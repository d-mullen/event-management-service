// Code generated by mockery v1.0.0. DO NOT EDIT.

package eventts

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockEventTSServiceClient is an autogenerated mock type for the EventTSServiceClient type
type MockEventTSServiceClient struct {
	mock.Mock
}

// GetEventCounts provides a mock function with given fields: ctx, in, opts
func (_m *MockEventTSServiceClient) GetEventCounts(ctx context.Context, in *EventTSCountsRequest, opts ...grpc.CallOption) (*EventTSCountsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *EventTSCountsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *EventTSCountsRequest, ...grpc.CallOption) *EventTSCountsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*EventTSCountsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *EventTSCountsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEventCountsStream provides a mock function with given fields: ctx, in, opts
func (_m *MockEventTSServiceClient) GetEventCountsStream(ctx context.Context, in *EventTSCountsRequest, opts ...grpc.CallOption) (EventTSService_GetEventCountsStreamClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 EventTSService_GetEventCountsStreamClient
	if rf, ok := ret.Get(0).(func(context.Context, *EventTSCountsRequest, ...grpc.CallOption) EventTSService_GetEventCountsStreamClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EventTSService_GetEventCountsStreamClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *EventTSCountsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEventFrequency provides a mock function with given fields: ctx, in, opts
func (_m *MockEventTSServiceClient) GetEventFrequency(ctx context.Context, in *EventTSFrequencyRequest, opts ...grpc.CallOption) (*EventTSFrequencyResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *EventTSFrequencyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *EventTSFrequencyRequest, ...grpc.CallOption) *EventTSFrequencyResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*EventTSFrequencyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *EventTSFrequencyRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEvents provides a mock function with given fields: ctx, in, opts
func (_m *MockEventTSServiceClient) GetEvents(ctx context.Context, in *EventTSRequest, opts ...grpc.CallOption) (*EventTSResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *EventTSResponse
	if rf, ok := ret.Get(0).(func(context.Context, *EventTSRequest, ...grpc.CallOption) *EventTSResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*EventTSResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *EventTSRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEventsStream provides a mock function with given fields: ctx, in, opts
func (_m *MockEventTSServiceClient) GetEventsStream(ctx context.Context, in *EventTSRequest, opts ...grpc.CallOption) (EventTSService_GetEventsStreamClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 EventTSService_GetEventsStreamClient
	if rf, ok := ret.Get(0).(func(context.Context, *EventTSRequest, ...grpc.CallOption) EventTSService_GetEventsStreamClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EventTSService_GetEventsStreamClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *EventTSRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
