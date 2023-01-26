package handler

import (
	"e-commerce/features/cart"
	"net/http"
	"strings"
)

type CartResp struct {
	ID         uint   `json:"id"`
	Quantity   int    `json:"qty"`
	TotalPrice int    `json:"total_price"`
	ProductID  uint   `json:"product_id"`
	Name       string `json:"product_name"`
	Image      string `json:"product_image"`
	UserID     uint   `json:"user_id"`
}

type CartRespUp struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"qty"`
}

func CoreToCartResp(data cart.Core) CartResp {
	return CartResp{
		ID:         data.ID,
		Quantity:   data.Quantity,
		TotalPrice: data.TotalPrice,
		ProductID:  data.ProductID,
		Name:       data.Name,
		Image:      data.Image,
		UserID:     data.UserID,
	}
}

type ListCartResp []CartResp

func ErrorResponse(msg string) (int, interface{}) {
	resp := map[string]interface{}{}
	code := http.StatusInternalServerError

	if msg != "" {
		resp["message"] = msg
	}

	switch true {
	case strings.Contains(msg, "server"):
		code = http.StatusInternalServerError
	case strings.Contains(msg, "format"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "not found"):
		code = http.StatusNotFound
	case strings.Contains(msg, "conflict"):
		code = http.StatusConflict
	case strings.Contains(msg, "duplicate"):
		code = http.StatusConflict
	case strings.Contains(msg, "bad request"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "validation"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "unmarshal"):
		resp["message"] = "bad request"
		code = http.StatusBadRequest
	case strings.Contains(msg, "upload"):
		code = http.StatusInternalServerError
	}

	return code, resp
}
