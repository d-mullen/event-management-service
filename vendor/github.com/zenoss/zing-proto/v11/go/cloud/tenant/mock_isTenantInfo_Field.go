// Code generated by mockery v2.14.0. DO NOT EDIT.

package tenant

import mock "github.com/stretchr/testify/mock"

// mockIsTenantInfo_Field is an autogenerated mock type for the isTenantInfo_Field type
type mockIsTenantInfo_Field struct {
	mock.Mock
}

type mockIsTenantInfo_Field_Expecter struct {
	mock *mock.Mock
}

func (_m *mockIsTenantInfo_Field) EXPECT() *mockIsTenantInfo_Field_Expecter {
	return &mockIsTenantInfo_Field_Expecter{mock: &_m.Mock}
}

// isTenantInfo_Field provides a mock function with given fields:
func (_m *mockIsTenantInfo_Field) isTenantInfo_Field() {
	_m.Called()
}

// mockIsTenantInfo_Field_isTenantInfo_Field_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'isTenantInfo_Field'
type mockIsTenantInfo_Field_isTenantInfo_Field_Call struct {
	*mock.Call
}

// isTenantInfo_Field is a helper method to define mock.On call
func (_e *mockIsTenantInfo_Field_Expecter) isTenantInfo_Field() *mockIsTenantInfo_Field_isTenantInfo_Field_Call {
	return &mockIsTenantInfo_Field_isTenantInfo_Field_Call{Call: _e.mock.On("isTenantInfo_Field")}
}

func (_c *mockIsTenantInfo_Field_isTenantInfo_Field_Call) Run(run func()) *mockIsTenantInfo_Field_isTenantInfo_Field_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mockIsTenantInfo_Field_isTenantInfo_Field_Call) Return() *mockIsTenantInfo_Field_isTenantInfo_Field_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTnewMockIsTenantInfo_Field interface {
	mock.TestingT
	Cleanup(func())
}

// newMockIsTenantInfo_Field creates a new instance of mockIsTenantInfo_Field. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockIsTenantInfo_Field(t mockConstructorTestingTnewMockIsTenantInfo_Field) *mockIsTenantInfo_Field {
	mock := &mockIsTenantInfo_Field{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
