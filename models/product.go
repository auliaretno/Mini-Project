package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	CategoryId  int     `json:"idcategory" form:"idcategory"`
	Category    Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductName string `json:"productname" form:"productname"`
	Stock       int    `json:"stock" form:"stock"`
	Price       int    `json:"price" form:"price"`
}
