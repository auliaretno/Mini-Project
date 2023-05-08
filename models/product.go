package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	CategoryID  int      `json:"category_id" form:"category_id"`
	Category    Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductName string   `json:"product_name" form:"product_name"`
	Stock       int      `json:"stock" form:"stock"`
	Price       int      `json:"price" form:"price"`
}
