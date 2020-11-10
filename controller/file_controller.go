package controller

import (
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/model"
	"cloudcomputing/webapp/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var bucketName string = tool.GetBucketName() //"webapp.jing.zhang"

func CreateQuestionFileAuth(c *gin.Context, userID string) {
	log.Info("creating question file")
	//get the current user
	var question entity.Question
	questionID := c.Params.ByName("question_id")
	if err := model.GetQuestionByID(&question, questionID); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	//verify the userID
	if question.UserID != userID {
		log.Error("the question is not created by the user")
		c.JSON(http.StatusForbidden, gin.H{
			"error": "the question is not created by the user!!!",
		})
		return
	}

	//get the image
	fileHeader, err := c.FormFile("image")
	if err != nil {
		log.Error(err)
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
	if err := tool.UploadFile(bucketName, fileHeader, file.S3ObjectName); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"info": "can't upload the file to s3",
			"err":  err.Error(),
		})
		return
	}

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
		log.Error(err)
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
	if err := model.CreateQuestionFile(&questionFile); err != nil {
		log.Error(err)
		fmt.Println("can't create the question file!")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, file)
	log.Info("question file created")
}

func DeleteQuestionFileAuth(c *gin.Context, userID string) {
	log.Info("question file deleting")
	//get question
	var question entity.Question
	questionID := c.Params.ByName("question_id")
	if err := model.GetQuestionByID(&question, questionID); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	//verify the user
	if question.UserID != userID {
		log.Error("the question is not created by the user")
		c.JSON(http.StatusForbidden, gin.H{
			"error": "the question is not created by the user!!!",
		})
		return
	}

	//get file
	fileID := c.Params.ByName("file_id")
	var file entity.File
	if err := model.GetFileByID(&file, fileID); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	//delete from s3
	if err := tool.DeleteFile(bucketName, file.S3ObjectName); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"info": "can't delete the file from s3",
			"err":  err.Error(),
		})
		return
	}

	//delete question file
	var questionFile entity.QuestionFile
	if err := model.DeleteQuestionFileByID(&questionFile, fileID, questionID); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	//delete file
	if err := model.DeleteFile(&file, fileID); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Info("question file deleted")
}

func CreateAnswerFileAuth(c *gin.Context, userID string) {
	log.Info("answer file creating")
	var question entity.Question
	questionID := c.Params.ByName("question_id")
	if err := model.GetQuestionByID(&question, questionID); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	var answer entity.Answer
	answerID := c.Params.ByName("answer_id")
	if err := model.GetAnswerByID(&answer, answerID); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	//verify the userID
	if question.UserID != userID {
		log.Error("the question is not created by the user")
		c.JSON(http.StatusForbidden, gin.H{
			"error": "the question is not created by the user!!!",
		})
		return
	}

	if question.UserID != answer.UserID {
		log.Error("the question and the answer are not created by the same user")
		c.JSON(http.StatusForbidden, gin.H{
			"error": "the question and the answer are not created by the same user!!!",
		})
		return
	}

	if question.ID != answer.QuestionID {
		log.Error("the answer is not for the question")
		c.JSON(http.StatusNotFound, gin.H{
			"error": "the answer is not for the question!!!",
		})
		return
	}

	//get the image
	fileHeader, err := c.FormFile("image")
	if err != nil {
		log.Error(err)
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
	if err := tool.UploadFile(bucketName, fileHeader, file.S3ObjectName); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"info": "can't upload the file to s3",
			"err":  err.Error(),
		})
		return
	}

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
		log.Error(err)
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
	if err := model.CreateAnswerFile(&answerFile); err != nil {
		log.Error(err)
		fmt.Println("can't create the answer file!")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, file)
	log.Info("answer file deleted")
}

func DeleteAnswerFileAuth(c *gin.Context, userID string) {
	log.Info("answer file deleting")
	//get question
	var question entity.Question
	questionID := c.Params.ByName("question_id")
	if err := model.GetQuestionByID(&question, questionID); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	var answer entity.Answer
	answerID := c.Params.ByName("answer_id")
	if err := model.GetAnswerByID(&answer, answerID); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	//verify the userID
	if question.UserID != userID {
		log.Error("the question is not created by the user")
		c.JSON(http.StatusForbidden, gin.H{
			"error": "the question is not created by the user!!!",
		})
		return
	}

	if question.UserID != answer.UserID {
		log.Error("the question and the answer are not created by the same user")
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
		log.Error("the question is not created by the user")
		c.JSON(http.StatusNotFound, gin.H{
			"error": "the question is not created by the user!!!",
		})
		return
	}

	//get file
	fileID := c.Params.ByName("file_id")
	var file entity.File
	if err := model.GetFileByID(&file, fileID); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	//delete from s3
	if err := tool.DeleteFile(bucketName, file.S3ObjectName); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"info": "can't delete the file from s3",
			"err":  err.Error(),
		})
		return
	}

	//delete answer file
	var answerFile entity.AnswerFile
	if err := model.DeleteAnswerFileByID(&answerFile, fileID, answerID); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	//delete file
	if err := model.DeleteFile(&file, fileID); err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Info("question file deleted")
}
