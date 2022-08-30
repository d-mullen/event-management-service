// Code generated by mockery v2.13.1. DO NOT EDIT.

package event_management

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockEventManagementClient is an autogenerated mock type for the EventManagementClient type
type MockEventManagementClient struct {
	mock.Mock
}

type MockEventManagementClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEventManagementClient) EXPECT() *MockEventManagementClient_Expecter {
	return &MockEventManagementClient_Expecter{mock: &_m.Mock}
}

// Annotate provides a mock function with given fields: ctx, in, opts
func (_m *MockEventManagementClient) Annotate(ctx context.Context, in *EventAnnotationRequest, opts ...grpc.CallOption) (*EventAnnotationResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *EventAnnotationResponse
	if rf, ok := ret.Get(0).(func(context.Context, *EventAnnotationRequest, ...grpc.CallOption) *EventAnnotationResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*EventAnnotationResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *EventAnnotationRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEventManagementClient_Annotate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Annotate'
type MockEventManagementClient_Annotate_Call struct {
	*mock.Call
}

// Annotate is a helper method to define mock.On call
//  - ctx context.Context
//  - in *EventAnnotationRequest
//  - opts ...grpc.CallOption
func (_e *MockEventManagementClient_Expecter) Annotate(ctx interface{}, in interface{}, opts ...interface{}) *MockEventManagementClient_Annotate_Call {
	return &MockEventManagementClient_Annotate_Call{Call: _e.mock.On("Annotate",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockEventManagementClient_Annotate_Call) Run(run func(ctx context.Context, in *EventAnnotationRequest, opts ...grpc.CallOption)) *MockEventManagementClient_Annotate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*EventAnnotationRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockEventManagementClient_Annotate_Call) Return(_a0 *EventAnnotationResponse, _a1 error) *MockEventManagementClient_Annotate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// DeleteAnnotations provides a mock function with given fields: ctx, in, opts
func (_m *MockEventManagementClient) DeleteAnnotations(ctx context.Context, in *EventAnnotationRequest, opts ...grpc.CallOption) (*EventAnnotationResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *EventAnnotationResponse
	if rf, ok := ret.Get(0).(func(context.Context, *EventAnnotationRequest, ...grpc.CallOption) *EventAnnotationResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*EventAnnotationResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *EventAnnotationRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEventManagementClient_DeleteAnnotations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAnnotations'
type MockEventManagementClient_DeleteAnnotations_Call struct {
	*mock.Call
}

// DeleteAnnotations is a helper method to define mock.On call
//  - ctx context.Context
//  - in *EventAnnotationRequest
//  - opts ...grpc.CallOption
func (_e *MockEventManagementClient_Expecter) DeleteAnnotations(ctx interface{}, in interface{}, opts ...interface{}) *MockEventManagementClient_DeleteAnnotations_Call {
	return &MockEventManagementClient_DeleteAnnotations_Call{Call: _e.mock.On("DeleteAnnotations",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockEventManagementClient_DeleteAnnotations_Call) Run(run func(ctx context.Context, in *EventAnnotationRequest, opts ...grpc.CallOption)) *MockEventManagementClient_DeleteAnnotations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*EventAnnotationRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockEventManagementClient_DeleteAnnotations_Call) Return(_a0 *EventAnnotationResponse, _a1 error) *MockEventManagementClient_DeleteAnnotations_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// SetStatus provides a mock function with given fields: ctx, in, opts
func (_m *MockEventManagementClient) SetStatus(ctx context.Context, in *EventStatusRequest, opts ...grpc.CallOption) (*EventStatusResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *EventStatusResponse
	if rf, ok := ret.Get(0).(func(context.Context, *EventStatusRequest, ...grpc.CallOption) *EventStatusResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*EventStatusResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *EventStatusRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEventManagementClient_SetStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetStatus'
type MockEventManagementClient_SetStatus_Call struct {
	*mock.Call
}

// SetStatus is a helper method to define mock.On call
//  - ctx context.Context
//  - in *EventStatusRequest
//  - opts ...grpc.CallOption
func (_e *MockEventManagementClient_Expecter) SetStatus(ctx interface{}, in interface{}, opts ...interface{}) *MockEventManagementClient_SetStatus_Call {
	return &MockEventManagementClient_SetStatus_Call{Call: _e.mock.On("SetStatus",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockEventManagementClient_SetStatus_Call) Run(run func(ctx context.Context, in *EventStatusRequest, opts ...grpc.CallOption)) *MockEventManagementClient_SetStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*EventStatusRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockEventManagementClient_SetStatus_Call) Return(_a0 *EventStatusResponse, _a1 error) *MockEventManagementClient_SetStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewMockEventManagementClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEventManagementClient creates a new instance of MockEventManagementClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEventManagementClient(t mockConstructorTestingTNewMockEventManagementClient) *MockEventManagementClient {
	mock := &MockEventManagementClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
