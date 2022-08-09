// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	redis "github.com/go-redis/redis/v8"

	time "time"
)

// RedisCommander is an autogenerated mock type for the RedisCommander type
type RedisCommander struct {
	mock.Mock
}

// Expire provides a mock function with given fields: ctx, key, expiration
func (_m *RedisCommander) Expire(ctx context.Context, key string, expiration time.Duration) *redis.BoolCmd {
	ret := _m.Called(ctx, key, expiration)

	var r0 *redis.BoolCmd
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration) *redis.BoolCmd); ok {
		r0 = rf(ctx, key, expiration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.BoolCmd)
		}
	}

	return r0
}

// Get provides a mock function with given fields: ctx, key
func (_m *RedisCommander) Get(ctx context.Context, key string) *redis.StringCmd {
	ret := _m.Called(ctx, key)

	var r0 *redis.StringCmd
	if rf, ok := ret.Get(0).(func(context.Context, string) *redis.StringCmd); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.StringCmd)
		}
	}

	return r0
}

// GetEx provides a mock function with given fields: ctx, key, expiration
func (_m *RedisCommander) GetEx(ctx context.Context, key string, expiration time.Duration) *redis.StringCmd {
	ret := _m.Called(ctx, key, expiration)

	var r0 *redis.StringCmd
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration) *redis.StringCmd); ok {
		r0 = rf(ctx, key, expiration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.StringCmd)
		}
	}

	return r0
}

// Set provides a mock function with given fields: ctx, key, value, expiration
func (_m *RedisCommander) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	ret := _m.Called(ctx, key, value, expiration)

	var r0 *redis.StatusCmd
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) *redis.StatusCmd); ok {
		r0 = rf(ctx, key, value, expiration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.StatusCmd)
		}
	}

	return r0
}

type mockConstructorTestingTNewRedisCommander interface {
	mock.TestingT
	Cleanup(func())
}

// NewRedisCommander creates a new instance of RedisCommander. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRedisCommander(t mockConstructorTestingTNewRedisCommander) *RedisCommander {
	mock := &RedisCommander{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
