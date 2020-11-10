package controller

import (
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetQuestionCategories ... Get all QuestionCategories
func GetQuestionCategories(c *gin.Context) {
	var questionCategories []entity.QuestionCategory
	err := model.GetAllQuestionCategories(&questionCategories)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, questionCategories)
}

//CreateQuestionCategory ... Create QuestionCategory
func CreateQuestionCategory(c *gin.Context) {
	var questionCategory entity.QuestionCategory
	c.BindJSON(&questionCategory)

	err := model.CreateQuestionCategory(&questionCategory)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, questionCategory)
	}
}

//GetQuestionCategoryByQuestionID ... Get the QuestionCategory by QuestionID
func GetQuestionCategoryByQuestionID(c *gin.Context) {
	questionID := c.Params.ByName("question_id")
	var questionCategory entity.QuestionCategory
	err := model.GetQuestionCategoryByQuestionID(&questionCategory, questionID)
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
func GetQuestionCategoryByCategoryID(c *gin.Context) {
	categoryID := c.Params.ByName("category_id")
	var questionCategory entity.QuestionCategory
	err := model.GetQuestionCategoryByCategoryID(&questionCategory, categoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, questionCategory)
	}
}

func UpdateQuestionCategory(c *gin.Context) {
	var questionCategory entity.QuestionCategory
	questionID := c.Params.ByName("question_id")
	categoryID := c.Params.ByName("category_id")
	err := model.GetQuestionCategoryByIDs(&questionCategory, questionID, categoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&questionCategory)
	err = model.UpdateQuestionCategory(&questionCategory)
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
func DeleteQuestionCategory(c *gin.Context) {
	var questionCategory entity.QuestionCategory
	questionID := c.Params.ByName("question_id")
	categoryID := c.Params.ByName("category_id")
	if err := model.GetQuestionCategoryByIDs(&questionCategory, questionID, categoryID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.DeleteQuestionCategory(&questionCategory, questionID, categoryID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"question_id" + questionID + "category_id" + categoryID: "is deleted"})
	}
}
