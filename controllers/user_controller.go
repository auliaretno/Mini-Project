package controllers

import (
	"net/http"
	"strconv"
	"project_petshop/config"
	"project_petshop/models"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var user models.User
	if err = config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get user by id",
		"user":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var user models.User
	if err := config.DB.First(&user, "id = ? ", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "User not found",
		})
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete user data",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully deleted user data",
		"data":    user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var user models.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{
			"message": "User not found",
		})
	}
	c.Bind(&user)
	if err := config.DB.Model(&user).Updates(user).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user data",
		"user":    user,
	})
}
