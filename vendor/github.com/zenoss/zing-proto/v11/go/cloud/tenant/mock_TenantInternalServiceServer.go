// Code generated by mockery v2.12.2. DO NOT EDIT.

package tenant

import (
	context "context"
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// MockTenantInternalServiceServer is an autogenerated mock type for the TenantInternalServiceServer type
type MockTenantInternalServiceServer struct {
	mock.Mock
}

type MockTenantInternalServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTenantInternalServiceServer) EXPECT() *MockTenantInternalServiceServer_Expecter {
	return &MockTenantInternalServiceServer_Expecter{mock: &_m.Mock}
}

// GetAllTenants provides a mock function with given fields: _a0, _a1
func (_m *MockTenantInternalServiceServer) GetAllTenants(_a0 context.Context, _a1 *GetAllTenantsRequest) (*GetAllTenantsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *GetAllTenantsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetAllTenantsRequest) *GetAllTenantsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetAllTenantsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetAllTenantsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantInternalServiceServer_GetAllTenants_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTenants'
type MockTenantInternalServiceServer_GetAllTenants_Call struct {
	*mock.Call
}

// GetAllTenants is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *GetAllTenantsRequest
func (_e *MockTenantInternalServiceServer_Expecter) GetAllTenants(_a0 interface{}, _a1 interface{}) *MockTenantInternalServiceServer_GetAllTenants_Call {
	return &MockTenantInternalServiceServer_GetAllTenants_Call{Call: _e.mock.On("GetAllTenants", _a0, _a1)}
}

func (_c *MockTenantInternalServiceServer_GetAllTenants_Call) Run(run func(_a0 context.Context, _a1 *GetAllTenantsRequest)) *MockTenantInternalServiceServer_GetAllTenants_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetAllTenantsRequest))
	})
	return _c
}

func (_c *MockTenantInternalServiceServer_GetAllTenants_Call) Return(_a0 *GetAllTenantsResponse, _a1 error) *MockTenantInternalServiceServer_GetAllTenants_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTenantDataId provides a mock function with given fields: _a0, _a1
func (_m *MockTenantInternalServiceServer) GetTenantDataId(_a0 context.Context, _a1 *GetTenantDataIdRequest) (*GetTenantDataIdResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *GetTenantDataIdResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetTenantDataIdRequest) *GetTenantDataIdResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetTenantDataIdResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetTenantDataIdRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantInternalServiceServer_GetTenantDataId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTenantDataId'
type MockTenantInternalServiceServer_GetTenantDataId_Call struct {
	*mock.Call
}

// GetTenantDataId is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *GetTenantDataIdRequest
func (_e *MockTenantInternalServiceServer_Expecter) GetTenantDataId(_a0 interface{}, _a1 interface{}) *MockTenantInternalServiceServer_GetTenantDataId_Call {
	return &MockTenantInternalServiceServer_GetTenantDataId_Call{Call: _e.mock.On("GetTenantDataId", _a0, _a1)}
}

func (_c *MockTenantInternalServiceServer_GetTenantDataId_Call) Run(run func(_a0 context.Context, _a1 *GetTenantDataIdRequest)) *MockTenantInternalServiceServer_GetTenantDataId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetTenantDataIdRequest))
	})
	return _c
}

func (_c *MockTenantInternalServiceServer_GetTenantDataId_Call) Return(_a0 *GetTenantDataIdResponse, _a1 error) *MockTenantInternalServiceServer_GetTenantDataId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTenantDataIds provides a mock function with given fields: _a0, _a1
func (_m *MockTenantInternalServiceServer) GetTenantDataIds(_a0 context.Context, _a1 *GetTenantDataIdsRequest) (*GetTenantDataIdsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *GetTenantDataIdsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetTenantDataIdsRequest) *GetTenantDataIdsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetTenantDataIdsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetTenantDataIdsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantInternalServiceServer_GetTenantDataIds_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTenantDataIds'
type MockTenantInternalServiceServer_GetTenantDataIds_Call struct {
	*mock.Call
}

