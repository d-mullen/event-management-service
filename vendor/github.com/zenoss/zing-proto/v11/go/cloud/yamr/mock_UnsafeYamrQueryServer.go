// Code generated by mockery v2.12.2. DO NOT EDIT.

package yamr

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// MockUnsafeYamrQueryServer is an autogenerated mock type for the UnsafeYamrQueryServer type
type MockUnsafeYamrQueryServer struct {
	mock.Mock
}

type MockUnsafeYamrQueryServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUnsafeYamrQueryServer) EXPECT() *MockUnsafeYamrQueryServer_Expecter {
	return &MockUnsafeYamrQueryServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedYamrQueryServer provides a mock function with given fields:
func (_m *MockUnsafeYamrQueryServer) mustEmbedUnimplementedYamrQueryServer() {
	_m.Called()
}

// MockUnsafeYamrQueryServer_mustEmbedUnimplementedYamrQueryServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedYamrQueryServer'
type MockUnsafeYamrQueryServer_mustEmbedUnimplementedYamrQueryServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedYamrQueryServer is a helper method to define mock.On call
func (_e *MockUnsafeYamrQueryServer_Expecter) mustEmbedUnimplementedYamrQueryServer() *MockUnsafeYamrQueryServer_mustEmbedUnimplementedYamrQueryServer_Call {
	return &MockUnsafeYamrQueryServer_mustEmbedUnimplementedYamrQueryServer_Call{Call: _e.mock.On("mustEmbedUnimplementedYamrQueryServer")}
}

func (_c *MockUnsafeYamrQueryServer_mustEmbedUnimplementedYamrQueryServer_Call) Run(run func()) *MockUnsafeYamrQueryServer_mustEmbedUnimplementedYamrQueryServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUnsafeYamrQueryServer_mustEmbedUnimplementedYamrQueryServer_Call) Return() *MockUnsafeYamrQueryServer_mustEmbedUnimplementedYamrQueryServer_Call {
	_c.Call.Return()
	return _c
}

// NewMockUnsafeYamrQueryServer creates a new instance of MockUnsafeYamrQueryServer. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUnsafeYamrQueryServer(t testing.TB) *MockUnsafeYamrQueryServer {
	mock := &MockUnsafeYamrQueryServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
