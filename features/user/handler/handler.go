package handler

import (
	"e-commerce/features/user"
	"e-commerce/helper"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type userControll struct {
	srv user.UserService
}

func New(srv user.UserService) user.UserHandler {
	return &userControll{
		srv: srv,
	}
}

func (uc *userControll) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}

		token, res, err := uc.srv.Login(input.Email, input.Password)
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		dataResp := ToResponses(res)
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil login", dataResp, token))
	}
}
func (uc *userControll) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		res, err := uc.srv.Register(*ToCore(input))
		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		dataResp := ToResponses(res)
		return c.JSON(helper.PrintSuccessReponse(http.StatusCreated, "berhasil mendaftar", dataResp))
	}
}
func (uc *userControll) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err := uc.srv.Profile(token)
		if err != nil {
			fmt.Println("ini error")
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		dataResp := ToResponse(res)
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil lihat profil", dataResp))
	}
}

func (uc *userControll) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		var profilePhoto *multipart.FileHeader

		updatedData := RegisterRequest{}
		if err := c.Bind(&updatedData); err != nil {
			return c.JSON(http.StatusBadRequest, "wrong input format")
		}
		file, err := c.FormFile("image")
		if file != nil && err == nil {
			profilePhoto = file
		} else if file != nil && err != nil {
			log.Println("error read profile_photo")
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong image input"))
		}

		res, err := uc.srv.Update(token, *ToCore(updatedData), profilePhoto)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				log.Println("user not found: ", err.Error())
				return c.JSON(http.StatusNotFound, helper.ErrorResponse("user not found"))
			} else {
				log.Println("error update service: ", err.Error())
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    ToResponse(res),
			"message": "success update user's data",
		})
	}
}

func (uc *userControll) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("user")
		err := uc.srv.Delete(tx)

		if err != nil {
			return c.JSON(helper.PrintErrorResponse(err.Error()))
		}
		return c.JSON(helper.PrintSuccessReponse(http.StatusOK, "berhasil hapus"))
	}
}
