package handler

import "e-commerce/features/order"

type OrderRequest struct {
	CartID  uint   `json:"cart_id" form:"cart_id"`
	Address string `json:"address" form:"address"`
}

func ToCore(data interface{}) *order.Core {
	res := order.Core{}

	switch data.(type) {
	case OrderRequest:
		cnv := data.(OrderRequest)
		res.CartID = cnv.CartID
		res.Address = cnv.Address
	}

	return &res
}
