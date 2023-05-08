package controllers

import (
	"net/http"
	"strconv"

	database "project_petshop/lib"
	"project_petshop/models"

	"github.com/labstack/echo/v4"
)

// get all products
func GetProductsController(c echo.Context) error {
	products, err := database.GetProducts()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all products",
		"products": products,
	})
}

// get product by id
func GetProductController(c echo.Context) error {
	Category, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product, err := database.GetProduct(Category)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get product",
		"products": product,
	})
}

// create new product
func CreateProductController(c echo.Context) error {
	product := models.Product{}
	c.Bind(&product)

	result, err := database.CreateProduct(product)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new product",
		"product": result,
	})
}

// update product by id
func UpdateProductController(c echo.Context) error {

	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product := models.Product{}
	c.Bind(&product)

	result, err := database.UpdateProduct(uint(productID), product)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update product",
		"products": result,
	})
}

// delete product by id
func DeleteProductController(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := database.DeleteProduct(productID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete product",
		"id":       result,
	})
}
