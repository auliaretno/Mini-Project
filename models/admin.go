package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Role     string `json:"role" form:"role" gorm:"type:enum('Admin', 'User');default:'Admin'; not-null"`
	Password string `json:"password" form:"password"`
}

type AdminResponse struct {
	ID    uint   `json:"id" form:"id"`
	Name  string `json:"name" form:"name" `
	Email string `json:"email" form:"email"`
	Token string `json:"token" form:"token"`
}
