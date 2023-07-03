// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FinishWorkOrderUC is an autogenerated mock type for the FinishWorkOrderUC type
type FinishWorkOrderUC struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ID
func (_m *FinishWorkOrderUC) Execute(ID string) error {
	ret := _m.Called(ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewFinishWorkOrderUC creates a new instance of FinishWorkOrderUC. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFinishWorkOrderUC(t interface {
	mock.TestingT
	Cleanup(func())
}) *FinishWorkOrderUC {
	mock := &FinishWorkOrderUC{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
