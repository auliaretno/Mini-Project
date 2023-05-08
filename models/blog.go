package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Id     string `json:"id" form:"id"`
	IdUser string `json:"iduser" form:"iduser"`
	Judul  string `json:"judul" form:"judul"`
	Konten string `json:"konten" form:"konten"`
}
