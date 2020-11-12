package controller

import (
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/alexcesaro/statsd.v2"
	"net/http"
)

//GetQuestionCategories ... Get all QuestionCategories
func GetQuestionCategories(c *gin.Context, client *statsd.Client) {
	var questionCategories []entity.QuestionCategory
	err := model.GetAllQuestionCategories(&questionCategories, client)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, questionCategories)
}

//CreateQuestionCategory ... Create QuestionCategory
func CreateQuestionCategory(c *gin.Context, client *statsd.Client) {
	var questionCategory entity.QuestionCategory
	c.BindJSON(&questionCategory)

	err := model.CreateQuestionCategory(&questionCategory, client)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, questionCategory)
	}
}

//GetQuestionCategoryByQuestionID ... Get the QuestionCategory by QuestionID
func GetQuestionCategoryByQuestionID(c *gin.Context, client *statsd.Client) {
	questionID := c.Params.ByName("question_id")
	var questionCategory entity.QuestionCategory
	err := model.GetQuestionCategoryByQuestionID(&questionCategory, questionID, client)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, questionCategory)
	}
}

//GetQuestionCategoryByCategoryID ... Get the QuestionCategory by CategoryID
func GetQuestionCategoryByCategoryID(c *gin.Context, client *statsd.Client) {
	categoryID := c.Params.ByName("category_id")
	var questionCategory entity.QuestionCategory
	err := model.GetQuestionCategoryByCategoryID(&questionCategory, categoryID, client)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, questionCategory)
	}
}

func UpdateQuestionCategory(c *gin.Context, client *statsd.Client) {
	var questionCategory entity.QuestionCategory
	questionID := c.Params.ByName("question_id")
	categoryID := c.Params.ByName("category_id")
	err := model.GetQuestionCategoryByIDs(&questionCategory, questionID, categoryID, client)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&questionCategory)
	err = model.UpdateQuestionCategory(&questionCategory, client)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, questionCategory)
	}
}

//DeleteQuestionCategory ... Delete the QuestionCategory
func DeleteQuestionCategory(c *gin.Context, client *statsd.Client) {
	var questionCategory entity.QuestionCategory
	questionID := c.Params.ByName("question_id")
	categoryID := c.Params.ByName("category_id")
	if err := model.GetQuestionCategoryByIDs(&questionCategory, questionID, categoryID, client); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.DeleteQuestionCategory(&questionCategory, questionID, categoryID, client); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"question_id" + questionID + "category_id" + categoryID: "is deleted"})
	}
}
