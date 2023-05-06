package controllers

import (
	"net/http"
	"strconv"
	"project_petshop/config"
	"project_petshop/models"

	"github.com/labstack/echo/v4"
)

// get all admins
func GetAdminsController(c echo.Context) error {
	var admins []models.Admin

	if err := config.DB.Find(&admins).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"admins":   admins,
	})
}

// get admin by id
func GetAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var admin models.Admin
	if err = config.DB.Where("id = ?", id).First(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get admin by id",
		"admin":    admin,
	})
}

// create new admin
func CreateAdminController(c echo.Context) error {
	admin := models.Admin{}
	c.Bind(&admin)

	if err := config.DB.Save(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new admin",
		"admin":    admin,
	})
}

// delete admin by id
func DeleteAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var admin models.Admin
	if err := config.DB.First(&admin, "id = ? ", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "admin not found",
		})
	}

	if err := config.DB.Delete(&admin).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete admin data",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully deleted admin data",
		"data":    admin,
	})
}

// update admin by id
func UpdateAdminController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var admin models.Admin
	if err := config.DB.Where("id = ?", id).First(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{
			"message": "admin not found",
		})
	}
	c.Bind(&admin)
	if err := config.DB.Model(&admin).Updates(admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update admin data",
		"admin":    admin,
	})
}
