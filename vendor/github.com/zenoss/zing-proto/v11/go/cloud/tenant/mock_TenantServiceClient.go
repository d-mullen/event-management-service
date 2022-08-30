// Code generated by mockery v2.13.1. DO NOT EDIT.

package tenant

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockTenantServiceClient is an autogenerated mock type for the TenantServiceClient type
type MockTenantServiceClient struct {
	mock.Mock
}

type MockTenantServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTenantServiceClient) EXPECT() *MockTenantServiceClient_Expecter {
	return &MockTenantServiceClient_Expecter{mock: &_m.Mock}
}

// CreateAuthInfo provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantServiceClient) CreateAuthInfo(ctx context.Context, in *CreateAuthRequest, opts ...grpc.CallOption) (*CreateAuthResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *CreateAuthResponse
	if rf, ok := ret.Get(0).(func(context.Context, *CreateAuthRequest, ...grpc.CallOption) *CreateAuthResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*CreateAuthResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *CreateAuthRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceClient_CreateAuthInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAuthInfo'
type MockTenantServiceClient_CreateAuthInfo_Call struct {
	*mock.Call
}

// CreateAuthInfo is a helper method to define mock.On call
//  - ctx context.Context
//  - in *CreateAuthRequest
//  - opts ...grpc.CallOption
func (_e *MockTenantServiceClient_Expecter) CreateAuthInfo(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantServiceClient_CreateAuthInfo_Call {
	return &MockTenantServiceClient_CreateAuthInfo_Call{Call: _e.mock.On("CreateAuthInfo",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantServiceClient_CreateAuthInfo_Call) Run(run func(ctx context.Context, in *CreateAuthRequest, opts ...grpc.CallOption)) *MockTenantServiceClient_CreateAuthInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*CreateAuthRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantServiceClient_CreateAuthInfo_Call) Return(_a0 *CreateAuthResponse, _a1 error) *MockTenantServiceClient_CreateAuthInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// DeleteAuthInfo provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantServiceClient) DeleteAuthInfo(ctx context.Context, in *DeleteAuthRequest, opts ...grpc.CallOption) (*DeleteAuthResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *DeleteAuthResponse
	if rf, ok := ret.Get(0).(func(context.Context, *DeleteAuthRequest, ...grpc.CallOption) *DeleteAuthResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DeleteAuthResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *DeleteAuthRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceClient_DeleteAuthInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAuthInfo'
type MockTenantServiceClient_DeleteAuthInfo_Call struct {
	*mock.Call
}

// DeleteAuthInfo is a helper method to define mock.On call
//  - ctx context.Context
//  - in *DeleteAuthRequest
//  - opts ...grpc.CallOption
func (_e *MockTenantServiceClient_Expecter) DeleteAuthInfo(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantServiceClient_DeleteAuthInfo_Call {
	return &MockTenantServiceClient_DeleteAuthInfo_Call{Call: _e.mock.On("DeleteAuthInfo",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantServiceClient_DeleteAuthInfo_Call) Run(run func(ctx context.Context, in *DeleteAuthRequest, opts ...grpc.CallOption)) *MockTenantServiceClient_DeleteAuthInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*DeleteAuthRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantServiceClient_DeleteAuthInfo_Call) Return(_a0 *DeleteAuthResponse, _a1 error) *MockTenantServiceClient_DeleteAuthInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetAuthInfo provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantServiceClient) GetAuthInfo(ctx context.Context, in *GetAuthRequest, opts ...grpc.CallOption) (*GetAuthResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *GetAuthResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetAuthRequest, ...grpc.CallOption) *GetAuthResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetAuthResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetAuthRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceClient_GetAuthInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuthInfo'
type MockTenantServiceClient_GetAuthInfo_Call struct {
	*mock.Call
}

// GetAuthInfo is a helper method to define mock.On call
//  - ctx context.Context
//  - in *GetAuthRequest
//  - opts ...grpc.CallOption
func (_e *MockTenantServiceClient_Expecter) GetAuthInfo(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantServiceClient_GetAuthInfo_Call {
	return &MockTenantServiceClient_GetAuthInfo_Call{Call: _e.mock.On("GetAuthInfo",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantServiceClient_GetAuthInfo_Call) Run(run func(ctx context.Context, in *GetAuthRequest, opts ...grpc.CallOption)) *MockTenantServiceClient_GetAuthInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*GetAuthRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantServiceClient_GetAuthInfo_Call) Return(_a0 *GetAuthResponse, _a1 error) *MockTenantServiceClient_GetAuthInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetCollectionZones provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantServiceClient) GetCollectionZones(ctx context.Context, in *GetCollectionZonesRequest, opts ...grpc.CallOption) (*GetCollectionZonesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *GetCollectionZonesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetCollectionZonesRequest, ...grpc.CallOption) *GetCollectionZonesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetCollectionZonesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetCollectionZonesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceClient_GetCollectionZones_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollectionZones'
type MockTenantServiceClient_GetCollectionZones_Call struct {
	*mock.Call
}

// GetCollectionZones is a helper method to define mock.On call
//  - ctx context.Context
//  - in *GetCollectionZonesRequest
//  - opts ...grpc.CallOption
func (_e *MockTenantServiceClient_Expecter) GetCollectionZones(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantServiceClient_GetCollectionZones_Call {
	return &MockTenantServiceClient_GetCollectionZones_Call{Call: _e.mock.On("GetCollectionZones",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantServiceClient_GetCollectionZones_Call) Run(run func(ctx context.Context, in *GetCollectionZonesRequest, opts ...grpc.CallOption)) *MockTenantServiceClient_GetCollectionZones_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*GetCollectionZonesRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantServiceClient_GetCollectionZones_Call) Return(_a0 *GetCollectionZonesResponse, _a1 error) *MockTenantServiceClient_GetCollectionZones_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetLandingPage provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantServiceClient) GetLandingPage(ctx context.Context, in *GetLandingPageRequest, opts ...grpc.CallOption) (*LandingPageResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *LandingPageResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetLandingPageRequest, ...grpc.CallOption) *LandingPageResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*LandingPageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetLandingPageRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceClient_GetLandingPage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLandingPage'
type MockTenantServiceClient_GetLandingPage_Call struct {
	*mock.Call
}

// GetLandingPage is a helper method to define mock.On call
//  - ctx context.Context
//  - in *GetLandingPageRequest
//  - opts ...grpc.CallOption
func (_e *MockTenantServiceClient_Expecter) GetLandingPage(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantServiceClient_GetLandingPage_Call {
	return &MockTenantServiceClient_GetLandingPage_Call{Call: _e.mock.On("GetLandingPage",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantServiceClient_GetLandingPage_Call) Run(run func(ctx context.Context, in *GetLandingPageRequest, opts ...grpc.CallOption)) *MockTenantServiceClient_GetLandingPage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*GetLandingPageRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantServiceClient_GetLandingPage_Call) Return(_a0 *LandingPageResponse, _a1 error) *MockTenantServiceClient_GetLandingPage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetLoginMessage provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantServiceClient) GetLoginMessage(ctx context.Context, in *GetLoginMessageRequest, opts ...grpc.CallOption) (*LoginMessageProto, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *LoginMessageProto
	if rf, ok := ret.Get(0).(func(context.Context, *GetLoginMessageRequest, ...grpc.CallOption) *LoginMessageProto); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*LoginMessageProto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetLoginMessageRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceClient_GetLoginMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLoginMessage'
type MockTenantServiceClient_GetLoginMessage_Call struct {
	*mock.Call
}

// GetLoginMessage is a helper method to define mock.On call
//  - ctx context.Context
//  - in *GetLoginMessageRequest
//  - opts ...grpc.CallOption
func (_e *MockTenantServiceClient_Expecter) GetLoginMessage(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantServiceClient_GetLoginMessage_Call {
	return &MockTenantServiceClient_GetLoginMessage_Call{Call: _e.mock.On("GetLoginMessage",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantServiceClient_GetLoginMessage_Call) Run(run func(ctx context.Context, in *GetLoginMessageRequest, opts ...grpc.CallOption)) *MockTenantServiceClient_GetLoginMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*GetLoginMessageRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantServiceClient_GetLoginMessage_Call) Return(_a0 *LoginMessageProto, _a1 error) *MockTenantServiceClient_GetLoginMessage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetLogo provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantServiceClient) GetLogo(ctx context.Context, in *GetLogoRequest, opts ...grpc.CallOption) (*LogoResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *LogoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetLogoRequest, ...grpc.CallOption) *LogoResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*LogoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetLogoRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceClient_GetLogo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLogo'
type MockTenantServiceClient_GetLogo_Call struct {
	*mock.Call
}

// GetLogo is a helper method to define mock.On call
//  - ctx context.Context
//  - in *GetLogoRequest
//  - opts ...grpc.CallOption
func (_e *MockTenantServiceClient_Expecter) GetLogo(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantServiceClient_GetLogo_Call {
	return &MockTenantServiceClient_GetLogo_Call{Call: _e.mock.On("GetLogo",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantServiceClient_GetLogo_Call) Run(run func(ctx context.Context, in *GetLogoRequest, opts ...grpc.CallOption)) *MockTenantServiceClient_GetLogo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*GetLogoRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantServiceClient_GetLogo_Call) Return(_a0 *LogoResponse, _a1 error) *MockTenantServiceClient_GetLogo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetSessionSettings provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantServiceClient) GetSessionSettings(ctx context.Context, in *GetSessionSettingsRequest, opts ...grpc.CallOption) (*GetSessionSettingsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *GetSessionSettingsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetSessionSettingsRequest, ...grpc.CallOption) *GetSessionSettingsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetSessionSettingsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetSessionSettingsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceClient_GetSessionSettings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSessionSettings'
type MockTenantServiceClient_GetSessionSettings_Call struct {
	*mock.Call
}

// GetSessionSettings is a helper method to define mock.On call
//  - ctx context.Context
//  - in *GetSessionSettingsRequest
//  - opts ...grpc.CallOption
func (_e *MockTenantServiceClient_Expecter) GetSessionSettings(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantServiceClient_GetSessionSettings_Call {
	return &MockTenantServiceClient_GetSessionSettings_Call{Call: _e.mock.On("GetSessionSettings",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantServiceClient_GetSessionSettings_Call) Run(run func(ctx context.Context, in *GetSessionSettingsRequest, opts ...grpc.CallOption)) *MockTenantServiceClient_GetSessionSettings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*GetSessionSettingsRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantServiceClient_GetSessionSettings_Call) Return(_a0 *GetSessionSettingsResponse, _a1 error) *MockTenantServiceClient_GetSessionSettings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTenant provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantServiceClient) GetTenant(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetTenantResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *GetTenantResponse
	if rf, ok := ret.Get(0).(func(context.Context, *Empty, ...grpc.CallOption) *GetTenantResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetTenantResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *Empty, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceClient_GetTenant_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTenant'
type MockTenantServiceClient_GetTenant_Call struct {
	*mock.Call
}

// GetTenant is a helper method to define mock.On call
//  - ctx context.Context
//  - in *Empty
//  - opts ...grpc.CallOption
func (_e *MockTenantServiceClient_Expecter) GetTenant(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantServiceClient_GetTenant_Call {
	return &MockTenantServiceClient_GetTenant_Call{Call: _e.mock.On("GetTenant",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantServiceClient_GetTenant_Call) Run(run func(ctx context.Context, in *Empty, opts ...grpc.CallOption)) *MockTenantServiceClient_GetTenant_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*Empty), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantServiceClient_GetTenant_Call) Return(_a0 *GetTenantResponse, _a1 error) *MockTenantServiceClient_GetTenant_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTenantTheme provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantServiceClient) GetTenantTheme(ctx context.Context, in *GetTenantThemeRequest, opts ...grpc.CallOption) (*TenantThemeResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *TenantThemeResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetTenantThemeRequest, ...grpc.CallOption) *TenantThemeResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TenantThemeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetTenantThemeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceClient_GetTenantTheme_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTenantTheme'
type MockTenantServiceClient_GetTenantTheme_Call struct {
	*mock.Call
}

// GetTenantTheme is a helper method to define mock.On call
//  - ctx context.Context
//  - in *GetTenantThemeRequest
//  - opts ...grpc.CallOption
func (_e *MockTenantServiceClient_Expecter) GetTenantTheme(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantServiceClient_GetTenantTheme_Call {
	return &MockTenantServiceClient_GetTenantTheme_Call{Call: _e.mock.On("GetTenantTheme",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantServiceClient_GetTenantTheme_Call) Run(run func(ctx context.Context, in *GetTenantThemeRequest, opts ...grpc.CallOption)) *MockTenantServiceClient_GetTenantTheme_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*GetTenantThemeRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantServiceClient_GetTenantTheme_Call) Return(_a0 *TenantThemeResponse, _a1 error) *MockTenantServiceClient_GetTenantTheme_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// LoadAuthInfo provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantServiceClient) LoadAuthInfo(ctx context.Context, in *LoadAuthRequest, opts ...grpc.CallOption) (*LoadAuthResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *LoadAuthResponse
	if rf, ok := ret.Get(0).(func(context.Context, *LoadAuthRequest, ...grpc.CallOption) *LoadAuthResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*LoadAuthResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *LoadAuthRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceClient_LoadAuthInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadAuthInfo'
type MockTenantServiceClient_LoadAuthInfo_Call struct {
	*mock.Call
}

// LoadAuthInfo is a helper method to define mock.On call
//  - ctx context.Context
//  - in *LoadAuthRequest
//  - opts ...grpc.CallOption
func (_e *MockTenantServiceClient_Expecter) LoadAuthInfo(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantServiceClient_LoadAuthInfo_Call {
	return &MockTenantServiceClient_LoadAuthInfo_Call{Call: _e.mock.On("LoadAuthInfo",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantServiceClient_LoadAuthInfo_Call) Run(run func(ctx context.Context, in *LoadAuthRequest, opts ...grpc.CallOption)) *MockTenantServiceClient_LoadAuthInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*LoadAuthRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantServiceClient_LoadAuthInfo_Call) Return(_a0 *LoadAuthResponse, _a1 error) *MockTenantServiceClient_LoadAuthInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdateAuthInfo provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantServiceClient) UpdateAuthInfo(ctx context.Context, in *UpdateAuthRequest, opts ...grpc.CallOption) (*UpdateAuthResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *UpdateAuthResponse
	if rf, ok := ret.Get(0).(func(context.Context, *UpdateAuthRequest, ...grpc.CallOption) *UpdateAuthResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UpdateAuthResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *UpdateAuthRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceClient_UpdateAuthInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAuthInfo'
type MockTenantServiceClient_UpdateAuthInfo_Call struct {
	*mock.Call
}

// UpdateAuthInfo is a helper method to define mock.On call
//  - ctx context.Context
//  - in *UpdateAuthRequest
//  - opts ...grpc.CallOption
func (_e *MockTenantServiceClient_Expecter) UpdateAuthInfo(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantServiceClient_UpdateAuthInfo_Call {
	return &MockTenantServiceClient_UpdateAuthInfo_Call{Call: _e.mock.On("UpdateAuthInfo",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantServiceClient_UpdateAuthInfo_Call) Run(run func(ctx context.Context, in *UpdateAuthRequest, opts ...grpc.CallOption)) *MockTenantServiceClient_UpdateAuthInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*UpdateAuthRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantServiceClient_UpdateAuthInfo_Call) Return(_a0 *UpdateAuthResponse, _a1 error) *MockTenantServiceClient_UpdateAuthInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdateLandingPage provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantServiceClient) UpdateLandingPage(ctx context.Context, in *LandingPageResponse, opts ...grpc.CallOption) (*LandingPageResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *LandingPageResponse
	if rf, ok := ret.Get(0).(func(context.Context, *LandingPageResponse, ...grpc.CallOption) *LandingPageResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*LandingPageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *LandingPageResponse, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceClient_UpdateLandingPage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateLandingPage'
type MockTenantServiceClient_UpdateLandingPage_Call struct {
	*mock.Call
}

// UpdateLandingPage is a helper method to define mock.On call
//  - ctx context.Context
//  - in *LandingPageResponse
//  - opts ...grpc.CallOption
func (_e *MockTenantServiceClient_Expecter) UpdateLandingPage(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantServiceClient_UpdateLandingPage_Call {
	return &MockTenantServiceClient_UpdateLandingPage_Call{Call: _e.mock.On("UpdateLandingPage",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantServiceClient_UpdateLandingPage_Call) Run(run func(ctx context.Context, in *LandingPageResponse, opts ...grpc.CallOption)) *MockTenantServiceClient_UpdateLandingPage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*LandingPageResponse), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantServiceClient_UpdateLandingPage_Call) Return(_a0 *LandingPageResponse, _a1 error) *MockTenantServiceClient_UpdateLandingPage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdateLoginMessage provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantServiceClient) UpdateLoginMessage(ctx context.Context, in *LoginMessageProto, opts ...grpc.CallOption) (*LoginMessageProto, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *LoginMessageProto
	if rf, ok := ret.Get(0).(func(context.Context, *LoginMessageProto, ...grpc.CallOption) *LoginMessageProto); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*LoginMessageProto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *LoginMessageProto, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceClient_UpdateLoginMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateLoginMessage'
type MockTenantServiceClient_UpdateLoginMessage_Call struct {
	*mock.Call
}

// UpdateLoginMessage is a helper method to define mock.On call
//  - ctx context.Context
//  - in *LoginMessageProto
//  - opts ...grpc.CallOption
func (_e *MockTenantServiceClient_Expecter) UpdateLoginMessage(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantServiceClient_UpdateLoginMessage_Call {
	return &MockTenantServiceClient_UpdateLoginMessage_Call{Call: _e.mock.On("UpdateLoginMessage",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantServiceClient_UpdateLoginMessage_Call) Run(run func(ctx context.Context, in *LoginMessageProto, opts ...grpc.CallOption)) *MockTenantServiceClient_UpdateLoginMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*LoginMessageProto), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantServiceClient_UpdateLoginMessage_Call) Return(_a0 *LoginMessageProto, _a1 error) *MockTenantServiceClient_UpdateLoginMessage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdateLogo provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantServiceClient) UpdateLogo(ctx context.Context, in *LogoResponse, opts ...grpc.CallOption) (*LogoResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *LogoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *LogoResponse, ...grpc.CallOption) *LogoResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*LogoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *LogoResponse, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceClient_UpdateLogo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateLogo'
type MockTenantServiceClient_UpdateLogo_Call struct {
	*mock.Call
}

// UpdateLogo is a helper method to define mock.On call
//  - ctx context.Context
//  - in *LogoResponse
//  - opts ...grpc.CallOption
func (_e *MockTenantServiceClient_Expecter) UpdateLogo(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantServiceClient_UpdateLogo_Call {
	return &MockTenantServiceClient_UpdateLogo_Call{Call: _e.mock.On("UpdateLogo",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantServiceClient_UpdateLogo_Call) Run(run func(ctx context.Context, in *LogoResponse, opts ...grpc.CallOption)) *MockTenantServiceClient_UpdateLogo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*LogoResponse), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantServiceClient_UpdateLogo_Call) Return(_a0 *LogoResponse, _a1 error) *MockTenantServiceClient_UpdateLogo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdateSessionSettings provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantServiceClient) UpdateSessionSettings(ctx context.Context, in *UpdateSessionSettingsRequest, opts ...grpc.CallOption) (*UpdateSessionSettingsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *UpdateSessionSettingsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *UpdateSessionSettingsRequest, ...grpc.CallOption) *UpdateSessionSettingsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UpdateSessionSettingsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *UpdateSessionSettingsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceClient_UpdateSessionSettings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSessionSettings'
type MockTenantServiceClient_UpdateSessionSettings_Call struct {
	*mock.Call
}

// UpdateSessionSettings is a helper method to define mock.On call
//  - ctx context.Context
//  - in *UpdateSessionSettingsRequest
//  - opts ...grpc.CallOption
func (_e *MockTenantServiceClient_Expecter) UpdateSessionSettings(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantServiceClient_UpdateSessionSettings_Call {
	return &MockTenantServiceClient_UpdateSessionSettings_Call{Call: _e.mock.On("UpdateSessionSettings",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantServiceClient_UpdateSessionSettings_Call) Run(run func(ctx context.Context, in *UpdateSessionSettingsRequest, opts ...grpc.CallOption)) *MockTenantServiceClient_UpdateSessionSettings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*UpdateSessionSettingsRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantServiceClient_UpdateSessionSettings_Call) Return(_a0 *UpdateSessionSettingsResponse, _a1 error) *MockTenantServiceClient_UpdateSessionSettings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdateTenantTheme provides a mock function with given fields: ctx, in, opts
func (_m *MockTenantServiceClient) UpdateTenantTheme(ctx context.Context, in *TenantThemeResponse, opts ...grpc.CallOption) (*TenantThemeResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *TenantThemeResponse
	if rf, ok := ret.Get(0).(func(context.Context, *TenantThemeResponse, ...grpc.CallOption) *TenantThemeResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TenantThemeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *TenantThemeResponse, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceClient_UpdateTenantTheme_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTenantTheme'
type MockTenantServiceClient_UpdateTenantTheme_Call struct {
	*mock.Call
}

// UpdateTenantTheme is a helper method to define mock.On call
//  - ctx context.Context
//  - in *TenantThemeResponse
//  - opts ...grpc.CallOption
func (_e *MockTenantServiceClient_Expecter) UpdateTenantTheme(ctx interface{}, in interface{}, opts ...interface{}) *MockTenantServiceClient_UpdateTenantTheme_Call {
	return &MockTenantServiceClient_UpdateTenantTheme_Call{Call: _e.mock.On("UpdateTenantTheme",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockTenantServiceClient_UpdateTenantTheme_Call) Run(run func(ctx context.Context, in *TenantThemeResponse, opts ...grpc.CallOption)) *MockTenantServiceClient_UpdateTenantTheme_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*TenantThemeResponse), variadicArgs...)
	})
	return _c
}

func (_c *MockTenantServiceClient_UpdateTenantTheme_Call) Return(_a0 *TenantThemeResponse, _a1 error) *MockTenantServiceClient_UpdateTenantTheme_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewMockTenantServiceClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockTenantServiceClient creates a new instance of MockTenantServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockTenantServiceClient(t mockConstructorTestingTNewMockTenantServiceClient) *MockTenantServiceClient {
	mock := &MockTenantServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
