// Code generated by mockery v2.14.0. DO NOT EDIT.

package tenant

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockTenantServiceServer is an autogenerated mock type for the TenantServiceServer type
type MockTenantServiceServer struct {
	mock.Mock
}

type MockTenantServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTenantServiceServer) EXPECT() *MockTenantServiceServer_Expecter {
	return &MockTenantServiceServer_Expecter{mock: &_m.Mock}
}

// CreateAuthInfo provides a mock function with given fields: _a0, _a1
func (_m *MockTenantServiceServer) CreateAuthInfo(_a0 context.Context, _a1 *CreateAuthRequest) (*CreateAuthResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *CreateAuthResponse
	if rf, ok := ret.Get(0).(func(context.Context, *CreateAuthRequest) *CreateAuthResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*CreateAuthResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *CreateAuthRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceServer_CreateAuthInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAuthInfo'
type MockTenantServiceServer_CreateAuthInfo_Call struct {
	*mock.Call
}

// CreateAuthInfo is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *CreateAuthRequest
func (_e *MockTenantServiceServer_Expecter) CreateAuthInfo(_a0 interface{}, _a1 interface{}) *MockTenantServiceServer_CreateAuthInfo_Call {
	return &MockTenantServiceServer_CreateAuthInfo_Call{Call: _e.mock.On("CreateAuthInfo", _a0, _a1)}
}

func (_c *MockTenantServiceServer_CreateAuthInfo_Call) Run(run func(_a0 context.Context, _a1 *CreateAuthRequest)) *MockTenantServiceServer_CreateAuthInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*CreateAuthRequest))
	})
	return _c
}

func (_c *MockTenantServiceServer_CreateAuthInfo_Call) Return(_a0 *CreateAuthResponse, _a1 error) *MockTenantServiceServer_CreateAuthInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// DeleteAuthInfo provides a mock function with given fields: _a0, _a1
func (_m *MockTenantServiceServer) DeleteAuthInfo(_a0 context.Context, _a1 *DeleteAuthRequest) (*DeleteAuthResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *DeleteAuthResponse
	if rf, ok := ret.Get(0).(func(context.Context, *DeleteAuthRequest) *DeleteAuthResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*DeleteAuthResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *DeleteAuthRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceServer_DeleteAuthInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAuthInfo'
type MockTenantServiceServer_DeleteAuthInfo_Call struct {
	*mock.Call
}

// DeleteAuthInfo is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *DeleteAuthRequest
func (_e *MockTenantServiceServer_Expecter) DeleteAuthInfo(_a0 interface{}, _a1 interface{}) *MockTenantServiceServer_DeleteAuthInfo_Call {
	return &MockTenantServiceServer_DeleteAuthInfo_Call{Call: _e.mock.On("DeleteAuthInfo", _a0, _a1)}
}

func (_c *MockTenantServiceServer_DeleteAuthInfo_Call) Run(run func(_a0 context.Context, _a1 *DeleteAuthRequest)) *MockTenantServiceServer_DeleteAuthInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*DeleteAuthRequest))
	})
	return _c
}

func (_c *MockTenantServiceServer_DeleteAuthInfo_Call) Return(_a0 *DeleteAuthResponse, _a1 error) *MockTenantServiceServer_DeleteAuthInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetAuthInfo provides a mock function with given fields: _a0, _a1
func (_m *MockTenantServiceServer) GetAuthInfo(_a0 context.Context, _a1 *GetAuthRequest) (*GetAuthResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *GetAuthResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetAuthRequest) *GetAuthResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetAuthResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetAuthRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceServer_GetAuthInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuthInfo'
type MockTenantServiceServer_GetAuthInfo_Call struct {
	*mock.Call
}

// GetAuthInfo is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *GetAuthRequest
func (_e *MockTenantServiceServer_Expecter) GetAuthInfo(_a0 interface{}, _a1 interface{}) *MockTenantServiceServer_GetAuthInfo_Call {
	return &MockTenantServiceServer_GetAuthInfo_Call{Call: _e.mock.On("GetAuthInfo", _a0, _a1)}
}

