package handler

import (
	"e-commerce/features/cart"
	"e-commerce/helper"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type cartHandler struct {
	srv cart.CartService
}

func New(srv *cart.CartService) cart.CartHandler {
	return &cartHandler{
		srv: *srv,
	}
}

func (ch *cartHandler) AddCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input CartReqAdd
		if err := c.Bind(&input); err != nil {
			return c.JSON(ErrorResponse(err.Error()))
		}

		input.UserID = uint(helper.ExtractToken(c.Get("user")))
		newCart := cart.Core{}
		copier.Copy(&newCart, &input)
		res, err := ch.srv.AddCartSrv(newCart)
		if err != nil {
			return c.JSON(ErrorResponse(err.Error()))
		}

		upResp := CartRespUp{}
		copier.Copy(&upResp, &res)
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "successfully add to cart",
			"data": upResp,
		})
	}
}

func (ch *cartHandler) MyCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := uint(helper.ExtractToken(c.Get("user")))
		
		res, err := ch.srv.MyCartSrv(uint(userID))
		if err != nil {
			return c.JSON(ErrorResponse(err.Error()))
		}
		dataRes := ListCartResp{}
		copier.Copy(&dataRes, &res)

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":"successfully get all cart",
			"data": dataRes,
		})
	}
}

func (ch *cartHandler) UpdateCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := uint(helper.ExtractToken(c.Get("user")))
		cartIDParam := c.Param("id")
		cartID, _ := strconv.Atoi(cartIDParam)

		var input CartReqUpdate
		if err := c.Bind(&input); err != nil {
			log.Println("Bind error")
			return c.JSON(ErrorResponse(err.Error()))
		}

		updateCart := cart.Core{}
		copier.Copy(&updateCart, &input)
		res, err := ch.srv.UpdateCartSrv(userID, uint(cartID), updateCart)
		if err != nil {
			return c.JSON(ErrorResponse(err.Error()))
		}

		upResp := CartRespUp{}
		copier.Copy(&upResp, &res)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":"successfully updated cart",
			"data": upResp,
		})
	}
}

func (ch *cartHandler) DeleteCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := uint(helper.ExtractToken(c.Get("user")))
		cartIDParam := c.Param("id")
		cartID, _ := strconv.Atoi(cartIDParam)

		err := ch.srv.DeleteCartSrv(userID, uint(cartID))
		if err != nil {
			return c.JSON(ErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "successfully delete cart",
		})
	}
}