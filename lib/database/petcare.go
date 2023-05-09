package database

import (
	"project_petshop/config"
	"project_petshop/models"
)

func GetPetcaresController() (interface{}, error) {
	var petcares []models.Petcare

	if err := config.DB.Find(&petcares).Error; err != nil {
		return nil, err
	}
	return petcares, nil
}

func GetPetcareController(PetcareID uint) (interface{}, error) {
	var petcare models.Petcare
	petcare.ID = PetcareID

	if err := config.DB.First(&petcare).Error; err != nil {
		return nil, err
	}

	return petcare, nil
}

func CreatePetcareController(u models.Petcare) (interface{}, error) {
	err := config.DB.Create(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func DeletePetcareController(PetcareID int) (interface{}, error) {
	err := config.DB.Delete(&models.Petcare{}, PetcareID).Error

	if err != nil {
		return nil, err
	}
	return PetcareID, nil
}

