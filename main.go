package main

import (
	"aqua-developer-assesment/config"
	"aqua-developer-assesment/router"

	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "aqua-developer-assesment/docs"
)

// @title ECommerce Management System API Documentation
// @version 1.0
// @description This is an API Documentation of ECommerce Management System.

// @contact.name Muhammad Asyrofi
// @contact.email asyrofimail@gmail.com

// @host localhost:1323
// @BasePath /api/v1

func main() {
	config.Database()
	config.Migrate()

	e := echo.New()

	router.ProductRouter(e)
	router.CartRouter(e)
	router.PaymentRouter(e)
	router.ImageRouter(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
