package service

import (
	"e-commerce/features/cart"
	"e-commerce/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCartSrv(t *testing.T) {
	repo := mocks.NewCartData(t)

	inputData := cart.Core{
		Quantity:  2,
		ProductID: 1,
	}

	resData := cart.Core{
		ProductID: 1,
		Quantity:  2,
		ID:        1,
	}

	t.Run("success add cart", func(t *testing.T) {
		// userID := uint(1)
		repo.On("AddCartData", inputData).Return(resData, nil).Once()
		repo.On("AddCHPData", inputData.ID, inputData).Return(resData, nil).Once()

		srv := New(repo)

		// _, token := helper.GenerateJWT(1)

		// pToken := token.(*jwt.Token)
		// pToken.Valid = true

		_, err := srv.AddCartSrv(inputData)
		assert.Nil(t, err)
		// assert.Equal(t, res.UserID, userID)
		repo.AssertExpectations(t)
	})
}

func TestDeleteCartSrv(t *testing.T) {
	repo := mocks.NewCartData(t)

	t.Run("success delete cart", func(t *testing.T) {
		cartID := uint(3)
		userID := uint(2)
		repo.On("DeleteCartData", userID, cartID).Return(nil).Once()

		srv := New(repo)

		err := srv.DeleteCartSrv(userID, cartID)
		assert.Nil(t, err)
		repo.AssertExpectations(t)

	})

	t.Run("data not found", func(t *testing.T) {
		cartID := uint(3)
		userID := uint(2)
		repo.On("DeleteCartData", userID, cartID).Return(errors.New("data not found")).Once()

		srv := New(repo)

		err := srv.DeleteCartSrv(userID, cartID)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		cartID := uint(3)
		userID := uint(2)
		repo.On("DeleteCartData", userID, cartID).Return(errors.New("server problem")).Once()

		srv := New(repo)

		err := srv.DeleteCartSrv(userID, cartID)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})
}

func TestUpdateCartSrv(t *testing.T) {
	repo := mocks.NewCartData(t)

	inputData := cart.Core{
		Quantity: 2,
	}
	resData := cart.Core{
		ProductID: 1,
		Quantity:  2,
	}

	t.Run("success update", func(t *testing.T) {
		userID := uint(1)
		cartID := uint(1)
		repo.On("UpdateCartData", userID, cartID, inputData).Return(resData, nil).Once()

		srv := New(repo)

		_, err := srv.UpdateCartSrv(userID, cartID, inputData)
		assert.Nil(t, err)
		// assert.Equal(t, res.UserID, userID)
		repo.AssertExpectations(t)
	})

	t.Run("cart not found", func(t *testing.T) {
		userID := uint(1)
		cartID := uint(1)
		repo.On("UpdateCartData", userID, cartID, inputData).Return(cart.Core{}, errors.New("cart not found")).Once()

		srv := New(repo)

		_, err := srv.UpdateCartSrv(userID, cartID, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		// assert.Equal(t, res.UserID, userID)
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		userID := uint(1)
		cartID := uint(1)
		repo.On("UpdateCartData", userID, cartID, inputData).Return(cart.Core{}, errors.New("server problem")).Once()

		srv := New(repo)

		_, err := srv.UpdateCartSrv(userID, cartID, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		// assert.Equal(t, res.UserID, userID)
		repo.AssertExpectations(t)
	})
}

func TestMyCartSrv(t *testing.T) {
	repo := mocks.NewCartData(t)

	resData := []cart.Core{
		{
			ID:         1,
			Quantity:   3,
			TotalPrice: 6000,
			ProductID:  4,
			Name:       "taro",
			Image:      "https://ecommercegroup7.s3.ap-southeast-1.amazonaws.com/files/product/2/taro.jpg",
			UserID:     2,
		},
		{
			ID:         2,
			Quantity:   1,
			TotalPrice: 3000,
			ProductID:  2,
			Name:       "taro",
			Image:      "https://ecommercegroup7.s3.ap-southeast-1.amazonaws.com/files/product/2/taro.jpg",
			UserID:     2,
		},
	}

	t.Run("success get mycart", func(t *testing.T) {
		userID := uint(2)
		repo.On("MyCartData", userID).Return(resData, nil).Once()

		srv := New(repo)

		res, err := srv.MyCartSrv(userID)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		userID := uint(2)
		repo.On("MyCartData", userID).Return([]cart.Core{}, errors.New("data not found")).Once()

		srv := New(repo)

		res, err := srv.MyCartSrv(userID)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		userID := uint(2)
		repo.On("MyCartData", userID).Return([]cart.Core{}, errors.New("server problem")).Once()

		srv := New(repo)

		res, err := srv.MyCartSrv(userID)
		assert.NotNil(t, err)
		assert.Empty(t, res)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})

}