func (_c *MockTenantServiceServer_GetAuthInfo_Call) Run(run func(_a0 context.Context, _a1 *GetAuthRequest)) *MockTenantServiceServer_GetAuthInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetAuthRequest))
	})
	return _c
}

func (_c *MockTenantServiceServer_GetAuthInfo_Call) Return(_a0 *GetAuthResponse, _a1 error) *MockTenantServiceServer_GetAuthInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetCollectionZones provides a mock function with given fields: _a0, _a1
func (_m *MockTenantServiceServer) GetCollectionZones(_a0 context.Context, _a1 *GetCollectionZonesRequest) (*GetCollectionZonesResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *GetCollectionZonesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetCollectionZonesRequest) *GetCollectionZonesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetCollectionZonesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetCollectionZonesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceServer_GetCollectionZones_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollectionZones'
type MockTenantServiceServer_GetCollectionZones_Call struct {
	*mock.Call
}

// GetCollectionZones is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *GetCollectionZonesRequest
func (_e *MockTenantServiceServer_Expecter) GetCollectionZones(_a0 interface{}, _a1 interface{}) *MockTenantServiceServer_GetCollectionZones_Call {
	return &MockTenantServiceServer_GetCollectionZones_Call{Call: _e.mock.On("GetCollectionZones", _a0, _a1)}
}

func (_c *MockTenantServiceServer_GetCollectionZones_Call) Run(run func(_a0 context.Context, _a1 *GetCollectionZonesRequest)) *MockTenantServiceServer_GetCollectionZones_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetCollectionZonesRequest))
	})
	return _c
}

func (_c *MockTenantServiceServer_GetCollectionZones_Call) Return(_a0 *GetCollectionZonesResponse, _a1 error) *MockTenantServiceServer_GetCollectionZones_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetLandingPage provides a mock function with given fields: _a0, _a1
func (_m *MockTenantServiceServer) GetLandingPage(_a0 context.Context, _a1 *GetLandingPageRequest) (*LandingPageResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *LandingPageResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetLandingPageRequest) *LandingPageResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*LandingPageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetLandingPageRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceServer_GetLandingPage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLandingPage'
type MockTenantServiceServer_GetLandingPage_Call struct {
	*mock.Call
}

// GetLandingPage is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *GetLandingPageRequest
func (_e *MockTenantServiceServer_Expecter) GetLandingPage(_a0 interface{}, _a1 interface{}) *MockTenantServiceServer_GetLandingPage_Call {
	return &MockTenantServiceServer_GetLandingPage_Call{Call: _e.mock.On("GetLandingPage", _a0, _a1)}
}

func (_c *MockTenantServiceServer_GetLandingPage_Call) Run(run func(_a0 context.Context, _a1 *GetLandingPageRequest)) *MockTenantServiceServer_GetLandingPage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetLandingPageRequest))
	})
	return _c
}

func (_c *MockTenantServiceServer_GetLandingPage_Call) Return(_a0 *LandingPageResponse, _a1 error) *MockTenantServiceServer_GetLandingPage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetLoginMessage provides a mock function with given fields: _a0, _a1
func (_m *MockTenantServiceServer) GetLoginMessage(_a0 context.Context, _a1 *GetLoginMessageRequest) (*LoginMessageProto, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *LoginMessageProto
	if rf, ok := ret.Get(0).(func(context.Context, *GetLoginMessageRequest) *LoginMessageProto); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*LoginMessageProto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetLoginMessageRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceServer_GetLoginMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLoginMessage'
type MockTenantServiceServer_GetLoginMessage_Call struct {
	*mock.Call
}

// GetLoginMessage is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *GetLoginMessageRequest
func (_e *MockTenantServiceServer_Expecter) GetLoginMessage(_a0 interface{}, _a1 interface{}) *MockTenantServiceServer_GetLoginMessage_Call {
	return &MockTenantServiceServer_GetLoginMessage_Call{Call: _e.mock.On("GetLoginMessage", _a0, _a1)}
}

func (_c *MockTenantServiceServer_GetLoginMessage_Call) Run(run func(_a0 context.Context, _a1 *GetLoginMessageRequest)) *MockTenantServiceServer_GetLoginMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetLoginMessageRequest))
	})
	return _c
}

