// Code generated by mockery v2.14.0. DO NOT EDIT.

package yamr

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockYamrQueryServer is an autogenerated mock type for the YamrQueryServer type
type MockYamrQueryServer struct {
	mock.Mock
}

type MockYamrQueryServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockYamrQueryServer) EXPECT() *MockYamrQueryServer_Expecter {
	return &MockYamrQueryServer_Expecter{mock: &_m.Mock}
}

// Count provides a mock function with given fields: _a0, _a1
func (_m *MockYamrQueryServer) Count(_a0 context.Context, _a1 *CountRequest) (*CountResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *CountResponse
	if rf, ok := ret.Get(0).(func(context.Context, *CountRequest) *CountResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*CountResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *CountRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockYamrQueryServer_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type MockYamrQueryServer_Count_Call struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *CountRequest
func (_e *MockYamrQueryServer_Expecter) Count(_a0 interface{}, _a1 interface{}) *MockYamrQueryServer_Count_Call {
	return &MockYamrQueryServer_Count_Call{Call: _e.mock.On("Count", _a0, _a1)}
}

func (_c *MockYamrQueryServer_Count_Call) Run(run func(_a0 context.Context, _a1 *CountRequest)) *MockYamrQueryServer_Count_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*CountRequest))
	})
	return _c
}

func (_c *MockYamrQueryServer_Count_Call) Return(_a0 *CountResponse, _a1 error) *MockYamrQueryServer_Count_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Frequency provides a mock function with given fields: _a0, _a1
func (_m *MockYamrQueryServer) Frequency(_a0 context.Context, _a1 *FrequencyRequest) (*FrequencyResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *FrequencyResponse
	if rf, ok := ret.Get(0).(func(context.Context, *FrequencyRequest) *FrequencyResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*FrequencyResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *FrequencyRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockYamrQueryServer_Frequency_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Frequency'
type MockYamrQueryServer_Frequency_Call struct {
	*mock.Call
}

// Frequency is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *FrequencyRequest
func (_e *MockYamrQueryServer_Expecter) Frequency(_a0 interface{}, _a1 interface{}) *MockYamrQueryServer_Frequency_Call {
	return &MockYamrQueryServer_Frequency_Call{Call: _e.mock.On("Frequency", _a0, _a1)}
}

func (_c *MockYamrQueryServer_Frequency_Call) Run(run func(_a0 context.Context, _a1 *FrequencyRequest)) *MockYamrQueryServer_Frequency_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*FrequencyRequest))
	})
	return _c
}

func (_c *MockYamrQueryServer_Frequency_Call) Return(_a0 *FrequencyResponse, _a1 error) *MockYamrQueryServer_Frequency_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Get provides a mock function with given fields: _a0, _a1
func (_m *MockYamrQueryServer) Get(_a0 context.Context, _a1 *GetRequest) (*GetResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *GetResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetRequest) *GetResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockYamrQueryServer_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockYamrQueryServer_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *GetRequest
func (_e *MockYamrQueryServer_Expecter) Get(_a0 interface{}, _a1 interface{}) *MockYamrQueryServer_Get_Call {
	return &MockYamrQueryServer_Get_Call{Call: _e.mock.On("Get", _a0, _a1)}
}

func (_c *MockYamrQueryServer_Get_Call) Run(run func(_a0 context.Context, _a1 *GetRequest)) *MockYamrQueryServer_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetRequest))
	})
	return _c
}

func (_c *MockYamrQueryServer_Get_Call) Return(_a0 *GetResponse, _a1 error) *MockYamrQueryServer_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetBulk provides a mock function with given fields: _a0, _a1
func (_m *MockYamrQueryServer) GetBulk(_a0 context.Context, _a1 *GetBulkRequest) (*GetBulkResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *GetBulkResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetBulkRequest) *GetBulkResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetBulkResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetBulkRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockYamrQueryServer_GetBulk_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBulk'
type MockYamrQueryServer_GetBulk_Call struct {
	*mock.Call
}

// GetBulk is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *GetBulkRequest
func (_e *MockYamrQueryServer_Expecter) GetBulk(_a0 interface{}, _a1 interface{}) *MockYamrQueryServer_GetBulk_Call {
	return &MockYamrQueryServer_GetBulk_Call{Call: _e.mock.On("GetBulk", _a0, _a1)}
}

func (_c *MockYamrQueryServer_GetBulk_Call) Run(run func(_a0 context.Context, _a1 *GetBulkRequest)) *MockYamrQueryServer_GetBulk_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetBulkRequest))
	})
	return _c
}

