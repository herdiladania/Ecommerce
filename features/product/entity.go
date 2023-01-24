package product

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID          uint
	Name        string
	Price       uint
	Quantity    uint
	Description string
	Image       string
	UserID      uint
	// CreatedAt   time.Time
}

type ProductHandler interface {
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	// MyPosts() echo.HandlerFunc
	// AllPosts() echo.HandlerFunc
	// GetPostById() echo.HandlerFunc
}

type ProductService interface {
	Add(token interface{}, newProduct Core, image *multipart.FileHeader) (Core, error)
	Update(token interface{}, productID uint, updateProduct Core, updateImage *multipart.FileHeader) (Core, error)
	Delete(token interface{}, productID uint) error
	// MyPosts(token interface{}) ([]MyPostsResp, error)
	// AllPosts() ([]MyPostsResp, error)
	// GetPostById(token interface{}, postID uint) (MyPostsResp, error)
}

type ProductData interface {
	Add(userID uint, newProduct Core) (Core, error)
	Update(productID uint, userID uint, updateProduct Core) (Core, error)
	Delete(productID uint, userID uint) error
	// MyPosts(userID uint) ([]MyPostsResp, error)
	// AllPosts() ([]MyPostsResp, error)
	// GetPostById(postID uint, userID uint) (MyPostsResp, error)
}
