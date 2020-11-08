package route

import (
	"cloudcomputing/webapp/auth"
	"cloudcomputing/webapp/controller"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	//public endpoint, auth is not required
	pub := r.Group("/v1")
	//create a new user
	pub.POST("user", controller.CreateUser)
	//get all the users
	pub.GET("users", controller.GetUsers)
	//pub.GET("/user/:email_address", controller.GetUserByEmail)
	//pub.GET("user/:id", controller.GetUserByID)
	pub.GET("user/:id", controller.GetUserByID)
	//delete the user given UserID
	pub.DELETE("user/:id", controller.DeleteUser)

	//get all the questions
	pub.GET("questions", controller.GetQuestions)
	//get a question according to given id
	pub.GET("question/:question_id", controller.GetQuestionByID)

	//get a question's answer
	pub.GET("question/:question_id/answer/:answer_id", controller.GetAnswerByID)

	//authorized endpoints, basic auth is required with username and password
	authorized := r.Group("/v1", auth.BasicAuth())
	//get the information of the authorized user
	authorized.GET("userself", func(c *gin.Context) {
		log.Trace("getting self user info")
		controller.GetUserByUsername(c, auth.GetCurrUsername())
	})

	//edit the information of the authorized user
	authorized.PUT("user/self", func(c *gin.Context) {
		log.Trace("updating user info with auth")
		controller.UpdateAuthorizedUser(c, auth.GetCurrUsername())
	})

	//post a new question
	authorized.POST("question", func(c *gin.Context) {
		log.Trace("creating a new question with auth")
		controller.CreateQuestionAuth(c, auth.GetCurrentUserID())
	})
	//update a question
	authorized.PUT("question/:question_id", func(c *gin.Context) {
		log.Trace("updating a question with auth")
		controller.UpdateQuestionAuth(c, auth.GetCurrentUserID())
	})
	//delete a question
	authorized.DELETE("question/:question_id", func(c *gin.Context) {
		log.Trace("deleting a question with auth")
		controller.DeleteQuestionAuth(c, auth.GetCurrentUserID())
	})

	//answer a question
	authorized.POST("question/:question_id/answer", func(c *gin.Context) {
		log.Trace("creating an answer for a question with auth")
		controller.CreateAnswer(c, auth.GetCurrentUserID())
	})
	//update a question's answer
	authorized.PUT("question/:question_id/answer/:answer_id", func(c *gin.Context) {
		log.Trace("updating an answer of a question with auth")
		controller.UpdateAnswer(c, auth.GetCurrentUserID())
	})
	//delete a question's answer
	authorized.DELETE("question/:question_id/answer/:answer_id", func(c *gin.Context) {
		log.Trace("deleting an answer of a question with auth")
		controller.DeleteAnswer(c, auth.GetCurrentUserID())
	})

	//upload a file to a question
	authorized.POST("question/:question_id/file", func(c *gin.Context) {
		log.Trace("creating an file for a question with auth")
		controller.CreateQuestionFileAuth(c, auth.GetCurrentUserID())
	})

	//delete a file to a question
	authorized.DELETE("question/:question_id/file/:file_id", func(c *gin.Context) {
		log.Trace("deleting an file of a question with auth")
		controller.DeleteQuestionFileAuth(c, auth.GetCurrentUserID())
	})

	//upload a file to an answer
	authorized.POST("question/:question_id/answer/:answer_id/file", func(c *gin.Context) {
		log.Trace("creating an file for an answer with auth")
		controller.CreateAnswerFileAuth(c, auth.GetCurrentUserID())
	})

	//delete a file to an answer
	authorized.DELETE("question/:question_id/answer/:answer_id/file/:file_id", func(c *gin.Context) {
		log.Trace("deleting an file of an answer with auth")
		controller.DeleteAnswerFileAuth(c, auth.GetCurrentUserID())
	})

	return r
}
