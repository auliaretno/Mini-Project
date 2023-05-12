package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"project_petshop/config"
	"project_petshop/models"
)

func TestGetCategoriesController(t *testing.T) {
	// Set up a new Echo instance
	e := echo.New()

	// Set up a new request and recorder
	req := httptest.NewRequest(http.MethodGet, "/categories", nil)
	rec := httptest.NewRecorder()

	// Create a new database connection and migrate the schema
	db := config.DB
	db.AutoMigrate(&models.Category{})

	// Add some test data to the database
	db.Create(&models.Category{Pet: "Cats"})
	db.Create(&models.Category{Pet: "Dogs"})
	db.Create(&models.Category{Pet: "Fish"})

	// Call the handler function
	c := e.NewContext(req, rec)
	err := GetCategoriesController(c)

	// Check the response status code
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Check the response body
	var response struct {
		Message    string             `json:"message"`
		Categories []models.Category `json:"categories"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success get all categories", response.Message)
	assert.Len(t, response.Categories, 3)
	assert.Equal(t, "Cats", response.Categories[0].Pet)
	assert.Equal(t, "Dogs", response.Categories[1].Pet)
	assert.Equal(t, "Fish", response.Categories[2].Pet)
}

func TestGetCategoryController(t *testing.T) {
	// Set up a new Echo instance
	e := echo.New()

	// Set up a new request and recorder
	req := httptest.NewRequest(http.MethodGet, "/categories/1", nil)
	rec := httptest.NewRecorder()

	// Create a new database connection and migrate the schema
	db := config.DB
	db.AutoMigrate(&models.Category{})

	// Add some test data to the database
	db.Create(&models.Category{Pet: "Cats"})
	db.Create(&models.Category{Pet: "Dogs"})
	db.Create(&models.Category{Pet: "Fish"})

	// Call the handler function
	c := e.NewContext(req, rec)
	c.SetPath("/categories/:id")
	c.SetParamValues("1")
	err := GetCategoryController(c)

	// Check the response status code
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	// Check the response body
	var response struct {
		Message  string          `json:"message"`
		Category models.Category `json:"category"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Success get category by id", response.Message)
	assert.Equal(t, "Cats", response.Category.Pet)
}

func TestCreateCategoryController(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/categories", strings.NewReader(`{"name": "Test Category"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call the controller function
	err := CreateCategoryController(c)

	// Assert the response
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteCategoryController(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/categories/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Call the controller function
	err := DeleteCategoryController(c)

	// Assert the response
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateCategoryController(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/categories/1", strings.NewReader(`{"name": "Updated Category"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Call the controller function
	err := UpdateCategoryController(c)

	// Assert the response
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}






