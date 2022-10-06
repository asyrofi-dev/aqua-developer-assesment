package router

import (
	"aqua-developer-assesment/config"
	handler "aqua-developer-assesment/handler/payment"
	repository "aqua-developer-assesment/repository/payment"
	usecase "aqua-developer-assesment/usecase/payment"

	"github.com/labstack/echo/v4"
)

func PaymentRouter(e *echo.Echo) {
	repository := repository.NewPaymentRepository(config.DB)
	usecase := usecase.NewPaymentUsecase(repository)
	handler := handler.NewPaymentHandler(usecase)

	base := e.Group("/api/v1")
	base.POST("/payments", handler.CreatePayment)
	base.GET("/payments/:id", handler.GetPayment)
}
