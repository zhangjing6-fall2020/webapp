package controller

import (
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetQuestions ... Get all Questions
func GetQuestions(c *gin.Context) {
	var questions []entity.Question
	err := model.GetAllQuestions(&questions)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": questions})
}

//CreateQuestion ... Create Question
func CreateQuestion(c *gin.Context) {
	var question entity.Question
	c.BindJSON(&question)

	err := model.CreateQuestion(&question)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": question})
	}
}

//GetQuestionByID ... Get the Question by id
func GetQuestionByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var question entity.Question
	err := model.GetQuestionByID(&question, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": question})
	}
}

func UpdateQuestion(c *gin.Context) {
	var question entity.Question
	id := c.Params.ByName("id")
	err := model.GetQuestionByID(&question, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&question)
	err = model.UpdateQuestion(&question, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": question})
	}
}

//DeleteQuestion ... Delete the Question
func DeleteQuestion(c *gin.Context) {
	var question entity.Question
	id := c.Params.ByName("id")
	err := model.DeleteQuestion(&question, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
