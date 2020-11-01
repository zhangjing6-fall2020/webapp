package controller

import (
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/model"
	"cloudcomputing/webapp/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type void struct{}

var member void

func getAllCategoriesByQuestion(question entity.Question) entity.Question {
	//get all questionCategories by the questionId
	var questionCategories []entity.QuestionCategory
	if err := model.GetAllQuestionCategoriesByQuestionID(&questionCategories, question.ID); err != nil {
		fmt.Println(err)
	}
	//get all the categories by questionCategories, and add the categories to the question
	for _, qc := range questionCategories {
		var category entity.Category
		if err := model.GetCategoryByID(&category, qc.CategoryID); err != nil {
			fmt.Println(err)
		}
		question.Categories = append(question.Categories, category)
	}
	return question
}

func getAllAnswersByQuestion(question entity.Question) entity.Question {
	var answers []entity.Answer
	err := model.GetAnswersByQuestionID(&answers, question.ID)
	if err != nil {
		fmt.Println(err)
	}else{
		for _, a := range answers {
			a = getAllFilesByAnswer(a)
			question.Answers = append(question.Answers, a)
		}
	}

	return question
}

func getAllFilesByQuestion(question entity.Question) entity.Question {
	var questionFiles []entity.QuestionFile
	if err := model.GetAllQuestionFilesByQuestionID(&questionFiles, question.ID); err != nil{
		fmt.Println(err)
	}
	for _, qf := range questionFiles {
		var file entity.File
		if err := model.GetFileByID(&file, qf.ID);err != nil{
			fmt.Println(err)
		}
		question.Attachments = append(question.Attachments, file)
	}

	return question
}

//GetQuestions ... Get all Questions
func GetQuestions(c *gin.Context) {
	var questions []entity.Question
	err := model.GetAllQuestions(&questions)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	var questionsWithCategoriesAndAnswers []entity.Question
	for _, q := range questions {
		q = getAllCategoriesByQuestion(q)
		q = getAllAnswersByQuestion(q)
		q = getAllFilesByQuestion(q)
		questionsWithCategoriesAndAnswers = append(questionsWithCategoriesAndAnswers, q)
	}

	c.JSON(http.StatusOK, questionsWithCategoriesAndAnswers)
}

//GetQuestionByID ... Get the Question by id
func GetQuestionByID(c *gin.Context) {
	id := c.Params.ByName("question_id")
	var question entity.Question
	err := model.GetQuestionByID(&question, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	question = getAllCategoriesByQuestion(question)
	question = getAllAnswersByQuestion(question)
	question = getAllFilesByQuestion(question)

	c.JSON(http.StatusOK, question)
}

//CreateQuestion ... Create Question for authorized user
//get the unique category
//create the category if not exist
//set the question's categories
//create the question
//create the questioncategory with the questionID and categoryID
func CreateQuestionAuth(c *gin.Context, userID string) {
	var question entity.Question
	c.BindJSON(&question)
	if question.Categories != nil {
		for i := range question.Categories {
			question.Categories[i].Category = strings.ToLower(question.Categories[i].Category)
		}
	}

	//get the current user
	question.UserID = userID
	model.GetUserByID(&question.User, userID)

	//add unique categories in the same question input to set
	set := make(map[entity.Category]void)
	for _, category := range question.Categories {
		_, exists := set[category]
		if !exists {
			set[category] = member
		}
	}

	//in order to drop duplicate categories in the question
	//1. clear the question.Categories
	question.Categories = append([]entity.Category{})
	//for each unique category:
	//a. check exists in the category or not, if not create
	//b. add to the question.Categories
	for category := range set {
		err := model.GetCategoryByName(&category, category.Category)
		//if the category isn't found, it will return `record not found`
		if err == nil {
			fmt.Println("the category: _" + category.Category + "_ already exists")
		} else if err.Error() == "record not found" {
			if err := model.CreateCategory(&category); err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		question.Categories = append(question.Categories, category)
	}

	//create the question
	err := model.CreateQuestion(&question)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	//question.User.Questions = append(question.User.Questions, question)

	//for each category, create questionCategory with questionID and categoryID
	//Attention: can't use set here because category in set doesn't have ID
	for _, c := range question.Categories {
		var questionCategory entity.QuestionCategory
		questionCategory.QuestionID = question.ID
		questionCategory.CategoryID = c.ID
		model.CreateQuestionCategory(&questionCategory)
		//question.QuestionCategories = append(question.QuestionCategories, questionCategory)
		//c.QuestionCategories = append(c.QuestionCategories, questionCategory)
	}

	c.JSON(http.StatusCreated, gin.H{
		"question_id":       question.ID,
		"created_timestamp": question.CreatedTimestamp,
		"updated_timestamp": question.UpdatedTimestamp,
		"user_id":           question.UserID,
		"question_text":     question.QuestionText,
		"categories":        question.Categories,
		"answers":           question.Answers,
	})
}

//UpdateQuestion ... Update Question for authorized user
//1. if both questiontext and categories are null, return StatusNotFound
//2. get the question according to the questionID
//3. update question text if the questiontext input is not null
//4. update the question categories if the question categories input is not null
//4.1 if categories are the same, return: didn't implement
//4.2 otherwise, remove the question's categories and all the questioncategories
//add all the categories(if not exist), questioncategories, the question's categories
func UpdateQuestionAuth(c *gin.Context, userID string) {
	var newQuestion entity.Question
	c.BindJSON(&newQuestion)

	if newQuestion.Categories != nil {
		for i := range newQuestion.Categories {
			newQuestion.Categories[i].Category = strings.ToLower(newQuestion.Categories[i].Category)
		}
	}

	//if the content of the update question is null, return 204: No Content
	if newQuestion.QuestionText == "" && newQuestion.Categories == nil {
		c.JSON(http.StatusNoContent, gin.H{
			"question_id":       newQuestion.ID,
			"created_timestamp": newQuestion.CreatedTimestamp,
			"updated_timestamp": newQuestion.UpdatedTimestamp,
			"user_id":           newQuestion.UserID,
			"question_text":     newQuestion.QuestionText,
			"categories":        newQuestion.Categories,
			"answers":           newQuestion.Answers,
		})
		return
	}

	//get the question id from context
	questionID, match := c.Params.Get("question_id")
	if !match {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "can't get the question id",
		})
		return
	}

	var currQuestion entity.Question
	//if the update question ID can't be found, return 404: Not Found
	if err := model.GetQuestionByID(&currQuestion, questionID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"question_id":       newQuestion.ID,
			"created_timestamp": newQuestion.CreatedTimestamp,
			"updated_timestamp": newQuestion.UpdatedTimestamp,
			"user_id":           newQuestion.UserID,
			"question_text":     newQuestion.QuestionText,
			"categories":        newQuestion.Categories,
			"answers":           newQuestion.Answers,
		})
		return
	}
	if currQuestion.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "only the user who posted the question can update the question!",
		})
		return
	}

	currQuestion = getAllCategoriesByQuestion(currQuestion)
	currQuestion = getAllAnswersByQuestion(currQuestion)
	currQuestion = getAllFilesByQuestion(currQuestion)

	//3. update question text if the questiontext input is not null
	if newQuestion.QuestionText != "" {
		currQuestion.QuestionText = newQuestion.QuestionText
		if err := model.UpdateQuestion(&currQuestion, newQuestion.ID); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	//if the question categories input is null, return
	if newQuestion.Categories == nil {
		c.JSON(http.StatusOK, gin.H{
			"question_id":       currQuestion.ID,
			"created_timestamp": currQuestion.CreatedTimestamp,
			"updated_timestamp": currQuestion.UpdatedTimestamp,
			"user_id":           currQuestion.UserID,
			"question_text":     currQuestion.QuestionText,
			"categories":        currQuestion.Categories,
			"answers":           currQuestion.Answers,
		})
		return
	}

	//others, the question categories input is not null:
	//add unique categories in the newQuestion input to set
	set := make(map[entity.Category]void)
	for _, category := range newQuestion.Categories {
		_, exists := set[category]
		if !exists {
			set[category] = member
		}
	}

	//4.2 otherwise, remove the question's categories and all the questioncategories
	if currQuestion.Categories != nil {
		var questionCategories []entity.QuestionCategory
		if err := model.GetAllQuestionCategoriesByQuestionID(&questionCategories, questionID); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		//get all the categories by questionCategories, and add the categories to the question
		for _, qc := range questionCategories {
			model.DeleteQuestionCategory(&qc, qc.QuestionID, qc.CategoryID)
		}
		currQuestion.Categories = append([]entity.Category{})
	}

	//add all the categories(if not exist), questioncategories, the question's categories
	for category := range set {
		err := model.GetCategoryByName(&category, category.Category)
		//if the category isn't found, it will return `record not found`
		if err == nil {
			fmt.Println("the category: _" + category.Category + "_ already exists, didn't create duplicate category")
		} else if err.Error() == "record not found" {
			if err := model.CreateCategory(&category); err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		currQuestion.Categories = append(currQuestion.Categories, category)
		//for each category, create questionCategory with questionID and categoryID
		var questionCategory entity.QuestionCategory
		questionCategory.QuestionID = currQuestion.ID
		questionCategory.CategoryID = category.ID
		model.UpdateQuestionCategory(&questionCategory)
	}

	//update the question
	if err := model.UpdateQuestion(&currQuestion, newQuestion.ID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"question_id":       currQuestion.ID,
		"created_timestamp": currQuestion.CreatedTimestamp,
		"updated_timestamp": currQuestion.UpdatedTimestamp,
		"user_id":           currQuestion.UserID,
		"question_text":     currQuestion.QuestionText,
		"categories":        currQuestion.Categories,
		"answers":           currQuestion.Answers,
	})
}

