package routes

import (
	"project_petshop/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	// Route users / to handler function
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	// Route admin / to handler function
	e.GET("/admin", controllers.GetAdminsController)
	e.GET("/admin/:id", controllers.GetAdminController)
	e.POST("/admin", controllers.CreateAdminController)
	e.DELETE("/admin/:id", controllers.DeleteAdminController)
	e.PUT("/admin/:id", controllers.UpdateAdminController)

		//categories routes
	e.GET("/categories", controllers.GetCategoriesController)
	e.GET("/categories/:id", controllers.GetCategoryController)
	e.POST("/categories", controllers.CreateCategoryController)
	e.PUT("/categories/:id", controllers.UpdateCategoryController)
	e.DELETE("/categories/:id", controllers.DeleteCategoryController)

	//products routes
	e.GET("/products", controllers.GetProductsController)
	e.GET("/products/:id", controllers.GetProductController)
	e.POST("/products", controllers.CreateProductController)
	e.PUT("/products/:id", controllers.UpdateProductController)
	e.DELETE("/products/:id", controllers.DeleteProductController)

	return e

}
