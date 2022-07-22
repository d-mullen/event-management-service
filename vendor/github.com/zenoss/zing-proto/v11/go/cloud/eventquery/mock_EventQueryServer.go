// Code generated by mockery v2.12.2. DO NOT EDIT.

package eventquery

import (
	context "context"
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// MockEventQueryServer is an autogenerated mock type for the EventQueryServer type
type MockEventQueryServer struct {
	mock.Mock
}

type MockEventQueryServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEventQueryServer) EXPECT() *MockEventQueryServer_Expecter {
	return &MockEventQueryServer_Expecter{mock: &_m.Mock}
}

// Count provides a mock function with given fields: _a0, _a1
func (_m *MockEventQueryServer) Count(_a0 context.Context, _a1 *CountRequest) (*CountResponse, error) {
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

// MockEventQueryServer_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type MockEventQueryServer_Count_Call struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *CountRequest
func (_e *MockEventQueryServer_Expecter) Count(_a0 interface{}, _a1 interface{}) *MockEventQueryServer_Count_Call {
	return &MockEventQueryServer_Count_Call{Call: _e.mock.On("Count", _a0, _a1)}
}

func (_c *MockEventQueryServer_Count_Call) Run(run func(_a0 context.Context, _a1 *CountRequest)) *MockEventQueryServer_Count_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*CountRequest))
	})
	return _c
}

func (_c *MockEventQueryServer_Count_Call) Return(_a0 *CountResponse, _a1 error) *MockEventQueryServer_Count_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// EventsWithCountsStream provides a mock function with given fields: _a0, _a1
func (_m *MockEventQueryServer) EventsWithCountsStream(_a0 *EventsWithCountsRequest, _a1 EventQuery_EventsWithCountsStreamServer) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*EventsWithCountsRequest, EventQuery_EventsWithCountsStreamServer) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventQueryServer_EventsWithCountsStream_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EventsWithCountsStream'
type MockEventQueryServer_EventsWithCountsStream_Call struct {
	*mock.Call
}

// EventsWithCountsStream is a helper method to define mock.On call
//  - _a0 *EventsWithCountsRequest
//  - _a1 EventQuery_EventsWithCountsStreamServer
func (_e *MockEventQueryServer_Expecter) EventsWithCountsStream(_a0 interface{}, _a1 interface{}) *MockEventQueryServer_EventsWithCountsStream_Call {
	return &MockEventQueryServer_EventsWithCountsStream_Call{Call: _e.mock.On("EventsWithCountsStream", _a0, _a1)}
}

func (_c *MockEventQueryServer_EventsWithCountsStream_Call) Run(run func(_a0 *EventsWithCountsRequest, _a1 EventQuery_EventsWithCountsStreamServer)) *MockEventQueryServer_EventsWithCountsStream_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*EventsWithCountsRequest), args[1].(EventQuery_EventsWithCountsStreamServer))
	})
	return _c
}

func (_c *MockEventQueryServer_EventsWithCountsStream_Call) Return(_a0 error) *MockEventQueryServer_EventsWithCountsStream_Call {
	_c.Call.Return(_a0)
	return _c
}

// Frequency provides a mock function with given fields: _a0, _a1
func (_m *MockEventQueryServer) Frequency(_a0 context.Context, _a1 *FrequencyRequest) (*FrequencyResponse, error) {
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

// MockEventQueryServer_Frequency_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Frequency'
type MockEventQueryServer_Frequency_Call struct {
	*mock.Call
}

// Frequency is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *FrequencyRequest
func (_e *MockEventQueryServer_Expecter) Frequency(_a0 interface{}, _a1 interface{}) *MockEventQueryServer_Frequency_Call {
	return &MockEventQueryServer_Frequency_Call{Call: _e.mock.On("Frequency", _a0, _a1)}
}

func (_c *MockEventQueryServer_Frequency_Call) Run(run func(_a0 context.Context, _a1 *FrequencyRequest)) *MockEventQueryServer_Frequency_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*FrequencyRequest))
	})
	return _c
}

func (_c *MockEventQueryServer_Frequency_Call) Return(_a0 *FrequencyResponse, _a1 error) *MockEventQueryServer_Frequency_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetEvent provides a mock function with given fields: _a0, _a1
func (_m *MockEventQueryServer) GetEvent(_a0 context.Context, _a1 *GetEventRequest) (*GetEventResponse, error) {
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

// MockEventQueryServer_GetEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEvent'
type MockEventQueryServer_GetEvent_Call struct {
	*mock.Call
}

// GetEvent is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *GetEventRequest
func (_e *MockEventQueryServer_Expecter) GetEvent(_a0 interface{}, _a1 interface{}) *MockEventQueryServer_GetEvent_Call {
	return &MockEventQueryServer_GetEvent_Call{Call: _e.mock.On("GetEvent", _a0, _a1)}
}

func (_c *MockEventQueryServer_GetEvent_Call) Run(run func(_a0 context.Context, _a1 *GetEventRequest)) *MockEventQueryServer_GetEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetEventRequest))
	})
	return _c
}

