package model

import (
	"cloudcomputing/webapp/config"
	"cloudcomputing/webapp/entity"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	guuid "github.com/google/uuid"
	"gopkg.in/alexcesaro/statsd.v2"
)

//GetAllCategories Fetch all category data
func GetAllCategories(Categories *[]entity.Category, client *statsd.Client) (err error) {
	t := client.NewTiming()
	if err = config.DB.Find(&Categories).Error; err != nil {
		return err
	}
	t.Send("get_all_categories.query_time")
	return nil
}

//CreateCategory ... Insert New category
func CreateCategory(category *entity.Category, client *statsd.Client) (err error) {
	t := client.NewTiming()
	category.ID = guuid.New().String()
	if err = config.DB.Create(&category).Error; err != nil {
		return err
	}
	t.Send("create_category.query_time")
	return nil
}

//GetCategoryByID ... Fetch only one Category by Id
func GetCategoryByID(category *entity.Category, id string, client *statsd.Client) (err error) {
	t := client.NewTiming()
	if err = config.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return err
	}
	t.Send("get_category_by_id.query_time")
	return nil
}

//GetCategoryByName ... Fetch only one Category by name
func GetCategoryByName(category *entity.Category, name string, client *statsd.Client) (err error) {
	t := client.NewTiming()
	if err = config.DB.Where("category = ?", name).First(&category).Error; err != nil {
		return err
	}
	t.Send("get_category_by_name.query_time")
	return nil
}

//UpdateCategory ... Update Category
func UpdateCategory(category *entity.Category, id string, client *statsd.Client) (err error) {
	t := client.NewTiming()
	config.DB.Save(&category)
	t.Send("update_category.query_time")
	return nil
}

//DeleteCategory ... Delete Category
func DeleteCategory(category *entity.Category, id string, client *statsd.Client) (err error) {
	//t := statsDClient.NewTiming()
	if config.DB.Where("id = ?", id).First(&category); category.ID == "" {
		return errors.New("the category doesn't exist!!!")
	}
	config.DB.Where("id = ?", id).Delete(&category)
	//t.Send("delete_category.query_time")
	return nil
}
