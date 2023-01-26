package handler

import (
	"e-commerce/features/order"
	"e-commerce/helper"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
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
		input := OrderRequest{}
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		res, err := oh.srv.Add(c.Get("user"), input.CartID, input.Address)
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

func (oh *OrderHandle) OrderHistory() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")

		res, err := oh.srv.OrderHistory(token)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("data not found"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    OrderHistoryResponse(res),
			"message": "successfully get order history",
		})
	}
}

func (oh *OrderHandle) UpdateOrderStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		updateData := UpdateStatusRequest{}
		if err := c.Bind(&updateData); err != nil {
			return c.JSON(http.StatusBadRequest, "format inputan salah")
		}
		err := oh.srv.UpdateOrderStatus(token, updateData.OrderID, updateData.Status)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("data not found"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    updateData.Status,
			"message": "successfully updated order status",
		})
	}
}