func (_c *MockTenantServiceServer_GetLoginMessage_Call) Return(_a0 *LoginMessageProto, _a1 error) *MockTenantServiceServer_GetLoginMessage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetLogo provides a mock function with given fields: _a0, _a1
func (_m *MockTenantServiceServer) GetLogo(_a0 context.Context, _a1 *GetLogoRequest) (*LogoResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *LogoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetLogoRequest) *LogoResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*LogoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetLogoRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceServer_GetLogo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLogo'
type MockTenantServiceServer_GetLogo_Call struct {
	*mock.Call
}

// GetLogo is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *GetLogoRequest
func (_e *MockTenantServiceServer_Expecter) GetLogo(_a0 interface{}, _a1 interface{}) *MockTenantServiceServer_GetLogo_Call {
	return &MockTenantServiceServer_GetLogo_Call{Call: _e.mock.On("GetLogo", _a0, _a1)}
}

func (_c *MockTenantServiceServer_GetLogo_Call) Run(run func(_a0 context.Context, _a1 *GetLogoRequest)) *MockTenantServiceServer_GetLogo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetLogoRequest))
	})
	return _c
}

func (_c *MockTenantServiceServer_GetLogo_Call) Return(_a0 *LogoResponse, _a1 error) *MockTenantServiceServer_GetLogo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetSessionSettings provides a mock function with given fields: _a0, _a1
func (_m *MockTenantServiceServer) GetSessionSettings(_a0 context.Context, _a1 *GetSessionSettingsRequest) (*GetSessionSettingsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *GetSessionSettingsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetSessionSettingsRequest) *GetSessionSettingsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetSessionSettingsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetSessionSettingsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceServer_GetSessionSettings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSessionSettings'
type MockTenantServiceServer_GetSessionSettings_Call struct {
	*mock.Call
}

// GetSessionSettings is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *GetSessionSettingsRequest
func (_e *MockTenantServiceServer_Expecter) GetSessionSettings(_a0 interface{}, _a1 interface{}) *MockTenantServiceServer_GetSessionSettings_Call {
	return &MockTenantServiceServer_GetSessionSettings_Call{Call: _e.mock.On("GetSessionSettings", _a0, _a1)}
}

func (_c *MockTenantServiceServer_GetSessionSettings_Call) Run(run func(_a0 context.Context, _a1 *GetSessionSettingsRequest)) *MockTenantServiceServer_GetSessionSettings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetSessionSettingsRequest))
	})
	return _c
}

func (_c *MockTenantServiceServer_GetSessionSettings_Call) Return(_a0 *GetSessionSettingsResponse, _a1 error) *MockTenantServiceServer_GetSessionSettings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTenant provides a mock function with given fields: _a0, _a1
func (_m *MockTenantServiceServer) GetTenant(_a0 context.Context, _a1 *Empty) (*GetTenantResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *GetTenantResponse
	if rf, ok := ret.Get(0).(func(context.Context, *Empty) *GetTenantResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetTenantResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *Empty) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceServer_GetTenant_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTenant'
type MockTenantServiceServer_GetTenant_Call struct {
	*mock.Call
}

// GetTenant is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *Empty
func (_e *MockTenantServiceServer_Expecter) GetTenant(_a0 interface{}, _a1 interface{}) *MockTenantServiceServer_GetTenant_Call {
	return &MockTenantServiceServer_GetTenant_Call{Call: _e.mock.On("GetTenant", _a0, _a1)}
}

func (_c *MockTenantServiceServer_GetTenant_Call) Run(run func(_a0 context.Context, _a1 *Empty)) *MockTenantServiceServer_GetTenant_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*Empty))
	})
	return _c
}

