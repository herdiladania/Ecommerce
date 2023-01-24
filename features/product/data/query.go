package data

import (
	"e-commerce/features/product"
	"errors"
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
	convert := CoreToData(updateProduct)
	qry := pd.db.Where("id = ? AND user_id = ?", productID, userID).Updates(&convert)
	if qry.RowsAffected <= 0 {
		log.Println("update product query error : data not found")
		return product.Core{}, errors.New("not found")
	}

	if err := qry.Error; err != nil {
		log.Println("update product query error :", err.Error())
		return product.Core{}, errors.New("not found")
	}

	return DataToCore(convert), nil
}

func (pd *productData) Delete(productID uint, userID uint) error {
	return nil
}

func (pd *productData) GetProductById(productID uint, userID uint) (product.Core, error) {
	res := Product{}

	err := pd.db.Where("id = ? AND user_id = ?", productID, userID).First(&res).Error
	if err != nil {
		log.Println("GetProductById query error")
		return product.Core{}, err
	}

	return DataToCore(res), nil
}
