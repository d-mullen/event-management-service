// Code generated by mockery v1.0.0. DO NOT EDIT.

package event_management

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockEventManagementClient is an autogenerated mock type for the EventManagementClient type
type MockEventManagementClient struct {
	mock.Mock
}

// Annotate provides a mock function with given fields: ctx, in, opts
func (_m *MockEventManagementClient) Annotate(ctx context.Context, in *EventAnnotationRequest, opts ...grpc.CallOption) (*EventAnnotationResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *EventAnnotationResponse
	if rf, ok := ret.Get(0).(func(context.Context, *EventAnnotationRequest, ...grpc.CallOption) *EventAnnotationResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*EventAnnotationResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *EventAnnotationRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllowedStates provides a mock function with given fields: ctx, in, opts
func (_m *MockEventManagementClient) GetAllowedStates(ctx context.Context, in *EventAllowedStatesRequest, opts ...grpc.CallOption) (*EventAllowedStatesResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *EventAllowedStatesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *EventAllowedStatesRequest, ...grpc.CallOption) *EventAllowedStatesResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*EventAllowedStatesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *EventAllowedStatesRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetAcknowledge provides a mock function with given fields: ctx, in, opts
func (_m *MockEventManagementClient) SetAcknowledge(ctx context.Context, in *EventAcknowledgeRequest, opts ...grpc.CallOption) (*EventManagementResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *EventManagementResponse
	if rf, ok := ret.Get(0).(func(context.Context, *EventAcknowledgeRequest, ...grpc.CallOption) *EventManagementResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*EventManagementResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *EventAcknowledgeRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetStatus provides a mock function with given fields: ctx, in, opts
func (_m *MockEventManagementClient) SetStatus(ctx context.Context, in *EventStatusRequest, opts ...grpc.CallOption) (*EventManagementResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *EventManagementResponse
	if rf, ok := ret.Get(0).(func(context.Context, *EventStatusRequest, ...grpc.CallOption) *EventManagementResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*EventManagementResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *EventStatusRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
