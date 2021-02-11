// Code generated by mockery v1.0.0. DO NOT EDIT.

package data_receiver

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockDataReceiverServiceClient is an autogenerated mock type for the DataReceiverServiceClient type
type MockDataReceiverServiceClient struct {
	mock.Mock
}

// PutEvent provides a mock function with given fields: ctx, opts
func (_m *MockDataReceiverServiceClient) PutEvent(ctx context.Context, opts ...grpc.CallOption) (DataReceiverService_PutEventClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 DataReceiverService_PutEventClient
	if rf, ok := ret.Get(0).(func(context.Context, ...grpc.CallOption) DataReceiverService_PutEventClient); ok {
		r0 = rf(ctx, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(DataReceiverService_PutEventClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutEvents provides a mock function with given fields: ctx, in, opts
func (_m *MockDataReceiverServiceClient) PutEvents(ctx context.Context, in *Events, opts ...grpc.CallOption) (*EventStatusResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *EventStatusResult
	if rf, ok := ret.Get(0).(func(context.Context, *Events, ...grpc.CallOption) *EventStatusResult); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*EventStatusResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *Events, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutMetric provides a mock function with given fields: ctx, opts
func (_m *MockDataReceiverServiceClient) PutMetric(ctx context.Context, opts ...grpc.CallOption) (DataReceiverService_PutMetricClient, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 DataReceiverService_PutMetricClient
	if rf, ok := ret.Get(0).(func(context.Context, ...grpc.CallOption) DataReceiverService_PutMetricClient); ok {
		r0 = rf(ctx, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(DataReceiverService_PutMetricClient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutMetrics provides a mock function with given fields: ctx, in, opts
func (_m *MockDataReceiverServiceClient) PutMetrics(ctx context.Context, in *Metrics, opts ...grpc.CallOption) (*StatusResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *StatusResult
	if rf, ok := ret.Get(0).(func(context.Context, *Metrics, ...grpc.CallOption) *StatusResult); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*StatusResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *Metrics, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutModels provides a mock function with given fields: ctx, in, opts
func (_m *MockDataReceiverServiceClient) PutModels(ctx context.Context, in *Models, opts ...grpc.CallOption) (*ModelStatusResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *ModelStatusResult
	if rf, ok := ret.Get(0).(func(context.Context, *Models, ...grpc.CallOption) *ModelStatusResult); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ModelStatusResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *Models, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
