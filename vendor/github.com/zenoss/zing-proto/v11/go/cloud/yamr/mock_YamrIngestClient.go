// Code generated by mockery v1.0.0. DO NOT EDIT.

package yamr

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// MockYamrIngestClient is an autogenerated mock type for the YamrIngestClient type
type MockYamrIngestClient struct {
	mock.Mock
}

// Put provides a mock function with given fields: ctx, in, opts
func (_m *MockYamrIngestClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *PutResponse
	if rf, ok := ret.Get(0).(func(context.Context, *PutRequest, ...grpc.CallOption) *PutResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PutResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *PutRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBulk provides a mock function with given fields: ctx, in, opts
func (_m *MockYamrIngestClient) PutBulk(ctx context.Context, in *PutBulkRequest, opts ...grpc.CallOption) (*PutBulkResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *PutBulkResponse
	if rf, ok := ret.Get(0).(func(context.Context, *PutBulkRequest, ...grpc.CallOption) *PutBulkResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PutBulkResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *PutBulkRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
