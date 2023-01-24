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

// For All Posts
type ProductHome struct {
	ID    uint
	Name  string
	Price uint
	Image string
}

func (dataModel *ProductHome) AllModelsToCore() product.Core {
	return product.Core{
		ID:    dataModel.ID,
		Name:  dataModel.Name,
		Price: dataModel.Price,
		Image: dataModel.Image,
	}
}

func ListAllModelsToCore(dataModels []ProductHome) []product.Core {
	var dataCore []product.Core
	for _, value := range dataModels {
		dataCore = append(dataCore, value.AllModelsToCore())
	}
	return dataCore
}
