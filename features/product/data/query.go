package data

import (
	"e-commerce/features/product"
	"log"

	"gorm.io/gorm"
)

type productData struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.ProductData {
	return &productData{
		db: db,
	}
}

func (pd *productData) Add(userID uint, newProduct product.Core) (product.Core, error) {
	convert := CoreToData(newProduct)
	convert.UserID = userID

	err := pd.db.Create(&convert).Error
	if err != nil {
		log.Println("add product query error", err.Error())
		return product.Core{}, err
	}

	newProduct.ID = convert.ID
	newProduct.UserID = convert.UserID

	return newProduct, nil
}

func (pd *productData) Update(productID uint, userID uint, updateProduct product.Core) (product.Core, error) {
	return product.Core{}, nil
}

func (pd *productData) Delete(productID uint, userID uint) error {
	return nil
}
