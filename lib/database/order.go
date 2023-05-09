package database

import (
	"project_petshop/config"
	"project_petshop/models"
)

func GetOrders() (interface{}, error) {
	var orders []models.Order

	if err := config.DB.Joins("Product").Joins("User").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func GetOrderController(OrderID uint) (interface{}, error) {
	var order models.Order
	order.ID = OrderID

	if err := config.DB.Preload("Product").Preload("User").Find(&order).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func CreateOrder(b models.Order) (interface{}, error) {
	err := config.DB.Create(&b).Error

	if err != nil {
		return nil, err
	}

	return b, nil
}

func DeleteOrder(orderID int) (interface{}, error) {
	err := config.DB.Delete(&models.Order{}, orderID).Error

	if err != nil {
		return nil, err
	}
	return orderID, nil
}