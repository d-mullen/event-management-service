// Code generated by mockery v2.14.0. DO NOT EDIT.

package yamr

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockYamrIngestClient is an autogenerated mock type for the YamrIngestClient type
type MockYamrIngestClient struct {
	mock.Mock
}

type MockYamrIngestClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockYamrIngestClient) EXPECT() *MockYamrIngestClient_Expecter {
	return &MockYamrIngestClient_Expecter{mock: &_m.Mock}
}

// Put provides a mock function with given fields: ctx, in, opts
func (_m *MockYamrIngestClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *PutResponse
	if rf, ok := ret.Get(0).(func(context.Context, *PutRequest, ...grpc.CallOption) *PutResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PutResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *PutRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockYamrIngestClient_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type MockYamrIngestClient_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//  - ctx context.Context
//  - in *PutRequest
//  - opts ...grpc.CallOption
func (_e *MockYamrIngestClient_Expecter) Put(ctx interface{}, in interface{}, opts ...interface{}) *MockYamrIngestClient_Put_Call {
	return &MockYamrIngestClient_Put_Call{Call: _e.mock.On("Put",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockYamrIngestClient_Put_Call) Run(run func(ctx context.Context, in *PutRequest, opts ...grpc.CallOption)) *MockYamrIngestClient_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*PutRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockYamrIngestClient_Put_Call) Return(_a0 *PutResponse, _a1 error) *MockYamrIngestClient_Put_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// PutBulk provides a mock function with given fields: ctx, in, opts
func (_m *MockYamrIngestClient) PutBulk(ctx context.Context, in *PutBulkRequest, opts ...grpc.CallOption) (*PutBulkResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *PutBulkResponse
	if rf, ok := ret.Get(0).(func(context.Context, *PutBulkRequest, ...grpc.CallOption) *PutBulkResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PutBulkResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *PutBulkRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockYamrIngestClient_PutBulk_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutBulk'
type MockYamrIngestClient_PutBulk_Call struct {
	*mock.Call
}

// PutBulk is a helper method to define mock.On call
//  - ctx context.Context
//  - in *PutBulkRequest
//  - opts ...grpc.CallOption
func (_e *MockYamrIngestClient_Expecter) PutBulk(ctx interface{}, in interface{}, opts ...interface{}) *MockYamrIngestClient_PutBulk_Call {
	return &MockYamrIngestClient_PutBulk_Call{Call: _e.mock.On("PutBulk",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockYamrIngestClient_PutBulk_Call) Run(run func(ctx context.Context, in *PutBulkRequest, opts ...grpc.CallOption)) *MockYamrIngestClient_PutBulk_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*PutBulkRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockYamrIngestClient_PutBulk_Call) Return(_a0 *PutBulkResponse, _a1 error) *MockYamrIngestClient_PutBulk_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewMockYamrIngestClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockYamrIngestClient creates a new instance of MockYamrIngestClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockYamrIngestClient(t mockConstructorTestingTNewMockYamrIngestClient) *MockYamrIngestClient {
	mock := &MockYamrIngestClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
