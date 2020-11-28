package controller

import (
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/model"
	"cloudcomputing/webapp/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alexcesaro/statsd.v2"
	"net/http"
)

func getAllFilesByAnswer(answer entity.Answer, client *statsd.Client) entity.Answer {
	var answerFiles []entity.AnswerFile
	if err := model.GetAllAnswerFilesByAnswerID(&answerFiles, answer.ID, client); err != nil {
		log.Error(err)
		fmt.Println(err)
	}
	for _, af := range answerFiles {
		var file entity.File
		if err := model.GetFileByID(&file, af.ID, client); err != nil {
			log.Error(err)
			fmt.Println(err)
		}
		answer.Attachments = append(answer.Attachments, file)
	}

	return answer
}

//GetAnswers ... Get all Answers
func GetAnswers(c *gin.Context, client *statsd.Client) {
	var answers []entity.Answer
	err := model.GetAllAnswers(&answers, client)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		for _, a := range answers {
			a = getAllFilesByAnswer(a, client)
		}
		c.JSON(http.StatusOK, answers)
	}
}

//CreateAnswer ... Create Answer
func CreateAnswer(c *gin.Context, userID string, client *statsd.Client) {
	log.Info("creating answer")
	var answer entity.Answer
	c.BindJSON(&answer)

	questionID, match := c.Params.Get("question_id")
	if !match {
		log.Error("can't get the question id from context")
		c.JSON(http.StatusNotFound, gin.H{
			"err": "can't get the question id",
		})
		return
	}

	answer.QuestionID = questionID
	if err := model.GetQuestionByID(&answer.Question, answer.QuestionID, client); err != nil {
		log.Error("can't get the question by id")
		c.JSON(http.StatusNotFound, gin.H{
			"err": "can't get the question id",
		})
		return
	}
	answer.UserID = userID
	if err := model.GetUserByID(&answer.User, userID, client); err != nil {
		log.Error("can't get the user by id")
		c.JSON(http.StatusNotFound, gin.H{
			"err": "can't get the question id",
		})
		return
	}

	err := model.CreateAnswer(&answer, client)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}

	log.Info("answer created")
	c.JSON(http.StatusCreated, answer)

	/*post a message on SNS topic, including:
	1.Question details such as ID, user's email address
	2.Answer details such as ID, answer text, etc
	3.HTTP link to the question & answer (created or updated):
	http://prod.bh7cw.me:80/v1/question/{questionId}/answer/{answerId}*/
	message01 := fmt.Sprintf("QuestionID: %v, QuestionText: %v, UserName: %v %v, UserEmail: %v",
		questionID, answer.Question.QuestionText, answer.User.FirstName, answer.User.LastName, answer.User.Username)

	message02 := fmt.Sprintf("AnswerID: %v, AnswerText: %v", answer.ID, answer.AnswerText)

	message03 := fmt.Sprintf("Link: http://prod.bh7cw.me:80/v1/question/%v/answer/%v", questionID, answer.ID)

	message := fmt.Sprintf("%v, %v, %v", message01, message02, message03)
	result, err := tool.PublishMessageOnSNS(message)
	if err != nil {
		log.Errorf("Publish error: %v", err)

	}

	log.Info("Publish message result: %v", result)
}

