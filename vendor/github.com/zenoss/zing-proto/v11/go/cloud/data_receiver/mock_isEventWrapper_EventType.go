// Code generated by mockery v2.13.1. DO NOT EDIT.

package data_receiver

import mock "github.com/stretchr/testify/mock"

// mockIsEventWrapper_EventType is an autogenerated mock type for the isEventWrapper_EventType type
type mockIsEventWrapper_EventType struct {
	mock.Mock
}

type mockIsEventWrapper_EventType_Expecter struct {
	mock *mock.Mock
}

func (_m *mockIsEventWrapper_EventType) EXPECT() *mockIsEventWrapper_EventType_Expecter {
	return &mockIsEventWrapper_EventType_Expecter{mock: &_m.Mock}
}

// isEventWrapper_EventType provides a mock function with given fields:
func (_m *mockIsEventWrapper_EventType) isEventWrapper_EventType() {
	_m.Called()
}

// mockIsEventWrapper_EventType_isEventWrapper_EventType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'isEventWrapper_EventType'
type mockIsEventWrapper_EventType_isEventWrapper_EventType_Call struct {
	*mock.Call
}

// isEventWrapper_EventType is a helper method to define mock.On call
func (_e *mockIsEventWrapper_EventType_Expecter) isEventWrapper_EventType() *mockIsEventWrapper_EventType_isEventWrapper_EventType_Call {
	return &mockIsEventWrapper_EventType_isEventWrapper_EventType_Call{Call: _e.mock.On("isEventWrapper_EventType")}
}

func (_c *mockIsEventWrapper_EventType_isEventWrapper_EventType_Call) Run(run func()) *mockIsEventWrapper_EventType_isEventWrapper_EventType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mockIsEventWrapper_EventType_isEventWrapper_EventType_Call) Return() *mockIsEventWrapper_EventType_isEventWrapper_EventType_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTnewMockIsEventWrapper_EventType interface {
	mock.TestingT
	Cleanup(func())
}

// newMockIsEventWrapper_EventType creates a new instance of mockIsEventWrapper_EventType. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockIsEventWrapper_EventType(t mockConstructorTestingTnewMockIsEventWrapper_EventType) *mockIsEventWrapper_EventType {
	mock := &mockIsEventWrapper_EventType{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