func (_c *MockTenantServiceServer_GetTenant_Call) Return(_a0 *GetTenantResponse, _a1 error) *MockTenantServiceServer_GetTenant_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTenantTheme provides a mock function with given fields: _a0, _a1
func (_m *MockTenantServiceServer) GetTenantTheme(_a0 context.Context, _a1 *GetTenantThemeRequest) (*TenantThemeResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *TenantThemeResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetTenantThemeRequest) *TenantThemeResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TenantThemeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetTenantThemeRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceServer_GetTenantTheme_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTenantTheme'
type MockTenantServiceServer_GetTenantTheme_Call struct {
	*mock.Call
}

// GetTenantTheme is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *GetTenantThemeRequest
func (_e *MockTenantServiceServer_Expecter) GetTenantTheme(_a0 interface{}, _a1 interface{}) *MockTenantServiceServer_GetTenantTheme_Call {
	return &MockTenantServiceServer_GetTenantTheme_Call{Call: _e.mock.On("GetTenantTheme", _a0, _a1)}
}

func (_c *MockTenantServiceServer_GetTenantTheme_Call) Run(run func(_a0 context.Context, _a1 *GetTenantThemeRequest)) *MockTenantServiceServer_GetTenantTheme_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetTenantThemeRequest))
	})
	return _c
}

func (_c *MockTenantServiceServer_GetTenantTheme_Call) Return(_a0 *TenantThemeResponse, _a1 error) *MockTenantServiceServer_GetTenantTheme_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// LoadAuthInfo provides a mock function with given fields: _a0, _a1
func (_m *MockTenantServiceServer) LoadAuthInfo(_a0 context.Context, _a1 *LoadAuthRequest) (*LoadAuthResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *LoadAuthResponse
	if rf, ok := ret.Get(0).(func(context.Context, *LoadAuthRequest) *LoadAuthResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*LoadAuthResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *LoadAuthRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceServer_LoadAuthInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadAuthInfo'
type MockTenantServiceServer_LoadAuthInfo_Call struct {
	*mock.Call
}

// LoadAuthInfo is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *LoadAuthRequest
func (_e *MockTenantServiceServer_Expecter) LoadAuthInfo(_a0 interface{}, _a1 interface{}) *MockTenantServiceServer_LoadAuthInfo_Call {
	return &MockTenantServiceServer_LoadAuthInfo_Call{Call: _e.mock.On("LoadAuthInfo", _a0, _a1)}
}

func (_c *MockTenantServiceServer_LoadAuthInfo_Call) Run(run func(_a0 context.Context, _a1 *LoadAuthRequest)) *MockTenantServiceServer_LoadAuthInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*LoadAuthRequest))
	})
	return _c
}

func (_c *MockTenantServiceServer_LoadAuthInfo_Call) Return(_a0 *LoadAuthResponse, _a1 error) *MockTenantServiceServer_LoadAuthInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdateAuthInfo provides a mock function with given fields: _a0, _a1
func (_m *MockTenantServiceServer) UpdateAuthInfo(_a0 context.Context, _a1 *UpdateAuthRequest) (*UpdateAuthResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *UpdateAuthResponse
	if rf, ok := ret.Get(0).(func(context.Context, *UpdateAuthRequest) *UpdateAuthResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UpdateAuthResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *UpdateAuthRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceServer_UpdateAuthInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateAuthInfo'
type MockTenantServiceServer_UpdateAuthInfo_Call struct {
	*mock.Call
}

// UpdateAuthInfo is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *UpdateAuthRequest
func (_e *MockTenantServiceServer_Expecter) UpdateAuthInfo(_a0 interface{}, _a1 interface{}) *MockTenantServiceServer_UpdateAuthInfo_Call {
	return &MockTenantServiceServer_UpdateAuthInfo_Call{Call: _e.mock.On("UpdateAuthInfo", _a0, _a1)}
}

func (_c *MockTenantServiceServer_UpdateAuthInfo_Call) Run(run func(_a0 context.Context, _a1 *UpdateAuthRequest)) *MockTenantServiceServer_UpdateAuthInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*UpdateAuthRequest))
	})
	return _c
}