//GetAnswerByID ... Get the answer by id
func GetAnswerByID(c *gin.Context, client *statsd.Client) {
	log.Info("getting answer by id")
	answerID := c.Params.ByName("answer_id")
	var answer entity.Answer
	if err := model.GetAnswerByID(&answer, answerID, client); err != nil {
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
	if err := model.GetQuestionByID(&question, questionID, client); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	answer = getAllFilesByAnswer(answer, client)

	c.JSON(http.StatusOK, answer)
}

func UpdateAnswer(c *gin.Context, userID string, client *statsd.Client) {
	log.Info("answer updating")
	var newAnswer entity.Answer
	c.BindJSON(&newAnswer)

	answerID := c.Params.ByName("answer_id")
	var currAnswer entity.Answer
	if err := model.GetAnswerByID(&currAnswer, answerID, client); err != nil {
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
	if err := model.GetQuestionByID(&question, questionID, client); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	currAnswer.UserID = userID
	if err := model.GetUserByID(&currAnswer.User, userID, client); err != nil {
		log.Error("can't get the question id")
		c.JSON(http.StatusNotFound, gin.H{
			"err": "can't get the question id",
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
	if err := model.UpdateAnswer(&currAnswer, answerID, client); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		currAnswer = getAllFilesByAnswer(currAnswer, client)
		c.JSON(http.StatusOK, currAnswer)
	}
	log.Info("answer updated")

	/*post a message on SNS topic, including:
	1.Question details such as ID, user's email address
	2.Answer details such as ID, answer text, etc
	3.HTTP link to the question & answer (created or updated):
	http://prod.bh7cw.me:80/v1/question/{questionId}/answer/{answerId}
	message content:
	QuestionID: %v, QuestionText: %v, UserName: %v %v, UserEmail: %v, AnswerID: %v, AnswerText: %v, link: http://prod.bh7cw.me:80/v1/question/%v/answer/%v
	*/
	message01 := fmt.Sprintf("QuestionID: %v, QuestionText: %v, UserName: %v %v, UserEmail: %v",
		questionID, currAnswer.Question.QuestionText, currAnswer.User.FirstName, currAnswer.User.LastName, currAnswer.User.Username)

	message02 := fmt.Sprintf("AnswerID: %v, AnswerText: %v", answerID, currAnswer.AnswerText)

	message03 := fmt.Sprintf("Link: http://prod.bh7cw.me:80/v1/question/%v/answer/%v", questionID, answerID)

	message := fmt.Sprintf("%v, %v, %v", message01, message02, message03)
	result, err := tool.PublishMessageOnSNS(message)
	if err != nil {
		log.Errorf("Publish error: %v", err)

	}

	log.Info("Publish message result: %v", result)
}

//DeleteAnswer ... Delete the Answer
func DeleteAnswer(c *gin.Context, userID string, client *statsd.Client) {
	log.Info("answer deleting")
	questionID := c.Params.ByName("question_id")
	var question entity.Question
	if err := model.GetQuestionByID(&question, questionID, client); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	answerID := c.Params.ByName("answer_id")
	var answer entity.Answer
	if err := model.GetAnswerByID(&answer, answerID, client); err != nil {
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

	//store the user/questions/answer information before delete the answer, so we can use them in the sns publish message
	answerText := answer.AnswerText
	questionText := question.QuestionText
	user := answer.User

	answer = getAllFilesByAnswer(answer, client)
	if len(answer.Attachments) != 0 {
		for _, a := range answer.Attachments {
			var answerFile entity.AnswerFile
			if err := model.DeleteAnswerFileByID(&answerFile, a.ID, answerID, client); err != nil {
				log.Error(err)
				c.JSON(http.StatusNotFound, gin.H{
					"info": "can't delete the answer file",
					"err":  err.Error(),
				})
				return
			}
			if err := tool.DeleteFile(tool.GetBucketName(), a.S3ObjectName, client); err != nil {
				log.Error(err)
				c.JSON(http.StatusNotFound, gin.H{
					"message": "can't delete the file from s3",
					"err":     err.Error(),
				})
				return
			}
			if err := model.DeleteFile(&a, a.ID, client); err != nil {
				log.Error(err)
				c.JSON(http.StatusNotFound, gin.H{
					"info": "can't delete the file",
					"err":  err.Error(),
				})
				return
			}
		}

		answer.Attachments = nil
	}

	err := model.DeleteAnswer(&answer, answerID, client)

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

	/*post a message on SNS topic, including:
	1.Question details such as ID, user's email address
	2.Answer details such as ID, answer text, etc*/
	message01 := fmt.Sprintf("QuestionID: %v, QuestionText: %v, UserName: %v %v, UserEmail: %v",
		questionID, questionText, user.FirstName, user.LastName, user.Username)
	message02 := fmt.Sprintf("AnswerID: %v, AnswerText: %v", answerID, answerText)

	message := fmt.Sprintf("%v, %v", message01, message02)
	result, err := tool.PublishMessageOnSNS(message)
	if err != nil {
		log.Errorf("Publish error: %v", err)

	}

	log.Info("Publish message result: %v", result)
}
