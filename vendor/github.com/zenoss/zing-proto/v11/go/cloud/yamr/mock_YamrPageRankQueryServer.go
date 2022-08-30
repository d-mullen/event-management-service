// Code generated by mockery v2.13.1. DO NOT EDIT.

package yamr

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockYamrPageRankQueryServer is an autogenerated mock type for the YamrPageRankQueryServer type
type MockYamrPageRankQueryServer struct {
	mock.Mock
}

type MockYamrPageRankQueryServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockYamrPageRankQueryServer) EXPECT() *MockYamrPageRankQueryServer_Expecter {
	return &MockYamrPageRankQueryServer_Expecter{mock: &_m.Mock}
}

// GetTopN provides a mock function with given fields: _a0, _a1
func (_m *MockYamrPageRankQueryServer) GetTopN(_a0 context.Context, _a1 *YamrPageRankGetTopNRequest) (*YamrPageRankGetTopNResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *YamrPageRankGetTopNResponse
	if rf, ok := ret.Get(0).(func(context.Context, *YamrPageRankGetTopNRequest) *YamrPageRankGetTopNResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*YamrPageRankGetTopNResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *YamrPageRankGetTopNRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockYamrPageRankQueryServer_GetTopN_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTopN'
type MockYamrPageRankQueryServer_GetTopN_Call struct {
	*mock.Call
}

// GetTopN is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *YamrPageRankGetTopNRequest
func (_e *MockYamrPageRankQueryServer_Expecter) GetTopN(_a0 interface{}, _a1 interface{}) *MockYamrPageRankQueryServer_GetTopN_Call {
	return &MockYamrPageRankQueryServer_GetTopN_Call{Call: _e.mock.On("GetTopN", _a0, _a1)}
}

func (_c *MockYamrPageRankQueryServer_GetTopN_Call) Run(run func(_a0 context.Context, _a1 *YamrPageRankGetTopNRequest)) *MockYamrPageRankQueryServer_GetTopN_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*YamrPageRankGetTopNRequest))
	})
	return _c
}

func (_c *MockYamrPageRankQueryServer_GetTopN_Call) Return(_a0 *YamrPageRankGetTopNResponse, _a1 error) *MockYamrPageRankQueryServer_GetTopN_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// mustEmbedUnimplementedYamrPageRankQueryServer provides a mock function with given fields:
func (_m *MockYamrPageRankQueryServer) mustEmbedUnimplementedYamrPageRankQueryServer() {
	_m.Called()
}

// MockYamrPageRankQueryServer_mustEmbedUnimplementedYamrPageRankQueryServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedYamrPageRankQueryServer'
type MockYamrPageRankQueryServer_mustEmbedUnimplementedYamrPageRankQueryServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedYamrPageRankQueryServer is a helper method to define mock.On call
func (_e *MockYamrPageRankQueryServer_Expecter) mustEmbedUnimplementedYamrPageRankQueryServer() *MockYamrPageRankQueryServer_mustEmbedUnimplementedYamrPageRankQueryServer_Call {
	return &MockYamrPageRankQueryServer_mustEmbedUnimplementedYamrPageRankQueryServer_Call{Call: _e.mock.On("mustEmbedUnimplementedYamrPageRankQueryServer")}
}

func (_c *MockYamrPageRankQueryServer_mustEmbedUnimplementedYamrPageRankQueryServer_Call) Run(run func()) *MockYamrPageRankQueryServer_mustEmbedUnimplementedYamrPageRankQueryServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockYamrPageRankQueryServer_mustEmbedUnimplementedYamrPageRankQueryServer_Call) Return() *MockYamrPageRankQueryServer_mustEmbedUnimplementedYamrPageRankQueryServer_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewMockYamrPageRankQueryServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockYamrPageRankQueryServer creates a new instance of MockYamrPageRankQueryServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockYamrPageRankQueryServer(t mockConstructorTestingTNewMockYamrPageRankQueryServer) *MockYamrPageRankQueryServer {
	mock := &MockYamrPageRankQueryServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
