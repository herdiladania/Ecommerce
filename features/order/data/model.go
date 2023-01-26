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

func ListOrderToCore(data []Order) []order.Core {
	var listOrder = []order.Core{}
	for _, order := range data {
		listOrder = append(listOrder, DataToCore(order))
	}

	return listOrder
}
