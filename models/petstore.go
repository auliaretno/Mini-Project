package models

import (
	"gorm.io/gorm"
)

type Petstore struct {
	gorm.Model
	IdCategory uint `json:"idcategory" form:"idcategory"`
	ProductName  string `json:"productname" form:"productname"`
	Stock int `json:"stock" form:"stock"`
	Price int `json:"price" form:"price"`
}
