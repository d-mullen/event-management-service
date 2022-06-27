// Code generated by mockery v2.12.2. DO NOT EDIT.

package yamr

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metadata "google.golang.org/grpc/metadata"

	testing "testing"
)

// MockYamrExecutor_ExecuteServer is an autogenerated mock type for the YamrExecutor_ExecuteServer type
type MockYamrExecutor_ExecuteServer struct {
	mock.Mock
}

type MockYamrExecutor_ExecuteServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockYamrExecutor_ExecuteServer) EXPECT() *MockYamrExecutor_ExecuteServer_Expecter {
	return &MockYamrExecutor_ExecuteServer_Expecter{mock: &_m.Mock}
}

// Context provides a mock function with given fields:
func (_m *MockYamrExecutor_ExecuteServer) Context() context.Context {
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

// MockYamrExecutor_ExecuteServer_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type MockYamrExecutor_ExecuteServer_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *MockYamrExecutor_ExecuteServer_Expecter) Context() *MockYamrExecutor_ExecuteServer_Context_Call {
	return &MockYamrExecutor_ExecuteServer_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *MockYamrExecutor_ExecuteServer_Context_Call) Run(run func()) *MockYamrExecutor_ExecuteServer_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockYamrExecutor_ExecuteServer_Context_Call) Return(_a0 context.Context) *MockYamrExecutor_ExecuteServer_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *MockYamrExecutor_ExecuteServer) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockYamrExecutor_ExecuteServer_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type MockYamrExecutor_ExecuteServer_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//  - m interface{}
func (_e *MockYamrExecutor_ExecuteServer_Expecter) RecvMsg(m interface{}) *MockYamrExecutor_ExecuteServer_RecvMsg_Call {
	return &MockYamrExecutor_ExecuteServer_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *MockYamrExecutor_ExecuteServer_RecvMsg_Call) Run(run func(m interface{})) *MockYamrExecutor_ExecuteServer_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockYamrExecutor_ExecuteServer_RecvMsg_Call) Return(_a0 error) *MockYamrExecutor_ExecuteServer_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

// Send provides a mock function with given fields: _a0
func (_m *MockYamrExecutor_ExecuteServer) Send(_a0 *YamrExecuteResponse) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*YamrExecuteResponse) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockYamrExecutor_ExecuteServer_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockYamrExecutor_ExecuteServer_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//  - _a0 *YamrExecuteResponse
func (_e *MockYamrExecutor_ExecuteServer_Expecter) Send(_a0 interface{}) *MockYamrExecutor_ExecuteServer_Send_Call {
	return &MockYamrExecutor_ExecuteServer_Send_Call{Call: _e.mock.On("Send", _a0)}
}

func (_c *MockYamrExecutor_ExecuteServer_Send_Call) Run(run func(_a0 *YamrExecuteResponse)) *MockYamrExecutor_ExecuteServer_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*YamrExecuteResponse))
	})
	return _c
}

func (_c *MockYamrExecutor_ExecuteServer_Send_Call) Return(_a0 error) *MockYamrExecutor_ExecuteServer_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

// SendHeader provides a mock function with given fields: _a0
func (_m *MockYamrExecutor_ExecuteServer) SendHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockYamrExecutor_ExecuteServer_SendHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendHeader'
type MockYamrExecutor_ExecuteServer_SendHeader_Call struct {
	*mock.Call
}

// SendHeader is a helper method to define mock.On call
//  - _a0 metadata.MD
func (_e *MockYamrExecutor_ExecuteServer_Expecter) SendHeader(_a0 interface{}) *MockYamrExecutor_ExecuteServer_SendHeader_Call {
	return &MockYamrExecutor_ExecuteServer_SendHeader_Call{Call: _e.mock.On("SendHeader", _a0)}
}

func (_c *MockYamrExecutor_ExecuteServer_SendHeader_Call) Run(run func(_a0 metadata.MD)) *MockYamrExecutor_ExecuteServer_SendHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockYamrExecutor_ExecuteServer_SendHeader_Call) Return(_a0 error) *MockYamrExecutor_ExecuteServer_SendHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *MockYamrExecutor_ExecuteServer) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockYamrExecutor_ExecuteServer_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type MockYamrExecutor_ExecuteServer_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//  - m interface{}
func (_e *MockYamrExecutor_ExecuteServer_Expecter) SendMsg(m interface{}) *MockYamrExecutor_ExecuteServer_SendMsg_Call {
	return &MockYamrExecutor_ExecuteServer_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *MockYamrExecutor_ExecuteServer_SendMsg_Call) Run(run func(m interface{})) *MockYamrExecutor_ExecuteServer_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockYamrExecutor_ExecuteServer_SendMsg_Call) Return(_a0 error) *MockYamrExecutor_ExecuteServer_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

// SetHeader provides a mock function with given fields: _a0
func (_m *MockYamrExecutor_ExecuteServer) SetHeader(_a0 metadata.MD) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(metadata.MD) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockYamrExecutor_ExecuteServer_SetHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetHeader'
type MockYamrExecutor_ExecuteServer_SetHeader_Call struct {
	*mock.Call
}

// SetHeader is a helper method to define mock.On call
//  - _a0 metadata.MD
func (_e *MockYamrExecutor_ExecuteServer_Expecter) SetHeader(_a0 interface{}) *MockYamrExecutor_ExecuteServer_SetHeader_Call {
	return &MockYamrExecutor_ExecuteServer_SetHeader_Call{Call: _e.mock.On("SetHeader", _a0)}
}

func (_c *MockYamrExecutor_ExecuteServer_SetHeader_Call) Run(run func(_a0 metadata.MD)) *MockYamrExecutor_ExecuteServer_SetHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockYamrExecutor_ExecuteServer_SetHeader_Call) Return(_a0 error) *MockYamrExecutor_ExecuteServer_SetHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

// SetTrailer provides a mock function with given fields: _a0
func (_m *MockYamrExecutor_ExecuteServer) SetTrailer(_a0 metadata.MD) {
	_m.Called(_a0)
}

// MockYamrExecutor_ExecuteServer_SetTrailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTrailer'
type MockYamrExecutor_ExecuteServer_SetTrailer_Call struct {
	*mock.Call
}

// SetTrailer is a helper method to define mock.On call
//  - _a0 metadata.MD
func (_e *MockYamrExecutor_ExecuteServer_Expecter) SetTrailer(_a0 interface{}) *MockYamrExecutor_ExecuteServer_SetTrailer_Call {
	return &MockYamrExecutor_ExecuteServer_SetTrailer_Call{Call: _e.mock.On("SetTrailer", _a0)}
}

func (_c *MockYamrExecutor_ExecuteServer_SetTrailer_Call) Run(run func(_a0 metadata.MD)) *MockYamrExecutor_ExecuteServer_SetTrailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(metadata.MD))
	})
	return _c
}

func (_c *MockYamrExecutor_ExecuteServer_SetTrailer_Call) Return() *MockYamrExecutor_ExecuteServer_SetTrailer_Call {
	_c.Call.Return()
	return _c
}

// NewMockYamrExecutor_ExecuteServer creates a new instance of MockYamrExecutor_ExecuteServer. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockYamrExecutor_ExecuteServer(t testing.TB) *MockYamrExecutor_ExecuteServer {
	mock := &MockYamrExecutor_ExecuteServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
