package cart

import "github.com/labstack/echo/v4"

type Core struct {
	ID         uint
	Quantity   int
	TotalPrice int
	ProductID  uint
	Name       string
	Image      string
	UserID     uint
}

type CartHandler interface {
	AddCart() echo.HandlerFunc
	MyCart() echo.HandlerFunc
	UpdateCart() echo.HandlerFunc
	DeleteCart() echo.HandlerFunc
}

type CartService interface {
	AddCartSrv(newCart Core) (Core, error)
	MyCartSrv(id uint) ([]Core, error)
	UpdateCartSrv(userID, cartID uint, updateCart Core) (Core, error)
	DeleteCartSrv(userID, cartID uint) error
}

type CartData interface {
	AddCartData(newCart Core) (Core, error)
	AddCHPData(cartID uint, newCart Core) (Core, error)
	MyCartData(id uint) ([]Core, error)
	UpdateCartData(userID, cartID uint, updateCart Core) (Core, error)
	DeleteCartData(userID, cartID uint) error
}
