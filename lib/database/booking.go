package database

import (
	"project_petshop/config"
	"project_petshop/models"
)

func GetBookings() (interface{}, error) {
	var bookings []models.Booking

	if err := config.DB.Joins("Petcare").Joins("User").Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

func GetBookingController(BookingID uint) (interface{}, error) {
	var booking models.Booking
	booking.ID = BookingID

	if err := config.DB.Preload("Petcare").Preload("User").Find(&booking).Error; err != nil {
		return nil, err
	}

	return booking, nil
}

func CreateBooking(b models.Booking) (interface{}, error) {
	err := config.DB.Create(&b).Error

	if err != nil {
		return nil, err
	}

	return b, nil
}

func DeleteBooking(bookingID int) (interface{}, error) {
	err := config.DB.Delete(&models.Booking{}, bookingID).Error

	if err != nil {
		return nil, err
	}
	return bookingID, nil
}
