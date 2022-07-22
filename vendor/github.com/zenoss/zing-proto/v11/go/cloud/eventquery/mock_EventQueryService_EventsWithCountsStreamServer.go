// Code generated by mockery v2.12.2. DO NOT EDIT.

package eventquery

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metadata "google.golang.org/grpc/metadata"

	testing "testing"
)

// MockEventQueryService_EventsWithCountsStreamServer is an autogenerated mock type for the EventQueryService_EventsWithCountsStreamServer type
type MockEventQueryService_EventsWithCountsStreamServer struct {
	mock.Mock
}

type MockEventQueryService_EventsWithCountsStreamServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEventQueryService_EventsWithCountsStreamServer) EXPECT() *MockEventQueryService_EventsWithCountsStreamServer_Expecter {
	return &MockEventQueryService_EventsWithCountsStreamServer_Expecter{mock: &_m.Mock}
}

// Context provides a mock function with given fields:
func (_m *MockEventQueryService_EventsWithCountsStreamServer) Context() context.Context {
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

// MockEventQueryService_EventsWithCountsStreamServer_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type MockEventQueryService_EventsWithCountsStreamServer_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *MockEventQueryService_EventsWithCountsStreamServer_Expecter) Context() *MockEventQueryService_EventsWithCountsStreamServer_Context_Call {
	return &MockEventQueryService_EventsWithCountsStreamServer_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *MockEventQueryService_EventsWithCountsStreamServer_Context_Call) Run(run func()) *MockEventQueryService_EventsWithCountsStreamServer_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEventQueryService_EventsWithCountsStreamServer_Context_Call) Return(_a0 context.Context) *MockEventQueryService_EventsWithCountsStreamServer_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *MockEventQueryService_EventsWithCountsStreamServer) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventQueryService_EventsWithCountsStreamServer_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type MockEventQueryService_EventsWithCountsStreamServer_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//  - m interface{}
func (_e *MockEventQueryService_EventsWithCountsStreamServer_Expecter) RecvMsg(m interface{}) *MockEventQueryService_EventsWithCountsStreamServer_RecvMsg_Call {
	return &MockEventQueryService_EventsWithCountsStreamServer_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *MockEventQueryService_EventsWithCountsStreamServer_RecvMsg_Call) Run(run func(m interface{})) *MockEventQueryService_EventsWithCountsStreamServer_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockEventQueryService_EventsWithCountsStreamServer_RecvMsg_Call) Return(_a0 error) *MockEventQueryService_EventsWithCountsStreamServer_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

// Send provides a mock function with given fields: _a0
func (_m *MockEventQueryService_EventsWithCountsStreamServer) Send(_a0 *EventsWithCountsStreamResponse) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*EventsWithCountsStreamResponse) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventQueryService_EventsWithCountsStreamServer_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockEventQueryService_EventsWithCountsStreamServer_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//  - _a0 *EventsWithCountsStreamResponse
func (_e *MockEventQueryService_EventsWithCountsStreamServer_Expecter) Send(_a0 interface{}) *MockEventQueryService_EventsWithCountsStreamServer_Send_Call {
	return &MockEventQueryService_EventsWithCountsStreamServer_Send_Call{Call: _e.mock.On("Send", _a0)}
}

func (_c *MockEventQueryService_EventsWithCountsStreamServer_Send_Call) Run(run func(_a0 *EventsWithCountsStreamResponse)) *MockEventQueryService_EventsWithCountsStreamServer_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*EventsWithCountsStreamResponse))
	})
	return _c
}

func (_c *MockEventQueryService_EventsWithCountsStreamServer_Send_Call) Return(_a0 error) *MockEventQueryService_EventsWithCountsStreamServer_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

// SendHeader provides a mock function with given fields: _a0
func (_m *MockEventQueryService_EventsWithCountsStreamServer) SendHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventQueryService_EventsWithCountsStreamServer_SendHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendHeader'
type MockEventQueryService_EventsWithCountsStreamServer_SendHeader_Call struct {
	*mock.Call
}

