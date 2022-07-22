// Code generated by mockery v2.12.2. DO NOT EDIT.

package tenant

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// MockUnsafeTenantServiceServer is an autogenerated mock type for the UnsafeTenantServiceServer type
type MockUnsafeTenantServiceServer struct {
	mock.Mock
}

type MockUnsafeTenantServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUnsafeTenantServiceServer) EXPECT() *MockUnsafeTenantServiceServer_Expecter {
	return &MockUnsafeTenantServiceServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedTenantServiceServer provides a mock function with given fields:
func (_m *MockUnsafeTenantServiceServer) mustEmbedUnimplementedTenantServiceServer() {
	_m.Called()
}

// MockUnsafeTenantServiceServer_mustEmbedUnimplementedTenantServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedTenantServiceServer'
type MockUnsafeTenantServiceServer_mustEmbedUnimplementedTenantServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedTenantServiceServer is a helper method to define mock.On call
func (_e *MockUnsafeTenantServiceServer_Expecter) mustEmbedUnimplementedTenantServiceServer() *MockUnsafeTenantServiceServer_mustEmbedUnimplementedTenantServiceServer_Call {
	return &MockUnsafeTenantServiceServer_mustEmbedUnimplementedTenantServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedTenantServiceServer")}
}

func (_c *MockUnsafeTenantServiceServer_mustEmbedUnimplementedTenantServiceServer_Call) Run(run func()) *MockUnsafeTenantServiceServer_mustEmbedUnimplementedTenantServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUnsafeTenantServiceServer_mustEmbedUnimplementedTenantServiceServer_Call) Return() *MockUnsafeTenantServiceServer_mustEmbedUnimplementedTenantServiceServer_Call {
	_c.Call.Return()
	return _c
}

// NewMockUnsafeTenantServiceServer creates a new instance of MockUnsafeTenantServiceServer. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUnsafeTenantServiceServer(t testing.TB) *MockUnsafeTenantServiceServer {
	mock := &MockUnsafeTenantServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
