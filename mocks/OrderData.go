// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	order "e-commerce/features/order"

	mock "github.com/stretchr/testify/mock"
)

// OrderData is an autogenerated mock type for the OrderData type
type OrderData struct {
	mock.Mock
}

// Add provides a mock function with given fields: userID, cartID, adrress
func (_m *OrderData) Add(userID uint, cartID uint, adrress string) (order.Core, error) {
	ret := _m.Called(userID, cartID, adrress)

	var r0 order.Core
	if rf, ok := ret.Get(0).(func(uint, uint, string) order.Core); ok {
		r0 = rf(userID, cartID, adrress)
	} else {
		r0 = ret.Get(0).(order.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint, string) error); ok {
		r1 = rf(userID, cartID, adrress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrderHistory provides a mock function with given fields: userId
func (_m *OrderData) OrderHistory(userId uint) ([]order.Core, error) {
	ret := _m.Called(userId)

	var r0 []order.Core
	if rf, ok := ret.Get(0).(func(uint) []order.Core); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]order.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOrderStatus provides a mock function with given fields: userID, orderID, updatedStatus
func (_m *OrderData) UpdateOrderStatus(userID uint, orderID uint, updatedStatus string) error {
	ret := _m.Called(userID, orderID, updatedStatus)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint, string) error); ok {
		r0 = rf(userID, orderID, updatedStatus)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewOrderData interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrderData creates a new instance of OrderData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrderData(t mockConstructorTestingTNewOrderData) *OrderData {
	mock := &OrderData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
