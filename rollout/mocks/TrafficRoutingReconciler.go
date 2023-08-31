// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1"
)

// TrafficRoutingReconciler is an autogenerated mock type for the TrafficRoutingReconciler type
type TrafficRoutingReconciler struct {
	mock.Mock
}

// RemoveManagedRoutes provides a mock function with given fields:
func (_m *TrafficRoutingReconciler) RemoveManagedRoutes() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetHeaderRoute provides a mock function with given fields: setHeaderRoute
func (_m *TrafficRoutingReconciler) SetHeaderRoute(setHeaderRoute *v1alpha1.SetHeaderRoute) error {
	ret := _m.Called(setHeaderRoute)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.SetHeaderRoute) error); ok {
		r0 = rf(setHeaderRoute)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetMirrorRoute provides a mock function with given fields: setMirrorRoute
func (_m *TrafficRoutingReconciler) SetMirrorRoute(setMirrorRoute *v1alpha1.SetMirrorRoute) error {
	ret := _m.Called(setMirrorRoute)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1alpha1.SetMirrorRoute) error); ok {
		r0 = rf(setMirrorRoute)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetWeight provides a mock function with given fields: desiredWeight, additionalDestinations
func (_m *TrafficRoutingReconciler) SetWeight(desiredWeight int32, additionalDestinations ...v1alpha1.WeightDestination) error {
	_va := make([]any, len(additionalDestinations))
	for _i := range additionalDestinations {
		_va[_i] = additionalDestinations[_i]
	}
	var _ca []any
	_ca = append(_ca, desiredWeight)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(int32, ...v1alpha1.WeightDestination) error); ok {
		r0 = rf(desiredWeight, additionalDestinations...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Type provides a mock function with given fields:
func (_m *TrafficRoutingReconciler) Type() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// UpdateHash provides a mock function with given fields: canaryHash, stableHash, additionalDestinations
func (_m *TrafficRoutingReconciler) UpdateHash(canaryHash string, stableHash string, additionalDestinations ...v1alpha1.WeightDestination) error {
	_va := make([]any, len(additionalDestinations))
	for _i := range additionalDestinations {
		_va[_i] = additionalDestinations[_i]
	}
	var _ca []any
	_ca = append(_ca, canaryHash, stableHash)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, ...v1alpha1.WeightDestination) error); ok {
		r0 = rf(canaryHash, stableHash, additionalDestinations...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyWeight provides a mock function with given fields: desiredWeight, additionalDestinations
func (_m *TrafficRoutingReconciler) VerifyWeight(desiredWeight int32, additionalDestinations ...v1alpha1.WeightDestination) (*bool, error) {
	_va := make([]any, len(additionalDestinations))
	for _i := range additionalDestinations {
		_va[_i] = additionalDestinations[_i]
	}
	var _ca []any
	_ca = append(_ca, desiredWeight)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int32, ...v1alpha1.WeightDestination) (*bool, error)); ok {
		return rf(desiredWeight, additionalDestinations...)
	}
	if rf, ok := ret.Get(0).(func(int32, ...v1alpha1.WeightDestination) *bool); ok {
		r0 = rf(desiredWeight, additionalDestinations...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bool)
		}
	}

	if rf, ok := ret.Get(1).(func(int32, ...v1alpha1.WeightDestination) error); ok {
		r1 = rf(desiredWeight, additionalDestinations...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTrafficRoutingReconciler creates a new instance of TrafficRoutingReconciler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTrafficRoutingReconciler(t interface {
	mock.TestingT
	Cleanup(func())
}) *TrafficRoutingReconciler {
	mock := &TrafficRoutingReconciler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
