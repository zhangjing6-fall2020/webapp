package controller

import (
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/model"
	"cloudcomputing/webapp/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const bucketName string = "webapp.jing.zhang"

func CreateQuestionFileAuth(c *gin.Context, userID string) {
	fileHeader, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	//get the current user
	var question entity.Question
	questionID := c.Params.ByName("question_id")
	if err := model.GetQuestionByID(&question, questionID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if question.UserID != userID {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "the question is not created by the user!!!",
		})
		return
	}

	//create the file data
	var file entity.File
	file.FileName = fileHeader.Filename
	file.IsQuestion = true
	if err := model.CreateFile(&file, questionID); err != nil {
		fmt.Println("can't create the file!")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//upload the file to s3
	tool.UploadFile(bucketName, fileHeader, file.S3ObjectName)

	//update Attachments in question
	question.Attachments = append(question.Attachments, file)

	//get S3 metadata

	//update file

	//create the question file data
	var questionFile entity.QuestionFile
	questionFile.ID = file.ID
	questionFile.File = file
	questionFile.QuestionID = questionID
	questionFile.Question = question
	if err := model.CreateQuestionFile(&questionFile);err != nil {
		fmt.Println("can't create the question file!")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, file)
}

func DeleteQuestionFileAuth(c *gin.Context, userID string) {
	//get question
	var question entity.Question
	questionID := c.Params.ByName("question_id")
	if err := model.GetQuestionByID(&question, questionID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if question.UserID != userID {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "the question is not created by the user!!!",
		})
		return
	}

	//get file
	fileID := c.Params.ByName("file_id")
	var file entity.File
	if err := model.GetFileByID(&file, fileID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	//delete from s3
	tool.DeleteFile(bucketName, file.S3ObjectName)

	//delete questionfile
	var questionFile entity.QuestionFile
	if err := model.DeleteQuestionFileByID(&questionFile, fileID, questionID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	//delete file
	if err := model.DeleteFile(&file, fileID); err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	//delete from question
	if len(question.Attachments) == 1 && question.Attachments[0].ID == fileID{
		question.Attachments = nil
	}
	for i, a := range question.Attachments {
		if a.ID == fileID {
			question.Attachments = append(question.Attachments[:i],question.Attachments[i+1:]...)
		}
	}
}

func CreateAnswerFileAuth(c *gin.Context, userID string) {
	//get the questionID, answerID, verify

	//verify userID

	//create the file

	//upload to S3

	//create answer file

	//update the attachments in answer

	//update the attachments in question
}

func DeleteAnswerFileAuth(c *gin.Context, userID string) {
	//get questionID, answerID, fileID, verify

	//verify userID

	//delete from S3

	//delete answer file

	//delete file

	//update from answer

	//update from question
}