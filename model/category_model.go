package model

import (
	"cloudcomputing/webapp/config"
	"cloudcomputing/webapp/entity"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	guuid "github.com/google/uuid"
)

//GetAllCategories Fetch all category data
func GetAllCategories(Categories *[]entity.Category) (err error) {
	if err = config.DB.Find(&Categories).Error; err != nil {
		return err
	}
	return nil
}

//CreateCategory ... Insert New category
func CreateCategory(category *entity.Category) (err error) {
	category.ID = guuid.New().String()
	if err = config.DB.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

//GetCategoryByID ... Fetch only one Category by Id
func GetCategoryByID(category *entity.Category, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return err
	}
	return nil
}

//GetCategoryByName ... Fetch only one Category by name
func GetCategoryByName(category *entity.Category, name string) (err error) {
	if err = config.DB.Where("category = ?", name).First(&category).Error; err != nil {
		return err
	}
	return nil
}

//UpdateCategory ... Update Category
func UpdateCategory(category *entity.Category, id string) (err error) {
	config.DB.Save(&category)
	return nil
}

//DeleteCategory ... Delete Category
func DeleteCategory(category *entity.Category, id string) (err error) {
	if config.DB.Where("id = ?", id).First(&category); category.ID == "" {
		return errors.New("the category doesn't exist!!!")
	}
	config.DB.Where("id = ?", id).Delete(&category)
	return nil
}
