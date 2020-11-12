package model

import (
	"cloudcomputing/webapp/config"
	"cloudcomputing/webapp/entity"
	"errors"
	"gopkg.in/alexcesaro/statsd.v2"
)

//GetAllQuestionCategories Fetch all QuestionCategories data
func GetAllQuestionCategories(questionCategories *[]entity.QuestionCategory, client *statsd.Client) (err error) {
	t := client.NewTiming()
	if err = config.DB.Find(&questionCategories).Error; err != nil {
		return err
	}
	t.Send("get_all_question_categories.query_time")
	return nil
}

//CreateQuestionCategory ... Insert New QuestionCategory
func CreateQuestionCategory(questionCategory *entity.QuestionCategory, client *statsd.Client) (err error) {
	t := client.NewTiming()
	if err = config.DB.Create(&questionCategory).Error; err != nil {
		return err
	}
	t.Send("create_question_category.query_time")
	return nil
}

//GetAllQuestionCategoriesByQuestionID ... Fetch all the QuestionCategories by QuestionId
func GetAllQuestionCategoriesByQuestionID(questionCategories *[]entity.QuestionCategory, questionId string, client *statsd.Client) (err error) {
	t := client.NewTiming()
	if err = config.DB.Where("question_id = ?", questionId).Find(&questionCategories).Error; err != nil {
		return err
	}
	t.Send("get_all_question_categories_by_question_id.query_time")
	return nil
}

//GetQuestionCategoryByQuestionID ... Fetch only one QuestionCategory by QuestionId
func GetQuestionCategoryByQuestionID(questionCategory *entity.QuestionCategory, questionId string, client *statsd.Client) (err error) {
	t := client.NewTiming()
	if err = config.DB.Where("question_id = ?", questionId).First(&questionCategory).Error; err != nil {
		return err
	}
	t.Send("get_question_category_by_question_id.query_time")
	return nil
}

//GetQuestionCategoryByCategoryID ... Fetch only one QuestionCategory by CategoryId
func GetQuestionCategoryByCategoryID(questionCategory *entity.QuestionCategory, categoryId string, client *statsd.Client) (err error) {
	t := client.NewTiming()
	if err = config.DB.Where("category_id = ?", categoryId).First(&questionCategory).Error; err != nil {
		return err
	}
	t.Send("get_question_category_by_category_id.query_time")
	return nil
}

//GetQuestionCategoryByIDs ... Fetch one QuestionCategory by CategoryId and QuestionId
func GetQuestionCategoryByIDs(questionCategory *entity.QuestionCategory, questionId string, categoryId string, client *statsd.Client) (err error) {
	t := client.NewTiming()
	if err = config.DB.Where("question_id = ? AND category_id = ?", questionId, categoryId).First(&questionCategory).Error; err != nil {
		return err
	}
	t.Send("get_question_category_by_ids.query_time")
	return nil
}

//UpdateQuestionCategory ... Update QuestionCategory
func UpdateQuestionCategory(questionCategory *entity.QuestionCategory, client *statsd.Client) (err error) {
	t := client.NewTiming()
	config.DB.Save(&questionCategory)
	t.Send("update_question_category.query_time")
	return nil
}

//DeleteQuestionCategory ... Delete QuestionCategory
func DeleteQuestionCategory(questionCategory *entity.QuestionCategory, questionId string, categoryId string, client *statsd.Client) (err error) {
	t := client.NewTiming()
	config.DB.Where("question_id = ? AND category_id = ?", questionId, categoryId).First(&questionCategory)
	if questionCategory.CategoryID == "" || questionCategory.QuestionID == "" {
		return errors.New("the QuestionCategory doesn't exist!!!")
	}
	config.DB.Where("question_id = ? AND category_id = ?", questionId, categoryId).Delete(&questionCategory)
	t.Send("delete_question_category.query_time")
	return nil
}

//DeleteQuestionCategory ... Delete QuestionCategory
func DeleteQuestionCategoryByQuestionId(questionCategory *entity.QuestionCategory, questionId string, client *statsd.Client) (err error) {
	t := client.NewTiming()
	config.DB.Where("question_id = ?", questionId).First(&questionCategory)
	if questionCategory.CategoryID == "" || questionCategory.QuestionID == "" {
		return errors.New("the QuestionCategory doesn't exist!!!")
	}
	config.DB.Where("question_id = ?", questionId).Delete(entity.QuestionCategory{})
	t.Send("delete_question_category_by_question_id.query_time")
	return nil
}
