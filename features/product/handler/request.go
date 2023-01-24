package handler

import "e-commerce/features/product"

type AddProductRequest struct {
	Name        string `json:"name" form:"name"`
	Price       uint   `json:"price" form:"price"`
	Quantity    uint   `json:"quantity" form:"quantity"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
}

type UpdateProductRequest struct {
	Name        string `json:"name" form:"name"`
	Price       uint   `json:"price" form:"price"`
	Quantity    uint   `json:"quantity" form:"quantity"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
}

func ToCore(data interface{}) *product.Core {
	res := product.Core{}

	switch data.(type) {
	case AddProductRequest:
		cnv := data.(AddProductRequest)
		res.Name = cnv.Name
		res.Price = cnv.Price
		res.Quantity = cnv.Quantity
		res.Description = cnv.Description
		res.Image = cnv.Image
	case UpdateProductRequest:
		cnv := data.(UpdateProductRequest)
		res.Name = cnv.Name
		res.Price = cnv.Price
		res.Quantity = cnv.Quantity
		res.Description = cnv.Description
		res.Image = cnv.Image
	default:
		return nil
	}

	return &res
}
