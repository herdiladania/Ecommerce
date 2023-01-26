package services

import (
	"e-commerce/features/user"
	helper "e-commerce/helper"
	"e-commerce/mocks"
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestRegister(t *testing.T) {
	data := mocks.NewUserData(t)

	newUser := user.Core{
		Name:     "Jhonny",
		HP:       "08124574452",
		Email:    "Jhonny@alta.id",
		Address:  "Jl.Merdeka 17, Jakarta",
		Password: "jhonny123",
	}
	expectedData := user.Core{
		Name:    "Jhonny",
		HP:      "08124574452",
		Email:   "Jhonny@alta.id",
		Address: "Jl.Merdeka 17, Jakarta",
	}

	t.Run("success register", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(expectedData, nil).Once()
		srv := New(data)
		res, err := srv.Register(newUser)
		assert.Nil(t, err)
		assert.Equal(t, expectedData.ID, res.ID)
		assert.Equal(t, expectedData.Name, res.Name)
		data.AssertExpectations(t)
	})

	t.Run("duplicate", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(user.Core{}, errors.New("duplicated")).Once()
		srv := New(data)
		res, err := srv.Register(newUser)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "data sudah terdaftar")
		assert.Equal(t, res.Name, "")
	})

	t.Run("server problem", func(t *testing.T) {
		data.On("Register", mock.Anything).Return(user.Core{}, errors.New("server error")).Once()
		srv := New(data)
		res, err := srv.Register(newUser)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, res.Name, "")
	})

}

func TestLogin(t *testing.T) {
	data := mocks.NewUserData(t)
	t.Run("succcess login", func(t *testing.T) {
		email := "Jhonny@alta.id"
		password := "Jhonny123"
		hashed, _ := bcrypt.GenerateFromPassword([]byte("Jhonny123"), bcrypt.DefaultCost)

		expectedData := user.Core{
			Email:    "Jhonny@alta.id",
			Name:     "Jhonny",
			HP:       "08124574452",
			Password: string(hashed),
			Address:  "Jl. Merdeka 17, Jakarta",
		}
		data.On("Login", email).Return(expectedData, nil).Once()
		srv := New(data)
		token, res, err := srv.Login(email, password)
		assert.Nil(t, err)
		assert.Equal(t, expectedData.Email, res.Email)
		assert.NotNil(t, token)
		data.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		inputEmail := "Jhonny@alta.id"
		data.On("Login", inputEmail).Return(user.Core{}, errors.New("data not found")).Once()

		srv := New(data)
		token, res, err := srv.Login(inputEmail, "be1422")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "data tidak ditemukan")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

	t.Run("wrong password", func(t *testing.T) {
		inputEmail := "Jhonny@alta.id"
		hashed, _ := bcrypt.GenerateFromPassword([]byte("Jhonny123"), bcrypt.DefaultCost)
		expData := user.Core{
			Email:    "Jhonny@alta.id",
			Name:     "Jhonny",
			HP:       "08124574452",
			Password: string(hashed),
		}
		data.On("Login", inputEmail).Return(expData, nil)

		srv := New(data)
		token, res, err := srv.Login(inputEmail, "be1423")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "password")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
}

func TestProfile(t *testing.T) {
	data := mocks.NewUserData(t)

	t.Run("Sukses lihat profile", func(t *testing.T) {
		resData := user.Core{ID: uint(1), Name: "jerry", Email: "jerry@alterra.id", HP: "08123456"}

		data.On("Profile", uint(1)).Return(resData, nil).Once()

		srv := New(data)

		claims := jwt.MapClaims{}
		claims["authorized"] = true
		claims["userID"] = 1
		claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		token.Valid = true

		res, err := srv.Profile(token)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		data.AssertExpectations(t)
	})

	t.Run("jwt tidak valid", func(t *testing.T) {
		srv := New(data)

		_, token := helper.GenerateJWT(1)

		res, err := srv.Profile(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "data tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("data tidak ditemukan", func(t *testing.T) {
		data.On("Profile", uint(4)).Return(user.Core{}, errors.New("data not found")).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "data tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})

	t.Run("masalah di server", func(t *testing.T) {
		data.On("Profile", mock.Anything).Return(user.Core{}, errors.New("terdapat masalah pada server")).Once()
		srv := New(data)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		data.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	data := mocks.NewUserData(t)
	t.Run("success delete", func(t *testing.T) {
		data.On("Delete", uint(1)).Return(nil).Once()

		srv := New(data)

		claims := jwt.MapClaims{}
		claims["authorized"] = true
		claims["userID"] = 1
		claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		token.Valid = true

		err := srv.Delete(token)
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})

	t.Run("jwt not valid", func(t *testing.T) {
		srv := New(data)

		_, token := helper.GenerateJWT(1)

		err := srv.Delete(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
	})

	t.Run("data not found", func(t *testing.T) {
		data.On("Delete", uint(4)).Return(errors.New("data not found")).Once()

		srv := New(data)

		_, token := helper.GenerateJWT(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		data.AssertExpectations(t)
	})

	t.Run("masalah di server", func(t *testing.T) {
		data.On("Delete", mock.Anything).Return(errors.New("terdapat masalah pada server")).Once()
		srv := New(data)

		_, token := helper.GenerateJWT(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		data.AssertExpectations(t)
	})
}
