// Code generated by mockery v2.14.0. DO NOT EDIT.

package event_context

import mock "github.com/stretchr/testify/mock"

// MockUnsafeEventContextQueryServer is an autogenerated mock type for the UnsafeEventContextQueryServer type
type MockUnsafeEventContextQueryServer struct {
	mock.Mock
}

type MockUnsafeEventContextQueryServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUnsafeEventContextQueryServer) EXPECT() *MockUnsafeEventContextQueryServer_Expecter {
	return &MockUnsafeEventContextQueryServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedEventContextQueryServer provides a mock function with given fields:
func (_m *MockUnsafeEventContextQueryServer) mustEmbedUnimplementedEventContextQueryServer() {
	_m.Called()
}

// MockUnsafeEventContextQueryServer_mustEmbedUnimplementedEventContextQueryServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedEventContextQueryServer'
type MockUnsafeEventContextQueryServer_mustEmbedUnimplementedEventContextQueryServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedEventContextQueryServer is a helper method to define mock.On call
func (_e *MockUnsafeEventContextQueryServer_Expecter) mustEmbedUnimplementedEventContextQueryServer() *MockUnsafeEventContextQueryServer_mustEmbedUnimplementedEventContextQueryServer_Call {
	return &MockUnsafeEventContextQueryServer_mustEmbedUnimplementedEventContextQueryServer_Call{Call: _e.mock.On("mustEmbedUnimplementedEventContextQueryServer")}
}

func (_c *MockUnsafeEventContextQueryServer_mustEmbedUnimplementedEventContextQueryServer_Call) Run(run func()) *MockUnsafeEventContextQueryServer_mustEmbedUnimplementedEventContextQueryServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUnsafeEventContextQueryServer_mustEmbedUnimplementedEventContextQueryServer_Call) Return() *MockUnsafeEventContextQueryServer_mustEmbedUnimplementedEventContextQueryServer_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewMockUnsafeEventContextQueryServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockUnsafeEventContextQueryServer creates a new instance of MockUnsafeEventContextQueryServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUnsafeEventContextQueryServer(t mockConstructorTestingTNewMockUnsafeEventContextQueryServer) *MockUnsafeEventContextQueryServer {
	mock := &MockUnsafeEventContextQueryServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
