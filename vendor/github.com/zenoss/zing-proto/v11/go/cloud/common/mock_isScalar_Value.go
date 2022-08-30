// Code generated by mockery v2.13.1. DO NOT EDIT.

package common

import mock "github.com/stretchr/testify/mock"

// mockIsScalar_Value is an autogenerated mock type for the isScalar_Value type
type mockIsScalar_Value struct {
	mock.Mock
}

type mockIsScalar_Value_Expecter struct {
	mock *mock.Mock
}

func (_m *mockIsScalar_Value) EXPECT() *mockIsScalar_Value_Expecter {
	return &mockIsScalar_Value_Expecter{mock: &_m.Mock}
}

// isScalar_Value provides a mock function with given fields:
func (_m *mockIsScalar_Value) isScalar_Value() {
	_m.Called()
}

// mockIsScalar_Value_isScalar_Value_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'isScalar_Value'
type mockIsScalar_Value_isScalar_Value_Call struct {
	*mock.Call
}

// isScalar_Value is a helper method to define mock.On call
func (_e *mockIsScalar_Value_Expecter) isScalar_Value() *mockIsScalar_Value_isScalar_Value_Call {
	return &mockIsScalar_Value_isScalar_Value_Call{Call: _e.mock.On("isScalar_Value")}
}

func (_c *mockIsScalar_Value_isScalar_Value_Call) Run(run func()) *mockIsScalar_Value_isScalar_Value_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mockIsScalar_Value_isScalar_Value_Call) Return() *mockIsScalar_Value_isScalar_Value_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTnewMockIsScalar_Value interface {
	mock.TestingT
	Cleanup(func())
}

// newMockIsScalar_Value creates a new instance of mockIsScalar_Value. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockIsScalar_Value(t mockConstructorTestingTnewMockIsScalar_Value) *mockIsScalar_Value {
	mock := &mockIsScalar_Value{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