func (_c *MockTenantServiceServer_UpdateAuthInfo_Call) Return(_a0 *UpdateAuthResponse, _a1 error) *MockTenantServiceServer_UpdateAuthInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdateLandingPage provides a mock function with given fields: _a0, _a1
func (_m *MockTenantServiceServer) UpdateLandingPage(_a0 context.Context, _a1 *LandingPageResponse) (*LandingPageResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *LandingPageResponse
	if rf, ok := ret.Get(0).(func(context.Context, *LandingPageResponse) *LandingPageResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*LandingPageResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *LandingPageResponse) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceServer_UpdateLandingPage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateLandingPage'
type MockTenantServiceServer_UpdateLandingPage_Call struct {
	*mock.Call
}

// UpdateLandingPage is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *LandingPageResponse
func (_e *MockTenantServiceServer_Expecter) UpdateLandingPage(_a0 interface{}, _a1 interface{}) *MockTenantServiceServer_UpdateLandingPage_Call {
	return &MockTenantServiceServer_UpdateLandingPage_Call{Call: _e.mock.On("UpdateLandingPage", _a0, _a1)}
}

func (_c *MockTenantServiceServer_UpdateLandingPage_Call) Run(run func(_a0 context.Context, _a1 *LandingPageResponse)) *MockTenantServiceServer_UpdateLandingPage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*LandingPageResponse))
	})
	return _c
}

func (_c *MockTenantServiceServer_UpdateLandingPage_Call) Return(_a0 *LandingPageResponse, _a1 error) *MockTenantServiceServer_UpdateLandingPage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdateLoginMessage provides a mock function with given fields: _a0, _a1
func (_m *MockTenantServiceServer) UpdateLoginMessage(_a0 context.Context, _a1 *LoginMessageProto) (*LoginMessageProto, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *LoginMessageProto
	if rf, ok := ret.Get(0).(func(context.Context, *LoginMessageProto) *LoginMessageProto); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*LoginMessageProto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *LoginMessageProto) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceServer_UpdateLoginMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateLoginMessage'
type MockTenantServiceServer_UpdateLoginMessage_Call struct {
	*mock.Call
}

// UpdateLoginMessage is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *LoginMessageProto
func (_e *MockTenantServiceServer_Expecter) UpdateLoginMessage(_a0 interface{}, _a1 interface{}) *MockTenantServiceServer_UpdateLoginMessage_Call {
	return &MockTenantServiceServer_UpdateLoginMessage_Call{Call: _e.mock.On("UpdateLoginMessage", _a0, _a1)}
}

func (_c *MockTenantServiceServer_UpdateLoginMessage_Call) Run(run func(_a0 context.Context, _a1 *LoginMessageProto)) *MockTenantServiceServer_UpdateLoginMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*LoginMessageProto))
	})
	return _c
}

func (_c *MockTenantServiceServer_UpdateLoginMessage_Call) Return(_a0 *LoginMessageProto, _a1 error) *MockTenantServiceServer_UpdateLoginMessage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdateLogo provides a mock function with given fields: _a0, _a1
func (_m *MockTenantServiceServer) UpdateLogo(_a0 context.Context, _a1 *LogoResponse) (*LogoResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *LogoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *LogoResponse) *LogoResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*LogoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *LogoResponse) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceServer_UpdateLogo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateLogo'
type MockTenantServiceServer_UpdateLogo_Call struct {
	*mock.Call
}

// UpdateLogo is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *LogoResponse
func (_e *MockTenantServiceServer_Expecter) UpdateLogo(_a0 interface{}, _a1 interface{}) *MockTenantServiceServer_UpdateLogo_Call {
	return &MockTenantServiceServer_UpdateLogo_Call{Call: _e.mock.On("UpdateLogo", _a0, _a1)}
}

func (_c *MockTenantServiceServer_UpdateLogo_Call) Run(run func(_a0 context.Context, _a1 *LogoResponse)) *MockTenantServiceServer_UpdateLogo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*LogoResponse))
	})
	return _c
}

