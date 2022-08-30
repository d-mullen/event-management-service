// Code generated by mockery v2.14.0. DO NOT EDIT.

package yamr

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockYamrExecutorClient is an autogenerated mock type for the YamrExecutorClient type
type MockYamrExecutorClient struct {
	mock.Mock
}

type MockYamrExecutorClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockYamrExecutorClient) EXPECT() *MockYamrExecutorClient_Expecter {
	return &MockYamrExecutorClient_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, in, opts
func (_m *MockYamrExecutorClient) Execute(ctx context.Context, in *YamrExecuteRequest, opts ...grpc.CallOption) (YamrExecutor_ExecuteClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 YamrExecutor_ExecuteClient
	if rf, ok := ret.Get(0).(func(context.Context, *YamrExecuteRequest, ...grpc.CallOption) YamrExecutor_ExecuteClient); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(YamrExecutor_ExecuteClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *YamrExecuteRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockYamrExecutorClient_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockYamrExecutorClient_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//  - ctx context.Context
//  - in *YamrExecuteRequest
//  - opts ...grpc.CallOption
func (_e *MockYamrExecutorClient_Expecter) Execute(ctx interface{}, in interface{}, opts ...interface{}) *MockYamrExecutorClient_Execute_Call {
	return &MockYamrExecutorClient_Execute_Call{Call: _e.mock.On("Execute",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockYamrExecutorClient_Execute_Call) Run(run func(ctx context.Context, in *YamrExecuteRequest, opts ...grpc.CallOption)) *MockYamrExecutorClient_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*YamrExecuteRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockYamrExecutorClient_Execute_Call) Return(_a0 YamrExecutor_ExecuteClient, _a1 error) *MockYamrExecutorClient_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewMockYamrExecutorClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockYamrExecutorClient creates a new instance of MockYamrExecutorClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockYamrExecutorClient(t mockConstructorTestingTNewMockYamrExecutorClient) *MockYamrExecutorClient {
	mock := &MockYamrExecutorClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
