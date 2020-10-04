package controller

import (
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderedMap map[string]string

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
func CreateQuestion(c *gin.Context, userID string) {
	var question entity.Question
	c.BindJSON(&question)

	//get the current user
	question.UserID = userID
	model.GetUserByID(&question.User, userID)

	//check each category exists or not:
	//if exist, continue
	//if doesn't exist, create the category
	var categoryIDs []string
	//in order to avoid duplicate categories in the same question input:
	//tbd
	for i, category := range question.Categories {
		err := model.GetCategoryByName(&category, category.Category)
		//if the category isn't found, it will return `record not found`
		if err == nil {
			categoryIDs = append(categoryIDs, category.ID)
		} else if err.Error() == "record not found" {
			if err := model.CreateCategory(&category); err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}
			categoryIDs = append(categoryIDs, category.ID)
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		question.Categories[i] = category
	}

	//create the question
	err := model.CreateQuestion(&question)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	//for each category, create questionCategory with questionID and categoryID
	for _, c := range categoryIDs {
		var questionCategory entity.QuestionCategory
		questionCategory.CategoryID = c
		questionCategory.QuestionID = question.ID
		model.CreateQuestionCategory(&questionCategory)
	}

	c.JSON(http.StatusCreated, gin.H{
		"question_id":       question.ID,
		"created_timestamp": question.CreatedTimestamp,
		"updated_timestamp": question.UpdatedTimestamp,
		"user_id":           question.UserID,
		"question_text":     question.QuestionText,
		"categories":        question.Categories,
		"answers":           question.Answers,
	})
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

func UpdateQuestion(c *gin.Context, userID string) {
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
