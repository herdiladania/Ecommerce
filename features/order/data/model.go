package data

import (
	"e-commerce/features/order"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID          uint
	SellerID        uint
	CartID          uint
	Quantity        int
	TotalPrice      float64
	Address         string
	PaymentUrl      string
	TransactionCode string
	TransactionDate time.Time
	Status          string
}

type userCart struct {
	ID         uint
	ProductID  uint
	Quantity   int
	TotalPrice float64
}

type userProduct struct {
	gorm.Model
	Name        string
	Price       uint
	Quantity    uint
	Description string
	Image       string
	UserID      uint
}

type orderProduct struct {
	ID         uint
	OrderID    uint
	ProductID  uint
	Quantity   int
	TotalPrice float64
}

func DataToCore(data Order) order.Core {
	return order.Core{
		ID:              data.ID,
		UserID:          data.UserID,
		CartID:          data.CartID,
		Address:         data.Address,
		Quantity:        data.Quantity,
		TotalPrice:      data.TotalPrice,
		PaymentUrl:      data.PaymentUrl,
		TransactionCode: data.TransactionCode,
		TransactionDate: data.TransactionDate,
		Status:          data.Status,
	}
}

func CoreToData(data order.Core) Order {
	return Order{
		Model:           gorm.Model{ID: data.ID},
		UserID:          data.UserID,
		CartID:          data.CartID,
		Quantity:        data.Quantity,
		TotalPrice:      data.TotalPrice,
		Address:         data.Address,
		PaymentUrl:      data.PaymentUrl,
		TransactionCode: data.TransactionCode,
		TransactionDate: data.TransactionDate,
		Status:          data.Status,
	}

}
