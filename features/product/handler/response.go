package handler

import "e-commerce/features/product"

type ProductResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Quantity    uint   `json:"quantity"`
	Description string `json:"description"`
	Image       string `json:"image"`
	UserID      uint   `json:"user_id"`
}

type AddProductResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Quantity    uint   `json:"quantity"`
	Description string `json:"description"`
	Image       string `json:"image"`
	UserID      uint   `json:"user_id"`
}

func AddProductToResponse(dataCore product.Core) AddProductResponse {
	return AddProductResponse{
		ID:          dataCore.ID,
		Name:        dataCore.Name,
		Price:       dataCore.Price,
		Quantity:    dataCore.Quantity,
		Description: dataCore.Description,
		Image:       dataCore.Image,
		UserID:      dataCore.UserID,
	}
}

type UpdateProductResponse struct {
	// ID          uint   `json:"id"`
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Quantity    uint   `json:"quantity"`
	Description string `json:"description"`
	Image       string `json:"image"`
	// UserID      uint   `json:"user_id"`
}

func UpdateProductToResponse(dataCore product.Core) AddProductResponse {
	return AddProductResponse{
		// ID:          dataCore.ID,
		Name:        dataCore.Name,
		Price:       dataCore.Price,
		Quantity:    dataCore.Quantity,
		Description: dataCore.Description,
		Image:       dataCore.Image,
		// UserID:      dataCore.UserID,
	}
}

//For AllProducts
type AllProductResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
	Image string `json:"image"`
}

func AllProductsToResponse(dataCore product.Core) AllProductResponse {
	return AllProductResponse{
		ID:    dataCore.ID,
		Name:  dataCore.Name,
		Price: dataCore.Price,
		Image: dataCore.Image,
	}
}
func ListAllProductsToResponse(dataCore []product.Core) []AllProductResponse {
	var DataResponse []AllProductResponse

	for _, value := range dataCore {
		DataResponse = append(DataResponse, AllProductsToResponse(value))
	}
	return DataResponse
}
