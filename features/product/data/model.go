package data

import (
	"e-commerce/features/product"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Price       uint
	Quantity    uint
	Description string
	Image       string
	UserID      uint
}

func DataToCore(data Product) product.Core {
	return product.Core{
		ID:          data.ID,
		Name:        data.Name,
		Price:       data.Price,
		Quantity:    data.Quantity,
		Description: data.Description,
		Image:       data.Image,
		UserID:      data.UserID,
	}
}

func CoreToData(data product.Core) Product {
	return Product{
		Model:       gorm.Model{ID: data.ID},
		Name:        data.Name,
		Price:       data.Price,
		Quantity:    data.Quantity,
		Description: data.Description,
		Image:       data.Image,
		UserID:      data.UserID,
	}
}
