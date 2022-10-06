package router

import (
	"aqua-developer-assesment/config"
	handler "aqua-developer-assesment/handler/cart"
	repository "aqua-developer-assesment/repository/cart"
	usecase "aqua-developer-assesment/usecase/cart"

	"github.com/labstack/echo/v4"
)

func CartRouter(e *echo.Echo) {
	repository := repository.NewCartRepository(config.DB)
	usecase := usecase.NewCartUsecase(repository)
	handler := handler.NewCartHandler(usecase)

	base := e.Group("/api/v1")
	base.POST("/carts", handler.CreateCart)
	base.GET("/carts/:id", handler.GetCart)
}
