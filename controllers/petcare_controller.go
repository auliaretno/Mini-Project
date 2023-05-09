package controllers

import (
	"net/http"
	"project_petshop/config"
	"project_petshop/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all petcare
func GetPetcaresController(c echo.Context) error {
	var petcares []models.Petcare

	if err := config.DB.Find(&petcares).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success get all petcares",
		"petcares": petcares,
	})
}

// get petcare by id
func GetPetcareController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var petcare models.Petcare
	if err = config.DB.Where("id = ?", id).First(&petcare).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Success get petcare by id",
		"petcare": petcare,
	})
}

// create new petcare
func CreatePetcareController(c echo.Context) error {
	petcare := models.Petcare{}
	c.Bind(&petcare)

	if err := config.DB.Save(&petcare).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success create new cetegory",
		"petcare": petcare,
	})
}

// delete petcare by id
func DeletePetcareController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var petcare models.Petcare
	if err := config.DB.First(&petcare, "id = ? ", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "petcare not found",
		})
	}

	if err := config.DB.Delete(&petcare).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete petcare",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully deleted petcare data",
		"data":    petcare,
	})
}

