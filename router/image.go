package router

import (
	handler "aqua-developer-assesment/handler/image"

	"github.com/labstack/echo/v4"
)

func ImageRouter(e *echo.Echo) {
	handler := handler.NewImageHandler()

	base := e.Group("/api/v1")
	base.GET("/images/:path/:filename", handler.GetImage)
}
