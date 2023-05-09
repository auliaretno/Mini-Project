package controllers

import (
	"net/http"
	"strconv"

	database "project_petshop/lib/database"
	"project_petshop/models"

	"github.com/labstack/echo/v4"
)

// get all bookings
func GetBookingsController(c echo.Context) error {
	bookings, err := database.GetBookings()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":  "success get all",
		"bookings": bookings,
	})
}

// get booking by id
func GetBookingController(c echo.Context) error {
	User, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	booking, err := database.GetBookingController(uint(User))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":  "success get by id",
		"bookings": booking,
	})
}

// create new booking
func CreateBookingController(c echo.Context) error {
	booking := models.Booking{}
	c.Bind(&booking)

	result, err := database.CreateBooking(booking)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new booking",
		"booking":    result,
	})
}


// delete booking by id
func DeleteBookingController(c echo.Context) error {
	bookingID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := database.DeleteBooking(bookingID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete product",
		"id":       result,
	})
}
