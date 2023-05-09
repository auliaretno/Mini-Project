package models

import (
	"gorm.io/gorm"
)

type Petcare struct {
	gorm.Model
	Pet     string `json:"pet" form:"pet"`
	Service string `json:"service" form:"service"`
	Price   int    `json:"price" form:"price"`
}
