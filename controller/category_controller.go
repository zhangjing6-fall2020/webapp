package controller

import (
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetCategories ... Get all Categories
func GetCategories(c *gin.Context) {
	var categories []entity.Category
	err := model.GetAllCategories(&categories)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

//CreateCategory ... Create Category
func CreateCategory(c *gin.Context) {
	var category entity.Category
	c.BindJSON(&category)

	err := model.CreateCategory(&category)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{"data": category})
	}
}

//GetCategoryByID ... Get the Category by id
func GetCategoryByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var category entity.Category
	err := model.GetCategoryByID(&category, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": category})
	}
}

func UpdateCategory(c *gin.Context) {
	var category entity.Category
	id := c.Params.ByName("id")
	err := model.GetCategoryByID(&category, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&category)
	err = model.UpdateCategory(&category, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": category})
	}
}

//DeleteCategory ... Delete the category
func DeleteCategory(c *gin.Context) {
	var category entity.Category
	id := c.Params.ByName("id")
	err := model.DeleteCategory(&category, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
