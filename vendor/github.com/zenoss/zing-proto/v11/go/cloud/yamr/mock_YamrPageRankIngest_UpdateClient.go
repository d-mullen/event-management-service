// Code generated by mockery v2.13.1. DO NOT EDIT.

package yamr

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metadata "google.golang.org/grpc/metadata"
)

// MockYamrPageRankIngest_UpdateClient is an autogenerated mock type for the YamrPageRankIngest_UpdateClient type
type MockYamrPageRankIngest_UpdateClient struct {
	mock.Mock
}

type MockYamrPageRankIngest_UpdateClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockYamrPageRankIngest_UpdateClient) EXPECT() *MockYamrPageRankIngest_UpdateClient_Expecter {
	return &MockYamrPageRankIngest_UpdateClient_Expecter{mock: &_m.Mock}
}

// CloseAndRecv provides a mock function with given fields:
func (_m *MockYamrPageRankIngest_UpdateClient) CloseAndRecv() (*YamrPageRankUpdateResponse, error) {
	ret := _m.Called()

	var r0 *YamrPageRankUpdateResponse
	if rf, ok := ret.Get(0).(func() *YamrPageRankUpdateResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*YamrPageRankUpdateResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockYamrPageRankIngest_UpdateClient_CloseAndRecv_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseAndRecv'
type MockYamrPageRankIngest_UpdateClient_CloseAndRecv_Call struct {
	*mock.Call
}

// CloseAndRecv is a helper method to define mock.On call
func (_e *MockYamrPageRankIngest_UpdateClient_Expecter) CloseAndRecv() *MockYamrPageRankIngest_UpdateClient_CloseAndRecv_Call {
	return &MockYamrPageRankIngest_UpdateClient_CloseAndRecv_Call{Call: _e.mock.On("CloseAndRecv")}
}

func (_c *MockYamrPageRankIngest_UpdateClient_CloseAndRecv_Call) Run(run func()) *MockYamrPageRankIngest_UpdateClient_CloseAndRecv_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockYamrPageRankIngest_UpdateClient_CloseAndRecv_Call) Return(_a0 *YamrPageRankUpdateResponse, _a1 error) *MockYamrPageRankIngest_UpdateClient_CloseAndRecv_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// CloseSend provides a mock function with given fields:
func (_m *MockYamrPageRankIngest_UpdateClient) CloseSend() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockYamrPageRankIngest_UpdateClient_CloseSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseSend'
type MockYamrPageRankIngest_UpdateClient_CloseSend_Call struct {
	*mock.Call
}

// CloseSend is a helper method to define mock.On call
func (_e *MockYamrPageRankIngest_UpdateClient_Expecter) CloseSend() *MockYamrPageRankIngest_UpdateClient_CloseSend_Call {
	return &MockYamrPageRankIngest_UpdateClient_CloseSend_Call{Call: _e.mock.On("CloseSend")}
}

func (_c *MockYamrPageRankIngest_UpdateClient_CloseSend_Call) Run(run func()) *MockYamrPageRankIngest_UpdateClient_CloseSend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockYamrPageRankIngest_UpdateClient_CloseSend_Call) Return(_a0 error) *MockYamrPageRankIngest_UpdateClient_CloseSend_Call {
	_c.Call.Return(_a0)
	return _c
}

// Context provides a mock function with given fields:
func (_m *MockYamrPageRankIngest_UpdateClient) Context() context.Context {
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

// MockYamrPageRankIngest_UpdateClient_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type MockYamrPageRankIngest_UpdateClient_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *MockYamrPageRankIngest_UpdateClient_Expecter) Context() *MockYamrPageRankIngest_UpdateClient_Context_Call {
	return &MockYamrPageRankIngest_UpdateClient_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *MockYamrPageRankIngest_UpdateClient_Context_Call) Run(run func()) *MockYamrPageRankIngest_UpdateClient_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockYamrPageRankIngest_UpdateClient_Context_Call) Return(_a0 context.Context) *MockYamrPageRankIngest_UpdateClient_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

