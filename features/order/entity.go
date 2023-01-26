package order

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID              uint
	UserID          uint
	SellerID        uint
	CartID          uint
	Address         string
	Quantity        int
	TotalPrice      float64
	PaymentUrl      string
	TransactionCode string
	TransactionDate time.Time
	Status          string
}

type OrderHandler interface {
	Add() echo.HandlerFunc
	OrderHistory() echo.HandlerFunc
	UpdateOrderStatus() echo.HandlerFunc
	// DeleteOrder() echo.HandlerFunc
}

type OrderService interface {
	Add(token interface{}, cartID uint, adrress string) (Core, error)
	OrderHistory(token interface{}) ([]Core, error)
	UpdateOrderStatus(token interface{}, orderID uint, updatedStatus string) error
	// DeleteOrder(token interface{}, orderID uint) error
}

type OrderData interface {
	Add(userID uint, cartID uint, adrress string) (Core, error)
	OrderHistory(userId uint) ([]Core, error)
	UpdateOrderStatus(userID uint, orderID uint, updatedStatus string) error
	// DeleteOrder(userID uint, orderID uint) error
}
