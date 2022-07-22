// Code generated by mockery v2.12.2. DO NOT EDIT.

package tenant

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// MockUnsafeTenantAdminServer is an autogenerated mock type for the UnsafeTenantAdminServer type
type MockUnsafeTenantAdminServer struct {
	mock.Mock
}

type MockUnsafeTenantAdminServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUnsafeTenantAdminServer) EXPECT() *MockUnsafeTenantAdminServer_Expecter {
	return &MockUnsafeTenantAdminServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedTenantAdminServer provides a mock function with given fields:
func (_m *MockUnsafeTenantAdminServer) mustEmbedUnimplementedTenantAdminServer() {
	_m.Called()
}

// MockUnsafeTenantAdminServer_mustEmbedUnimplementedTenantAdminServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedTenantAdminServer'
type MockUnsafeTenantAdminServer_mustEmbedUnimplementedTenantAdminServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedTenantAdminServer is a helper method to define mock.On call
func (_e *MockUnsafeTenantAdminServer_Expecter) mustEmbedUnimplementedTenantAdminServer() *MockUnsafeTenantAdminServer_mustEmbedUnimplementedTenantAdminServer_Call {
	return &MockUnsafeTenantAdminServer_mustEmbedUnimplementedTenantAdminServer_Call{Call: _e.mock.On("mustEmbedUnimplementedTenantAdminServer")}
}

func (_c *MockUnsafeTenantAdminServer_mustEmbedUnimplementedTenantAdminServer_Call) Run(run func()) *MockUnsafeTenantAdminServer_mustEmbedUnimplementedTenantAdminServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUnsafeTenantAdminServer_mustEmbedUnimplementedTenantAdminServer_Call) Return() *MockUnsafeTenantAdminServer_mustEmbedUnimplementedTenantAdminServer_Call {
	_c.Call.Return()
	return _c
}

// NewMockUnsafeTenantAdminServer creates a new instance of MockUnsafeTenantAdminServer. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUnsafeTenantAdminServer(t testing.TB) *MockUnsafeTenantAdminServer {
	mock := &MockUnsafeTenantAdminServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}