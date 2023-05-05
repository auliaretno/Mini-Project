package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Pet   string `json:"pet" form:"pet"`
	Category   string `json:"category" form:"category"`
}
