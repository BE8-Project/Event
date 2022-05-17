package category

import (
	"errors"
	"event/entity"

	"gorm.io/gorm"
)

type categoryModel struct {
	DB *gorm.DB
}

func NewCategotyModel(db *gorm.DB) *categoryModel {
	return &categoryModel{
		DB: db,
	}
}
func (c *categoryModel) Insert(categ entity.Category, id uint) (string, error) {
	var user entity.User
	if err := c.DB.Where("id=?", id).First(&user).Error; err != nil {
		return "", err
	}
	if user.Role != 1 {
		return "", errors.New("You can not Access")
	}
	if err := c.DB.Create(&categ).Error; err != nil {
		return "", errors.New("Can not Create category")
	}
	return "create category Success", nil
}
func (c *categoryModel) Get() ([]entity.Category, error) {
	var categ []entity.Category
	if err := c.DB.Find(&categ).Error; err != nil {
		return []entity.Category{}, err
	}
	return categ, nil
}
func (c *categoryModel) Delete(id_user, id_categ uint) (string, error) {
	var user entity.User
	if err := c.DB.Where("id=?", id_user).First(&user).Error; err != nil {
		return "", err
	}
	if user.Role != 1 {
		return "", errors.New("You can not Access")
	}
	var categ entity.Category
	if err := c.DB.Delete(&categ).Error; err != nil {
		return "", err
	}
	return "success delete Category", nil
}
