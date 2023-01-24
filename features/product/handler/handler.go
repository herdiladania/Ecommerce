package handler

import (
	"e-commerce/features/product"
	"e-commerce/helper"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type productHandle struct {
	srv product.ProductService
}

func New(ps product.ProductService) product.ProductHandler {
	return &productHandle{
		srv: ps,
	}
}

func (ph *productHandle) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := AddProductRequest{}
		var image *multipart.FileHeader
		if err := c.Bind(&input); err != nil {
			log.Println("add product req body scan error")
			return c.JSON(http.StatusBadRequest, "wrong input format")
		}

		file, err := c.FormFile("image")
		if file != nil && err == nil {
			image = file
		} else if file != nil && err != nil {
			log.Println("error read product image")
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong image input"))
		}

		res, err := ph.srv.Add(token, *ToCore(input), image)
		if err != nil {
			log.Println("error running add product service")
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    AddProductToResponse(res),
			"message": "success add product",
		})
	}
}

func (ph *productHandle) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		var updateImage *multipart.FileHeader

		productID := c.Param("id")
		cnv, err := strconv.Atoi(productID)
		if err != nil {
			log.Println("update product param error")
			return c.JSON(http.StatusBadRequest, "wrong product ID")
		}

		input := UpdateProductRequest{}
		err2 := c.Bind(&input)
		if err2 != nil {
			log.Println("update product body scan error")
			return c.JSON(http.StatusBadRequest, "wrong input format")
		}

		file, err := c.FormFile("image")
		if file != nil && err == nil {
			updateImage = file
		} else if file != nil && err != nil {
			log.Println("error read update post photo")
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("wrong image input"))
		}

		res, err := ph.srv.Update(token, uint(cnv), *ToCore(input), updateImage)
		if err != nil {
			log.Println("error running update product service")
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    UpdateProductToResponse(res),
			"message": "success update product",
		})
	}
}

func (ph *productHandle) Delete() echo.HandlerFunc {
	return nil
}

func (ph *productHandle) GetProductById() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user")
		input := c.Param("id")
		cnv, err := strconv.Atoi(input)
		if err != nil {
			log.Println("GetProductById param error")
			return c.JSON(http.StatusBadRequest, "wrong product id")
		}

		res, err := ph.srv.GetProductById(token, uint(cnv))
		if err != nil {
			log.Println("error running GetProductById service")
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("server problem"))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    AddProductToResponse(res),
			"message": "success show product detail",
		})
	}
}