// SendHeader is a helper method to define mock.On call
//  - _a0 metadata.MD
func (_e *MockEventQueryService_EventsWithCountsStreamServer_Expecter) SendHeader(_a0 interface{}) *MockEventQueryService_EventsWithCountsStreamServer_SendHeader_Call {
	return &MockEventQueryService_EventsWithCountsStreamServer_SendHeader_Call{Call: _e.mock.On("SendHeader", _a0)}
}

func (_c *MockEventQueryService_EventsWithCountsStreamServer_SendHeader_Call) Run(run func(_a0 metadata.MD)) *MockEventQueryService_EventsWithCountsStreamServer_SendHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockEventQueryService_EventsWithCountsStreamServer_SendHeader_Call) Return(_a0 error) *MockEventQueryService_EventsWithCountsStreamServer_SendHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *MockEventQueryService_EventsWithCountsStreamServer) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventQueryService_EventsWithCountsStreamServer_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type MockEventQueryService_EventsWithCountsStreamServer_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//  - m interface{}
func (_e *MockEventQueryService_EventsWithCountsStreamServer_Expecter) SendMsg(m interface{}) *MockEventQueryService_EventsWithCountsStreamServer_SendMsg_Call {
	return &MockEventQueryService_EventsWithCountsStreamServer_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *MockEventQueryService_EventsWithCountsStreamServer_SendMsg_Call) Run(run func(m interface{})) *MockEventQueryService_EventsWithCountsStreamServer_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockEventQueryService_EventsWithCountsStreamServer_SendMsg_Call) Return(_a0 error) *MockEventQueryService_EventsWithCountsStreamServer_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

// SetHeader provides a mock function with given fields: _a0
func (_m *MockEventQueryService_EventsWithCountsStreamServer) SetHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventQueryService_EventsWithCountsStreamServer_SetHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetHeader'
type MockEventQueryService_EventsWithCountsStreamServer_SetHeader_Call struct {
	*mock.Call
}

// SetHeader is a helper method to define mock.On call
//  - _a0 metadata.MD
func (_e *MockEventQueryService_EventsWithCountsStreamServer_Expecter) SetHeader(_a0 interface{}) *MockEventQueryService_EventsWithCountsStreamServer_SetHeader_Call {
	return &MockEventQueryService_EventsWithCountsStreamServer_SetHeader_Call{Call: _e.mock.On("SetHeader", _a0)}
}

func (_c *MockEventQueryService_EventsWithCountsStreamServer_SetHeader_Call) Run(run func(_a0 metadata.MD)) *MockEventQueryService_EventsWithCountsStreamServer_SetHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockEventQueryService_EventsWithCountsStreamServer_SetHeader_Call) Return(_a0 error) *MockEventQueryService_EventsWithCountsStreamServer_SetHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

// SetTrailer provides a mock function with given fields: _a0
func (_m *MockEventQueryService_EventsWithCountsStreamServer) SetTrailer(_a0 metadata.MD) {
	_m.Called(_a0)
}

// MockEventQueryService_EventsWithCountsStreamServer_SetTrailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTrailer'
type MockEventQueryService_EventsWithCountsStreamServer_SetTrailer_Call struct {
	*mock.Call
}

// SetTrailer is a helper method to define mock.On call
//  - _a0 metadata.MD
func (_e *MockEventQueryService_EventsWithCountsStreamServer_Expecter) SetTrailer(_a0 interface{}) *MockEventQueryService_EventsWithCountsStreamServer_SetTrailer_Call {
	return &MockEventQueryService_EventsWithCountsStreamServer_SetTrailer_Call{Call: _e.mock.On("SetTrailer", _a0)}
}

func (_c *MockEventQueryService_EventsWithCountsStreamServer_SetTrailer_Call) Run(run func(_a0 metadata.MD)) *MockEventQueryService_EventsWithCountsStreamServer_SetTrailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockEventQueryService_EventsWithCountsStreamServer_SetTrailer_Call) Return() *MockEventQueryService_EventsWithCountsStreamServer_SetTrailer_Call {
	_c.Call.Return()
	return _c
}

// NewMockEventQueryService_EventsWithCountsStreamServer creates a new instance of MockEventQueryService_EventsWithCountsStreamServer. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEventQueryService_EventsWithCountsStreamServer(t testing.TB) *MockEventQueryService_EventsWithCountsStreamServer {
	mock := &MockEventQueryService_EventsWithCountsStreamServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
