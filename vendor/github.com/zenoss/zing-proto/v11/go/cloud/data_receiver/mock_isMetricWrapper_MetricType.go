// Code generated by mockery v2.14.0. DO NOT EDIT.

package data_receiver

import mock "github.com/stretchr/testify/mock"

// mockIsMetricWrapper_MetricType is an autogenerated mock type for the isMetricWrapper_MetricType type
type mockIsMetricWrapper_MetricType struct {
	mock.Mock
}

type mockIsMetricWrapper_MetricType_Expecter struct {
	mock *mock.Mock
}

func (_m *mockIsMetricWrapper_MetricType) EXPECT() *mockIsMetricWrapper_MetricType_Expecter {
	return &mockIsMetricWrapper_MetricType_Expecter{mock: &_m.Mock}
}

// isMetricWrapper_MetricType provides a mock function with given fields:
func (_m *mockIsMetricWrapper_MetricType) isMetricWrapper_MetricType() {
	_m.Called()
}

// mockIsMetricWrapper_MetricType_isMetricWrapper_MetricType_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'isMetricWrapper_MetricType'
type mockIsMetricWrapper_MetricType_isMetricWrapper_MetricType_Call struct {
	*mock.Call
}

// isMetricWrapper_MetricType is a helper method to define mock.On call
func (_e *mockIsMetricWrapper_MetricType_Expecter) isMetricWrapper_MetricType() *mockIsMetricWrapper_MetricType_isMetricWrapper_MetricType_Call {
	return &mockIsMetricWrapper_MetricType_isMetricWrapper_MetricType_Call{Call: _e.mock.On("isMetricWrapper_MetricType")}
}

func (_c *mockIsMetricWrapper_MetricType_isMetricWrapper_MetricType_Call) Run(run func()) *mockIsMetricWrapper_MetricType_isMetricWrapper_MetricType_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *mockIsMetricWrapper_MetricType_isMetricWrapper_MetricType_Call) Return() *mockIsMetricWrapper_MetricType_isMetricWrapper_MetricType_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTnewMockIsMetricWrapper_MetricType interface {
	mock.TestingT
	Cleanup(func())
}

// newMockIsMetricWrapper_MetricType creates a new instance of mockIsMetricWrapper_MetricType. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockIsMetricWrapper_MetricType(t mockConstructorTestingTnewMockIsMetricWrapper_MetricType) *mockIsMetricWrapper_MetricType {
	mock := &mockIsMetricWrapper_MetricType{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
