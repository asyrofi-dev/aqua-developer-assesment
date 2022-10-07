package handler

import (
	"aqua-developer-assesment/entity"
	response "aqua-developer-assesment/entity/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Create Cart godoc
// @Summary      Create new cart
// @Description  Create new cart
// @Tags         Cart
// @Accept       json
// @Produce      json
// @Param        cart   body      entity.CartRequest  true "new cart"
// @Success      201  {object}  response.SuccessResponse{data=entity.CartResponse}
// @Failure      400  {object}  response.ErrorResponse
// @Router       /carts [post]
func (handler CartHandler) CreateCart(c echo.Context) error {
	successResponse := response.SuccessResponse{Error: false}
	errorResponse := response.ErrorResponse{Error: true}

	var request entity.CartRequest

	err := c.Bind(&request)

	if err != nil {
		errorResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	result, err := handler.CartUC.CreateCart(request)

	if err != nil {
		errorResponse.Message = err.Error()
		return c.JSON(http.StatusBadRequest, errorResponse)
	}

	successResponse.Data = result
	successResponse.Message = "Create Cart Succeed"

	return c.JSON(http.StatusCreated, successResponse)
}
