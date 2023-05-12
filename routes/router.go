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

	e.GET("/cookie", controllers.GetCookieHandler)

	// USER ROUTES
	// Not authenticated
	eUser := e.Group("/users")
	eUser.POST("/register", controllers.CreateUserController)
	eUser.POST("/login", controllers.LoginUserController)
	// 	Authenticated with JWT
	eUserJwt := eUser.Group("")
	eUserJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eUserJwt.GET("", controllers.GetUsersController)
	eUserJwt.GET("/:id", controllers.GetUserController)
	eUserJwt.PUT("/:id", controllers.UpdateUserController)
	eUserJwt.DELETE("/:id", controllers.DeleteUserController)

	// ADMIN ROUTES
	// Not authenticated
	eAdmin := e.Group("/admins")
	eAdmin.POST("/register", controllers.CreateAdminController)
	eAdmin.POST("/login", controllers.LoginAdminController)
	// 	Authenticated with JWT
	eAdminJwt := eAdmin.Group("")
	eAdminJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eAdminJwt.GET("", controllers.GetAdminsController)
	eAdminJwt.GET("/:id", controllers.GetAdminController)
	eAdminJwt.PUT("/:id", controllers.UpdateAdminController)
	eAdminJwt.DELETE("/:id", controllers.DeleteAdminController)

	//categories routes
	eCategory := e.Group("/categories")
	eCategory.GET("", controllers.GetCategoriesController)
	eCategory.GET("/:id", controllers.GetCategoryController)
	//categories routes admin
	eCategoryJwt := eCategory.Group("/admin")
	eCategoryJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eCategoryJwt.POST("", controllers.CreateCategoryController,)
	eCategoryJwt.PUT("/:id", controllers.UpdateCategoryController)
	eCategoryJwt.DELETE("/:id", controllers.DeleteCategoryController)

	//products routes
	eProducts := e.Group("/products")
	eProducts.GET("", controllers.GetProductsController)
	eProducts.GET("/:id", controllers.GetProductController)
	//products routes admin
	eProductsJwt := eProducts.Group("/admin")
	eProductsJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eProductsJwt.POST("", controllers.CreateProductController)
	eProductsJwt.PUT("/:id", controllers.UpdateProductController)
	eProductsJwt.DELETE("/:id", controllers.DeleteProductController)

	//Orders routes admin
	eOrders := e.Group("/orders")
	eOrders.POST("", controllers.CreateOrderController)
	eOrders.GET("/:id", controllers.GetOrderController)
	//Orders routes user
	eOrdersJwt := eOrders.Group("/admin")
	eOrdersJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eOrdersJwt.GET("", controllers.GetOrdersController)
	eOrdersJwt.DELETE("/:id", controllers.DeleteOrderController)

	//petcare routes
	ePetcare := e.Group("/petcare")
	ePetcare.GET("", controllers.GetCategoriesController)
	ePetcare.GET("/:id", controllers.GetPetcareController)
	//Petcare routes admin
	ePetcareJwt := ePetcare.Group("/admin")
	ePetcareJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	ePetcareJwt.POST("", controllers.CreatePetcareController)
	ePetcareJwt.DELETE("/:id", controllers.DeletePetcareController)

	//Booking routes admin
	eBooking := e.Group("/booking")
	eBooking.POST("", controllers.CreateBookingController)
	eBooking.GET("/:id", controllers.GetBookingController)
	//Booking routes user
	eBookingJwt := eBooking.Group("/admin")
	eBookingJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eBookingJwt.GET("", controllers.GetBookingController)
	eBookingJwt.DELETE("/:id", controllers.DeleteBookingController)

	return e
}
