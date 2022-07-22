// Code generated by mockery v2.12.2. DO NOT EDIT.

package event_context

import (
	context "context"
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// MockEventContextIngestServer is an autogenerated mock type for the EventContextIngestServer type
type MockEventContextIngestServer struct {
	mock.Mock
}

type MockEventContextIngestServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEventContextIngestServer) EXPECT() *MockEventContextIngestServer_Expecter {
	return &MockEventContextIngestServer_Expecter{mock: &_m.Mock}
}

// DeleteTenantData provides a mock function with given fields: _a0, _a1
func (_m *MockEventContextIngestServer) DeleteTenantData(_a0 context.Context, _a1 *DeleteTenantDataRequest) (*DeleteTenantDataResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *DeleteTenantDataResponse
	if rf, ok := ret.Get(0).(func(context.Context, *DeleteTenantDataRequest) *DeleteTenantDataResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DeleteTenantDataResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *DeleteTenantDataRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEventContextIngestServer_DeleteTenantData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteTenantData'
type MockEventContextIngestServer_DeleteTenantData_Call struct {
	*mock.Call
}

// DeleteTenantData is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *DeleteTenantDataRequest
func (_e *MockEventContextIngestServer_Expecter) DeleteTenantData(_a0 interface{}, _a1 interface{}) *MockEventContextIngestServer_DeleteTenantData_Call {
	return &MockEventContextIngestServer_DeleteTenantData_Call{Call: _e.mock.On("DeleteTenantData", _a0, _a1)}
}

func (_c *MockEventContextIngestServer_DeleteTenantData_Call) Run(run func(_a0 context.Context, _a1 *DeleteTenantDataRequest)) *MockEventContextIngestServer_DeleteTenantData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*DeleteTenantDataRequest))
	})
	return _c
}

func (_c *MockEventContextIngestServer_DeleteTenantData_Call) Return(_a0 *DeleteTenantDataResponse, _a1 error) *MockEventContextIngestServer_DeleteTenantData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// PutEvent provides a mock function with given fields: _a0, _a1
func (_m *MockEventContextIngestServer) PutEvent(_a0 context.Context, _a1 *PutEventRequest) (*PutEventResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *PutEventResponse
	if rf, ok := ret.Get(0).(func(context.Context, *PutEventRequest) *PutEventResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PutEventResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *PutEventRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEventContextIngestServer_PutEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutEvent'
type MockEventContextIngestServer_PutEvent_Call struct {
	*mock.Call
}

// PutEvent is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *PutEventRequest
func (_e *MockEventContextIngestServer_Expecter) PutEvent(_a0 interface{}, _a1 interface{}) *MockEventContextIngestServer_PutEvent_Call {
	return &MockEventContextIngestServer_PutEvent_Call{Call: _e.mock.On("PutEvent", _a0, _a1)}
}

func (_c *MockEventContextIngestServer_PutEvent_Call) Run(run func(_a0 context.Context, _a1 *PutEventRequest)) *MockEventContextIngestServer_PutEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*PutEventRequest))
	})
	return _c
}

func (_c *MockEventContextIngestServer_PutEvent_Call) Return(_a0 *PutEventResponse, _a1 error) *MockEventContextIngestServer_PutEvent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// PutEventBulk provides a mock function with given fields: _a0, _a1
func (_m *MockEventContextIngestServer) PutEventBulk(_a0 context.Context, _a1 *PutEventBulkRequest) (*PutEventBulkResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *PutEventBulkResponse
	if rf, ok := ret.Get(0).(func(context.Context, *PutEventBulkRequest) *PutEventBulkResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PutEventBulkResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *PutEventBulkRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEventContextIngestServer_PutEventBulk_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutEventBulk'
type MockEventContextIngestServer_PutEventBulk_Call struct {
	*mock.Call
}

// PutEventBulk is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *PutEventBulkRequest
func (_e *MockEventContextIngestServer_Expecter) PutEventBulk(_a0 interface{}, _a1 interface{}) *MockEventContextIngestServer_PutEventBulk_Call {
	return &MockEventContextIngestServer_PutEventBulk_Call{Call: _e.mock.On("PutEventBulk", _a0, _a1)}
}

func (_c *MockEventContextIngestServer_PutEventBulk_Call) Run(run func(_a0 context.Context, _a1 *PutEventBulkRequest)) *MockEventContextIngestServer_PutEventBulk_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*PutEventBulkRequest))
	})
	return _c
}

func (_c *MockEventContextIngestServer_PutEventBulk_Call) Return(_a0 *PutEventBulkResponse, _a1 error) *MockEventContextIngestServer_PutEventBulk_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdateEvent provides a mock function with given fields: _a0, _a1
func (_m *MockEventContextIngestServer) UpdateEvent(_a0 context.Context, _a1 *UpdateEventRequest) (*UpdateEventResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *UpdateEventResponse
	if rf, ok := ret.Get(0).(func(context.Context, *UpdateEventRequest) *UpdateEventResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UpdateEventResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *UpdateEventRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEventContextIngestServer_UpdateEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateEvent'
type MockEventContextIngestServer_UpdateEvent_Call struct {
	*mock.Call
}

// UpdateEvent is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *UpdateEventRequest
func (_e *MockEventContextIngestServer_Expecter) UpdateEvent(_a0 interface{}, _a1 interface{}) *MockEventContextIngestServer_UpdateEvent_Call {
	return &MockEventContextIngestServer_UpdateEvent_Call{Call: _e.mock.On("UpdateEvent", _a0, _a1)}
}

func (_c *MockEventContextIngestServer_UpdateEvent_Call) Run(run func(_a0 context.Context, _a1 *UpdateEventRequest)) *MockEventContextIngestServer_UpdateEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*UpdateEventRequest))
	})
	return _c
}

func (_c *MockEventContextIngestServer_UpdateEvent_Call) Return(_a0 *UpdateEventResponse, _a1 error) *MockEventContextIngestServer_UpdateEvent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// mustEmbedUnimplementedEventContextIngestServer provides a mock function with given fields:
func (_m *MockEventContextIngestServer) mustEmbedUnimplementedEventContextIngestServer() {
	_m.Called()
}

// MockEventContextIngestServer_mustEmbedUnimplementedEventContextIngestServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedEventContextIngestServer'
type MockEventContextIngestServer_mustEmbedUnimplementedEventContextIngestServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedEventContextIngestServer is a helper method to define mock.On call
func (_e *MockEventContextIngestServer_Expecter) mustEmbedUnimplementedEventContextIngestServer() *MockEventContextIngestServer_mustEmbedUnimplementedEventContextIngestServer_Call {
	return &MockEventContextIngestServer_mustEmbedUnimplementedEventContextIngestServer_Call{Call: _e.mock.On("mustEmbedUnimplementedEventContextIngestServer")}
}

func (_c *MockEventContextIngestServer_mustEmbedUnimplementedEventContextIngestServer_Call) Run(run func()) *MockEventContextIngestServer_mustEmbedUnimplementedEventContextIngestServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEventContextIngestServer_mustEmbedUnimplementedEventContextIngestServer_Call) Return() *MockEventContextIngestServer_mustEmbedUnimplementedEventContextIngestServer_Call {
	_c.Call.Return()
	return _c
}

// NewMockEventContextIngestServer creates a new instance of MockEventContextIngestServer. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEventContextIngestServer(t testing.TB) *MockEventContextIngestServer {
	mock := &MockEventContextIngestServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
