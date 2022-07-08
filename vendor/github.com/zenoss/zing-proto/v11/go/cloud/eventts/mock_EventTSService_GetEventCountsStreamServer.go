// Code generated by mockery v2.12.2. DO NOT EDIT.

package eventts

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metadata "google.golang.org/grpc/metadata"

	testing "testing"
)

// MockEventTSService_GetEventCountsStreamServer is an autogenerated mock type for the EventTSService_GetEventCountsStreamServer type
type MockEventTSService_GetEventCountsStreamServer struct {
	mock.Mock
}

type MockEventTSService_GetEventCountsStreamServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEventTSService_GetEventCountsStreamServer) EXPECT() *MockEventTSService_GetEventCountsStreamServer_Expecter {
	return &MockEventTSService_GetEventCountsStreamServer_Expecter{mock: &_m.Mock}
}

// Context provides a mock function with given fields:
func (_m *MockEventTSService_GetEventCountsStreamServer) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// MockEventTSService_GetEventCountsStreamServer_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type MockEventTSService_GetEventCountsStreamServer_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *MockEventTSService_GetEventCountsStreamServer_Expecter) Context() *MockEventTSService_GetEventCountsStreamServer_Context_Call {
	return &MockEventTSService_GetEventCountsStreamServer_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *MockEventTSService_GetEventCountsStreamServer_Context_Call) Run(run func()) *MockEventTSService_GetEventCountsStreamServer_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEventTSService_GetEventCountsStreamServer_Context_Call) Return(_a0 context.Context) *MockEventTSService_GetEventCountsStreamServer_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *MockEventTSService_GetEventCountsStreamServer) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventTSService_GetEventCountsStreamServer_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type MockEventTSService_GetEventCountsStreamServer_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//  - m interface{}
func (_e *MockEventTSService_GetEventCountsStreamServer_Expecter) RecvMsg(m interface{}) *MockEventTSService_GetEventCountsStreamServer_RecvMsg_Call {
	return &MockEventTSService_GetEventCountsStreamServer_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *MockEventTSService_GetEventCountsStreamServer_RecvMsg_Call) Run(run func(m interface{})) *MockEventTSService_GetEventCountsStreamServer_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockEventTSService_GetEventCountsStreamServer_RecvMsg_Call) Return(_a0 error) *MockEventTSService_GetEventCountsStreamServer_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

// Send provides a mock function with given fields: _a0
func (_m *MockEventTSService_GetEventCountsStreamServer) Send(_a0 *EventTSCountsResponse) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*EventTSCountsResponse) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventTSService_GetEventCountsStreamServer_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockEventTSService_GetEventCountsStreamServer_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//  - _a0 *EventTSCountsResponse
func (_e *MockEventTSService_GetEventCountsStreamServer_Expecter) Send(_a0 interface{}) *MockEventTSService_GetEventCountsStreamServer_Send_Call {
	return &MockEventTSService_GetEventCountsStreamServer_Send_Call{Call: _e.mock.On("Send", _a0)}
}

func (_c *MockEventTSService_GetEventCountsStreamServer_Send_Call) Run(run func(_a0 *EventTSCountsResponse)) *MockEventTSService_GetEventCountsStreamServer_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*EventTSCountsResponse))
	})
	return _c
}

func (_c *MockEventTSService_GetEventCountsStreamServer_Send_Call) Return(_a0 error) *MockEventTSService_GetEventCountsStreamServer_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

// SendHeader provides a mock function with given fields: _a0
func (_m *MockEventTSService_GetEventCountsStreamServer) SendHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventTSService_GetEventCountsStreamServer_SendHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendHeader'
type MockEventTSService_GetEventCountsStreamServer_SendHeader_Call struct {
	*mock.Call
}

// SendHeader is a helper method to define mock.On call
//  - _a0 metadata.MD
func (_e *MockEventTSService_GetEventCountsStreamServer_Expecter) SendHeader(_a0 interface{}) *MockEventTSService_GetEventCountsStreamServer_SendHeader_Call {
	return &MockEventTSService_GetEventCountsStreamServer_SendHeader_Call{Call: _e.mock.On("SendHeader", _a0)}
}

