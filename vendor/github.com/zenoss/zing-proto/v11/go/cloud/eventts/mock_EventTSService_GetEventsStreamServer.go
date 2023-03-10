// Code generated by mockery v2.14.0. DO NOT EDIT.

package eventts

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metadata "google.golang.org/grpc/metadata"
)

// MockEventTSService_GetEventsStreamServer is an autogenerated mock type for the EventTSService_GetEventsStreamServer type
type MockEventTSService_GetEventsStreamServer struct {
	mock.Mock
}

type MockEventTSService_GetEventsStreamServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEventTSService_GetEventsStreamServer) EXPECT() *MockEventTSService_GetEventsStreamServer_Expecter {
	return &MockEventTSService_GetEventsStreamServer_Expecter{mock: &_m.Mock}
}

// Context provides a mock function with given fields:
func (_m *MockEventTSService_GetEventsStreamServer) Context() context.Context {
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

// MockEventTSService_GetEventsStreamServer_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type MockEventTSService_GetEventsStreamServer_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *MockEventTSService_GetEventsStreamServer_Expecter) Context() *MockEventTSService_GetEventsStreamServer_Context_Call {
	return &MockEventTSService_GetEventsStreamServer_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *MockEventTSService_GetEventsStreamServer_Context_Call) Run(run func()) *MockEventTSService_GetEventsStreamServer_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEventTSService_GetEventsStreamServer_Context_Call) Return(_a0 context.Context) *MockEventTSService_GetEventsStreamServer_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *MockEventTSService_GetEventsStreamServer) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventTSService_GetEventsStreamServer_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type MockEventTSService_GetEventsStreamServer_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//  - m interface{}
func (_e *MockEventTSService_GetEventsStreamServer_Expecter) RecvMsg(m interface{}) *MockEventTSService_GetEventsStreamServer_RecvMsg_Call {
	return &MockEventTSService_GetEventsStreamServer_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *MockEventTSService_GetEventsStreamServer_RecvMsg_Call) Run(run func(m interface{})) *MockEventTSService_GetEventsStreamServer_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockEventTSService_GetEventsStreamServer_RecvMsg_Call) Return(_a0 error) *MockEventTSService_GetEventsStreamServer_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

// Send provides a mock function with given fields: _a0
func (_m *MockEventTSService_GetEventsStreamServer) Send(_a0 *EventTSResponse) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*EventTSResponse) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventTSService_GetEventsStreamServer_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockEventTSService_GetEventsStreamServer_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//  - _a0 *EventTSResponse
func (_e *MockEventTSService_GetEventsStreamServer_Expecter) Send(_a0 interface{}) *MockEventTSService_GetEventsStreamServer_Send_Call {
	return &MockEventTSService_GetEventsStreamServer_Send_Call{Call: _e.mock.On("Send", _a0)}
}

func (_c *MockEventTSService_GetEventsStreamServer_Send_Call) Run(run func(_a0 *EventTSResponse)) *MockEventTSService_GetEventsStreamServer_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*EventTSResponse))
	})
	return _c
}

func (_c *MockEventTSService_GetEventsStreamServer_Send_Call) Return(_a0 error) *MockEventTSService_GetEventsStreamServer_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

// SendHeader provides a mock function with given fields: _a0
func (_m *MockEventTSService_GetEventsStreamServer) SendHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventTSService_GetEventsStreamServer_SendHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendHeader'
type MockEventTSService_GetEventsStreamServer_SendHeader_Call struct {
	*mock.Call
}

// SendHeader is a helper method to define mock.On call
//  - _a0 metadata.MD
func (_e *MockEventTSService_GetEventsStreamServer_Expecter) SendHeader(_a0 interface{}) *MockEventTSService_GetEventsStreamServer_SendHeader_Call {
	return &MockEventTSService_GetEventsStreamServer_SendHeader_Call{Call: _e.mock.On("SendHeader", _a0)}
}

func (_c *MockEventTSService_GetEventsStreamServer_SendHeader_Call) Run(run func(_a0 metadata.MD)) *MockEventTSService_GetEventsStreamServer_SendHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockEventTSService_GetEventsStreamServer_SendHeader_Call) Return(_a0 error) *MockEventTSService_GetEventsStreamServer_SendHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *MockEventTSService_GetEventsStreamServer) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventTSService_GetEventsStreamServer_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type MockEventTSService_GetEventsStreamServer_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//  - m interface{}
func (_e *MockEventTSService_GetEventsStreamServer_Expecter) SendMsg(m interface{}) *MockEventTSService_GetEventsStreamServer_SendMsg_Call {
	return &MockEventTSService_GetEventsStreamServer_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *MockEventTSService_GetEventsStreamServer_SendMsg_Call) Run(run func(m interface{})) *MockEventTSService_GetEventsStreamServer_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockEventTSService_GetEventsStreamServer_SendMsg_Call) Return(_a0 error) *MockEventTSService_GetEventsStreamServer_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

// SetHeader provides a mock function with given fields: _a0
func (_m *MockEventTSService_GetEventsStreamServer) SetHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventTSService_GetEventsStreamServer_SetHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetHeader'
type MockEventTSService_GetEventsStreamServer_SetHeader_Call struct {
	*mock.Call
}

// SetHeader is a helper method to define mock.On call
//  - _a0 metadata.MD
func (_e *MockEventTSService_GetEventsStreamServer_Expecter) SetHeader(_a0 interface{}) *MockEventTSService_GetEventsStreamServer_SetHeader_Call {
	return &MockEventTSService_GetEventsStreamServer_SetHeader_Call{Call: _e.mock.On("SetHeader", _a0)}
}

func (_c *MockEventTSService_GetEventsStreamServer_SetHeader_Call) Run(run func(_a0 metadata.MD)) *MockEventTSService_GetEventsStreamServer_SetHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockEventTSService_GetEventsStreamServer_SetHeader_Call) Return(_a0 error) *MockEventTSService_GetEventsStreamServer_SetHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

// SetTrailer provides a mock function with given fields: _a0
func (_m *MockEventTSService_GetEventsStreamServer) SetTrailer(_a0 metadata.MD) {
	_m.Called(_a0)
}

// MockEventTSService_GetEventsStreamServer_SetTrailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTrailer'
type MockEventTSService_GetEventsStreamServer_SetTrailer_Call struct {
	*mock.Call
}

// SetTrailer is a helper method to define mock.On call
//  - _a0 metadata.MD
func (_e *MockEventTSService_GetEventsStreamServer_Expecter) SetTrailer(_a0 interface{}) *MockEventTSService_GetEventsStreamServer_SetTrailer_Call {
	return &MockEventTSService_GetEventsStreamServer_SetTrailer_Call{Call: _e.mock.On("SetTrailer", _a0)}
}

func (_c *MockEventTSService_GetEventsStreamServer_SetTrailer_Call) Run(run func(_a0 metadata.MD)) *MockEventTSService_GetEventsStreamServer_SetTrailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockEventTSService_GetEventsStreamServer_SetTrailer_Call) Return() *MockEventTSService_GetEventsStreamServer_SetTrailer_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewMockEventTSService_GetEventsStreamServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEventTSService_GetEventsStreamServer creates a new instance of MockEventTSService_GetEventsStreamServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEventTSService_GetEventsStreamServer(t mockConstructorTestingTNewMockEventTSService_GetEventsStreamServer) *MockEventTSService_GetEventsStreamServer {
	mock := &MockEventTSService_GetEventsStreamServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
