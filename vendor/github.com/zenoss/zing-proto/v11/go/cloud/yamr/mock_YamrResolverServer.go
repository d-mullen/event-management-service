// Code generated by mockery v2.14.0. DO NOT EDIT.

package yamr

import mock "github.com/stretchr/testify/mock"

// MockYamrResolverServer is an autogenerated mock type for the YamrResolverServer type
type MockYamrResolverServer struct {
	mock.Mock
}

type MockYamrResolverServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockYamrResolverServer) EXPECT() *MockYamrResolverServer_Expecter {
	return &MockYamrResolverServer_Expecter{mock: &_m.Mock}
}

// Resolve provides a mock function with given fields: _a0, _a1
func (_m *MockYamrResolverServer) Resolve(_a0 *YamrResolveRequest, _a1 YamrResolver_ResolveServer) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*YamrResolveRequest, YamrResolver_ResolveServer) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockYamrResolverServer_Resolve_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Resolve'
type MockYamrResolverServer_Resolve_Call struct {
	*mock.Call
}

// Resolve is a helper method to define mock.On call
//  - _a0 *YamrResolveRequest
//  - _a1 YamrResolver_ResolveServer
func (_e *MockYamrResolverServer_Expecter) Resolve(_a0 interface{}, _a1 interface{}) *MockYamrResolverServer_Resolve_Call {
	return &MockYamrResolverServer_Resolve_Call{Call: _e.mock.On("Resolve", _a0, _a1)}
}

func (_c *MockYamrResolverServer_Resolve_Call) Run(run func(_a0 *YamrResolveRequest, _a1 YamrResolver_ResolveServer)) *MockYamrResolverServer_Resolve_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*YamrResolveRequest), args[1].(YamrResolver_ResolveServer))
	})
	return _c
}

func (_c *MockYamrResolverServer_Resolve_Call) Return(_a0 error) *MockYamrResolverServer_Resolve_Call {
	_c.Call.Return(_a0)
	return _c
}

// mustEmbedUnimplementedYamrResolverServer provides a mock function with given fields:
func (_m *MockYamrResolverServer) mustEmbedUnimplementedYamrResolverServer() {
	_m.Called()
}

// MockYamrResolverServer_mustEmbedUnimplementedYamrResolverServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedYamrResolverServer'
type MockYamrResolverServer_mustEmbedUnimplementedYamrResolverServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedYamrResolverServer is a helper method to define mock.On call
func (_e *MockYamrResolverServer_Expecter) mustEmbedUnimplementedYamrResolverServer() *MockYamrResolverServer_mustEmbedUnimplementedYamrResolverServer_Call {
	return &MockYamrResolverServer_mustEmbedUnimplementedYamrResolverServer_Call{Call: _e.mock.On("mustEmbedUnimplementedYamrResolverServer")}
}

func (_c *MockYamrResolverServer_mustEmbedUnimplementedYamrResolverServer_Call) Run(run func()) *MockYamrResolverServer_mustEmbedUnimplementedYamrResolverServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockYamrResolverServer_mustEmbedUnimplementedYamrResolverServer_Call) Return() *MockYamrResolverServer_mustEmbedUnimplementedYamrResolverServer_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewMockYamrResolverServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockYamrResolverServer creates a new instance of MockYamrResolverServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockYamrResolverServer(t mockConstructorTestingTNewMockYamrResolverServer) *MockYamrResolverServer {
	mock := &MockYamrResolverServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
