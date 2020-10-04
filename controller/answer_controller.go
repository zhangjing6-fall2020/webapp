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

	c.JSON(http.StatusOK, answers)
}

//CreateAnswer ... Create Answer
func CreateAnswer(c *gin.Context, userID string) {
	var answer entity.Answer
	c.BindJSON(&answer)

	questionID, match := c.Params.Get("question_id")
	if !match {
		c.JSON(http.StatusNotFound, gin.H{
			"err":       "can't get the question id",
		})
		return
	}

	answer.QuestionID = questionID
	if err := model.GetQuestionByID(&answer.Question, answer.QuestionID); err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"err":       "can't get the question id",
		})
		return
	}
	answer.UserID = userID
	if err := model.GetUserByID(&answer.User, userID); err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"err":       "can't get the question id",
		})
		return
	}

	err := model.CreateAnswer(&answer)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, answer)
	}


}

//GetAnswerByID ... Get the answer by id
func GetAnswerByID(c *gin.Context) {
	questionID := c.Params.ByName("question_id")
	var question entity.Question
	model.GetQuestionByID(&question, questionID)

	answerID := c.Params.ByName("answer_id")
	var answer entity.Answer
	err := model.GetAnswerByID(&answer, answerID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, answer)
	}
}

func UpdateAnswer(c *gin.Context, userID string) {
	var newAnswer entity.Answer
	c.BindJSON(&newAnswer)

	questionID := c.Params.ByName("question_id")
	var question entity.Question
	if err := model.GetQuestionByID(&question, questionID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	answerID := c.Params.ByName("answer_id")
	var currAnswer entity.Answer
	if err := model.GetAnswerByID(&currAnswer, answerID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	currAnswer.UserID = userID
	if err := model.GetUserByID(&currAnswer.User, userID); err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"err":       "can't get the question id",
		})
		return
	}

	if currAnswer.AnswerText == newAnswer.AnswerText {
		c.JSON(http.StatusOK, gin.H{
			"err": "the answerText is the same, no need to update1",
		})
	}
	currAnswer.AnswerText = newAnswer.AnswerText
	if err := model.UpdateAnswer(&currAnswer, answerID);err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, currAnswer)
	}
}

//DeleteAnswer ... Delete the Answer
func DeleteAnswer(c *gin.Context, userID string) {
	questionID := c.Params.ByName("question_id")
	var question entity.Question
	if err := model.GetQuestionByID(&question, questionID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	answerID := c.Params.ByName("answer_id")
	var answer entity.Answer
	if err := model.GetAnswerByID(&answer, answerID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	answer.UserID = userID
	if err := model.GetUserByID(&answer.User, userID); err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"err":       "can't get the question id",
		})
		return
	}

	err := model.DeleteAnswer(&answer, answerID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + answerID: "is deleted"})
	}
}
