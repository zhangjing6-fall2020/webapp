package model

import (
	"cloudcomputing/webapp/config"
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/monitor"
	"errors"
	"gopkg.in/alexcesaro/statsd.v2"
)

var statsDClient *statsd.Client = monitor.SetUpStatsD()

//GetAllAnswerFiles Fetch all AnswerFile data
func GetAllAnswerFiles(answerFile *[]entity.AnswerFile) (err error) {
	if err = config.DB.Find(&answerFile).Error; err != nil {
		return err
	}
	return nil
}

//CreateAnswerFile ... Insert New data
func CreateAnswerFile(answerFile *entity.AnswerFile) (err error) {
	t := statsDClient.NewTiming()
	if err = config.DB.Create(&answerFile).Error; err != nil {
		return err
	}
	t.Send("create_answer_file.query_time")
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

func GetAllAnswerFilesByAnswerID(answerFiles *[]entity.AnswerFile, answerID string) (err error) {
	t := statsDClient.NewTiming()
	if err = config.DB.Where("answer_id = ?", answerID).Find(&answerFiles).Error; err != nil {
		return err
	}
	t.Send("get_all_answer_files_by_answer_id.query_time")
	return nil
}

//DeleteAnswerFileByID ... Delete AnswerFile by ID
func DeleteAnswerFileByID(answerFile *entity.AnswerFile, fileID string, answerID string) (err error) {
	t := statsDClient.NewTiming()
	config.DB.Where("id = ? AND answer_id = ?", fileID, answerID).First(&answerFile)
	if answerFile.ID == "" || answerFile.AnswerID == "" {
		return errors.New("the AnswerFile doesn't exist!!!")
	}
	config.DB.Where("id = ? AND answer_id = ?", fileID, answerID).Delete(&answerFile)
	t.Send("delete_answer_file_by_id.query_time")
	return nil
}