// GetTenantDataIds is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *GetTenantDataIdsRequest
func (_e *MockTenantInternalServiceServer_Expecter) GetTenantDataIds(_a0 interface{}, _a1 interface{}) *MockTenantInternalServiceServer_GetTenantDataIds_Call {
	return &MockTenantInternalServiceServer_GetTenantDataIds_Call{Call: _e.mock.On("GetTenantDataIds", _a0, _a1)}
}

func (_c *MockTenantInternalServiceServer_GetTenantDataIds_Call) Run(run func(_a0 context.Context, _a1 *GetTenantDataIdsRequest)) *MockTenantInternalServiceServer_GetTenantDataIds_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetTenantDataIdsRequest))
	})
	return _c
}

func (_c *MockTenantInternalServiceServer_GetTenantDataIds_Call) Return(_a0 *GetTenantDataIdsResponse, _a1 error) *MockTenantInternalServiceServer_GetTenantDataIds_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTenantId provides a mock function with given fields: _a0, _a1
func (_m *MockTenantInternalServiceServer) GetTenantId(_a0 context.Context, _a1 *GetTenantIdRequest) (*GetTenantIdResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *GetTenantIdResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetTenantIdRequest) *GetTenantIdResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetTenantIdResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetTenantIdRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantInternalServiceServer_GetTenantId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTenantId'
type MockTenantInternalServiceServer_GetTenantId_Call struct {
	*mock.Call
}

// GetTenantId is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *GetTenantIdRequest
func (_e *MockTenantInternalServiceServer_Expecter) GetTenantId(_a0 interface{}, _a1 interface{}) *MockTenantInternalServiceServer_GetTenantId_Call {
	return &MockTenantInternalServiceServer_GetTenantId_Call{Call: _e.mock.On("GetTenantId", _a0, _a1)}
}

func (_c *MockTenantInternalServiceServer_GetTenantId_Call) Run(run func(_a0 context.Context, _a1 *GetTenantIdRequest)) *MockTenantInternalServiceServer_GetTenantId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetTenantIdRequest))
	})
	return _c
}

func (_c *MockTenantInternalServiceServer_GetTenantId_Call) Return(_a0 *GetTenantIdResponse, _a1 error) *MockTenantInternalServiceServer_GetTenantId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTenantIds provides a mock function with given fields: _a0, _a1
func (_m *MockTenantInternalServiceServer) GetTenantIds(_a0 context.Context, _a1 *GetTenantIdsRequest) (*GetTenantIdsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *GetTenantIdsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetTenantIdsRequest) *GetTenantIdsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetTenantIdsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetTenantIdsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantInternalServiceServer_GetTenantIds_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTenantIds'
type MockTenantInternalServiceServer_GetTenantIds_Call struct {
	*mock.Call
}

// GetTenantIds is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *GetTenantIdsRequest
func (_e *MockTenantInternalServiceServer_Expecter) GetTenantIds(_a0 interface{}, _a1 interface{}) *MockTenantInternalServiceServer_GetTenantIds_Call {
	return &MockTenantInternalServiceServer_GetTenantIds_Call{Call: _e.mock.On("GetTenantIds", _a0, _a1)}
}

func (_c *MockTenantInternalServiceServer_GetTenantIds_Call) Run(run func(_a0 context.Context, _a1 *GetTenantIdsRequest)) *MockTenantInternalServiceServer_GetTenantIds_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetTenantIdsRequest))
	})
	return _c
}

func (_c *MockTenantInternalServiceServer_GetTenantIds_Call) Return(_a0 *GetTenantIdsResponse, _a1 error) *MockTenantInternalServiceServer_GetTenantIds_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTenantName provides a mock function with given fields: _a0, _a1
func (_m *MockTenantInternalServiceServer) GetTenantName(_a0 context.Context, _a1 *GetTenantNameRequest) (*GetTenantNameResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *GetTenantNameResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetTenantNameRequest) *GetTenantNameResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetTenantNameResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetTenantNameRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantInternalServiceServer_GetTenantName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTenantName'
type MockTenantInternalServiceServer_GetTenantName_Call struct {
	*mock.Call
}