func (_c *MockEventTSService_GetEventCountsStreamServer_SendHeader_Call) Run(run func(_a0 metadata.MD)) *MockEventTSService_GetEventCountsStreamServer_SendHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockEventTSService_GetEventCountsStreamServer_SendHeader_Call) Return(_a0 error) *MockEventTSService_GetEventCountsStreamServer_SendHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *MockEventTSService_GetEventCountsStreamServer) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventTSService_GetEventCountsStreamServer_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type MockEventTSService_GetEventCountsStreamServer_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//  - m interface{}
func (_e *MockEventTSService_GetEventCountsStreamServer_Expecter) SendMsg(m interface{}) *MockEventTSService_GetEventCountsStreamServer_SendMsg_Call {
	return &MockEventTSService_GetEventCountsStreamServer_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *MockEventTSService_GetEventCountsStreamServer_SendMsg_Call) Run(run func(m interface{})) *MockEventTSService_GetEventCountsStreamServer_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockEventTSService_GetEventCountsStreamServer_SendMsg_Call) Return(_a0 error) *MockEventTSService_GetEventCountsStreamServer_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

// SetHeader provides a mock function with given fields: _a0
func (_m *MockEventTSService_GetEventCountsStreamServer) SetHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventTSService_GetEventCountsStreamServer_SetHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetHeader'
type MockEventTSService_GetEventCountsStreamServer_SetHeader_Call struct {
	*mock.Call
}

// SetHeader is a helper method to define mock.On call
//  - _a0 metadata.MD
func (_e *MockEventTSService_GetEventCountsStreamServer_Expecter) SetHeader(_a0 interface{}) *MockEventTSService_GetEventCountsStreamServer_SetHeader_Call {
	return &MockEventTSService_GetEventCountsStreamServer_SetHeader_Call{Call: _e.mock.On("SetHeader", _a0)}
}

func (_c *MockEventTSService_GetEventCountsStreamServer_SetHeader_Call) Run(run func(_a0 metadata.MD)) *MockEventTSService_GetEventCountsStreamServer_SetHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockEventTSService_GetEventCountsStreamServer_SetHeader_Call) Return(_a0 error) *MockEventTSService_GetEventCountsStreamServer_SetHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

// SetTrailer provides a mock function with given fields: _a0
func (_m *MockEventTSService_GetEventCountsStreamServer) SetTrailer(_a0 metadata.MD) {
	_m.Called(_a0)
}

// MockEventTSService_GetEventCountsStreamServer_SetTrailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTrailer'
type MockEventTSService_GetEventCountsStreamServer_SetTrailer_Call struct {
	*mock.Call
}

// SetTrailer is a helper method to define mock.On call
//  - _a0 metadata.MD
func (_e *MockEventTSService_GetEventCountsStreamServer_Expecter) SetTrailer(_a0 interface{}) *MockEventTSService_GetEventCountsStreamServer_SetTrailer_Call {
	return &MockEventTSService_GetEventCountsStreamServer_SetTrailer_Call{Call: _e.mock.On("SetTrailer", _a0)}
}

func (_c *MockEventTSService_GetEventCountsStreamServer_SetTrailer_Call) Run(run func(_a0 metadata.MD)) *MockEventTSService_GetEventCountsStreamServer_SetTrailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockEventTSService_GetEventCountsStreamServer_SetTrailer_Call) Return() *MockEventTSService_GetEventCountsStreamServer_SetTrailer_Call {
	_c.Call.Return()
	return _c
}

// NewMockEventTSService_GetEventCountsStreamServer creates a new instance of MockEventTSService_GetEventCountsStreamServer. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEventTSService_GetEventCountsStreamServer(t testing.TB) *MockEventTSService_GetEventCountsStreamServer {
	mock := &MockEventTSService_GetEventCountsStreamServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}