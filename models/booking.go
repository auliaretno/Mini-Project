package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	PetcareID   int       `json:"petcare_id" form:"petcare_id"`
	Petcare     Petcare   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID      int       `json:"user_id" form:"user_id"`
	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BookingName string    `json:"booking_name" form:"booking_name"`
	BookingDate time.Time `json:"booking_date" form:"booking_date"`
}
