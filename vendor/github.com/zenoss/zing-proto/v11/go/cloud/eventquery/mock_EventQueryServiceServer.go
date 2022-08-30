// Code generated by mockery v2.14.0. DO NOT EDIT.

package eventquery

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockEventQueryServiceServer is an autogenerated mock type for the EventQueryServiceServer type
type MockEventQueryServiceServer struct {
	mock.Mock
}

type MockEventQueryServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEventQueryServiceServer) EXPECT() *MockEventQueryServiceServer_Expecter {
	return &MockEventQueryServiceServer_Expecter{mock: &_m.Mock}
}

// Count provides a mock function with given fields: _a0, _a1
func (_m *MockEventQueryServiceServer) Count(_a0 context.Context, _a1 *CountRequest) (*CountResponse, error) {
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

// MockEventQueryServiceServer_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type MockEventQueryServiceServer_Count_Call struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *CountRequest
func (_e *MockEventQueryServiceServer_Expecter) Count(_a0 interface{}, _a1 interface{}) *MockEventQueryServiceServer_Count_Call {
	return &MockEventQueryServiceServer_Count_Call{Call: _e.mock.On("Count", _a0, _a1)}
}

func (_c *MockEventQueryServiceServer_Count_Call) Run(run func(_a0 context.Context, _a1 *CountRequest)) *MockEventQueryServiceServer_Count_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*CountRequest))
	})
	return _c
}

func (_c *MockEventQueryServiceServer_Count_Call) Return(_a0 *CountResponse, _a1 error) *MockEventQueryServiceServer_Count_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// EventsWithCountsStream provides a mock function with given fields: _a0, _a1
func (_m *MockEventQueryServiceServer) EventsWithCountsStream(_a0 *EventsWithCountsStreamRequest, _a1 EventQueryService_EventsWithCountsStreamServer) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*EventsWithCountsStreamRequest, EventQueryService_EventsWithCountsStreamServer) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventQueryServiceServer_EventsWithCountsStream_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EventsWithCountsStream'
type MockEventQueryServiceServer_EventsWithCountsStream_Call struct {
	*mock.Call
}

// EventsWithCountsStream is a helper method to define mock.On call
//  - _a0 *EventsWithCountsStreamRequest
//  - _a1 EventQueryService_EventsWithCountsStreamServer
func (_e *MockEventQueryServiceServer_Expecter) EventsWithCountsStream(_a0 interface{}, _a1 interface{}) *MockEventQueryServiceServer_EventsWithCountsStream_Call {
	return &MockEventQueryServiceServer_EventsWithCountsStream_Call{Call: _e.mock.On("EventsWithCountsStream", _a0, _a1)}
}

func (_c *MockEventQueryServiceServer_EventsWithCountsStream_Call) Run(run func(_a0 *EventsWithCountsStreamRequest, _a1 EventQueryService_EventsWithCountsStreamServer)) *MockEventQueryServiceServer_EventsWithCountsStream_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*EventsWithCountsStreamRequest), args[1].(EventQueryService_EventsWithCountsStreamServer))
	})
	return _c
}

func (_c *MockEventQueryServiceServer_EventsWithCountsStream_Call) Return(_a0 error) *MockEventQueryServiceServer_EventsWithCountsStream_Call {
	_c.Call.Return(_a0)
	return _c
}

// Frequency provides a mock function with given fields: _a0, _a1
func (_m *MockEventQueryServiceServer) Frequency(_a0 context.Context, _a1 *FrequencyRequest) (*FrequencyResponse, error) {
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

// MockEventQueryServiceServer_Frequency_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Frequency'
type MockEventQueryServiceServer_Frequency_Call struct {
	*mock.Call
}

// Frequency is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *FrequencyRequest
func (_e *MockEventQueryServiceServer_Expecter) Frequency(_a0 interface{}, _a1 interface{}) *MockEventQueryServiceServer_Frequency_Call {
	return &MockEventQueryServiceServer_Frequency_Call{Call: _e.mock.On("Frequency", _a0, _a1)}
}

func (_c *MockEventQueryServiceServer_Frequency_Call) Run(run func(_a0 context.Context, _a1 *FrequencyRequest)) *MockEventQueryServiceServer_Frequency_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*FrequencyRequest))
	})
	return _c
}

func (_c *MockEventQueryServiceServer_Frequency_Call) Return(_a0 *FrequencyResponse, _a1 error) *MockEventQueryServiceServer_Frequency_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetEvent provides a mock function with given fields: _a0, _a1
func (_m *MockEventQueryServiceServer) GetEvent(_a0 context.Context, _a1 *GetEventRequest) (*GetEventResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *GetEventResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetEventRequest) *GetEventResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetEventResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetEventRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEventQueryServiceServer_GetEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEvent'
type MockEventQueryServiceServer_GetEvent_Call struct {
	*mock.Call
}

// GetEvent is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *GetEventRequest
func (_e *MockEventQueryServiceServer_Expecter) GetEvent(_a0 interface{}, _a1 interface{}) *MockEventQueryServiceServer_GetEvent_Call {
	return &MockEventQueryServiceServer_GetEvent_Call{Call: _e.mock.On("GetEvent", _a0, _a1)}
}

func (_c *MockEventQueryServiceServer_GetEvent_Call) Run(run func(_a0 context.Context, _a1 *GetEventRequest)) *MockEventQueryServiceServer_GetEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetEventRequest))
	})
	return _c
}

