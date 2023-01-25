package service

import (
	"e-commerce/features/product"
	"e-commerce/helper"
	"e-commerce/mocks"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	repo := mocks.NewProductData(t)

	inputData := product.Core{
		Name:        "aqua",
		Price:       3000,
		Quantity:    24,
		Description: "air minum",
		Image:       "aqua.jpg",
	}

	resData := product.Core{
		ID:          1,
		Name:        "aqua",
		Price:       3000,
		Quantity:    24,
		Description: "air minum",
		Image:       "https://ecommercegroup7.s3.ap-southeast-1.amazonaws.com/files/product/1/aqua.jpg",
		UserID:      1,
	}

	var productImage *multipart.FileHeader

	t.Run("success add product", func(t *testing.T) {
		repo.On("Add", uint(1), inputData).Return(resData, nil).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData, productImage)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("product not found", func(t *testing.T) {
		userID := uint(1)

		repo.On("Add", userID, inputData).Return(product.Core{}, errors.New("data not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData, productImage)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})

	t.Run("server problem", func(t *testing.T) {
		userID := uint(2)

		repo.On("Add", userID, inputData).Return(product.Core{}, errors.New("server problem")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData, productImage)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res.UserID, uint(0))
		repo.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateJWT(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Add(pToken, inputData, productImage)
		assert.NotNil(t, err)
		assert.Equal(t, res.UserID, uint(0))
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewProductData(t)

	inputData := product.Core{
		Name:        "aqua",
		Price:       3000,
		Quantity:    24,
		Description: "air minum",
		Image:       "aqua.jpg",
	}

	resData := product.Core{
		ID:          1,
		Name:        "aqua",
		Price:       3000,
		Quantity:    24,
		Description: "air minum",
		Image:       "https://ecommercegroup7.s3.ap-southeast-1.amazonaws.com/files/product/1/aqua.jpg",
		UserID:      1,
	}

	var productImage *multipart.FileHeader

	t.Run("success update product", func(t *testing.T) {
		userID := uint(1)
		productId := uint(1)

		repo.On("Update", productId, userID, inputData).Return(resData, nil).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Update(pToken, productId, inputData, productImage)
		assert.Nil(t, err)
		assert.Equal(t, res.ID, resData.ID)
		repo.AssertExpectations(t)
	})

	// t.Run("error upload post photo", func(t *testing.T) {
	// 	postID := uint(1)
	// 	userID := uint(1)
	// 	postPhoto := &multipart.FileHeader{
	// 		Filename: "a",
	// 		Size:     10,
	// 	}
	// 	srv := Isolation(repo)

	// 	_, token := helper.GenerateJWT(1)

	// 	pToken := token.(*jwt.Token)
	// 	pToken.Valid = true

	// 	res, err := srv.Update(pToken, postID, inputData, postPhoto)
	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, res.UserID, userID)
	// })

	t.Run("product not found", func(t *testing.T) {
		userID := uint(1)
		productId := uint(1)

		repo.On("Update", productId, userID, inputData).Return(product.Core{}, errors.New("data not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Update(pToken, productId, inputData, productImage)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)

	})

	t.Run("server problem", func(t *testing.T) {
		userID := uint(2)
		productId := uint(3)

		repo.On("Update", productId, userID, inputData).Return(product.Core{}, errors.New("server problem")).Once()

		srv := New(repo)

		_, token := helper.GenerateJWT(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Update(pToken, productId, inputData, productImage)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("jwt not valid", func(t *testing.T) {
		productId := uint(2)

		srv := New(repo)

		_, token := helper.GenerateJWT(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Update(pToken, productId, inputData, productImage)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, uint(0), res.UserID)
		repo.AssertExpectations(t)
	})
}
