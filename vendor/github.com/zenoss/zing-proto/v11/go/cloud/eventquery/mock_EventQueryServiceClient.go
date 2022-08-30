// Code generated by mockery v2.13.1. DO NOT EDIT.

package eventquery

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockEventQueryServiceClient is an autogenerated mock type for the EventQueryServiceClient type
type MockEventQueryServiceClient struct {
	mock.Mock
}

type MockEventQueryServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEventQueryServiceClient) EXPECT() *MockEventQueryServiceClient_Expecter {
	return &MockEventQueryServiceClient_Expecter{mock: &_m.Mock}
}

// Count provides a mock function with given fields: ctx, in, opts
func (_m *MockEventQueryServiceClient) Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error) {
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

// MockEventQueryServiceClient_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type MockEventQueryServiceClient_Count_Call struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
//  - ctx context.Context
//  - in *CountRequest
//  - opts ...grpc.CallOption
func (_e *MockEventQueryServiceClient_Expecter) Count(ctx interface{}, in interface{}, opts ...interface{}) *MockEventQueryServiceClient_Count_Call {
	return &MockEventQueryServiceClient_Count_Call{Call: _e.mock.On("Count",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockEventQueryServiceClient_Count_Call) Run(run func(ctx context.Context, in *CountRequest, opts ...grpc.CallOption)) *MockEventQueryServiceClient_Count_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*CountRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockEventQueryServiceClient_Count_Call) Return(_a0 *CountResponse, _a1 error) *MockEventQueryServiceClient_Count_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// EventsWithCountsStream provides a mock function with given fields: ctx, in, opts
func (_m *MockEventQueryServiceClient) EventsWithCountsStream(ctx context.Context, in *EventsWithCountsStreamRequest, opts ...grpc.CallOption) (EventQueryService_EventsWithCountsStreamClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 EventQueryService_EventsWithCountsStreamClient
	if rf, ok := ret.Get(0).(func(context.Context, *EventsWithCountsStreamRequest, ...grpc.CallOption) EventQueryService_EventsWithCountsStreamClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EventQueryService_EventsWithCountsStreamClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *EventsWithCountsStreamRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEventQueryServiceClient_EventsWithCountsStream_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EventsWithCountsStream'
type MockEventQueryServiceClient_EventsWithCountsStream_Call struct {
	*mock.Call
}

// EventsWithCountsStream is a helper method to define mock.On call
//  - ctx context.Context
//  - in *EventsWithCountsStreamRequest
//  - opts ...grpc.CallOption
func (_e *MockEventQueryServiceClient_Expecter) EventsWithCountsStream(ctx interface{}, in interface{}, opts ...interface{}) *MockEventQueryServiceClient_EventsWithCountsStream_Call {
	return &MockEventQueryServiceClient_EventsWithCountsStream_Call{Call: _e.mock.On("EventsWithCountsStream",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockEventQueryServiceClient_EventsWithCountsStream_Call) Run(run func(ctx context.Context, in *EventsWithCountsStreamRequest, opts ...grpc.CallOption)) *MockEventQueryServiceClient_EventsWithCountsStream_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*EventsWithCountsStreamRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockEventQueryServiceClient_EventsWithCountsStream_Call) Return(_a0 EventQueryService_EventsWithCountsStreamClient, _a1 error) *MockEventQueryServiceClient_EventsWithCountsStream_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Frequency provides a mock function with given fields: ctx, in, opts
func (_m *MockEventQueryServiceClient) Frequency(ctx context.Context, in *FrequencyRequest, opts ...grpc.CallOption) (*FrequencyResponse, error) {
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

// MockEventQueryServiceClient_Frequency_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Frequency'
type MockEventQueryServiceClient_Frequency_Call struct {
	*mock.Call
}

// Frequency is a helper method to define mock.On call
//  - ctx context.Context
//  - in *FrequencyRequest
//  - opts ...grpc.CallOption
func (_e *MockEventQueryServiceClient_Expecter) Frequency(ctx interface{}, in interface{}, opts ...interface{}) *MockEventQueryServiceClient_Frequency_Call {
	return &MockEventQueryServiceClient_Frequency_Call{Call: _e.mock.On("Frequency",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockEventQueryServiceClient_Frequency_Call) Run(run func(ctx context.Context, in *FrequencyRequest, opts ...grpc.CallOption)) *MockEventQueryServiceClient_Frequency_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*FrequencyRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockEventQueryServiceClient_Frequency_Call) Return(_a0 *FrequencyResponse, _a1 error) *MockEventQueryServiceClient_Frequency_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetEvent provides a mock function with given fields: ctx, in, opts
func (_m *MockEventQueryServiceClient) GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*GetEventResponse, error) {
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

// MockEventQueryServiceClient_GetEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEvent'
type MockEventQueryServiceClient_GetEvent_Call struct {
	*mock.Call
}

// GetEvent is a helper method to define mock.On call
//  - ctx context.Context
//  - in *GetEventRequest
//  - opts ...grpc.CallOption
func (_e *MockEventQueryServiceClient_Expecter) GetEvent(ctx interface{}, in interface{}, opts ...interface{}) *MockEventQueryServiceClient_GetEvent_Call {
	return &MockEventQueryServiceClient_GetEvent_Call{Call: _e.mock.On("GetEvent",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockEventQueryServiceClient_GetEvent_Call) Run(run func(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption)) *MockEventQueryServiceClient_GetEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*GetEventRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockEventQueryServiceClient_GetEvent_Call) Return(_a0 *GetEventResponse, _a1 error) *MockEventQueryServiceClient_GetEvent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetEvents provides a mock function with given fields: ctx, in, opts
func (_m *MockEventQueryServiceClient) GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error) {
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

// MockEventQueryServiceClient_GetEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEvents'
type MockEventQueryServiceClient_GetEvents_Call struct {
	*mock.Call
}

// GetEvents is a helper method to define mock.On call
//  - ctx context.Context
//  - in *GetEventsRequest
//  - opts ...grpc.CallOption
func (_e *MockEventQueryServiceClient_Expecter) GetEvents(ctx interface{}, in interface{}, opts ...interface{}) *MockEventQueryServiceClient_GetEvents_Call {
	return &MockEventQueryServiceClient_GetEvents_Call{Call: _e.mock.On("GetEvents",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockEventQueryServiceClient_GetEvents_Call) Run(run func(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption)) *MockEventQueryServiceClient_GetEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*GetEventsRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockEventQueryServiceClient_GetEvents_Call) Return(_a0 *GetEventsResponse, _a1 error) *MockEventQueryServiceClient_GetEvents_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Search provides a mock function with given fields: ctx, in, opts
func (_m *MockEventQueryServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
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

// MockEventQueryServiceClient_Search_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Search'
type MockEventQueryServiceClient_Search_Call struct {
	*mock.Call
}

// Search is a helper method to define mock.On call
//  - ctx context.Context
//  - in *SearchRequest
//  - opts ...grpc.CallOption
func (_e *MockEventQueryServiceClient_Expecter) Search(ctx interface{}, in interface{}, opts ...interface{}) *MockEventQueryServiceClient_Search_Call {
	return &MockEventQueryServiceClient_Search_Call{Call: _e.mock.On("Search",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockEventQueryServiceClient_Search_Call) Run(run func(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption)) *MockEventQueryServiceClient_Search_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*SearchRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockEventQueryServiceClient_Search_Call) Return(_a0 *SearchResponse, _a1 error) *MockEventQueryServiceClient_Search_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// SearchStream provides a mock function with given fields: ctx, in, opts
func (_m *MockEventQueryServiceClient) SearchStream(ctx context.Context, in *SearchStreamRequest, opts ...grpc.CallOption) (EventQueryService_SearchStreamClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 EventQueryService_SearchStreamClient
	if rf, ok := ret.Get(0).(func(context.Context, *SearchStreamRequest, ...grpc.CallOption) EventQueryService_SearchStreamClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EventQueryService_SearchStreamClient)
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

// MockEventQueryServiceClient_SearchStream_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchStream'
type MockEventQueryServiceClient_SearchStream_Call struct {
	*mock.Call
}

// SearchStream is a helper method to define mock.On call
//  - ctx context.Context
//  - in *SearchStreamRequest
//  - opts ...grpc.CallOption
func (_e *MockEventQueryServiceClient_Expecter) SearchStream(ctx interface{}, in interface{}, opts ...interface{}) *MockEventQueryServiceClient_SearchStream_Call {
	return &MockEventQueryServiceClient_SearchStream_Call{Call: _e.mock.On("SearchStream",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockEventQueryServiceClient_SearchStream_Call) Run(run func(ctx context.Context, in *SearchStreamRequest, opts ...grpc.CallOption)) *MockEventQueryServiceClient_SearchStream_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*SearchStreamRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockEventQueryServiceClient_SearchStream_Call) Return(_a0 EventQueryService_SearchStreamClient, _a1 error) *MockEventQueryServiceClient_SearchStream_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewMockEventQueryServiceClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEventQueryServiceClient creates a new instance of MockEventQueryServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEventQueryServiceClient(t mockConstructorTestingTNewMockEventQueryServiceClient) *MockEventQueryServiceClient {
	mock := &MockEventQueryServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
