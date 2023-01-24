package main

import (
	"e-commerce/config"
	"e-commerce/features/user/data"
	"e-commerce/features/user/handler"
	services "e-commerce/features/user/service"

	pdata "e-commerce/features/product/data"
	phdl "e-commerce/features/product/handler"
	psrv "e-commerce/features/product/service"

	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	userData := data.New(db)
	userSrv := services.New(userData)
	userHdl := handler.New(userSrv)

	prodData := pdata.New(db)
	prodSrv := psrv.New(prodData)
	prodHdl := phdl.New(prodSrv)

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

	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