//DeleteQuestion ... Delete the Question for authorized user
//1. get the question according to questionID
//2. if the question has any answers, can't delete the question
//3. delete the questioncategory with the questionID
//3. delete the question
func DeleteQuestionAuth(c *gin.Context, userID string) {
	//get the question id from context
	questionID, match := c.Params.Get("question_id")
	if !match {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "can't get the question id",
		})
		return
	}

	var question entity.Question
	if err := model.GetQuestionByID(&question, questionID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"question_id":       question.ID,
			"created_timestamp": question.CreatedTimestamp,
			"updated_timestamp": question.UpdatedTimestamp,
			"user_id":           question.UserID,
			"question_text":     question.QuestionText,
			"categories":        question.Categories,
			"answers":           question.Answers,
		})
		return
	}
	if question.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "only the user who posted the question can delete the question!",
		})
		return
	}

	question = getAllAnswersByQuestion(question)
	if len(question.Answers) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "the question with answers can't be deleted!",
		})
		return
	}

	var questionCategory entity.QuestionCategory
	if err := model.GetQuestionCategoryByQuestionID(&questionCategory, questionID); err == nil {
		if err := model.DeleteQuestionCategoryByQuestionId(&questionCategory, questionID); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"err": err.Error(),
			})
			return
		}
	} else {
		fmt.Println("the question doesn't have any categories!")
	}

	question = getAllFilesByQuestion(question)
	if len(question.Attachments) != 0 {
		for _,q := range question.Attachments {
			var questionFile entity.QuestionFile
			if err := model.DeleteQuestionFileByID(&questionFile,q.ID,questionID);err != nil{
				c.JSON(http.StatusNotFound, gin.H{
					"info": "can't delete the question file",
					"err": err.Error(),
				})
				return
			}
			tool.DeleteFile(tool.GetBucketName(), q.S3ObjectName)
			if err := model.DeleteFile(&q,q.ID);err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"info": "can't delete the file",
					"err": err.Error(),
				})
				return
			}
		}
		question.Attachments = nil
	}

	if err := model.DeleteQuestion(&question, questionID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + questionID: "is deleted"})
	}
}
