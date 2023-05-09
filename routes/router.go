package routes

import (
	"project_petshop/constants"
	"project_petshop/controllers"
	m "project_petshop/middlewares"

	mid "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)

	eUser := e.Group("/users")
	eUser.POST("", controllers.CreateUserController)
	eUser.POST("/login", controllers.LoginUserController)
	// 	Authenticated with JWT
	eUserJwt := eUser.Group("")
	eUserJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eUserJwt.GET("", controllers.GetUsersController)
	eUserJwt.GET("/:id", controllers.GetUserController)
	eUserJwt.PUT("/:id", controllers.UpdateUserController)
	eUserJwt.DELETE("/:id", controllers.DeleteUserController)

	// blogs routes
	eBlog := e.Group("/blogs")
	eBlog.GET("", controllers.GetBlogsController)
	eBlog.GET("/:id", controllers.GetBlogController)
	// Authenticated with JWT
	eBlogJwt := eBlog.Group("")
	eBlogJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eBlogJwt.POST("", controllers.CreateBlogController)
	eBlogJwt.PUT("/:id", controllers.UpdateBlogController)
	eBlogJwt.DELETE("/:id", controllers.DeleteBlogController)

	//categories routes
	e.GET("/categories", controllers.GetCategoriesController)
	e.GET("/categories/:id", controllers.GetCategoryController)

	//categories routes admin
	e.POST("admin/category", controllers.CreateCategoryController)
	e.PUT("admin/categories/:id", controllers.UpdateCategoryController)
	e.DELETE("admin/categories/:id", controllers.DeleteCategoryController)

	//products routes
	e.GET("/products", controllers.GetProductsController)
	e.GET("/products/:id", controllers.GetProductController)

	//products routes admin
	e.POST("admin/products", controllers.CreateProductController)
	e.PUT("admin/products/:id", controllers.UpdateProductController)
	e.DELETE("admin/products/:id", controllers.DeleteProductController)

	return e
}
