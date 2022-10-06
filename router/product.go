package router

import (
	"aqua-developer-assesment/config"
	handler "aqua-developer-assesment/handler/product"
	repository "aqua-developer-assesment/repository/product"
	usecase "aqua-developer-assesment/usecase/product"

	"github.com/labstack/echo/v4"
)

func ProductRouter(e *echo.Echo) {
	repository := repository.NewProductRepository(config.DB)
	usecase := usecase.NewProductUsecase(repository)
	handler := handler.NewProductHandler(usecase)

	base := e.Group("/api/v1")
	base.POST("/products", handler.CreateProduct)
	base.GET("/products", handler.ListProducts)
	base.GET("/products/:id", handler.GetProduct)
}
