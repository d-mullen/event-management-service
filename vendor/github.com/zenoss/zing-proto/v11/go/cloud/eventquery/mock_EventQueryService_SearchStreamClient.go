// Code generated by mockery v2.12.2. DO NOT EDIT.

package eventquery

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	metadata "google.golang.org/grpc/metadata"

	testing "testing"
)

// MockEventQueryService_SearchStreamClient is an autogenerated mock type for the EventQueryService_SearchStreamClient type
type MockEventQueryService_SearchStreamClient struct {
	mock.Mock
}

type MockEventQueryService_SearchStreamClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEventQueryService_SearchStreamClient) EXPECT() *MockEventQueryService_SearchStreamClient_Expecter {
	return &MockEventQueryService_SearchStreamClient_Expecter{mock: &_m.Mock}
}

// CloseSend provides a mock function with given fields:
func (_m *MockEventQueryService_SearchStreamClient) CloseSend() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventQueryService_SearchStreamClient_CloseSend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloseSend'
type MockEventQueryService_SearchStreamClient_CloseSend_Call struct {
	*mock.Call
}

// CloseSend is a helper method to define mock.On call
func (_e *MockEventQueryService_SearchStreamClient_Expecter) CloseSend() *MockEventQueryService_SearchStreamClient_CloseSend_Call {
	return &MockEventQueryService_SearchStreamClient_CloseSend_Call{Call: _e.mock.On("CloseSend")}
}

func (_c *MockEventQueryService_SearchStreamClient_CloseSend_Call) Run(run func()) *MockEventQueryService_SearchStreamClient_CloseSend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEventQueryService_SearchStreamClient_CloseSend_Call) Return(_a0 error) *MockEventQueryService_SearchStreamClient_CloseSend_Call {
	_c.Call.Return(_a0)
	return _c
}

// Context provides a mock function with given fields:
func (_m *MockEventQueryService_SearchStreamClient) Context() context.Context {
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

// MockEventQueryService_SearchStreamClient_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type MockEventQueryService_SearchStreamClient_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *MockEventQueryService_SearchStreamClient_Expecter) Context() *MockEventQueryService_SearchStreamClient_Context_Call {
	return &MockEventQueryService_SearchStreamClient_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *MockEventQueryService_SearchStreamClient_Context_Call) Run(run func()) *MockEventQueryService_SearchStreamClient_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEventQueryService_SearchStreamClient_Context_Call) Return(_a0 context.Context) *MockEventQueryService_SearchStreamClient_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

// Header provides a mock function with given fields:
func (_m *MockEventQueryService_SearchStreamClient) Header() (metadata.MD, error) {
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

// MockEventQueryService_SearchStreamClient_Header_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Header'
type MockEventQueryService_SearchStreamClient_Header_Call struct {
	*mock.Call
}

// Header is a helper method to define mock.On call
func (_e *MockEventQueryService_SearchStreamClient_Expecter) Header() *MockEventQueryService_SearchStreamClient_Header_Call {
	return &MockEventQueryService_SearchStreamClient_Header_Call{Call: _e.mock.On("Header")}
}

func (_c *MockEventQueryService_SearchStreamClient_Header_Call) Run(run func()) *MockEventQueryService_SearchStreamClient_Header_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEventQueryService_SearchStreamClient_Header_Call) Return(_a0 metadata.MD, _a1 error) *MockEventQueryService_SearchStreamClient_Header_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Recv provides a mock function with given fields:
func (_m *MockEventQueryService_SearchStreamClient) Recv() (*SearchStreamResponse, error) {
	ret := _m.Called()

	var r0 *SearchStreamResponse
	if rf, ok := ret.Get(0).(func() *SearchStreamResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*SearchStreamResponse)
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

// MockEventQueryService_SearchStreamClient_Recv_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Recv'
type MockEventQueryService_SearchStreamClient_Recv_Call struct {
	*mock.Call
}

// Recv is a helper method to define mock.On call
func (_e *MockEventQueryService_SearchStreamClient_Expecter) Recv() *MockEventQueryService_SearchStreamClient_Recv_Call {
	return &MockEventQueryService_SearchStreamClient_Recv_Call{Call: _e.mock.On("Recv")}
}

func (_c *MockEventQueryService_SearchStreamClient_Recv_Call) Run(run func()) *MockEventQueryService_SearchStreamClient_Recv_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEventQueryService_SearchStreamClient_Recv_Call) Return(_a0 *SearchStreamResponse, _a1 error) *MockEventQueryService_SearchStreamClient_Recv_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// RecvMsg provides a mock function with given fields: m
func (_m *MockEventQueryService_SearchStreamClient) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventQueryService_SearchStreamClient_RecvMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecvMsg'
type MockEventQueryService_SearchStreamClient_RecvMsg_Call struct {
	*mock.Call
}

