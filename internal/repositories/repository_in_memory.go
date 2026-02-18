package repositories

import "portal/internal/models"

var Products []models.Product

func ListProducts() []models.Product {
	return Products
}

func GetProductById(id string) models.Product {
	for _, p := range Products {
		if p.Id == id {
			return p
		}
	}
	return models.Product{}
}

func CreateProduct(p models.Product) {
	Products = append(Products, p)
}

func UpdateProduct(id string, product models.Product) models.Product {
	for i, p := range Products {
		if p.Id == id {
			Products[i].Name = product.Name
			Products[i].Price = product.Price
			return Products[i]
		}
	}
	return models.Product{}
}

func DeleteProduct(product models.Product) {
	for i, p := range Products {
		if product.Id == p.Id {
			Products = append(Products[:i], Products[i+1:]...)
			return
		}
	}
}
