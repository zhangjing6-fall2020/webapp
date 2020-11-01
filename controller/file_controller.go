package controller

import (
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/model"
	"cloudcomputing/webapp/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	"net/http"
)

var bucketName string = tool.GetBucketName()//"webapp.jing.zhang"

func CreateQuestionFileAuth(c *gin.Context, userID string) {
	//get the current user
	var question entity.Question
	questionID := c.Params.ByName("question_id")
	if err := model.GetQuestionByID(&question, questionID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	//verify the userID
	if question.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "the question is not created by the user!!!",
		})
		return
	}

	//get the image
	fileHeader, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	//upload the file to s3
	var file entity.File
	file.ID = guuid.New().String()
	file.FileName = fileHeader.Filename
	file.S3ObjectName = fmt.Sprintf("%s/%s/%s", questionID, file.ID, file.FileName)
	tool.UploadFile(bucketName, fileHeader, file.S3ObjectName)

	//get S3 metadata
	metadata := tool.GetObjectMetaData(bucketName, file.S3ObjectName)
	file.AcceptRanges = *metadata.AcceptRanges
	file.ContentLength = *metadata.ContentLength
	file.ContentType = *metadata.ContentType
	file.ETag = *metadata.ETag
	file.LastModified = *metadata.LastModified

	//create the file data
	file.IsQuestion = true
	if err := model.CreateFileAuth(&file); err != nil {
		fmt.Println("can't create the file!")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

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

	//verify the user
	if question.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{
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

	//delete question file
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
}

func CreateAnswerFileAuth(c *gin.Context, userID string) {
	var question entity.Question
	questionID := c.Params.ByName("question_id")
	if err := model.GetQuestionByID(&question, questionID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	var answer entity.Answer
	answerID := c.Params.ByName("answer_id")
	if err := model.GetAnswerByID(&answer, answerID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	//verify the userID
	if question.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "the question is not created by the user!!!",
		})
		return
	}

	if question.UserID != answer.UserID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "the question and the answer are not created by the same user!!!",
		})
		return
	}

	if question.ID != answer.QuestionID {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "the answer is not for the question!!!",
		})
		return
	}

	//get the image
	fileHeader, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	//upload the file to s3
	var file entity.File
	file.ID = guuid.New().String()
	file.FileName = fileHeader.Filename
	file.S3ObjectName = fmt.Sprintf("%s/%s/%s", answerID, file.ID, file.FileName)
	tool.UploadFile(bucketName, fileHeader, file.S3ObjectName)

	//get S3 metadata
	metadata := tool.GetObjectMetaData(bucketName, file.S3ObjectName)
	file.AcceptRanges = *metadata.AcceptRanges
	file.ContentLength = *metadata.ContentLength
	file.ContentType = *metadata.ContentType
	file.ETag = *metadata.ETag
	file.LastModified = *metadata.LastModified

	//create the file data
	file.IsQuestion = false
	if err := model.CreateFileAuth(&file); err != nil {
		fmt.Println("can't create the file!")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//create the answer file data
	var answerFile entity.AnswerFile
	answerFile.ID = file.ID
	answerFile.File = file
	answerFile.AnswerID = answerID
	answerFile.Answer = answer
	if err := model.CreateAnswerFile(&answerFile);err != nil {
		fmt.Println("can't create the answer file!")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, file)
}

func DeleteAnswerFileAuth(c *gin.Context, userID string) {
	//get question
	var question entity.Question
	questionID := c.Params.ByName("question_id")
	if err := model.GetQuestionByID(&question, questionID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	var answer entity.Answer
	answerID := c.Params.ByName("answer_id")
	if err := model.GetAnswerByID(&answer, answerID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	//verify the userID
	if question.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "the question is not created by the user!!!",
		})
		return
	}

	if question.UserID != answer.UserID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "the question and the answer are not created by the same user!!!",
		})
		return
	}

	if question.ID != answer.QuestionID {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "the answer is not for the question!!!",
		})
		return
	}

	//verify the user
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

	//delete answer file
	var answerFile entity.AnswerFile
	if err := model.DeleteAnswerFileByID(&answerFile, fileID, answerID); err != nil {
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
}