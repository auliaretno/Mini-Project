package controllers

import (
	"net/http"
	"strconv"

	database "project_petshop/lib/database"
	"project_petshop/models"

	"github.com/labstack/echo/v4"
)

// get all orders
func GetOrdersController(c echo.Context) error {
	orders, err := database.GetOrders()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":  "success get all products",
		"orders": orders,
	})
}

// get order by id
func GetOrderController(c echo.Context) error {
	User, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	order, err := database.GetOrderController(uint(User))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":  "success get product",
		"orders": order,
	})
}

// create new order
func CreateOrderController(c echo.Context) error {
	order := models.Order{}
	c.Bind(&order)

	result, err := database.CreateOrder(order)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new order",
		"order":    result,
	})
}


// delete order by id
func DeleteOrderController(c echo.Context) error {
	orderID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := database.DeleteOrder(orderID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete product",
		"id":       result,
	})
}
