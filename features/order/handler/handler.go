package handler

import (
	"e-commerce/features/order"
	"e-commerce/helper"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type OrderHandle struct {
	srv order.OrderService
}

func New(os order.OrderService) order.OrderHandler {
	return &OrderHandle{
		srv: os,
	}
}
func (oh *OrderHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {

		res, err := oh.srv.Add(c.Get("user"))
		if err != nil {
			log.Println("error running add product service")
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    ToResponse(res),
			"message": "order payment created",
		})
	}
}
