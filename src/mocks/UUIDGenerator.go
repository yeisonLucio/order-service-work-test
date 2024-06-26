// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
)

// UUIDGenerator is an autogenerated mock type for the UUIDGenerator type
type UUIDGenerator struct {
	mock.Mock
}

// Generate provides a mock function with given fields:
func (_m *UUIDGenerator) Generate() uuid.UUID {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Generate")
	}

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func() uuid.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	return r0
}

// NewUUIDGenerator creates a new instance of UUIDGenerator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUUIDGenerator(t interface {
	mock.TestingT
	Cleanup(func())
}) *UUIDGenerator {
	mock := &UUIDGenerator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
