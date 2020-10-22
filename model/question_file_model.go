package model

import (
	"cloudcomputing/webapp/config"
	"cloudcomputing/webapp/entity"
	"errors"
)

//GetAllQuestionFiles Fetch all QuestionFile data
func GetAllQuestionFiles(questionFiles *[]entity.QuestionFile) (err error) {
	if err = config.DB.Find(&questionFiles).Error; err != nil {
		return err
	}
	return nil
}

//CreateQuestionFile ... Insert New data
func CreateQuestionFile(questionFile *entity.QuestionFile) (err error) {
	if err = config.DB.Create(&questionFile).Error; err != nil {
		return err
	}
	return nil
}

//GetQuestionFileByID ... Fetch only one QuestionFile by Id
func GetQuestionFileByFileID(questionFile *entity.QuestionFile, fileID string) (err error) {
	if err = config.DB.Where("id = ?", fileID).First(&questionFile).Error; err != nil {
		return err
	}
	return nil
}

//GetQuestionFileByQuestionID ... Fetch only one QuestionFile by QuestionId
func GetQuestionFileByQuestionID(questionFile *entity.QuestionFile, questionID string) (err error) {
	if err = config.DB.Where("question_id = ?", questionID).First(&questionFile).Error; err != nil {
		return err
	}
	return nil
}

func GetAllQuestionFilesByQuestionID(questionFiles *[]entity.QuestionFile, questionID string) (err error) {
	if err = config.DB.Where("question_id = ?", questionID).Find(&questionFiles).Error; err != nil {
		return err
	}
	return nil
}

//DeleteQuestionFileByID ... Delete QuestionFile by ID
func DeleteQuestionFileByID(questionFile *entity.QuestionFile, fileID string, questionID string) (err error) {
	config.DB.Where("id = ? AND question_id = ?", fileID, questionID).First(&questionFile)
	if questionFile.ID == "" || questionFile.QuestionID == "" {
		return errors.New("the QuestionFile doesn't exist!!!")
	}
	config.DB.Where("id = ? AND question_id = ?", fileID, questionID).Delete(&questionFile)
	return nil
}