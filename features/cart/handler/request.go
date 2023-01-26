package handler

// import "e-commerce/features/cart"

type CartReqAdd struct {
	Quantity  int  `json:"quantity" form:"quantity"`
	ProductID uint `json:"product_id" form:"product_id"`
	UserID    uint `json:"user_id" form:"user_id"`
}

type CartReqUpdate struct {
	Quantity  int  `json:"quantity" form:"quantity"`
}