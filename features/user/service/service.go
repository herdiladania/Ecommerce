package services

import (
	"e-commerce/features/user"
	helper "e-commerce/helper"
	"errors"
	"log"
	"mime/multipart"
	"strings"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	qry user.UserData
	vld *validator.Validate
}

func New(ud user.UserData) user.UserService {
	return &userUseCase{
		qry: ud,
		vld: validator.New(),
	}
}

func (uuc *userUseCase) Login(email, password string) (string, user.Core, error) {
	res, err := uuc.qry.Login(email)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return "", user.Core{}, errors.New(msg)
	}

	if err := helper.CheckPassword(res.Password, password); err != nil {
		log.Println("login compare", err.Error())
		return "", user.Core{}, errors.New("password tidak sesuai " + res.Password)
	}

	token, _ := helper.GenerateJWT(int(res.ID))

	return token, res, nil

}

func (uuc *userUseCase) Register(newUser user.Core) (user.Core, error) {

	err := uuc.vld.Struct(newUser)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println(err)
		}
		return user.Core{}, errors.New("validation error")
	}
	hashed, err := helper.GeneratePassword(newUser.Password)
	if err != nil {
		log.Println("bcrypt error ", err.Error())
		return user.Core{}, errors.New("password process error")
	}

	newUser.Password = hashed
	res, err := uuc.qry.Register(newUser)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "data sudah terdaftar"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}

func (uuc *userUseCase) Profile(token interface{}) (user.Core, error) {
	id := helper.ExtractToken(token)
	if id <= 0 {
		return user.Core{}, errors.New("data tidak ditemukan")
	}
	res, err := uuc.qry.Profile(uint(id))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return user.Core{}, errors.New(msg)
	}
	return res, nil
}

func (uuc *userUseCase) Update(token interface{}, updatedData user.Core, profilePhoto *multipart.FileHeader) (user.Core, error) {
	userId := helper.ExtractToken(token)
	if userId <= 0 {
		log.Println("extract token error")
		return user.Core{}, errors.New("extract token error")
	}
	if updatedData.Password != "" {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(updatedData.Password), bcrypt.DefaultCost)
		updatedData.Password = string(hashed)
	}

	res, err := uuc.qry.Profile(uint(userId))
	if err != nil {
		errmsg := ""
		if strings.Contains(err.Error(), "not found") {
			errmsg = "data not found"
		} else {
			errmsg = "server problem"
		}
		log.Println("error profile query: ", err.Error())
		return user.Core{}, errors.New(errmsg)
	}

	if profilePhoto != nil {
		path, _ := helper.UploadProfilePhotoS3(*profilePhoto, res.Email)
		updatedData.Image = path
	}

	res, err = uuc.qry.Update(uint(userId), updatedData)
	if err != nil {
		errmsg := ""
		if strings.Contains(err.Error(), "not found") {
			errmsg = "data not found"
		} else {
			errmsg = "server problem"
		}
		log.Println("error update query: ", err.Error())
		return user.Core{}, errors.New(errmsg)
	}
	return res, nil
}

func (uuc *userUseCase) Delete(token interface{}) (user.Core, error) {

	id := helper.ExtractToken(token)
	if id <= 0 {
		return user.Core{}, errors.New("user not found")
	}
	data, err := uuc.qry.Delete(uint(id))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "internal server error"
		}
		return user.Core{}, errors.New(msg)
	}
	return data, nil

}
