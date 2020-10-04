package route

import (
	"cloudcomputing/webapp/auth"
	"cloudcomputing/webapp/controller"
	"github.com/gin-gonic/gin"
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
	//pub.PUT("user/:id", controller.UpdateUser)
	//delete the user given UserID
	pub.DELETE("user/:id", controller.DeleteUser)

	//answer a question
	pub.POST("question/:id")
	//get all the questions
	pub.GET("questions", controller.GetQuestions)
	//get a question according to given id
	pub.GET("question/:id", func(c *gin.Context) {
		controller.GetQuestionByID(c)
	})

	//update a question's answer
	pub.PUT("question/:id/answer/:id")
	//delete a question's answer
	pub.DELETE("question/:id/answer/:id")
	//get a question's answer
	pub.GET("question/:id/answer/:id")


	//authorized endpoints, basic auth is required with username and password
	authorized := r.Group("/v1", auth.BasicAuth())
	//get the information of the authorized user
	authorized.GET("/user/self", func(c *gin.Context) {
		controller.GetUserByUsername(c, auth.GetCurrUsername())
	})
	//edit the information of the authorized user
	authorized.PUT("/user/self", func(c *gin.Context) {
		controller.UpdateAuthorizedUser(c, auth.GetCurrUsername())
	})

	//post a new question
	authorized.POST("question", func(c *gin.Context) {
		controller.CreateQuestionAuth(c, auth.GetCurrentUserID())
	})
	//update a question
	authorized.PUT("question/:id", func(c *gin.Context) {
		controller.UpdateQuestionAuth(c, auth.GetCurrentUserID())
	})
	//delete a question
	authorized.DELETE("question/:id", func(c *gin.Context) {
		controller.DeleteQuestionAuth(c, auth.GetCurrentUserID())
	})

	return r
}
