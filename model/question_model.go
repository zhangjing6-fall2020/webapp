package model

import (
"cloudcomputing/webapp/config"
"cloudcomputing/webapp/entity"
"errors"
_ "github.com/go-sql-driver/mysql"
guuid "github.com/google/uuid"
"time"
)

//GetAllQuestions Fetch all question data
func GetAllQuestions(questions *[]entity.Question) (err error) {
	if err = config.DB.Find(&questions).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New question
func CreateQuestion(question *entity.Question) (err error) {
	question.ID = guuid.New().String()
	question.CreatedTimestamp = time.Now()
	question.UpdatedTimestamp = time.Now()
	if err = config.DB.Create(&question).Error; err != nil {
		return err
	}
	return nil
}

//GetQuestionByID ... Fetch only one Question by Id
func GetQuestionByID(question *entity.Question, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(&question).Error; err != nil {
		return err
	}
	return nil
}

//UpdateQuestion ... Update question
func UpdateQuestion(question *entity.Question, id string) (err error) {
	question.UpdatedTimestamp = time.Now()
	config.DB.Save(&question)
	return nil
}

//DeleteUser ... Delete user
func DeleteQuestion(question *entity.Question, id string) (err error) {
	if config.DB.Where("id = ?", id).First(&question); question.ID == "" {
		return errors.New("the question doesn't exist!!!")
	}
	config.DB.Where("id = ?", id).Delete(&question)
	return nil
}

