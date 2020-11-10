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
	t := statsDClient.NewTiming()
	if err = config.DB.Create(&questionFile).Error; err != nil {
		return err
	}
	t.Send("create_question_file.query_time")
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
	t := statsDClient.NewTiming()
	if err = config.DB.Where("question_id = ?", questionID).Find(&questionFiles).Error; err != nil {
		return err
	}
	t.Send("get_all_question_files_by_question_id.query_time")
	return nil
}

//DeleteQuestionFileByID ... Delete QuestionFile by ID
func DeleteQuestionFileByID(questionFile *entity.QuestionFile, fileID string, questionID string) (err error) {
	t := statsDClient.NewTiming()
	config.DB.Where("id = ? AND question_id = ?", fileID, questionID).First(&questionFile)
	if questionFile.ID == "" || questionFile.QuestionID == "" {
		return errors.New("the QuestionFile doesn't exist!!!")
	}
	config.DB.Where("id = ? AND question_id = ?", fileID, questionID).Delete(&questionFile)
	t.Send("delete_question_file_by_id.query_time")
	return nil
}