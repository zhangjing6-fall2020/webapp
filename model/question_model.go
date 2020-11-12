package model

import (
	"cloudcomputing/webapp/config"
	"cloudcomputing/webapp/entity"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	guuid "github.com/google/uuid"
	"gopkg.in/alexcesaro/statsd.v2"
	"time"
)

//GetAllQuestions Fetch all question data
func GetAllQuestions(questions *[]entity.Question, client *statsd.Client) (err error) {
	t := client.NewTiming()
	if err = config.DB.Find(&questions).Error; err != nil {
		return err
	}
	t.Send("get_all_questions.query_time")
	return nil
}

//CreateUser ... Insert New question
func CreateQuestion(question *entity.Question, client *statsd.Client) (err error) {
	t := client.NewTiming()
	question.ID = guuid.New().String()
	question.CreatedTimestamp = time.Now()
	question.UpdatedTimestamp = time.Now()
	if err = config.DB.Create(&question).Error; err != nil {
		return err
	}
	t.Send("create_question.query_time")
	return nil
}

//GetQuestionByID ... Fetch only one Question by Id
func GetQuestionByID(question *entity.Question, id string, client *statsd.Client) (err error) {
	t := client.NewTiming()
	if err = config.DB.Where("id = ?", id).First(&question).Error; err != nil {
		return err
	}
	t.Send("get_question_by_id.query_time")
	return nil
}

//UpdateQuestion ... Update question
func UpdateQuestion(question *entity.Question, id string, client *statsd.Client) (err error) {
	t := client.NewTiming()
	question.UpdatedTimestamp = time.Now()
	config.DB.Save(&question)
	t.Send("update_question.query_time")
	return nil
}

//DeleteUser ... Delete user
func DeleteQuestion(question *entity.Question, id string, client *statsd.Client) (err error) {
	t := client.NewTiming()
	if config.DB.Where("id = ?", id).First(&question); question.ID == "" {
		return errors.New("the question doesn't exist!!!")
	}
	config.DB.Where("id = ?", id).Delete(&question)
	t.Send("delete_question.query_time")
	return nil
}
