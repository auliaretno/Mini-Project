package database

import (
	"project_petshop/config"
	"project_petshop/models"
)

func GetProducts() (interface{}, error) {
	var products []models.Product

	if err := config.DB.Joins("Category = ?").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProduct(productID int) (interface{}, error) {
	var product models.Product
	product.ID = uint(productID)

	if err := config.DB.Joins("Category = ?").Find(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func CreateProduct(b models.Product) (interface{}, error) {
	if err := config.DB.Create(&b).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Joins("Category = ?").Find(&b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func UpdateProduct(productID uint, b models.Product) (interface{}, error) {
	product := models.Product{}
	product.ID = productID
	if err := config.DB.Joins("Category = ?").Find(&product).Error; err != nil {
		return nil, err
	}

	product.CategoryId = b.CategoryId
	product.ProductName = b.ProductName
	product.Stock = b.Stock
	product.Price = b.Price

	if err := config.DB.Save(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func DeleteProduct(productID int) (interface{}, error) {
	err := config.DB.Delete(&models.Product{}, productID).Error

	if err != nil {
		return nil, err
	}
	return productID, nil
}
