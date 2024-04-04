// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	commondtos "lucio.com/order-service/src/domain/common/dtos"

	dtos "lucio.com/order-service/src/domain/workorder/dtos"

	entities "lucio.com/order-service/src/domain/workorder/entities"

	mock "github.com/stretchr/testify/mock"
)

// UpdateWorkOrderUC is an autogenerated mock type for the UpdateWorkOrderUC type
type UpdateWorkOrderUC struct {
	mock.Mock
}

// Execute provides a mock function with given fields: updateWorkOrder
func (_m *UpdateWorkOrderUC) Execute(updateWorkOrder entities.WorkOrder) (*dtos.UpdatedWorkOrderResponse, *commondtos.CustomError) {
	ret := _m.Called(updateWorkOrder)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 *dtos.UpdatedWorkOrderResponse
	var r1 *commondtos.CustomError
	if rf, ok := ret.Get(0).(func(entities.WorkOrder) (*dtos.UpdatedWorkOrderResponse, *commondtos.CustomError)); ok {
		return rf(updateWorkOrder)
	}
	if rf, ok := ret.Get(0).(func(entities.WorkOrder) *dtos.UpdatedWorkOrderResponse); ok {
		r0 = rf(updateWorkOrder)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.UpdatedWorkOrderResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(entities.WorkOrder) *commondtos.CustomError); ok {
		r1 = rf(updateWorkOrder)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*commondtos.CustomError)
		}
	}

	return r0, r1
}

// NewUpdateWorkOrderUC creates a new instance of UpdateWorkOrderUC. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUpdateWorkOrderUC(t interface {
	mock.TestingT
	Cleanup(func())
}) *UpdateWorkOrderUC {
	mock := &UpdateWorkOrderUC{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
