package main

import (
	"aqua-developer-assesment/config"
	"aqua-developer-assesment/router"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Database()
	config.Migrate()

	e := echo.New()

	router.ProductRouter(e)
	router.CartRouter(e)
	router.PaymentRouter(e)

	e.Logger.Fatal(e.Start(":1323"))
}
