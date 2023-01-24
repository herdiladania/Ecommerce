package service

import (
	"e-commerce/features/product"
	"e-commerce/helper"
	"errors"
	"mime/multipart"
	"strings"
)

type productSrv struct {
	data product.ProductData
}

func New(d product.ProductData) product.ProductService {
	return &productSrv{
		data: d,
	}
}

func (ps *productSrv) Add(token interface{}, newProduct product.Core, image *multipart.FileHeader) (product.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return product.Core{}, errors.New("user not found")
	}

	if image != nil {
		path, err := helper.UploadProductImageS3(*image, helper.ExtractToken(token))
		if err != nil {
			return product.Core{}, err
		}
		newProduct.Image = path
	}

	res, err := ps.data.Add(uint(userID), newProduct)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "product not found"
		} else {
			msg = "server problem"
		}
		return product.Core{}, errors.New(msg)
	}

	return res, nil
}

func (ps *productSrv) Update(token interface{}, productID uint, updateProduct product.Core, updateImage *multipart.FileHeader) (product.Core, error) {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return product.Core{}, errors.New("user not found")
	}

	if updateImage != nil {
		path, _ := helper.UploadProductImageS3(*updateImage, helper.ExtractToken(token))
		updateProduct.Image = path
	}

	res, err := ps.data.Update(productID, uint(userID), updateProduct)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "product data not found"
		} else {
			msg = "server problem"
		}
		return product.Core{}, errors.New(msg)
	}

	return res, nil
}

func (ps *productSrv) Delete(token interface{}, productID uint) error {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return errors.New("user not found")
	}

	err := ps.data.Delete(productID, uint(userID))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "product not found"
		} else {
			msg = "server problem"
		}
		return errors.New(msg)
	}
	return nil
}

func (ps *productSrv) GetProductById(token interface{}, productID uint) (product.Core, error) {
	userID := helper.ExtractToken(token)
	res, err := ps.data.GetProductById(productID, uint(userID))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		return product.Core{}, errors.New(msg)
	}
	return res, nil
}
