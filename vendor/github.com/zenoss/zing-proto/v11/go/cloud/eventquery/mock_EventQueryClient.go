// Code generated by mockery v1.0.0. DO NOT EDIT.

package eventquery

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockEventQueryClient is an autogenerated mock type for the EventQueryClient type
type MockEventQueryClient struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, in, opts
func (_m *MockEventQueryClient) Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *CountResponse
	if rf, ok := ret.Get(0).(func(context.Context, *CountRequest, ...grpc.CallOption) *CountResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*CountResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *CountRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EventsWithCountsStream provides a mock function with given fields: ctx, in, opts
func (_m *MockEventQueryClient) EventsWithCountsStream(ctx context.Context, in *EventsWithCountsRequest, opts ...grpc.CallOption) (EventQuery_EventsWithCountsStreamClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 EventQuery_EventsWithCountsStreamClient
	if rf, ok := ret.Get(0).(func(context.Context, *EventsWithCountsRequest, ...grpc.CallOption) EventQuery_EventsWithCountsStreamClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EventQuery_EventsWithCountsStreamClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *EventsWithCountsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Frequency provides a mock function with given fields: ctx, in, opts
func (_m *MockEventQueryClient) Frequency(ctx context.Context, in *FrequencyRequest, opts ...grpc.CallOption) (*FrequencyResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *FrequencyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *FrequencyRequest, ...grpc.CallOption) *FrequencyResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*FrequencyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *FrequencyRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEvent provides a mock function with given fields: ctx, in, opts
func (_m *MockEventQueryClient) GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*GetEventResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *GetEventResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetEventRequest, ...grpc.CallOption) *GetEventResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetEventResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetEventRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEvents provides a mock function with given fields: ctx, in, opts
func (_m *MockEventQueryClient) GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *GetEventsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetEventsRequest, ...grpc.CallOption) *GetEventsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetEventsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetEventsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: ctx, in, opts
func (_m *MockEventQueryClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *SearchResponse
	if rf, ok := ret.Get(0).(func(context.Context, *SearchRequest, ...grpc.CallOption) *SearchResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*SearchResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *SearchRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchStream provides a mock function with given fields: ctx, in, opts
func (_m *MockEventQueryClient) SearchStream(ctx context.Context, in *SearchStreamRequest, opts ...grpc.CallOption) (EventQuery_SearchStreamClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 EventQuery_SearchStreamClient
	if rf, ok := ret.Get(0).(func(context.Context, *SearchStreamRequest, ...grpc.CallOption) EventQuery_SearchStreamClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EventQuery_SearchStreamClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *SearchStreamRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
