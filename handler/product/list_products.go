package handler

import (
	"aqua-developer-assesment/entity"
	response "aqua-developer-assesment/entity/response"
	"fmt"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

// Get All Product godoc
// @Summary      Get all product
// @Description  Get all product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param        category   query      string  false "product category"
// @Param        min_price   query      int  false "product minimum price"
// @Param        max_price   query      int  false "product maximum price"
// @Success      201  {object}  response.SuccessResponse{data=[]entity.ProductResponse}
// @Failure      400  {object}  response.ErrorResponse
// @Router       /products [get]
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
