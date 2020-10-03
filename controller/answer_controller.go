package controller

import (
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetAnswers ... Get all Answers
func GetAnswers(c *gin.Context) {
	var answers []entity.Answer
	err := model.GetAllAnswers(&answers)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": answers})
}

//CreateAnswer ... Create Answer
func CreateAnswer(c *gin.Context) {
	var answer entity.Answer
	c.BindJSON(&answer)

	err := model.CreateAnswer(&answer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": answer})
	}
}

//GetAnswerByID ... Get the answer by id
func GetAnswerByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var answer entity.Answer
	err := model.GetAnswerByID(&answer, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": answer})
	}
}

func UpdateAnswer(c *gin.Context) {
	var answer entity.Answer
	id := c.Params.ByName("id")
	err := model.GetAnswerByID(&answer, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&answer)
	err = model.UpdateAnswer(&answer, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": answer})
	}
}

//DeleteAnswer ... Delete the Answer
func DeleteAnswer(c *gin.Context) {
	var answer entity.Answer
	id := c.Params.ByName("id")
	err := model.DeleteAnswer(&answer, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}