func (_c *MockEventQueryServiceServer_GetEvent_Call) Return(_a0 *GetEventResponse, _a1 error) *MockEventQueryServiceServer_GetEvent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetEvents provides a mock function with given fields: _a0, _a1
func (_m *MockEventQueryServiceServer) GetEvents(_a0 context.Context, _a1 *GetEventsRequest) (*GetEventsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *GetEventsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *GetEventsRequest) *GetEventsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetEventsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *GetEventsRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockEventQueryServiceServer_GetEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEvents'
type MockEventQueryServiceServer_GetEvents_Call struct {
	*mock.Call
}

// GetEvents is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *GetEventsRequest
func (_e *MockEventQueryServiceServer_Expecter) GetEvents(_a0 interface{}, _a1 interface{}) *MockEventQueryServiceServer_GetEvents_Call {
	return &MockEventQueryServiceServer_GetEvents_Call{Call: _e.mock.On("GetEvents", _a0, _a1)}
}

func (_c *MockEventQueryServiceServer_GetEvents_Call) Run(run func(_a0 context.Context, _a1 *GetEventsRequest)) *MockEventQueryServiceServer_GetEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetEventsRequest))
	})
	return _c
}

func (_c *MockEventQueryServiceServer_GetEvents_Call) Return(_a0 *GetEventsResponse, _a1 error) *MockEventQueryServiceServer_GetEvents_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Search provides a mock function with given fields: _a0, _a1
func (_m *MockEventQueryServiceServer) Search(_a0 context.Context, _a1 *SearchRequest) (*SearchResponse, error) {
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

// MockEventQueryServiceServer_Search_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Search'
type MockEventQueryServiceServer_Search_Call struct {
	*mock.Call
}

// Search is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *SearchRequest
func (_e *MockEventQueryServiceServer_Expecter) Search(_a0 interface{}, _a1 interface{}) *MockEventQueryServiceServer_Search_Call {
	return &MockEventQueryServiceServer_Search_Call{Call: _e.mock.On("Search", _a0, _a1)}
}

func (_c *MockEventQueryServiceServer_Search_Call) Run(run func(_a0 context.Context, _a1 *SearchRequest)) *MockEventQueryServiceServer_Search_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*SearchRequest))
	})
	return _c
}

func (_c *MockEventQueryServiceServer_Search_Call) Return(_a0 *SearchResponse, _a1 error) *MockEventQueryServiceServer_Search_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// SearchStream provides a mock function with given fields: _a0, _a1
func (_m *MockEventQueryServiceServer) SearchStream(_a0 *SearchStreamRequest, _a1 EventQueryService_SearchStreamServer) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*SearchStreamRequest, EventQueryService_SearchStreamServer) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventQueryServiceServer_SearchStream_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchStream'
type MockEventQueryServiceServer_SearchStream_Call struct {
	*mock.Call
}

// SearchStream is a helper method to define mock.On call
//  - _a0 *SearchStreamRequest
//  - _a1 EventQueryService_SearchStreamServer
func (_e *MockEventQueryServiceServer_Expecter) SearchStream(_a0 interface{}, _a1 interface{}) *MockEventQueryServiceServer_SearchStream_Call {
	return &MockEventQueryServiceServer_SearchStream_Call{Call: _e.mock.On("SearchStream", _a0, _a1)}
}

func (_c *MockEventQueryServiceServer_SearchStream_Call) Run(run func(_a0 *SearchStreamRequest, _a1 EventQueryService_SearchStreamServer)) *MockEventQueryServiceServer_SearchStream_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*SearchStreamRequest), args[1].(EventQueryService_SearchStreamServer))
	})
	return _c
}

func (_c *MockEventQueryServiceServer_SearchStream_Call) Return(_a0 error) *MockEventQueryServiceServer_SearchStream_Call {
	_c.Call.Return(_a0)
	return _c
}

// mustEmbedUnimplementedEventQueryServiceServer provides a mock function with given fields:
func (_m *MockEventQueryServiceServer) mustEmbedUnimplementedEventQueryServiceServer() {
	_m.Called()
}

// MockEventQueryServiceServer_mustEmbedUnimplementedEventQueryServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedEventQueryServiceServer'
type MockEventQueryServiceServer_mustEmbedUnimplementedEventQueryServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedEventQueryServiceServer is a helper method to define mock.On call
func (_e *MockEventQueryServiceServer_Expecter) mustEmbedUnimplementedEventQueryServiceServer() *MockEventQueryServiceServer_mustEmbedUnimplementedEventQueryServiceServer_Call {
	return &MockEventQueryServiceServer_mustEmbedUnimplementedEventQueryServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedEventQueryServiceServer")}
}

func (_c *MockEventQueryServiceServer_mustEmbedUnimplementedEventQueryServiceServer_Call) Run(run func()) *MockEventQueryServiceServer_mustEmbedUnimplementedEventQueryServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEventQueryServiceServer_mustEmbedUnimplementedEventQueryServiceServer_Call) Return() *MockEventQueryServiceServer_mustEmbedUnimplementedEventQueryServiceServer_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewMockEventQueryServiceServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEventQueryServiceServer creates a new instance of MockEventQueryServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEventQueryServiceServer(t mockConstructorTestingTNewMockEventQueryServiceServer) *MockEventQueryServiceServer {
	mock := &MockEventQueryServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
