package services

import (
	"portal/internal/models"
	"portal/internal/repositories"
	"portal/internal/validators"

	"github.com/google/uuid"
)

func ListProducts() []models.Product {
	return repositories.ListProducts()
}

func GetProductById(id string) models.Product {
	return repositories.GetProductById(id)
}

func CreateProduct(p models.Product) (models.Product, error) {
	err := validators.ProductValidate(p)

	if err != nil {
		return p, err
	}

	product := models.Product{
		Id:    uuid.New().String(),
		Name:  p.Name,
		Price: p.Price,
	}

	repositories.CreateProduct(product)
	return product, nil
}

func UpdateProduct(id string, p models.Product) (models.Product, error) {
	err := validators.ProductValidate(p)

	if err != nil {
		return p, err
	}

	product := GetProductById(id)
	product.Name = p.Name
	product.Price = p.Price
	response := repositories.UpdateProduct(id, product)
	return response, nil
}

func DeleteProduct(id string) error {
	p := GetProductById(id)
	repositories.DeleteProduct(p)
	return nil
}
