package service

import (
	"e-commerce/features/order"
	"e-commerce/helper"
	"e-commerce/mocks"
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	repo := mocks.NewOrderData(t)

	cartID := uint(1)
	adrress := "Jl.Merdeka 17, Jakarta"

	resData := order.Core{
		ID:              1,
		UserID:          1,
		SellerID:        8,
		CartID:          cartID,
		Address:         "Jl.Merdeka 17, Jakarta",
		Quantity:        2,
		TotalPrice:      20000,
		PaymentUrl:      "https://app.sandbox.midtrans.com/snap/v3/redirection/269bc1f7-da56-4154-8f82-81f5acb187ee",
		TransactionCode: "Transaction-8",
		TransactionDate: time.Time{},
		Status:          "waiting for payment",
	}

	t.Run("Successfully add order", func(t *testing.T) {
		repo.On("Add", uint(1), cartID, adrress).Return(resData, nil).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		srv := New(repo)
		res, err := srv.Add(pToken, cartID, adrress)
		assert.Nil(t, err)
		assert.Equal(t, res, resData)
		assert.Equal(t, resData.PaymentUrl, res.PaymentUrl)
		repo.AssertExpectations(t)
	})

	t.Run("User not found", func(t *testing.T) {

		srv := New(repo)

		_, token := helper.GenerateJWT(0)
		res, err := srv.Add(token, cartID, adrress)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("order not found", func(t *testing.T) {
		repo.On("Add", uint(1), cartID, adrress).Return(order.Core{}, errors.New("not found")).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		srv := New(repo)

		res, err := srv.Add(pToken, cartID, adrress)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.Contains(t, err.Error(), "not found")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("Add", uint(1), cartID, adrress).Return(order.Core{}, errors.New("server problem")).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		srv := New(repo)

		res, err := srv.Add(pToken, cartID, adrress)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.Contains(t, err.Error(), "server problem")
		repo.AssertExpectations(t)
	})

}

func TestOrderHistory(t *testing.T) {
	repo := mocks.NewOrderData(t)
	resData := []order.Core{
		{
			ID:              1,
			UserID:          1,
			SellerID:        8,
			CartID:          1,
			Address:         "Jl.Merdeka 17, Jakarta",
			Quantity:        2,
			TotalPrice:      20000,
			PaymentUrl:      "https://app.sandbox.midtrans.com/snap/v3/redirection/269bc1f7-da56-4154-8f82-81f5acb187ee",
			TransactionCode: "Transaction-8",
			TransactionDate: time.Time{},
			Status:          "waiting for payment",
		},
	}

	t.Run("successfully get order history", func(t *testing.T) {
		repo.On("OrderHistory", uint(1)).Return(resData, nil).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.OrderHistory(pToken)
		assert.Nil(t, err)
		assert.Equal(t, len(res), len(resData))
		repo.AssertExpectations(t)
	})

	t.Run("User not found", func(t *testing.T) {

		srv := New(repo)

		_, token := helper.GenerateJWT(0)
		res, err := srv.OrderHistory(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, []order.Core{}, res)
		repo.AssertExpectations(t)
	})

	t.Run("order history not found", func(t *testing.T) {
		repo.On("OrderHistory", uint(1)).Return([]order.Core{}, errors.New("not found")).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		srv := New(repo)

		res, err := srv.OrderHistory(pToken)
		assert.NotNil(t, err)
		assert.Equal(t, []order.Core{}, res)
		assert.Contains(t, err.Error(), "not found")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("OrderHistory", uint(1)).Return([]order.Core{}, errors.New("server problem")).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		srv := New(repo)

		res, err := srv.OrderHistory(pToken)
		assert.NotNil(t, err)
		assert.Equal(t, []order.Core{}, res)
		assert.Contains(t, err.Error(), "server problem")
		repo.AssertExpectations(t)
	})

}

func TestUpdateOrderStatus(t *testing.T) {
	repo := mocks.NewOrderData(t)
	updateStatus := "canceled"
	orderID := uint(1)

	t.Run("successfully update order status", func(t *testing.T) {
		repo.On("UpdateOrderStatus", uint(1), orderID, updateStatus).Return(nil).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		srv := New(repo)

		err := srv.UpdateOrderStatus(pToken, orderID, updateStatus)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("User not found", func(t *testing.T) {

		srv := New(repo)

		_, token := helper.GenerateJWT(0)
		err := srv.UpdateOrderStatus(token, orderID, updateStatus)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("UpdateOrderStatus", uint(1), orderID, updateStatus).Return(errors.New("not found")).Once()

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		srv := New(repo)
		err := srv.UpdateOrderStatus(pToken, orderID, updateStatus)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "not found")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		repo.On("UpdateOrderStatus", uint(1), orderID, updateStatus).Return(errors.New("server problem")).Once()
		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		srv := New(repo)
		err := srv.UpdateOrderStatus(pToken, orderID, updateStatus)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "server problem")
		repo.AssertExpectations(t)
	})
}
