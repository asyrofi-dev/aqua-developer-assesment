package handler

import (
	response "aqua-developer-assesment/entity/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (handler ProductHandler) ListProducts(c echo.Context) error {
	successResponse := response.SuccessResponse{Error: false}
	errorResponse := response.ErrorResponse{Error: true}

	result, err := handler.ProductUC.ListBooks()

	if err != nil {
		errorResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	successResponse.Message = "Get List Products Succeed!"
	successResponse.Data = result

	return c.JSON(http.StatusOK, successResponse)
}
