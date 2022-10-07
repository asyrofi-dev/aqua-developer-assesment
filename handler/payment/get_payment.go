package handler

import (
	response "aqua-developer-assesment/entity/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Get Single Payment godoc
// @Summary      Get single payment
// @Description  Get single payment
// @Tags         Payment
// @Accept       json
// @Produce      json
// @Param        id   path      int  true "payment_id"
// @Success      201  {object}  response.SuccessResponse{data=entity.PaymentResponse}
// @Failure      400  {object}  response.ErrorResponse
// @Router       /payments/{id} [get]
func (handler PaymentHandler) GetPayment(c echo.Context) error {
	successResponse := response.SuccessResponse{Error: false}
	errorResponse := response.ErrorResponse{Error: true}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		errorResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	result, err := handler.PaymentUC.GetPayment(id)

	if err != nil {
		errorResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	successResponse.Message = "Get Data Succeed"
	successResponse.Data = result

	return c.JSON(http.StatusOK, successResponse)
}
