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

//GetAllAnswers Fetch all Answer data
func GetAllAnswers(answers *[]entity.Answer, client *statsd.Client) (err error) {
	t := client.NewTiming()
	if err = config.DB.Find(&answers).Error; err != nil {
		return err
	}
	t.Send("get_all_answers.query_time")
	return nil
}

//CreateAnswer ... Insert New Answer
func CreateAnswer(answer *entity.Answer, client *statsd.Client) (err error) {
	t := client.NewTiming()
	answer.ID = guuid.New().String()
	answer.CreatedTimestamp = time.Now()
	answer.UpdatedTimestamp = time.Now()
	if err = config.DB.Create(&answer).Error; err != nil {
		return err
	}
	t.Send("create_answer.query_time")
	return nil
}

//GetAnswerByID ... Fetch only one Answer by Id
func GetAnswerByID(answer *entity.Answer, id string, client *statsd.Client) (err error) {
	t := client.NewTiming()
	if err = config.DB.Where("id = ?", id).First(&answer).Error; err != nil {
		return err
	}
	t.Send("get_answer_by_id.query_time")
	return nil
}

func GetAnswersByQuestionID(answers *[]entity.Answer, questionID string, client *statsd.Client) (err error) {
	t := client.NewTiming()
	if err = config.DB.Where("question_id = ?", questionID).Find(&answers).Error; err != nil {
		return err
	}
	t.Send("get_answers_by_question_id.query_time")
	return nil
}

//UpdateAnswer ... Update Answer
func UpdateAnswer(answer *entity.Answer, id string, client *statsd.Client) (err error) {
	t := client.NewTiming()
	answer.UpdatedTimestamp = time.Now()
	config.DB.Save(&answer)
	t.Send("update_answer.query_time")
	return nil
}

//DeleteAnswer ... Delete Answer
func DeleteAnswer(answer *entity.Answer, id string, client *statsd.Client) (err error) {
	t := client.NewTiming()
	if config.DB.Where("id = ?", id).First(&answer); answer.ID == "" {
		return errors.New("the answer doesn't exist!!!")
	}
	config.DB.Where("id = ?", id).Delete(&answer)
	t.Send("delete_answer.query_time")
	return nil
}
