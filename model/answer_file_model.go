package model

import (
	"cloudcomputing/webapp/config"
	"cloudcomputing/webapp/entity"
	"errors"
)

//GetAllAnswerFiles Fetch all AnswerFile data
func GetAllAnswerFiles(answerFile *[]entity.AnswerFile) (err error) {
	if err = config.DB.Find(&answerFile).Error; err != nil {
		return err
	}
	return nil
}

//CreateAnswerFile ... Insert New data
func CreateAnswerFile(answerFile *entity.AnswerFile) (err error) {
	if err = config.DB.Create(&answerFile).Error; err != nil {
		return err
	}
	return nil
}

//GetAnswerFileByID ... Fetch only one AnswerFile by Id
func GetAnswerFileByFileID(answerFile *entity.AnswerFile, fileID string) (err error) {
	if err = config.DB.Where("id = ?", fileID).First(&answerFile).Error; err != nil {
		return err
	}
	return nil
}

//GetAnswerFileByAnswerID ... Fetch only one AnswerFile by AnswerId
func GetAnswerFileByAnswerID(answerFile *entity.AnswerFile, answerID string) (err error) {
	if err = config.DB.Where("answer_id = ?", answerID).First(&answerFile).Error; err != nil {
		return err
	}
	return nil
}

//DeleteAnswerFileByID ... Delete AnswerFile by ID
func DeleteAnswerFileByID(answerFile *entity.AnswerFile, fileID string, answerID string) (err error) {
	if config.DB.Where("id = ?", fileID).First(&answerFile); answerFile.ID == "" {
		return errors.New("the file doesn't exist!!!")
	}
	config.DB.Where("id = ?", fileID).Delete(&answerFile)
	return nil

	config.DB.Where("id = ? AND answer_id = ?", fileID, answerID).First(&answerFile)
	if answerFile.ID == "" || answerFile.AnswerID == "" {
		return errors.New("the AnswerFile doesn't exist!!!")
	}
	config.DB.Where("id = ? AND answer_id = ?", fileID, answerID).Delete(&answerFile)
	return nil
}

