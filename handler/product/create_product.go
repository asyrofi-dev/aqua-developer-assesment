package handler

import (
	"aqua-developer-assesment/entity"
	response "aqua-developer-assesment/entity/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (handler ProductHandler) CreateProduct(c echo.Context) error {
	successResponse := response.SuccessResponse{Error: false}
	errorResponse := response.ErrorResponse{Error: true}

	var request entity.ProductRequest

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

	result, err := handler.ProductUC.CreateProduct(request)

	if err != nil {
		errorResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	successResponse.Data = result

	return c.JSON(http.StatusCreated, successResponse)
}
