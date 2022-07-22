// Code generated by mockery v2.12.2. DO NOT EDIT.

package tenant

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// MockUnsafeTenantInternalServiceServer is an autogenerated mock type for the UnsafeTenantInternalServiceServer type
type MockUnsafeTenantInternalServiceServer struct {
	mock.Mock
}

type MockUnsafeTenantInternalServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUnsafeTenantInternalServiceServer) EXPECT() *MockUnsafeTenantInternalServiceServer_Expecter {
	return &MockUnsafeTenantInternalServiceServer_Expecter{mock: &_m.Mock}
}

// mustEmbedUnimplementedTenantInternalServiceServer provides a mock function with given fields:
func (_m *MockUnsafeTenantInternalServiceServer) mustEmbedUnimplementedTenantInternalServiceServer() {
	_m.Called()
}

// MockUnsafeTenantInternalServiceServer_mustEmbedUnimplementedTenantInternalServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedTenantInternalServiceServer'
type MockUnsafeTenantInternalServiceServer_mustEmbedUnimplementedTenantInternalServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedTenantInternalServiceServer is a helper method to define mock.On call
func (_e *MockUnsafeTenantInternalServiceServer_Expecter) mustEmbedUnimplementedTenantInternalServiceServer() *MockUnsafeTenantInternalServiceServer_mustEmbedUnimplementedTenantInternalServiceServer_Call {
	return &MockUnsafeTenantInternalServiceServer_mustEmbedUnimplementedTenantInternalServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedTenantInternalServiceServer")}
}

func (_c *MockUnsafeTenantInternalServiceServer_mustEmbedUnimplementedTenantInternalServiceServer_Call) Run(run func()) *MockUnsafeTenantInternalServiceServer_mustEmbedUnimplementedTenantInternalServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUnsafeTenantInternalServiceServer_mustEmbedUnimplementedTenantInternalServiceServer_Call) Return() *MockUnsafeTenantInternalServiceServer_mustEmbedUnimplementedTenantInternalServiceServer_Call {
	_c.Call.Return()
	return _c
}

// NewMockUnsafeTenantInternalServiceServer creates a new instance of MockUnsafeTenantInternalServiceServer. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUnsafeTenantInternalServiceServer(t testing.TB) *MockUnsafeTenantInternalServiceServer {
	mock := &MockUnsafeTenantInternalServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}