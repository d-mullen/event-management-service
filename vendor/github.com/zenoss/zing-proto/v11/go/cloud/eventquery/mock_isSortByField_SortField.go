// Code generated by mockery v2.13.1. DO NOT EDIT.

package eventquery

import mock "github.com/stretchr/testify/mock"

// mockIsSortByField_SortField is an autogenerated mock type for the isSortByField_SortField type
type mockIsSortByField_SortField struct {
	mock.Mock
}

type mockIsSortByField_SortField_Expecter struct {
	mock *mock.Mock
}

func (_m *mockIsSortByField_SortField) EXPECT() *mockIsSortByField_SortField_Expecter {
	return &mockIsSortByField_SortField_Expecter{mock: &_m.Mock}
}

// isSortByField_SortField provides a mock function with given fields:
func (_m *mockIsSortByField_SortField) isSortByField_SortField() {
	_m.Called()
}

// mockIsSortByField_SortField_isSortByField_SortField_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'isSortByField_SortField'
type mockIsSortByField_SortField_isSortByField_SortField_Call struct {
	*mock.Call
}

// isSortByField_SortField is a helper method to define mock.On call
func (_e *mockIsSortByField_SortField_Expecter) isSortByField_SortField() *mockIsSortByField_SortField_isSortByField_SortField_Call {
	return &mockIsSortByField_SortField_isSortByField_SortField_Call{Call: _e.mock.On("isSortByField_SortField")}
}

func (_c *mockIsSortByField_SortField_isSortByField_SortField_Call) Run(run func()) *mockIsSortByField_SortField_isSortByField_SortField_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mockIsSortByField_SortField_isSortByField_SortField_Call) Return() *mockIsSortByField_SortField_isSortByField_SortField_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTnewMockIsSortByField_SortField interface {
	mock.TestingT
	Cleanup(func())
}

// newMockIsSortByField_SortField creates a new instance of mockIsSortByField_SortField. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockIsSortByField_SortField(t mockConstructorTestingTnewMockIsSortByField_SortField) *mockIsSortByField_SortField {
	mock := &mockIsSortByField_SortField{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
