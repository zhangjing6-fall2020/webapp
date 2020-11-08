package controller

import (
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/model"
	"cloudcomputing/webapp/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func getAllFilesByAnswer(answer entity.Answer) entity.Answer {
	var answerFiles []entity.AnswerFile
	if err := model.GetAllAnswerFilesByAnswerID(&answerFiles, answer.ID); err != nil{
		log.Error(err)
		fmt.Println(err)
	}
	for _, af := range answerFiles {
		var file entity.File
		if err := model.GetFileByID(&file, af.ID);err != nil{
			log.Error(err)
			fmt.Println(err)
		}
		answer.Attachments = append(answer.Attachments, file)
	}

	return answer
}

//GetAnswers ... Get all Answers
func GetAnswers(c *gin.Context) {
	var answers []entity.Answer
	err := model.GetAllAnswers(&answers)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}else{
		for _, a := range answers {
			a = getAllFilesByAnswer(a)
		}
		c.JSON(http.StatusOK, answers)
	}
}

//CreateAnswer ... Create Answer
func CreateAnswer(c *gin.Context, userID string) {
	log.Info("creating answer")
	var answer entity.Answer
	c.BindJSON(&answer)

	questionID, match := c.Params.Get("question_id")
	if !match {
		log.Error("can't get the question id from context")
		c.JSON(http.StatusNotFound, gin.H{
			"err":       "can't get the question id",
		})
		return
	}

	answer.QuestionID = questionID
	if err := model.GetQuestionByID(&answer.Question, answer.QuestionID); err != nil{
		log.Error("can't get the question by id")
		c.JSON(http.StatusNotFound, gin.H{
			"err":       "can't get the question id",
		})
		return
	}
	answer.UserID = userID
	if err := model.GetUserByID(&answer.User, userID); err != nil{
		log.Error("can't get the user by id")
		c.JSON(http.StatusNotFound, gin.H{
			"err":       "can't get the question id",
		})
		return
	}

	err := model.CreateAnswer(&answer)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	} else {
		log.Info("answer created")
		c.JSON(http.StatusCreated, answer)
	}
}

//GetAnswerByID ... Get the answer by id
func GetAnswerByID(c *gin.Context) {
	log.Info("getting answer by id")
	answerID := c.Params.ByName("answer_id")
	var answer entity.Answer
	if err := model.GetAnswerByID(&answer, answerID); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	questionID := c.Params.ByName("question_id")
	if answer.QuestionID != questionID {
		log.Error("The answer id and question id don't match")
		c.JSON(http.StatusNotFound, gin.H{
			"error": "The answer id and question id don't match!!!",
		})
		return
	}

	var question entity.Question
	if err := model.GetQuestionByID(&question, questionID); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	answer = getAllFilesByAnswer(answer)

	c.JSON(http.StatusOK, answer)
}

func UpdateAnswer(c *gin.Context, userID string) {
	log.Info("answer updating")
	var newAnswer entity.Answer
	c.BindJSON(&newAnswer)

	answerID := c.Params.ByName("answer_id")
	var currAnswer entity.Answer
	if err := model.GetAnswerByID(&currAnswer, answerID); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if currAnswer.UserID != userID {
		log.Error("only the user who post the answer can delete the answer")
		c.JSON(http.StatusForbidden, gin.H{
			"error": "only the user who post the answer can delete the answer",
		})
		return
	}

	questionID := c.Params.ByName("question_id")
	if currAnswer.QuestionID != questionID {
		log.Error("The answer id and question id don't match")
		c.JSON(http.StatusNotFound, gin.H{
			"error": "The answer id and question id don't match",
		})
		return
	}

	var question entity.Question
	if err := model.GetQuestionByID(&question, questionID); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	currAnswer.UserID = userID
	if err := model.GetUserByID(&currAnswer.User, userID); err != nil{
		log.Error("can't get the question id")
		c.JSON(http.StatusNotFound, gin.H{
			"err":       "can't get the question id",
		})
		return
	}

	if currAnswer.AnswerText == newAnswer.AnswerText {
		log.Error("the answerText is the same")
		c.JSON(http.StatusOK, gin.H{
			"err": "the answerText is the same, no need to update!",
		})
		return
	}
	currAnswer.AnswerText = newAnswer.AnswerText
	if err := model.UpdateAnswer(&currAnswer, answerID);err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		currAnswer = getAllFilesByAnswer(currAnswer)
		c.JSON(http.StatusOK, currAnswer)
	}
	log.Info("answer updated")
}

//DeleteAnswer ... Delete the Answer
func DeleteAnswer(c *gin.Context, userID string) {
	log.Info("answer deleting")
	questionID := c.Params.ByName("question_id")
	var question entity.Question
	if err := model.GetQuestionByID(&question, questionID); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	answerID := c.Params.ByName("answer_id")
	var answer entity.Answer
	if err := model.GetAnswerByID(&answer, answerID); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if answer.UserID != userID {
		log.Error("only the user who post the answer can delete the answer")
		c.JSON(http.StatusForbidden, gin.H{
			"error": "only the user who post the answer can delete the answer!",
		})
		return
	}

	if answer.QuestionID != questionID {
		log.Error("The answer id and question id don't match")
		c.JSON(http.StatusNotFound, gin.H{
			"error": "The answer id and question id don't match!!!",
		})
		return
	}

	answer = getAllFilesByAnswer(answer)
	if len(answer.Attachments) != 0 {
		for _,a := range answer.Attachments{
			var answerFile entity.AnswerFile
			if err := model.DeleteAnswerFileByID(&answerFile,a.ID,answerID);err != nil{
				log.Error(err)
				c.JSON(http.StatusNotFound, gin.H{
					"info": "can't delete the answer file",
					"err": err.Error(),
				})
				return
			}
			if err := tool.DeleteFile(tool.GetBucketName(), a.S3ObjectName); err != nil {
				log.Error(err)
				c.JSON(http.StatusNotFound, gin.H{
					"message": "can't delete the file from s3",
					"err": err.Error(),
				})
				return
			}
			if err := model.DeleteFile(&a,a.ID);err != nil {
				log.Error(err)
				c.JSON(http.StatusNotFound, gin.H{
					"info": "can't delete the file",
					"err": err.Error(),
				})
				return
			}
		}

		answer.Attachments = nil
	}

	err := model.DeleteAnswer(&answer, answerID)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + answerID: "is deleted"})
	}

	log.Info("answer deleted")
}
