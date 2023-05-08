package controllers

import (
	"net/http"
	"strconv"
	"project_petshop/config"
	"project_petshop/models"

	"github.com/labstack/echo/v4"
)

func GetBlogsController(c echo.Context) error {
	var blogs []models.Blog

	if err := config.DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get data blogs",
		"data":    blogs,
	})
}

// get book by id
func GetBlogController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var blogs models.Blog
	if err = config.DB.Where("id = ?", id).First(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get blog by id",
		"user":    blogs,
	})
}

// create new book
func CreateBlogController(c echo.Context) error {
	blogs := models.Blog{}
	c.Bind(&blogs)

	if err := config.DB.Save(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog",
		"user":    blogs,
	})
}

// delete book by id
func DeleteBlogController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var blogs models.Blog
	if err := config.DB.First(&blogs, "id = ? ", id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Blog not found",
		})
	}

	if err := config.DB.Delete(&blogs).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to delete blog data",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully deleted blog data",
		"data":    blogs,
	})
}

// update book by id
func UpdateBlogController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Invalid Id",
		})
	}

	var blogs models.Blog
	if err := config.DB.Where("id = ?", id).First(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{
			"message": "Blog not found",
		})
	}
	c.Bind(&blogs)
	if err := config.DB.Model(&blogs).Updates(blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog data",
		"user":    blogs,
	})
}