// Header provides a mock function with given fields:
func (_m *MockYamrPageRankIngest_UpdateClient) Header() (metadata.MD, error) {
	ret := _m.Called()

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockYamrPageRankIngest_UpdateClient_Header_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Header'
type MockYamrPageRankIngest_UpdateClient_Header_Call struct {
	*mock.Call
}

// Header is a helper method to define mock.On call
func (_e *MockYamrPageRankIngest_UpdateClient_Expecter) Header() *MockYamrPageRankIngest_UpdateClient_Header_Call {
	return &MockYamrPageRankIngest_UpdateClient_Header_Call{Call: _e.mock.On("Header")}
}

func (_c *MockYamrPageRankIngest_UpdateClient_Header_Call) Run(run func()) *MockYamrPageRankIngest_UpdateClient_Header_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockYamrPageRankIngest_UpdateClient_Header_Call) Return(_a0 metadata.MD, _a1 error) *MockYamrPageRankIngest_UpdateClient_Header_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *MockYamrPageRankIngest_UpdateClient) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockYamrPageRankIngest_UpdateClient_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type MockYamrPageRankIngest_UpdateClient_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//  - m interface{}
func (_e *MockYamrPageRankIngest_UpdateClient_Expecter) RecvMsg(m interface{}) *MockYamrPageRankIngest_UpdateClient_RecvMsg_Call {
	return &MockYamrPageRankIngest_UpdateClient_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *MockYamrPageRankIngest_UpdateClient_RecvMsg_Call) Run(run func(m interface{})) *MockYamrPageRankIngest_UpdateClient_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockYamrPageRankIngest_UpdateClient_RecvMsg_Call) Return(_a0 error) *MockYamrPageRankIngest_UpdateClient_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

// Send provides a mock function with given fields: _a0
func (_m *MockYamrPageRankIngest_UpdateClient) Send(_a0 *YamrPageRankUpdateRequest) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*YamrPageRankUpdateRequest) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockYamrPageRankIngest_UpdateClient_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockYamrPageRankIngest_UpdateClient_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//  - _a0 *YamrPageRankUpdateRequest
func (_e *MockYamrPageRankIngest_UpdateClient_Expecter) Send(_a0 interface{}) *MockYamrPageRankIngest_UpdateClient_Send_Call {
	return &MockYamrPageRankIngest_UpdateClient_Send_Call{Call: _e.mock.On("Send", _a0)}
}

func (_c *MockYamrPageRankIngest_UpdateClient_Send_Call) Run(run func(_a0 *YamrPageRankUpdateRequest)) *MockYamrPageRankIngest_UpdateClient_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*YamrPageRankUpdateRequest))
	})
	return _c
}

func (_c *MockYamrPageRankIngest_UpdateClient_Send_Call) Return(_a0 error) *MockYamrPageRankIngest_UpdateClient_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *MockYamrPageRankIngest_UpdateClient) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockYamrPageRankIngest_UpdateClient_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type MockYamrPageRankIngest_UpdateClient_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//  - m interface{}
func (_e *MockYamrPageRankIngest_UpdateClient_Expecter) SendMsg(m interface{}) *MockYamrPageRankIngest_UpdateClient_SendMsg_Call {
	return &MockYamrPageRankIngest_UpdateClient_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *MockYamrPageRankIngest_UpdateClient_SendMsg_Call) Run(run func(m interface{})) *MockYamrPageRankIngest_UpdateClient_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockYamrPageRankIngest_UpdateClient_SendMsg_Call) Return(_a0 error) *MockYamrPageRankIngest_UpdateClient_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

// Trailer provides a mock function with given fields:
func (_m *MockYamrPageRankIngest_UpdateClient) Trailer() metadata.MD {
	ret := _m.Called()

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	return r0
}

// MockYamrPageRankIngest_UpdateClient_Trailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trailer'
type MockYamrPageRankIngest_UpdateClient_Trailer_Call struct {
	*mock.Call
}

// Trailer is a helper method to define mock.On call
func (_e *MockYamrPageRankIngest_UpdateClient_Expecter) Trailer() *MockYamrPageRankIngest_UpdateClient_Trailer_Call {
	return &MockYamrPageRankIngest_UpdateClient_Trailer_Call{Call: _e.mock.On("Trailer")}
}

func (_c *MockYamrPageRankIngest_UpdateClient_Trailer_Call) Run(run func()) *MockYamrPageRankIngest_UpdateClient_Trailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockYamrPageRankIngest_UpdateClient_Trailer_Call) Return(_a0 metadata.MD) *MockYamrPageRankIngest_UpdateClient_Trailer_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewMockYamrPageRankIngest_UpdateClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockYamrPageRankIngest_UpdateClient creates a new instance of MockYamrPageRankIngest_UpdateClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockYamrPageRankIngest_UpdateClient(t mockConstructorTestingTNewMockYamrPageRankIngest_UpdateClient) *MockYamrPageRankIngest_UpdateClient {
	mock := &MockYamrPageRankIngest_UpdateClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