// RecvMsg is a helper method to define mock.On call
//  - m interface{}
func (_e *MockEventQueryService_SearchStreamClient_Expecter) RecvMsg(m interface{}) *MockEventQueryService_SearchStreamClient_RecvMsg_Call {
	return &MockEventQueryService_SearchStreamClient_RecvMsg_Call{Call: _e.mock.On("RecvMsg", m)}
}

func (_c *MockEventQueryService_SearchStreamClient_RecvMsg_Call) Run(run func(m interface{})) *MockEventQueryService_SearchStreamClient_RecvMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockEventQueryService_SearchStreamClient_RecvMsg_Call) Return(_a0 error) *MockEventQueryService_SearchStreamClient_RecvMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

// SendMsg provides a mock function with given fields: m
func (_m *MockEventQueryService_SearchStreamClient) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventQueryService_SearchStreamClient_SendMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMsg'
type MockEventQueryService_SearchStreamClient_SendMsg_Call struct {
	*mock.Call
}

// SendMsg is a helper method to define mock.On call
//  - m interface{}
func (_e *MockEventQueryService_SearchStreamClient_Expecter) SendMsg(m interface{}) *MockEventQueryService_SearchStreamClient_SendMsg_Call {
	return &MockEventQueryService_SearchStreamClient_SendMsg_Call{Call: _e.mock.On("SendMsg", m)}
}

func (_c *MockEventQueryService_SearchStreamClient_SendMsg_Call) Run(run func(m interface{})) *MockEventQueryService_SearchStreamClient_SendMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockEventQueryService_SearchStreamClient_SendMsg_Call) Return(_a0 error) *MockEventQueryService_SearchStreamClient_SendMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

// Trailer provides a mock function with given fields:
func (_m *MockEventQueryService_SearchStreamClient) Trailer() metadata.MD {
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

// MockEventQueryService_SearchStreamClient_Trailer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trailer'
type MockEventQueryService_SearchStreamClient_Trailer_Call struct {
	*mock.Call
}

// Trailer is a helper method to define mock.On call
func (_e *MockEventQueryService_SearchStreamClient_Expecter) Trailer() *MockEventQueryService_SearchStreamClient_Trailer_Call {
	return &MockEventQueryService_SearchStreamClient_Trailer_Call{Call: _e.mock.On("Trailer")}
}

func (_c *MockEventQueryService_SearchStreamClient_Trailer_Call) Run(run func()) *MockEventQueryService_SearchStreamClient_Trailer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEventQueryService_SearchStreamClient_Trailer_Call) Return(_a0 metadata.MD) *MockEventQueryService_SearchStreamClient_Trailer_Call {
	_c.Call.Return(_a0)
	return _c
}

// NewMockEventQueryService_SearchStreamClient creates a new instance of MockEventQueryService_SearchStreamClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEventQueryService_SearchStreamClient(t testing.TB) *MockEventQueryService_SearchStreamClient {
	mock := &MockEventQueryService_SearchStreamClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
