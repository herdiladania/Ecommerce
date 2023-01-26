package main

import (
	"e-commerce/config"
	"e-commerce/features/user/data"
	"e-commerce/features/user/handler"
	services "e-commerce/features/user/service"
	"e-commerce/migration"

	pdata "e-commerce/features/product/data"
	phdl "e-commerce/features/product/handler"
	psrv "e-commerce/features/product/service"

	cartData "e-commerce/features/cart/data"
	cartHandler "e-commerce/features/cart/handler"
	cartService "e-commerce/features/cart/service"

	orderData "e-commerce/features/order/data"
	orderHandler "e-commerce/features/order/handler"
	orderService "e-commerce/features/order/service"

	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	migration.Migrate(db)

	userData := data.New(db)
	userSrv := services.New(userData)
	userHdl := handler.New(userSrv)

	prodData := pdata.New(db)
	prodSrv := psrv.New(prodData)
	prodHdl := phdl.New(prodSrv)

	cartData := cartData.New(db)
	cartSrv := cartService.New(cartData)
	cartHdl := cartHandler.New(&cartSrv)

	orderData := orderData.New(db)
	orderSrv := orderService.New(orderData)
	orderHdl := orderHandler.New(orderSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	e.POST("/register", userHdl.Register())
	e.POST("/login", userHdl.Login())
	e.GET("/users", userHdl.Profile(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/users", userHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/users", userHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/products", prodHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/products/:id", prodHdl.Update(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/products/:id", prodHdl.GetProductById(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/products/:id", prodHdl.Delete(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/products", prodHdl.AllProducts())

	e.POST("/carts", cartHdl.AddCart(), middleware.JWT([]byte(config.JWT_KEY)))
	e.GET("/carts", cartHdl.MyCart(), middleware.JWT([]byte(config.JWT_KEY)))
	e.PUT("/carts/:id", cartHdl.UpdateCart(), middleware.JWT([]byte(config.JWT_KEY)))
	e.DELETE("/carts/:id", cartHdl.DeleteCart(), middleware.JWT([]byte(config.JWT_KEY)))

	e.POST("/orders", orderHdl.Add(), middleware.JWT([]byte(config.JWT_KEY)))

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
