package handler

import (
	response "aqua-developer-assesment/entity/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (handler CartHandler) GetCart(c echo.Context) error {
	successResponse := response.SuccessResponse{Error: false}
	errorResponse := response.ErrorResponse{Error: true}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		errorResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	result, err := handler.CartUC.GetCart(id)

	if err != nil {
		errorResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	successResponse.Message = "Get Data Succeed"
	successResponse.Data = result

	return c.JSON(http.StatusOK, successResponse)
}
