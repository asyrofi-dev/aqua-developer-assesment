package usecase

import (
	"aqua-developer-assesment/entity"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/jinzhu/copier"
)

func (usecase ProductUsecase) CreateProduct(product entity.ProductRequest) (entity.ProductResponse, error) {
	var newProduct entity.Product
	var response entity.ProductResponse

	copier.Copy(&newProduct, &product)

	result, err := usecase.ProductRepo.CreateProduct(newProduct)

	if err != nil {
		return entity.ProductResponse{}, err
	}

	// Upload Photo

	src, err := product.Photo.Open()

	if err != nil {
		return entity.ProductResponse{}, err
	}

	defer src.Close()

	filename := fmt.Sprintf("Product_%d%s", result.ID, filepath.Ext(product.Photo.Filename))

	path := filepath.Join("public", "images", "products", filename)

	dst, err := os.Create(path)

	if err != nil {
		return entity.ProductResponse{}, err
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return entity.ProductResponse{}, err
	}

	// update result

	result.Photo = fmt.Sprintf("/api/v1/images/products/%s", filename)

	result, err = usecase.ProductRepo.UpdateProduct(result)

	if err != nil {
		return entity.ProductResponse{}, err
	}

	copier.Copy(&response, &result)

	return response, nil
}
