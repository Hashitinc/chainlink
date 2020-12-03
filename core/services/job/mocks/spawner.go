// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import (
	context "context"

	job "github.com/smartcontractkit/chainlink/core/services/job"
	mock "github.com/stretchr/testify/mock"

	null "gopkg.in/guregu/null.v4"
)

// Spawner is an autogenerated mock type for the Spawner type
type Spawner struct {
	mock.Mock
}

// CreateJob provides a mock function with given fields: ctx, spec, name
func (_m *Spawner) CreateJob(ctx context.Context, spec job.Spec, name null.String) (int32, error) {
	ret := _m.Called(ctx, spec, name)

	var r0 int32
	if rf, ok := ret.Get(0).(func(context.Context, job.Spec, null.String) int32); ok {
		r0 = rf(ctx, spec, name)
	} else {
		r0 = ret.Get(0).(int32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, job.Spec, null.String) error); ok {
		r1 = rf(ctx, spec, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteJob provides a mock function with given fields: ctx, jobID
func (_m *Spawner) DeleteJob(ctx context.Context, jobID int32) error {
	ret := _m.Called(ctx, jobID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(ctx, jobID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterDelegate provides a mock function with given fields: delegate
func (_m *Spawner) RegisterDelegate(delegate job.Delegate) {
	_m.Called(delegate)
}

// Start provides a mock function with given fields:
func (_m *Spawner) Start() {
	_m.Called()
}

// Stop provides a mock function with given fields:
func (_m *Spawner) Stop() {
	_m.Called()
}
