// Code generated by mockery v2.13.1. DO NOT EDIT.

package yamr

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockYamrPageRankIngestClient is an autogenerated mock type for the YamrPageRankIngestClient type
type MockYamrPageRankIngestClient struct {
	mock.Mock
}

type MockYamrPageRankIngestClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockYamrPageRankIngestClient) EXPECT() *MockYamrPageRankIngestClient_Expecter {
	return &MockYamrPageRankIngestClient_Expecter{mock: &_m.Mock}
}

// Update provides a mock function with given fields: ctx, opts
func (_m *MockYamrPageRankIngestClient) Update(ctx context.Context, opts ...grpc.CallOption) (YamrPageRankIngest_UpdateClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 YamrPageRankIngest_UpdateClient
	if rf, ok := ret.Get(0).(func(context.Context, ...grpc.CallOption) YamrPageRankIngest_UpdateClient); ok {
		r0 = rf(ctx, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(YamrPageRankIngest_UpdateClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockYamrPageRankIngestClient_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockYamrPageRankIngestClient_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//  - ctx context.Context
//  - opts ...grpc.CallOption
func (_e *MockYamrPageRankIngestClient_Expecter) Update(ctx interface{}, opts ...interface{}) *MockYamrPageRankIngestClient_Update_Call {
	return &MockYamrPageRankIngestClient_Update_Call{Call: _e.mock.On("Update",
		append([]interface{}{ctx}, opts...)...)}
}

func (_c *MockYamrPageRankIngestClient_Update_Call) Run(run func(ctx context.Context, opts ...grpc.CallOption)) *MockYamrPageRankIngestClient_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *MockYamrPageRankIngestClient_Update_Call) Return(_a0 YamrPageRankIngest_UpdateClient, _a1 error) *MockYamrPageRankIngestClient_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewMockYamrPageRankIngestClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockYamrPageRankIngestClient creates a new instance of MockYamrPageRankIngestClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockYamrPageRankIngestClient(t mockConstructorTestingTNewMockYamrPageRankIngestClient) *MockYamrPageRankIngestClient {
	mock := &MockYamrPageRankIngestClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