func (_c *MockTenantServiceServer_UpdateLogo_Call) Return(_a0 *LogoResponse, _a1 error) *MockTenantServiceServer_UpdateLogo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdateSessionSettings provides a mock function with given fields: _a0, _a1
func (_m *MockTenantServiceServer) UpdateSessionSettings(_a0 context.Context, _a1 *UpdateSessionSettingsRequest) (*UpdateSessionSettingsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *UpdateSessionSettingsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *UpdateSessionSettingsRequest) *UpdateSessionSettingsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*UpdateSessionSettingsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *UpdateSessionSettingsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceServer_UpdateSessionSettings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSessionSettings'
type MockTenantServiceServer_UpdateSessionSettings_Call struct {
	*mock.Call
}

// UpdateSessionSettings is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *UpdateSessionSettingsRequest
func (_e *MockTenantServiceServer_Expecter) UpdateSessionSettings(_a0 interface{}, _a1 interface{}) *MockTenantServiceServer_UpdateSessionSettings_Call {
	return &MockTenantServiceServer_UpdateSessionSettings_Call{Call: _e.mock.On("UpdateSessionSettings", _a0, _a1)}
}

func (_c *MockTenantServiceServer_UpdateSessionSettings_Call) Run(run func(_a0 context.Context, _a1 *UpdateSessionSettingsRequest)) *MockTenantServiceServer_UpdateSessionSettings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*UpdateSessionSettingsRequest))
	})
	return _c
}

func (_c *MockTenantServiceServer_UpdateSessionSettings_Call) Return(_a0 *UpdateSessionSettingsResponse, _a1 error) *MockTenantServiceServer_UpdateSessionSettings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdateTenantTheme provides a mock function with given fields: _a0, _a1
func (_m *MockTenantServiceServer) UpdateTenantTheme(_a0 context.Context, _a1 *TenantThemeResponse) (*TenantThemeResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *TenantThemeResponse
	if rf, ok := ret.Get(0).(func(context.Context, *TenantThemeResponse) *TenantThemeResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TenantThemeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *TenantThemeResponse) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTenantServiceServer_UpdateTenantTheme_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTenantTheme'
type MockTenantServiceServer_UpdateTenantTheme_Call struct {
	*mock.Call
}

// UpdateTenantTheme is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *TenantThemeResponse
func (_e *MockTenantServiceServer_Expecter) UpdateTenantTheme(_a0 interface{}, _a1 interface{}) *MockTenantServiceServer_UpdateTenantTheme_Call {
	return &MockTenantServiceServer_UpdateTenantTheme_Call{Call: _e.mock.On("UpdateTenantTheme", _a0, _a1)}
}

func (_c *MockTenantServiceServer_UpdateTenantTheme_Call) Run(run func(_a0 context.Context, _a1 *TenantThemeResponse)) *MockTenantServiceServer_UpdateTenantTheme_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*TenantThemeResponse))
	})
	return _c
}

func (_c *MockTenantServiceServer_UpdateTenantTheme_Call) Return(_a0 *TenantThemeResponse, _a1 error) *MockTenantServiceServer_UpdateTenantTheme_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// mustEmbedUnimplementedTenantServiceServer provides a mock function with given fields:
func (_m *MockTenantServiceServer) mustEmbedUnimplementedTenantServiceServer() {
	_m.Called()
}

// MockTenantServiceServer_mustEmbedUnimplementedTenantServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedTenantServiceServer'
type MockTenantServiceServer_mustEmbedUnimplementedTenantServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedTenantServiceServer is a helper method to define mock.On call
func (_e *MockTenantServiceServer_Expecter) mustEmbedUnimplementedTenantServiceServer() *MockTenantServiceServer_mustEmbedUnimplementedTenantServiceServer_Call {
	return &MockTenantServiceServer_mustEmbedUnimplementedTenantServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedTenantServiceServer")}
}

func (_c *MockTenantServiceServer_mustEmbedUnimplementedTenantServiceServer_Call) Run(run func()) *MockTenantServiceServer_mustEmbedUnimplementedTenantServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTenantServiceServer_mustEmbedUnimplementedTenantServiceServer_Call) Return() *MockTenantServiceServer_mustEmbedUnimplementedTenantServiceServer_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewMockTenantServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockTenantServiceServer creates a new instance of MockTenantServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockTenantServiceServer(t mockConstructorTestingTNewMockTenantServiceServer) *MockTenantServiceServer {
	mock := &MockTenantServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
