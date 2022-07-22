// Code generated by mockery v2.12.2. DO NOT EDIT.

package event_management

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// MockUnsafeEventManagementServer is an autogenerated mock type for the UnsafeEventManagementServer type
type MockUnsafeEventManagementServer struct {
	mock.Mock
}

type MockUnsafeEventManagementServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUnsafeEventManagementServer) EXPECT() *MockUnsafeEventManagementServer_Expecter {
	return &MockUnsafeEventManagementServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedEventManagementServer provides a mock function with given fields:
func (_m *MockUnsafeEventManagementServer) mustEmbedUnimplementedEventManagementServer() {
	_m.Called()
}

// MockUnsafeEventManagementServer_mustEmbedUnimplementedEventManagementServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedEventManagementServer'
type MockUnsafeEventManagementServer_mustEmbedUnimplementedEventManagementServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedEventManagementServer is a helper method to define mock.On call
func (_e *MockUnsafeEventManagementServer_Expecter) mustEmbedUnimplementedEventManagementServer() *MockUnsafeEventManagementServer_mustEmbedUnimplementedEventManagementServer_Call {
	return &MockUnsafeEventManagementServer_mustEmbedUnimplementedEventManagementServer_Call{Call: _e.mock.On("mustEmbedUnimplementedEventManagementServer")}
}

func (_c *MockUnsafeEventManagementServer_mustEmbedUnimplementedEventManagementServer_Call) Run(run func()) *MockUnsafeEventManagementServer_mustEmbedUnimplementedEventManagementServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUnsafeEventManagementServer_mustEmbedUnimplementedEventManagementServer_Call) Return() *MockUnsafeEventManagementServer_mustEmbedUnimplementedEventManagementServer_Call {
	_c.Call.Return()
	return _c
}

// NewMockUnsafeEventManagementServer creates a new instance of MockUnsafeEventManagementServer. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUnsafeEventManagementServer(t testing.TB) *MockUnsafeEventManagementServer {
	mock := &MockUnsafeEventManagementServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}