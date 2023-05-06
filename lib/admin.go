package database

import (
	"project_petshop/config"
	"project_petshop/models"
)

func GetAdminsController() (interface{}, error) {
	var admins []models.Admin

	if err := config.DB.Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}

func GetAdminController(adminID uint) (interface{}, error) {
	var admin models.Admin
	admin.ID = adminID

	if err := config.DB.First(&admin).Error; err != nil {
		return nil, err
	}

	return admin, nil
}

func CreateAdminController(u models.Admin) (interface{}, error) {
	err := config.DB.Create(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func UpdateAdminController(adminID uint, u models.Admin) (interface{}, error) {
	admin := models.Admin{}
	admin.ID = adminID
	config.DB.First(&admin)

	admin.Name = u.Name
	admin.Email = u.Email
	admin.Password = u.Password

	err := config.DB.Save(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func DeleteAdminController(adminID int) (interface{}, error) {
	err := config.DB.Delete(&models.Admin{}, adminID).Error

	if err != nil {
		return nil, err
	}
	return adminID, nil
}
