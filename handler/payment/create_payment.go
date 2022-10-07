package handler

import (
	"aqua-developer-assesment/entity"
	response "aqua-developer-assesment/entity/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Create Payment godoc
// @Summary      Create new payment
// @Description  Create new payment
// @Tags         Payment
// @Accept       mpfd
// @Produce      json
// @Param        payment   formData      entity.PaymentRequest  true "new payment"
// @Success      201  {object}  response.SuccessResponse{data=entity.PaymentResponse}
// @Failure      400  {object}  response.ErrorResponse
// @Router       /payments [post]
func (handler PaymentHandler) CreatePayment(c echo.Context) error {
	successResponse := response.SuccessResponse{Error: false}
	errorResponse := response.ErrorResponse{Error: true}

	var request entity.PaymentRequest

	err := c.Bind(&request)

	if err != nil {
		errorResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	photo, err := c.FormFile("photo")

	if err != nil {
		errorResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	request.Photo = photo

	result, err := handler.PaymentUC.CreatePayment(request)

	if err != nil {
		errorResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	successResponse.Data = result
	successResponse.Message = "Create Payment Succeed"

	return c.JSON(http.StatusCreated, successResponse)
}
