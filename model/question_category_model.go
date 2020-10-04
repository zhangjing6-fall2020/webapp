package model

import (
	"cloudcomputing/webapp/config"
	"cloudcomputing/webapp/entity"
	"errors"
)

//GetAllQuestionCategories Fetch all QuestionCategories data
func GetAllQuestionCategories(questionCategories *[]entity.QuestionCategory) (err error) {
	if err = config.DB.Find(&questionCategories).Error; err != nil {
		return err
	}
	return nil
}

//CreateQuestionCategory ... Insert New QuestionCategory
func CreateQuestionCategory(questionCategory *entity.QuestionCategory) (err error) {
	if err = config.DB.Create(&questionCategory).Error; err != nil {
		return err
	}
	return nil
}

//GetAllQuestionCategoriesByQuestionID ... Fetch all the QuestionCategories by QuestionId
func GetAllQuestionCategoriesByQuestionID(questionCategories *[]entity.QuestionCategory, questionId string) (err error) {
	if err = config.DB.Where("question_id = ?", questionId).Find(&questionCategories).Error; err != nil {
		return err
	}
	return nil
}

//GetQuestionCategoryByQuestionID ... Fetch only one QuestionCategory by QuestionId
func GetQuestionCategoryByQuestionID(questionCategory *entity.QuestionCategory, questionId string) (err error) {
	if err = config.DB.Where("question_id = ?", questionId).First(&questionCategory).Error; err != nil {
		return err
	}
	return nil
}

//GetQuestionCategoryByCategoryID ... Fetch only one QuestionCategory by CategoryId
func GetQuestionCategoryByCategoryID(questionCategory *entity.QuestionCategory, categoryId string) (err error) {
	if err = config.DB.Where("category_id = ?", categoryId).First(&questionCategory).Error; err != nil {
		return err
	}
	return nil
}

//GetQuestionCategoryByIDs ... Fetch one QuestionCategory by CategoryId and QuestionId
func GetQuestionCategoryByIDs(questionCategory *entity.QuestionCategory, questionId string, categoryId string) (err error) {
	if err = config.DB.Where("question_id = ? AND category_id = ?", questionId, categoryId).First(&questionCategory).Error; err != nil {
		return err
	}
	return nil
}

//UpdateQuestionCategory ... Update QuestionCategory
func UpdateQuestionCategory(questionCategory *entity.QuestionCategory) (err error) {
	config.DB.Save(&questionCategory)
	return nil
}

//DeleteQuestionCategory ... Delete QuestionCategory
func DeleteQuestionCategory(questionCategory *entity.QuestionCategory, questionId string, categoryId string) (err error) {
	config.DB.Where("question_id = ? AND category_id = ?", questionId, categoryId).First(&questionCategory)
	if questionCategory.CategoryID == "" || questionCategory.QuestionID == "" {
		return errors.New("the QuestionCategory doesn't exist!!!")
	}
	config.DB.Where("question_id = ? AND category_id = ?", questionId, categoryId).Delete(&questionCategory)
	return nil
}

//DeleteQuestionCategory ... Delete QuestionCategory
func DeleteQuestionCategoryByQuestionId(questionCategory *entity.QuestionCategory, questionId string) (err error) {
	config.DB.Where("question_id = ?", questionId).First(&questionCategory)
	if questionCategory.CategoryID == "" || questionCategory.QuestionID == "" {
		return errors.New("the QuestionCategory doesn't exist!!!")
	}
	config.DB.Where("question_id = ?", questionId).Delete(entity.QuestionCategory{})
	return nil
}
