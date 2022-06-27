// Code generated by mockery v2.12.2. DO NOT EDIT.

package tenant

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// MockTenantInternalServiceClient is an autogenerated mock type for the TenantInternalServiceClient type
type MockTenantInternalServiceClient struct {
	mock.Mock
}

type MockTenantInternalServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTenantInternalServiceClient) EXPECT() *MockTenantInternalServiceClient_Expecter {
	return &MockTenantInternalServiceClient_Expecter{mock: &_m.Mock}
}

// GetAllTenants provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantInternalServiceClient) GetAllTenants(ctx context.Context, in *GetAllTenantsRequest, opts ...grpc.CallOption) (*GetAllTenantsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *GetAllTenantsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetAllTenantsRequest, ...grpc.CallOption) *GetAllTenantsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetAllTenantsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetAllTenantsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantInternalServiceClient_GetAllTenants_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTenants'
type MockTenantInternalServiceClient_GetAllTenants_Call struct {
	*mock.Call
}

// GetAllTenants is a helper method to define mock.On call
//  - ctx context.Context
//  - in *GetAllTenantsRequest
//  - opts ...grpc.CallOption
func (_e *MockTenantInternalServiceClient_Expecter) GetAllTenants(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantInternalServiceClient_GetAllTenants_Call {
	return &MockTenantInternalServiceClient_GetAllTenants_Call{Call: _e.mock.On("GetAllTenants",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantInternalServiceClient_GetAllTenants_Call) Run(run func(ctx context.Context, in *GetAllTenantsRequest, opts ...grpc.CallOption)) *MockTenantInternalServiceClient_GetAllTenants_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*GetAllTenantsRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantInternalServiceClient_GetAllTenants_Call) Return(_a0 *GetAllTenantsResponse, _a1 error) *MockTenantInternalServiceClient_GetAllTenants_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTenantDataId provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantInternalServiceClient) GetTenantDataId(ctx context.Context, in *GetTenantDataIdRequest, opts ...grpc.CallOption) (*GetTenantDataIdResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *GetTenantDataIdResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetTenantDataIdRequest, ...grpc.CallOption) *GetTenantDataIdResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetTenantDataIdResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetTenantDataIdRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantInternalServiceClient_GetTenantDataId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTenantDataId'
type MockTenantInternalServiceClient_GetTenantDataId_Call struct {
	*mock.Call
}

// GetTenantDataId is a helper method to define mock.On call
//  - ctx context.Context
//  - in *GetTenantDataIdRequest
//  - opts ...grpc.CallOption
func (_e *MockTenantInternalServiceClient_Expecter) GetTenantDataId(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantInternalServiceClient_GetTenantDataId_Call {
	return &MockTenantInternalServiceClient_GetTenantDataId_Call{Call: _e.mock.On("GetTenantDataId",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantInternalServiceClient_GetTenantDataId_Call) Run(run func(ctx context.Context, in *GetTenantDataIdRequest, opts ...grpc.CallOption)) *MockTenantInternalServiceClient_GetTenantDataId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*GetTenantDataIdRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantInternalServiceClient_GetTenantDataId_Call) Return(_a0 *GetTenantDataIdResponse, _a1 error) *MockTenantInternalServiceClient_GetTenantDataId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTenantDataIds provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantInternalServiceClient) GetTenantDataIds(ctx context.Context, in *GetTenantDataIdsRequest, opts ...grpc.CallOption) (*GetTenantDataIdsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *GetTenantDataIdsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetTenantDataIdsRequest, ...grpc.CallOption) *GetTenantDataIdsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetTenantDataIdsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetTenantDataIdsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantInternalServiceClient_GetTenantDataIds_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTenantDataIds'
type MockTenantInternalServiceClient_GetTenantDataIds_Call struct {
	*mock.Call
}

// GetTenantDataIds is a helper method to define mock.On call
//  - ctx context.Context
//  - in *GetTenantDataIdsRequest
//  - opts ...grpc.CallOption
func (_e *MockTenantInternalServiceClient_Expecter) GetTenantDataIds(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantInternalServiceClient_GetTenantDataIds_Call {
	return &MockTenantInternalServiceClient_GetTenantDataIds_Call{Call: _e.mock.On("GetTenantDataIds",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantInternalServiceClient_GetTenantDataIds_Call) Run(run func(ctx context.Context, in *GetTenantDataIdsRequest, opts ...grpc.CallOption)) *MockTenantInternalServiceClient_GetTenantDataIds_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*GetTenantDataIdsRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantInternalServiceClient_GetTenantDataIds_Call) Return(_a0 *GetTenantDataIdsResponse, _a1 error) *MockTenantInternalServiceClient_GetTenantDataIds_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTenantId provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantInternalServiceClient) GetTenantId(ctx context.Context, in *GetTenantIdRequest, opts ...grpc.CallOption) (*GetTenantIdResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *GetTenantIdResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetTenantIdRequest, ...grpc.CallOption) *GetTenantIdResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetTenantIdResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetTenantIdRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantInternalServiceClient_GetTenantId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTenantId'
type MockTenantInternalServiceClient_GetTenantId_Call struct {
	*mock.Call
}

// GetTenantId is a helper method to define mock.On call
//  - ctx context.Context
//  - in *GetTenantIdRequest
//  - opts ...grpc.CallOption
func (_e *MockTenantInternalServiceClient_Expecter) GetTenantId(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantInternalServiceClient_GetTenantId_Call {
	return &MockTenantInternalServiceClient_GetTenantId_Call{Call: _e.mock.On("GetTenantId",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantInternalServiceClient_GetTenantId_Call) Run(run func(ctx context.Context, in *GetTenantIdRequest, opts ...grpc.CallOption)) *MockTenantInternalServiceClient_GetTenantId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*GetTenantIdRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantInternalServiceClient_GetTenantId_Call) Return(_a0 *GetTenantIdResponse, _a1 error) *MockTenantInternalServiceClient_GetTenantId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTenantIds provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantInternalServiceClient) GetTenantIds(ctx context.Context, in *GetTenantIdsRequest, opts ...grpc.CallOption) (*GetTenantIdsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *GetTenantIdsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetTenantIdsRequest, ...grpc.CallOption) *GetTenantIdsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetTenantIdsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetTenantIdsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantInternalServiceClient_GetTenantIds_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTenantIds'
type MockTenantInternalServiceClient_GetTenantIds_Call struct {
	*mock.Call
}

// GetTenantIds is a helper method to define mock.On call
//  - ctx context.Context
//  - in *GetTenantIdsRequest
//  - opts ...grpc.CallOption
func (_e *MockTenantInternalServiceClient_Expecter) GetTenantIds(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantInternalServiceClient_GetTenantIds_Call {
	return &MockTenantInternalServiceClient_GetTenantIds_Call{Call: _e.mock.On("GetTenantIds",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantInternalServiceClient_GetTenantIds_Call) Run(run func(ctx context.Context, in *GetTenantIdsRequest, opts ...grpc.CallOption)) *MockTenantInternalServiceClient_GetTenantIds_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*GetTenantIdsRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantInternalServiceClient_GetTenantIds_Call) Return(_a0 *GetTenantIdsResponse, _a1 error) *MockTenantInternalServiceClient_GetTenantIds_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTenantName provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantInternalServiceClient) GetTenantName(ctx context.Context, in *GetTenantNameRequest, opts ...grpc.CallOption) (*GetTenantNameResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *GetTenantNameResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetTenantNameRequest, ...grpc.CallOption) *GetTenantNameResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetTenantNameResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetTenantNameRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantInternalServiceClient_GetTenantName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTenantName'
type MockTenantInternalServiceClient_GetTenantName_Call struct {
	*mock.Call
}

// GetTenantName is a helper method to define mock.On call
//  - ctx context.Context
//  - in *GetTenantNameRequest
//  - opts ...grpc.CallOption
func (_e *MockTenantInternalServiceClient_Expecter) GetTenantName(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantInternalServiceClient_GetTenantName_Call {
	return &MockTenantInternalServiceClient_GetTenantName_Call{Call: _e.mock.On("GetTenantName",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantInternalServiceClient_GetTenantName_Call) Run(run func(ctx context.Context, in *GetTenantNameRequest, opts ...grpc.CallOption)) *MockTenantInternalServiceClient_GetTenantName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*GetTenantNameRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantInternalServiceClient_GetTenantName_Call) Return(_a0 *GetTenantNameResponse, _a1 error) *MockTenantInternalServiceClient_GetTenantName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTenantNames provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantInternalServiceClient) GetTenantNames(ctx context.Context, in *GetTenantNamesRequest, opts ...grpc.CallOption) (*GetTenantNamesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *GetTenantNamesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetTenantNamesRequest, ...grpc.CallOption) *GetTenantNamesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetTenantNamesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetTenantNamesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantInternalServiceClient_GetTenantNames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTenantNames'
type MockTenantInternalServiceClient_GetTenantNames_Call struct {
	*mock.Call
}

// GetTenantNames is a helper method to define mock.On call
//  - ctx context.Context
//  - in *GetTenantNamesRequest
//  - opts ...grpc.CallOption
func (_e *MockTenantInternalServiceClient_Expecter) GetTenantNames(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantInternalServiceClient_GetTenantNames_Call {
	return &MockTenantInternalServiceClient_GetTenantNames_Call{Call: _e.mock.On("GetTenantNames",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantInternalServiceClient_GetTenantNames_Call) Run(run func(ctx context.Context, in *GetTenantNamesRequest, opts ...grpc.CallOption)) *MockTenantInternalServiceClient_GetTenantNames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*GetTenantNamesRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantInternalServiceClient_GetTenantNames_Call) Return(_a0 *GetTenantNamesResponse, _a1 error) *MockTenantInternalServiceClient_GetTenantNames_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ListTenants provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantInternalServiceClient) ListTenants(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ListTenantsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ListTenantsRequest, ...grpc.CallOption) *ListTenantsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ListTenantsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ListTenantsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantInternalServiceClient_ListTenants_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListTenants'
type MockTenantInternalServiceClient_ListTenants_Call struct {
	*mock.Call
}

// ListTenants is a helper method to define mock.On call
//  - ctx context.Context
//  - in *ListTenantsRequest
//  - opts ...grpc.CallOption
func (_e *MockTenantInternalServiceClient_Expecter) ListTenants(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantInternalServiceClient_ListTenants_Call {
	return &MockTenantInternalServiceClient_ListTenants_Call{Call: _e.mock.On("ListTenants",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantInternalServiceClient_ListTenants_Call) Run(run func(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption)) *MockTenantInternalServiceClient_ListTenants_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*ListTenantsRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantInternalServiceClient_ListTenants_Call) Return(_a0 *ListTenantsResponse, _a1 error) *MockTenantInternalServiceClient_ListTenants_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// NewMockTenantInternalServiceClient creates a new instance of MockTenantInternalServiceClient. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockTenantInternalServiceClient(t testing.TB) *MockTenantInternalServiceClient {
	mock := &MockTenantInternalServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
