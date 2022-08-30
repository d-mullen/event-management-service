// Code generated by mockery v2.14.0. DO NOT EDIT.

package event_context

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockEventContextIngestClient is an autogenerated mock type for the EventContextIngestClient type
type MockEventContextIngestClient struct {
	mock.Mock
}

type MockEventContextIngestClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEventContextIngestClient) EXPECT() *MockEventContextIngestClient_Expecter {
	return &MockEventContextIngestClient_Expecter{mock: &_m.Mock}
}

// DeleteTenantData provides a mock function with given fields: ctx, in, opts
func (_m *MockEventContextIngestClient) DeleteTenantData(ctx context.Context, in *DeleteTenantDataRequest, opts ...grpc.CallOption) (*DeleteTenantDataResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *DeleteTenantDataResponse
	if rf, ok := ret.Get(0).(func(context.Context, *DeleteTenantDataRequest, ...grpc.CallOption) *DeleteTenantDataResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DeleteTenantDataResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *DeleteTenantDataRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEventContextIngestClient_DeleteTenantData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteTenantData'
type MockEventContextIngestClient_DeleteTenantData_Call struct {
	*mock.Call
}

// DeleteTenantData is a helper method to define mock.On call
//  - ctx context.Context
//  - in *DeleteTenantDataRequest
//  - opts ...grpc.CallOption
func (_e *MockEventContextIngestClient_Expecter) DeleteTenantData(ctx interface{}, in interface{}, opts ...interface{}) *MockEventContextIngestClient_DeleteTenantData_Call {
	return &MockEventContextIngestClient_DeleteTenantData_Call{Call: _e.mock.On("DeleteTenantData",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockEventContextIngestClient_DeleteTenantData_Call) Run(run func(ctx context.Context, in *DeleteTenantDataRequest, opts ...grpc.CallOption)) *MockEventContextIngestClient_DeleteTenantData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*DeleteTenantDataRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockEventContextIngestClient_DeleteTenantData_Call) Return(_a0 *DeleteTenantDataResponse, _a1 error) *MockEventContextIngestClient_DeleteTenantData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// PutEvent provides a mock function with given fields: ctx, in, opts
func (_m *MockEventContextIngestClient) PutEvent(ctx context.Context, in *PutEventRequest, opts ...grpc.CallOption) (*PutEventResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *PutEventResponse
	if rf, ok := ret.Get(0).(func(context.Context, *PutEventRequest, ...grpc.CallOption) *PutEventResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PutEventResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *PutEventRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEventContextIngestClient_PutEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutEvent'
type MockEventContextIngestClient_PutEvent_Call struct {
	*mock.Call
}

// PutEvent is a helper method to define mock.On call
//  - ctx context.Context
//  - in *PutEventRequest
//  - opts ...grpc.CallOption
func (_e *MockEventContextIngestClient_Expecter) PutEvent(ctx interface{}, in interface{}, opts ...interface{}) *MockEventContextIngestClient_PutEvent_Call {
	return &MockEventContextIngestClient_PutEvent_Call{Call: _e.mock.On("PutEvent",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockEventContextIngestClient_PutEvent_Call) Run(run func(ctx context.Context, in *PutEventRequest, opts ...grpc.CallOption)) *MockEventContextIngestClient_PutEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*PutEventRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockEventContextIngestClient_PutEvent_Call) Return(_a0 *PutEventResponse, _a1 error) *MockEventContextIngestClient_PutEvent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// PutEventBulk provides a mock function with given fields: ctx, in, opts
func (_m *MockEventContextIngestClient) PutEventBulk(ctx context.Context, in *PutEventBulkRequest, opts ...grpc.CallOption) (*PutEventBulkResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *PutEventBulkResponse
	if rf, ok := ret.Get(0).(func(context.Context, *PutEventBulkRequest, ...grpc.CallOption) *PutEventBulkResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PutEventBulkResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *PutEventBulkRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEventContextIngestClient_PutEventBulk_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutEventBulk'
type MockEventContextIngestClient_PutEventBulk_Call struct {
	*mock.Call
}

// PutEventBulk is a helper method to define mock.On call
//  - ctx context.Context
//  - in *PutEventBulkRequest
//  - opts ...grpc.CallOption
func (_e *MockEventContextIngestClient_Expecter) PutEventBulk(ctx interface{}, in interface{}, opts ...interface{}) *MockEventContextIngestClient_PutEventBulk_Call {
	return &MockEventContextIngestClient_PutEventBulk_Call{Call: _e.mock.On("PutEventBulk",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockEventContextIngestClient_PutEventBulk_Call) Run(run func(ctx context.Context, in *PutEventBulkRequest, opts ...grpc.CallOption)) *MockEventContextIngestClient_PutEventBulk_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*PutEventBulkRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockEventContextIngestClient_PutEventBulk_Call) Return(_a0 *PutEventBulkResponse, _a1 error) *MockEventContextIngestClient_PutEventBulk_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdateEvent provides a mock function with given fields: ctx, in, opts
func (_m *MockEventContextIngestClient) UpdateEvent(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption) (*UpdateEventResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *UpdateEventResponse
	if rf, ok := ret.Get(0).(func(context.Context, *UpdateEventRequest, ...grpc.CallOption) *UpdateEventResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UpdateEventResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *UpdateEventRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEventContextIngestClient_UpdateEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateEvent'
type MockEventContextIngestClient_UpdateEvent_Call struct {
	*mock.Call
}

// UpdateEvent is a helper method to define mock.On call
//  - ctx context.Context
//  - in *UpdateEventRequest
//  - opts ...grpc.CallOption
func (_e *MockEventContextIngestClient_Expecter) UpdateEvent(ctx interface{}, in interface{}, opts ...interface{}) *MockEventContextIngestClient_UpdateEvent_Call {
	return &MockEventContextIngestClient_UpdateEvent_Call{Call: _e.mock.On("UpdateEvent",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockEventContextIngestClient_UpdateEvent_Call) Run(run func(ctx context.Context, in *UpdateEventRequest, opts ...grpc.CallOption)) *MockEventContextIngestClient_UpdateEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*UpdateEventRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockEventContextIngestClient_UpdateEvent_Call) Return(_a0 *UpdateEventResponse, _a1 error) *MockEventContextIngestClient_UpdateEvent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewMockEventContextIngestClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEventContextIngestClient creates a new instance of MockEventContextIngestClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEventContextIngestClient(t mockConstructorTestingTNewMockEventContextIngestClient) *MockEventContextIngestClient {
	mock := &MockEventContextIngestClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
