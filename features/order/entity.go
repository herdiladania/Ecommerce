package order

import "github.com/labstack/echo"

type Core struct {
	ID              uint
	UserID          uint
	ProductID       uint
	ProductName     string
	Quantity        int
	TotalPrice      int
	PaymentUrl      string
	TransactionCode string
	TransactionDate string
	Status          int
}

type OrderHandler interface {
	Add() echo.HandlerFunc
	OrderHistory() echo.HandlerFunc
	SellingHistory() echo.HandlerFunc
	UpdateOrderStatus() echo.HandlerFunc
	DeleteOrder() echo.HandlerFunc
}

type OrderService interface {
	Add(token interface{}, totalPrice float64) (Core, string, error)
	OrderHistory(token interface{}) ([]Core, error)
	SellingHistory(token interface{}) ([]Core, error)
	UpdateOrderStatus(token interface{}, orderID uint, updatedStatus int) error
	DeleteOrder(token interface{}, orderID uint) error
}

type OrderData interface {
	Add(userId uint, totalPrice float64) (Core, string, error)
	OrderHistory(userId uint) ([]Core, error)
	SellingHistory(userId uint) ([]Core, error)
	UpdateOrderStatus(userID uint, orderID uint, updatedStatus int) error
	DeleteOrder(userID uint, orderID uint) error
}
