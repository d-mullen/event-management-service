// Code generated by mockery v1.0.0. DO NOT EDIT.

package event_context

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockEventContextQueryClient is an autogenerated mock type for the EventContextQueryClient type
type MockEventContextQueryClient struct {
	mock.Mock
}

// GetActiveEvents provides a mock function with given fields: ctx, in, opts
func (_m *MockEventContextQueryClient) GetActiveEvents(ctx context.Context, in *ECGetActiveEventsRequest, opts ...grpc.CallOption) (*ECGetActiveEventsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ECGetActiveEventsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ECGetActiveEventsRequest, ...grpc.CallOption) *ECGetActiveEventsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ECGetActiveEventsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ECGetActiveEventsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBulk provides a mock function with given fields: ctx, in, opts
func (_m *MockEventContextQueryClient) GetBulk(ctx context.Context, in *ECGetBulkRequest, opts ...grpc.CallOption) (*ECGetBulkResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ECGetBulkResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ECGetBulkRequest, ...grpc.CallOption) *ECGetBulkResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ECGetBulkResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ECGetBulkRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: ctx, in, opts
func (_m *MockEventContextQueryClient) Search(ctx context.Context, in *ECSearchRequest, opts ...grpc.CallOption) (*ECSearchResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ECSearchResponse
	if rf, ok := ret.Get(0).(func(context.Context, *ECSearchRequest, ...grpc.CallOption) *ECSearchResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ECSearchResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *ECSearchRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}