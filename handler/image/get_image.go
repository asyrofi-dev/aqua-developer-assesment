package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func (handler ImageHandler) GetImage(c echo.Context) error {
	path := c.Param("path")
	filename := c.Param("filename")

	imagePath := fmt.Sprintf("public/images/%s/%s", path, filename)

	return c.File(imagePath)
}
