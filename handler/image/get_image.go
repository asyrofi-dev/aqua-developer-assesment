package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// Get Image Cart godoc
// @Summary      Get image
// @Description  Get image
// @Tags         Images
// @Accept       json
// @Produce      image/jpeg
// @Produce      image/png
// @Param        path   path      string  true "file path"
// @Param        filename   path      string  true "file name"
// @Router       /images/{path}/{filename} [get]
func (handler ImageHandler) GetImage(c echo.Context) error {
	path := c.Param("path")
	filename := c.Param("filename")

	imagePath := fmt.Sprintf("public/images/%s/%s", path, filename)

	return c.File(imagePath)
}
