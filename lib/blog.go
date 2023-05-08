package database

import (
	"project_petshop/config"
	"project_petshop/models"
)

func GetBlogsController() (interface{}, error) {
	var blogs []models.Blog

	if err := config.DB.Joins("User").Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func GetBlogController(blogID int) (interface{}, error) {
	var blog models.Blog
	blog.ID = uint(blogID)

	if err := config.DB.Joins("User").Find(&blog).Error; err != nil {
		return nil, err
	}

	return blog, nil
}

func CreateBlogController(b models.Blog) (interface{}, error) {
	if err := config.DB.Create(&b).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Joins("User").Find(&b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func UpdateBlogController(blogID uint, b models.Blog) (interface{}, error) {
	blog := models.Blog{}
	blog.ID = blogID
	if err := config.DB.Joins("User").Find(&blog).Error; err != nil {
		return nil, err
	}

	blog.Judul = b.Judul
	blog.Konten = b.Konten
	blog.IdUser = b.IdUser

	if err := config.DB.Save(&blog).Error; err != nil {
		return nil, err
	}

	return blog, nil
}

func DeleteBlogController(blogID int) (interface{}, error) {
	err := config.DB.Delete(&models.Blog{}, blogID).Error

	if err != nil {
		return nil, err
	}
	return blogID, nil
}
