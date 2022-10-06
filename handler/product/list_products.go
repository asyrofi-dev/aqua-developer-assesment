package handler

import (
	"aqua-developer-assesment/entity"
	response "aqua-developer-assesment/entity/response"
	"fmt"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

func (handler ProductHandler) ListProducts(c echo.Context) error {
	successResponse := response.SuccessResponse{Error: false}
	errorResponse := response.ErrorResponse{Error: true}

	// var filter entity.ProductFilter
	var result []entity.ProductResponse

	var filter entity.ProductFilter
	defaultFilter := entity.ProductFilter{Category: "", MinPrice: 0, MaxPrice: 999999999}

	err := c.Bind(&filter)

	if err != nil {
		errorResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	if filter == (entity.ProductFilter{}) {
		result, err = handler.ProductUC.ListProducts()
	} else {
		copier.CopyWithOption(&defaultFilter, &filter, copier.Option{IgnoreEmpty: true})

		defaultFilter.Category = fmt.Sprintf("%%%s%%", defaultFilter.Category)
		result, err = handler.ProductUC.SearchProducts(defaultFilter)
	}

	if err != nil {
		errorResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	successResponse.Message = "Get List Products Succeed!"
	successResponse.Data = result

	return c.JSON(http.StatusOK, successResponse)
}
