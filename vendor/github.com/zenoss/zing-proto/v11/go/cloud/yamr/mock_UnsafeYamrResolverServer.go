// Code generated by mockery v2.12.2. DO NOT EDIT.

package yamr

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// MockUnsafeYamrResolverServer is an autogenerated mock type for the UnsafeYamrResolverServer type
type MockUnsafeYamrResolverServer struct {
	mock.Mock
}

type MockUnsafeYamrResolverServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUnsafeYamrResolverServer) EXPECT() *MockUnsafeYamrResolverServer_Expecter {
	return &MockUnsafeYamrResolverServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedYamrResolverServer provides a mock function with given fields:
func (_m *MockUnsafeYamrResolverServer) mustEmbedUnimplementedYamrResolverServer() {
	_m.Called()
}

// MockUnsafeYamrResolverServer_mustEmbedUnimplementedYamrResolverServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedYamrResolverServer'
type MockUnsafeYamrResolverServer_mustEmbedUnimplementedYamrResolverServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedYamrResolverServer is a helper method to define mock.On call
func (_e *MockUnsafeYamrResolverServer_Expecter) mustEmbedUnimplementedYamrResolverServer() *MockUnsafeYamrResolverServer_mustEmbedUnimplementedYamrResolverServer_Call {
	return &MockUnsafeYamrResolverServer_mustEmbedUnimplementedYamrResolverServer_Call{Call: _e.mock.On("mustEmbedUnimplementedYamrResolverServer")}
}

func (_c *MockUnsafeYamrResolverServer_mustEmbedUnimplementedYamrResolverServer_Call) Run(run func()) *MockUnsafeYamrResolverServer_mustEmbedUnimplementedYamrResolverServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUnsafeYamrResolverServer_mustEmbedUnimplementedYamrResolverServer_Call) Return() *MockUnsafeYamrResolverServer_mustEmbedUnimplementedYamrResolverServer_Call {
	_c.Call.Return()
	return _c
}

// NewMockUnsafeYamrResolverServer creates a new instance of MockUnsafeYamrResolverServer. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUnsafeYamrResolverServer(t testing.TB) *MockUnsafeYamrResolverServer {
	mock := &MockUnsafeYamrResolverServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}