package handler

import (
	response "aqua-developer-assesment/entity/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Get Single Product godoc
// @Summary      Get single product
// @Description  Get single product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        id   path      int  true "product_id"
// @Success      201  {object}  response.SuccessResponse{data=entity.ProductResponse}
// @Failure      400  {object}  response.ErrorResponse
// @Router       /products/{id} [get]
func (handler ProductHandler) GetProduct(c echo.Context) error {
	successResponse := response.SuccessResponse{Error: false}
	errorResponse := response.ErrorResponse{Error: true}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		errorResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	result, err := handler.ProductUC.GetProduct(id)

	if err != nil {
		errorResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	successResponse.Message = "Get Data Succeed"
	successResponse.Data = result

	return c.JSON(http.StatusOK, successResponse)
}