func (_c *MockYamrQueryServer_GetBulk_Call) Return(_a0 *GetBulkResponse, _a1 error) *MockYamrQueryServer_GetBulk_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Search provides a mock function with given fields: _a0, _a1
func (_m *MockYamrQueryServer) Search(_a0 context.Context, _a1 *SearchRequest) (*SearchResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *SearchResponse
	if rf, ok := ret.Get(0).(func(context.Context, *SearchRequest) *SearchResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*SearchResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *SearchRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockYamrQueryServer_Search_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Search'
type MockYamrQueryServer_Search_Call struct {
	*mock.Call
}

// Search is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *SearchRequest
func (_e *MockYamrQueryServer_Expecter) Search(_a0 interface{}, _a1 interface{}) *MockYamrQueryServer_Search_Call {
	return &MockYamrQueryServer_Search_Call{Call: _e.mock.On("Search", _a0, _a1)}
}

func (_c *MockYamrQueryServer_Search_Call) Run(run func(_a0 context.Context, _a1 *SearchRequest)) *MockYamrQueryServer_Search_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*SearchRequest))
	})
	return _c
}

func (_c *MockYamrQueryServer_Search_Call) Return(_a0 *SearchResponse, _a1 error) *MockYamrQueryServer_Search_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// StreamingSearch provides a mock function with given fields: _a0, _a1
func (_m *MockYamrQueryServer) StreamingSearch(_a0 *SearchRequest, _a1 YamrQuery_StreamingSearchServer) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*SearchRequest, YamrQuery_StreamingSearchServer) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockYamrQueryServer_StreamingSearch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StreamingSearch'
type MockYamrQueryServer_StreamingSearch_Call struct {
	*mock.Call
}

// StreamingSearch is a helper method to define mock.On call
//  - _a0 *SearchRequest
//  - _a1 YamrQuery_StreamingSearchServer
func (_e *MockYamrQueryServer_Expecter) StreamingSearch(_a0 interface{}, _a1 interface{}) *MockYamrQueryServer_StreamingSearch_Call {
	return &MockYamrQueryServer_StreamingSearch_Call{Call: _e.mock.On("StreamingSearch", _a0, _a1)}
}

func (_c *MockYamrQueryServer_StreamingSearch_Call) Run(run func(_a0 *SearchRequest, _a1 YamrQuery_StreamingSearchServer)) *MockYamrQueryServer_StreamingSearch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*SearchRequest), args[1].(YamrQuery_StreamingSearchServer))
	})
	return _c
}

func (_c *MockYamrQueryServer_StreamingSearch_Call) Return(_a0 error) *MockYamrQueryServer_StreamingSearch_Call {
	_c.Call.Return(_a0)
	return _c
}

// Suggest provides a mock function with given fields: _a0, _a1
func (_m *MockYamrQueryServer) Suggest(_a0 context.Context, _a1 *SuggestRequest) (*SuggestResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *SuggestResponse
	if rf, ok := ret.Get(0).(func(context.Context, *SuggestRequest) *SuggestResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*SuggestResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *SuggestRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockYamrQueryServer_Suggest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Suggest'
type MockYamrQueryServer_Suggest_Call struct {
	*mock.Call
}

// Suggest is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *SuggestRequest
func (_e *MockYamrQueryServer_Expecter) Suggest(_a0 interface{}, _a1 interface{}) *MockYamrQueryServer_Suggest_Call {
	return &MockYamrQueryServer_Suggest_Call{Call: _e.mock.On("Suggest", _a0, _a1)}
}

func (_c *MockYamrQueryServer_Suggest_Call) Run(run func(_a0 context.Context, _a1 *SuggestRequest)) *MockYamrQueryServer_Suggest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*SuggestRequest))
	})
	return _c
}

func (_c *MockYamrQueryServer_Suggest_Call) Return(_a0 *SuggestResponse, _a1 error) *MockYamrQueryServer_Suggest_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// mustEmbedUnimplementedYamrQueryServer provides a mock function with given fields:
func (_m *MockYamrQueryServer) mustEmbedUnimplementedYamrQueryServer() {
	_m.Called()
}

// MockYamrQueryServer_mustEmbedUnimplementedYamrQueryServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedYamrQueryServer'
type MockYamrQueryServer_mustEmbedUnimplementedYamrQueryServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedYamrQueryServer is a helper method to define mock.On call
func (_e *MockYamrQueryServer_Expecter) mustEmbedUnimplementedYamrQueryServer() *MockYamrQueryServer_mustEmbedUnimplementedYamrQueryServer_Call {
	return &MockYamrQueryServer_mustEmbedUnimplementedYamrQueryServer_Call{Call: _e.mock.On("mustEmbedUnimplementedYamrQueryServer")}
}

func (_c *MockYamrQueryServer_mustEmbedUnimplementedYamrQueryServer_Call) Run(run func()) *MockYamrQueryServer_mustEmbedUnimplementedYamrQueryServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockYamrQueryServer_mustEmbedUnimplementedYamrQueryServer_Call) Return() *MockYamrQueryServer_mustEmbedUnimplementedYamrQueryServer_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewMockYamrQueryServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockYamrQueryServer creates a new instance of MockYamrQueryServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockYamrQueryServer(t mockConstructorTestingTNewMockYamrQueryServer) *MockYamrQueryServer {
	mock := &MockYamrQueryServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
