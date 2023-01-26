package data

import (
	"e-commerce/features/order"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID          uint
	CartID          uint
	Quantity        int
	TotalPrice      int
	Address         string
	PaymentUrl      string
	TransactionCode string
	TransactionDate string
	Status          string
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