func (_c *MockEventQueryServer_GetEvent_Call) Return(_a0 *GetEventResponse, _a1 error) *MockEventQueryServer_GetEvent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetEvents provides a mock function with given fields: _a0, _a1
func (_m *MockEventQueryServer) GetEvents(_a0 context.Context, _a1 *GetEventsRequest) (*GetEventsResponse, error) {
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

// MockEventQueryServer_GetEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetEvents'
type MockEventQueryServer_GetEvents_Call struct {
	*mock.Call
}

// GetEvents is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *GetEventsRequest
func (_e *MockEventQueryServer_Expecter) GetEvents(_a0 interface{}, _a1 interface{}) *MockEventQueryServer_GetEvents_Call {
	return &MockEventQueryServer_GetEvents_Call{Call: _e.mock.On("GetEvents", _a0, _a1)}
}

func (_c *MockEventQueryServer_GetEvents_Call) Run(run func(_a0 context.Context, _a1 *GetEventsRequest)) *MockEventQueryServer_GetEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*GetEventsRequest))
	})
	return _c
}

func (_c *MockEventQueryServer_GetEvents_Call) Return(_a0 *GetEventsResponse, _a1 error) *MockEventQueryServer_GetEvents_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Search provides a mock function with given fields: _a0, _a1
func (_m *MockEventQueryServer) Search(_a0 context.Context, _a1 *SearchRequest) (*SearchResponse, error) {
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

// MockEventQueryServer_Search_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Search'
type MockEventQueryServer_Search_Call struct {
	*mock.Call
}

// Search is a helper method to define mock.On call
//  - _a0 context.Context
//  - _a1 *SearchRequest
func (_e *MockEventQueryServer_Expecter) Search(_a0 interface{}, _a1 interface{}) *MockEventQueryServer_Search_Call {
	return &MockEventQueryServer_Search_Call{Call: _e.mock.On("Search", _a0, _a1)}
}

func (_c *MockEventQueryServer_Search_Call) Run(run func(_a0 context.Context, _a1 *SearchRequest)) *MockEventQueryServer_Search_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*SearchRequest))
	})
	return _c
}

func (_c *MockEventQueryServer_Search_Call) Return(_a0 *SearchResponse, _a1 error) *MockEventQueryServer_Search_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// SearchStream provides a mock function with given fields: _a0, _a1
func (_m *MockEventQueryServer) SearchStream(_a0 *SearchStreamRequest, _a1 EventQuery_SearchStreamServer) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*SearchStreamRequest, EventQuery_SearchStreamServer) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEventQueryServer_SearchStream_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchStream'
type MockEventQueryServer_SearchStream_Call struct {
	*mock.Call
}

// SearchStream is a helper method to define mock.On call
//  - _a0 *SearchStreamRequest
//  - _a1 EventQuery_SearchStreamServer
func (_e *MockEventQueryServer_Expecter) SearchStream(_a0 interface{}, _a1 interface{}) *MockEventQueryServer_SearchStream_Call {
	return &MockEventQueryServer_SearchStream_Call{Call: _e.mock.On("SearchStream", _a0, _a1)}
}

func (_c *MockEventQueryServer_SearchStream_Call) Run(run func(_a0 *SearchStreamRequest, _a1 EventQuery_SearchStreamServer)) *MockEventQueryServer_SearchStream_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*SearchStreamRequest), args[1].(EventQuery_SearchStreamServer))
	})
	return _c
}

func (_c *MockEventQueryServer_SearchStream_Call) Return(_a0 error) *MockEventQueryServer_SearchStream_Call {
	_c.Call.Return(_a0)
	return _c
}

// mustEmbedUnimplementedEventQueryServer provides a mock function with given fields:
func (_m *MockEventQueryServer) mustEmbedUnimplementedEventQueryServer() {
	_m.Called()
}

// MockEventQueryServer_mustEmbedUnimplementedEventQueryServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedEventQueryServer'
type MockEventQueryServer_mustEmbedUnimplementedEventQueryServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedEventQueryServer is a helper method to define mock.On call
func (_e *MockEventQueryServer_Expecter) mustEmbedUnimplementedEventQueryServer() *MockEventQueryServer_mustEmbedUnimplementedEventQueryServer_Call {
	return &MockEventQueryServer_mustEmbedUnimplementedEventQueryServer_Call{Call: _e.mock.On("mustEmbedUnimplementedEventQueryServer")}
}

func (_c *MockEventQueryServer_mustEmbedUnimplementedEventQueryServer_Call) Run(run func()) *MockEventQueryServer_mustEmbedUnimplementedEventQueryServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEventQueryServer_mustEmbedUnimplementedEventQueryServer_Call) Return() *MockEventQueryServer_mustEmbedUnimplementedEventQueryServer_Call {
	_c.Call.Return()
	return _c
}

// NewMockEventQueryServer creates a new instance of MockEventQueryServer. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEventQueryServer(t testing.TB) *MockEventQueryServer {
	mock := &MockEventQueryServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