// GetTenantName is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *GetTenantNameRequest
func (_e *MockTenantInternalServiceServer_Expecter) GetTenantName(_a0 interface{}, _a1 interface{}) *MockTenantInternalServiceServer_GetTenantName_Call {
	return &MockTenantInternalServiceServer_GetTenantName_Call{Call: _e.mock.On("GetTenantName", _a0, _a1)}
}

func (_c *MockTenantInternalServiceServer_GetTenantName_Call) Run(run func(_a0 context.Context, _a1 *GetTenantNameRequest)) *MockTenantInternalServiceServer_GetTenantName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetTenantNameRequest))
	})
	return _c
}

func (_c *MockTenantInternalServiceServer_GetTenantName_Call) Return(_a0 *GetTenantNameResponse, _a1 error) *MockTenantInternalServiceServer_GetTenantName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTenantNames provides a mock function with given fields: _a0, _a1
func (_m *MockTenantInternalServiceServer) GetTenantNames(_a0 context.Context, _a1 *GetTenantNamesRequest) (*GetTenantNamesResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *GetTenantNamesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetTenantNamesRequest) *GetTenantNamesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetTenantNamesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetTenantNamesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantInternalServiceServer_GetTenantNames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTenantNames'
type MockTenantInternalServiceServer_GetTenantNames_Call struct {
	*mock.Call
}

// GetTenantNames is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *GetTenantNamesRequest
func (_e *MockTenantInternalServiceServer_Expecter) GetTenantNames(_a0 interface{}, _a1 interface{}) *MockTenantInternalServiceServer_GetTenantNames_Call {
	return &MockTenantInternalServiceServer_GetTenantNames_Call{Call: _e.mock.On("GetTenantNames", _a0, _a1)}
}

func (_c *MockTenantInternalServiceServer_GetTenantNames_Call) Run(run func(_a0 context.Context, _a1 *GetTenantNamesRequest)) *MockTenantInternalServiceServer_GetTenantNames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetTenantNamesRequest))
	})
	return _c
}

func (_c *MockTenantInternalServiceServer_GetTenantNames_Call) Return(_a0 *GetTenantNamesResponse, _a1 error) *MockTenantInternalServiceServer_GetTenantNames_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListTenants provides a mock function with given fields: _a0, _a1
func (_m *MockTenantInternalServiceServer) ListTenants(_a0 context.Context, _a1 *ListTenantsRequest) (*ListTenantsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *ListTenantsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ListTenantsRequest) *ListTenantsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ListTenantsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ListTenantsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantInternalServiceServer_ListTenants_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListTenants'
type MockTenantInternalServiceServer_ListTenants_Call struct {
	*mock.Call
}

// ListTenants is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *ListTenantsRequest
func (_e *MockTenantInternalServiceServer_Expecter) ListTenants(_a0 interface{}, _a1 interface{}) *MockTenantInternalServiceServer_ListTenants_Call {
	return &MockTenantInternalServiceServer_ListTenants_Call{Call: _e.mock.On("ListTenants", _a0, _a1)}
}

func (_c *MockTenantInternalServiceServer_ListTenants_Call) Run(run func(_a0 context.Context, _a1 *ListTenantsRequest)) *MockTenantInternalServiceServer_ListTenants_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*ListTenantsRequest))
	})
	return _c
}

func (_c *MockTenantInternalServiceServer_ListTenants_Call) Return(_a0 *ListTenantsResponse, _a1 error) *MockTenantInternalServiceServer_ListTenants_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// NewMockTenantInternalServiceServer creates a new instance of MockTenantInternalServiceServer. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockTenantInternalServiceServer(t testing.TB) *MockTenantInternalServiceServer {
	mock := &MockTenantInternalServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}