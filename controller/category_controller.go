package controller

import (
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/alexcesaro/statsd.v2"
	"net/http"
)

//GetCategories ... Get all Categories
func GetCategories(c *gin.Context, client *statsd.Client) {
	var categories []entity.Category
	err := model.GetAllCategories(&categories, client)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, categories)
}

//CreateCategory ... Create Category
func CreateCategory(c *gin.Context, client *statsd.Client) {
	var category entity.Category
	c.BindJSON(&category)

	err := model.CreateCategory(&category, client)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, category)
	}
}

//GetCategoryByID ... Get the Category by id
func GetCategoryByID(c *gin.Context, client *statsd.Client) {
	id := c.Params.ByName("id")
	var category entity.Category
	err := model.GetCategoryByID(&category, id, client)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, category)
	}
}

func UpdateCategory(c *gin.Context, client *statsd.Client) {
	var category entity.Category
	id := c.Params.ByName("id")
	err := model.GetCategoryByID(&category, id, client)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&category)
	err = model.UpdateCategory(&category, id, client)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, category)
	}
}

//DeleteCategory ... Delete the category
func DeleteCategory(c *gin.Context, client *statsd.Client) {
	var category entity.Category
	id := c.Params.ByName("id")
	err := model.DeleteCategory(&category, id, client)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
