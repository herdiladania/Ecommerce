package handler

import (
	"e-commerce/features/order"
	"time"
)

type OrderResponse struct {
	ID              uint      `json:"id"`
	UserID          uint      `json:"user_id"`
	Address         string    `json:"address"`
	Quantity        int       `json:"qty"`
	TotalPrice      float64   `json:"total_price"`
	PaymentUrl      string    `json:"payment_url"`
	TransactionCode string    `json:"transaction_code"`
	TransactionDate time.Time `json:"transaction_date"`
	Status          string    `json:"status"`
}

func ToResponse(data order.Core) OrderResponse {
	return OrderResponse{
		ID:              data.ID,
		UserID:          data.UserID,
		Address:         data.Address,
		Quantity:        data.Quantity,
		TotalPrice:      data.TotalPrice,
		PaymentUrl:      data.PaymentUrl,
		TransactionCode: data.TransactionCode,
		TransactionDate: data.TransactionDate,
		Status:          data.Status,
	}

}

func OrderHistoryResponse(data []order.Core) []OrderResponse {
	var listOrder = []OrderResponse{}
	for _, order := range data {
		listOrder = append(listOrder, ToResponse(order))
	}
	return listOrder
}
