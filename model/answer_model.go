package model

import (
	"cloudcomputing/webapp/config"
	"cloudcomputing/webapp/entity"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	guuid "github.com/google/uuid"
	"time"
)

//GetAllAnswers Fetch all Answer data
func GetAllAnswers(answers *[]entity.Answer) (err error) {
	if err = config.DB.Find(&answers).Error; err != nil {
		return err
	}
	return nil
}

//CreateAnswer ... Insert New Answer
func CreateAnswer(answer *entity.Answer) (err error) {
	answer.ID = guuid.New().String()
	answer.CreatedTimestamp = time.Now()
	answer.UpdatedTimestamp = time.Now()
	if err = config.DB.Create(&answer).Error; err != nil {
		return err
	}
	return nil
}

//GetAnswerByID ... Fetch only one Answer by Id
func GetAnswerByID(answer *entity.Answer, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(&answer).Error; err != nil {
		return err
	}
	return nil
}

//UpdateAnswer ... Update Answer
func UpdateAnswer(answer *entity.Answer, id string) (err error) {
	answer.UpdatedTimestamp = time.Now()
	config.DB.Save(&answer)
	return nil
}

//DeleteAnswer ... Delete Answer
func DeleteAnswer(answer *entity.Answer, id string) (err error) {
	if config.DB.Where("id = ?", id).First(&answer); answer.ID == "" {
		return errors.New("the answer doesn't exist!!!")
	}
	config.DB.Where("id = ?", id).Delete(&answer)
	return nil
}
