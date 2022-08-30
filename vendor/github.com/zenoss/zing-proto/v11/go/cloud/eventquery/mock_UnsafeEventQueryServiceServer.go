// Code generated by mockery v2.14.0. DO NOT EDIT.

package eventquery

import mock "github.com/stretchr/testify/mock"

// MockUnsafeEventQueryServiceServer is an autogenerated mock type for the UnsafeEventQueryServiceServer type
type MockUnsafeEventQueryServiceServer struct {
	mock.Mock
}

type MockUnsafeEventQueryServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUnsafeEventQueryServiceServer) EXPECT() *MockUnsafeEventQueryServiceServer_Expecter {
	return &MockUnsafeEventQueryServiceServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedEventQueryServiceServer provides a mock function with given fields:
func (_m *MockUnsafeEventQueryServiceServer) mustEmbedUnimplementedEventQueryServiceServer() {
	_m.Called()
}

// MockUnsafeEventQueryServiceServer_mustEmbedUnimplementedEventQueryServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedEventQueryServiceServer'
type MockUnsafeEventQueryServiceServer_mustEmbedUnimplementedEventQueryServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedEventQueryServiceServer is a helper method to define mock.On call
func (_e *MockUnsafeEventQueryServiceServer_Expecter) mustEmbedUnimplementedEventQueryServiceServer() *MockUnsafeEventQueryServiceServer_mustEmbedUnimplementedEventQueryServiceServer_Call {
	return &MockUnsafeEventQueryServiceServer_mustEmbedUnimplementedEventQueryServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedEventQueryServiceServer")}
}

func (_c *MockUnsafeEventQueryServiceServer_mustEmbedUnimplementedEventQueryServiceServer_Call) Run(run func()) *MockUnsafeEventQueryServiceServer_mustEmbedUnimplementedEventQueryServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUnsafeEventQueryServiceServer_mustEmbedUnimplementedEventQueryServiceServer_Call) Return() *MockUnsafeEventQueryServiceServer_mustEmbedUnimplementedEventQueryServiceServer_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewMockUnsafeEventQueryServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockUnsafeEventQueryServiceServer creates a new instance of MockUnsafeEventQueryServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUnsafeEventQueryServiceServer(t mockConstructorTestingTNewMockUnsafeEventQueryServiceServer) *MockUnsafeEventQueryServiceServer {
	mock := &MockUnsafeEventQueryServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
