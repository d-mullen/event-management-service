// Code generated by mockery v2.14.0. DO NOT EDIT.

package event_context

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockEventContextQueryClient is an autogenerated mock type for the EventContextQueryClient type
type MockEventContextQueryClient struct {
	mock.Mock
}

type MockEventContextQueryClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEventContextQueryClient) EXPECT() *MockEventContextQueryClient_Expecter {
	return &MockEventContextQueryClient_Expecter{mock: &_m.Mock}
}

// GetActiveEvents provides a mock function with given fields: ctx, in, opts
func (_m *MockEventContextQueryClient) GetActiveEvents(ctx context.Context, in *ECGetActiveEventsRequest, opts ...grpc.CallOption) (*ECGetActiveEventsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ECGetActiveEventsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ECGetActiveEventsRequest, ...grpc.CallOption) *ECGetActiveEventsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ECGetActiveEventsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ECGetActiveEventsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEventContextQueryClient_GetActiveEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetActiveEvents'
type MockEventContextQueryClient_GetActiveEvents_Call struct {
	*mock.Call
}

// GetActiveEvents is a helper method to define mock.On call
//  - ctx context.Context
//  - in *ECGetActiveEventsRequest
//  - opts ...grpc.CallOption
func (_e *MockEventContextQueryClient_Expecter) GetActiveEvents(ctx interface{}, in interface{}, opts ...interface{}) *MockEventContextQueryClient_GetActiveEvents_Call {
	return &MockEventContextQueryClient_GetActiveEvents_Call{Call: _e.mock.On("GetActiveEvents",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockEventContextQueryClient_GetActiveEvents_Call) Run(run func(ctx context.Context, in *ECGetActiveEventsRequest, opts ...grpc.CallOption)) *MockEventContextQueryClient_GetActiveEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*ECGetActiveEventsRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockEventContextQueryClient_GetActiveEvents_Call) Return(_a0 *ECGetActiveEventsResponse, _a1 error) *MockEventContextQueryClient_GetActiveEvents_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetBulk provides a mock function with given fields: ctx, in, opts
func (_m *MockEventContextQueryClient) GetBulk(ctx context.Context, in *ECGetBulkRequest, opts ...grpc.CallOption) (*ECGetBulkResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ECGetBulkResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ECGetBulkRequest, ...grpc.CallOption) *ECGetBulkResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ECGetBulkResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ECGetBulkRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEventContextQueryClient_GetBulk_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBulk'
type MockEventContextQueryClient_GetBulk_Call struct {
	*mock.Call
}

// GetBulk is a helper method to define mock.On call
//  - ctx context.Context
//  - in *ECGetBulkRequest
//  - opts ...grpc.CallOption
func (_e *MockEventContextQueryClient_Expecter) GetBulk(ctx interface{}, in interface{}, opts ...interface{}) *MockEventContextQueryClient_GetBulk_Call {
	return &MockEventContextQueryClient_GetBulk_Call{Call: _e.mock.On("GetBulk",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockEventContextQueryClient_GetBulk_Call) Run(run func(ctx context.Context, in *ECGetBulkRequest, opts ...grpc.CallOption)) *MockEventContextQueryClient_GetBulk_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*ECGetBulkRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockEventContextQueryClient_GetBulk_Call) Return(_a0 *ECGetBulkResponse, _a1 error) *MockEventContextQueryClient_GetBulk_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Search provides a mock function with given fields: ctx, in, opts
func (_m *MockEventContextQueryClient) Search(ctx context.Context, in *ECSearchRequest, opts ...grpc.CallOption) (*ECSearchResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ECSearchResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ECSearchRequest, ...grpc.CallOption) *ECSearchResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ECSearchResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ECSearchRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEventContextQueryClient_Search_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Search'
type MockEventContextQueryClient_Search_Call struct {
	*mock.Call
}

// Search is a helper method to define mock.On call
//  - ctx context.Context
//  - in *ECSearchRequest
//  - opts ...grpc.CallOption
func (_e *MockEventContextQueryClient_Expecter) Search(ctx interface{}, in interface{}, opts ...interface{}) *MockEventContextQueryClient_Search_Call {
	return &MockEventContextQueryClient_Search_Call{Call: _e.mock.On("Search",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockEventContextQueryClient_Search_Call) Run(run func(ctx context.Context, in *ECSearchRequest, opts ...grpc.CallOption)) *MockEventContextQueryClient_Search_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*ECSearchRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockEventContextQueryClient_Search_Call) Return(_a0 *ECSearchResponse, _a1 error) *MockEventContextQueryClient_Search_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// StreamingSearch provides a mock function with given fields: ctx, in, opts
func (_m *MockEventContextQueryClient) StreamingSearch(ctx context.Context, in *ECStreamingSearchRequest, opts ...grpc.CallOption) (EventContextQuery_StreamingSearchClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 EventContextQuery_StreamingSearchClient
	if rf, ok := ret.Get(0).(func(context.Context, *ECStreamingSearchRequest, ...grpc.CallOption) EventContextQuery_StreamingSearchClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EventContextQuery_StreamingSearchClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ECStreamingSearchRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEventContextQueryClient_StreamingSearch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StreamingSearch'
type MockEventContextQueryClient_StreamingSearch_Call struct {
	*mock.Call
}

// StreamingSearch is a helper method to define mock.On call
//  - ctx context.Context
//  - in *ECStreamingSearchRequest
//  - opts ...grpc.CallOption
func (_e *MockEventContextQueryClient_Expecter) StreamingSearch(ctx interface{}, in interface{}, opts ...interface{}) *MockEventContextQueryClient_StreamingSearch_Call {
	return &MockEventContextQueryClient_StreamingSearch_Call{Call: _e.mock.On("StreamingSearch",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockEventContextQueryClient_StreamingSearch_Call) Run(run func(ctx context.Context, in *ECStreamingSearchRequest, opts ...grpc.CallOption)) *MockEventContextQueryClient_StreamingSearch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*ECStreamingSearchRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockEventContextQueryClient_StreamingSearch_Call) Return(_a0 EventContextQuery_StreamingSearchClient, _a1 error) *MockEventContextQueryClient_StreamingSearch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewMockEventContextQueryClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEventContextQueryClient creates a new instance of MockEventContextQueryClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEventContextQueryClient(t mockConstructorTestingTNewMockEventContextQueryClient) *MockEventContextQueryClient {
	mock := &MockEventContextQueryClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
