package handler

import (
	"e-commerce/features/order"
)

type OrderRequest struct {
	CartID  uint   `json:"cart_id" form:"cart_id"`
	Address string `json:"address" form:"address"`
}

type UpdateStatusRequest struct {
	OrderID uint   `json:"order_id" form:"order_id"`
	Status  string `json:"status" form:"status"`
}

func ToCore(data interface{}) *order.Core {
	res := order.Core{}

	switch data.(type) {
	case OrderRequest:
		cnv := data.(OrderRequest)
		res.CartID = cnv.CartID
		res.Address = cnv.Address
	case UpdateStatusRequest:
		cnv := data.(UpdateStatusRequest)
		res.ID = cnv.OrderID
		res.Status = cnv.Status
	}

	return &res
}
