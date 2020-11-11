package route

import (
	"cloudcomputing/webapp/auth"
	"cloudcomputing/webapp/controller"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alexcesaro/statsd.v2"
)

//SetupRouter ... Configure routes
func SetupRouter(client *statsd.Client) *gin.Engine {
	r := gin.Default()
	//public endpoint, auth is not required
	pub := r.Group("/v1")
	i := 0

	pub.POST("user", func(c *gin.Context) {
		i++
		client.Count("webapp.post_user", i)
		defer client.NewTiming().Send("post_user.response_time")
		controller.CreateUser(c)
	})

	client.Count("webapp.get_users", 10)
	//get all the users
	pub.GET("users", func(c *gin.Context) {
		//client.Count("webapp.get_users", 0)
		client.Count("webapp.get_users", 0)
		t := client.NewTiming()
		controller.GetUsers(c)
		t.Send("get_users.response_time")
	})
	//pub.GET("/user/:email_address", controller.GetUserByEmail)
	//pub.GET("user/:id", controller.GetUserByID)
	pub.GET("user/:id", func(c *gin.Context) {
		client.Increment("webapp.get_user")
		t := client.NewTiming()
		controller.GetUserByID(c)
		t.Send("get_user.response_time")
	})
	//delete the user given UserID
	pub.DELETE("user/:id", func(c *gin.Context) {
		client.Increment("webapp.delete_user")
		t := client.NewTiming()
		controller.DeleteUser(c)
		t.Send("delete_user.response_time")
	})

	//get all the questions
	pub.GET("questions", func(c *gin.Context) {
		client.Increment("webapp.get_questions")
		t := client.NewTiming()
		controller.GetQuestions(c)
		t.Send("get_questions.response_time")
	})
	//get a question according to given id
	pub.GET("question/:question_id", func(c *gin.Context) {
		client.Increment("webapp.get_question")
		t := client.NewTiming()
		controller.GetQuestionByID(c)
		t.Send("get_question.response_time")
	})

	//get a question's answer
	pub.GET("question/:question_id/answer/:answer_id", func(c *gin.Context) {
		client.Increment("webapp.get_answer")
		t := client.NewTiming()
		controller.GetAnswerByID(c)
		t.Send("get_answer.response_time")
	})

	//authorized endpoints, basic auth is required with username and password
	authorized := r.Group("/v1", auth.BasicAuth())
	//get the information of the authorized user
	authorized.GET("userself", func(c *gin.Context) {
		// Increment a counter.
		client.Increment("webapp.get_user_self")
		log.Info("getting self user info")
		t := client.NewTiming()
		controller.GetUserByUsername(c, auth.GetCurrUsername())
		t.Send("get_user_self.response_time")
	})

	//edit the information of the authorized user
	authorized.PUT("user/self", func(c *gin.Context) {
		client.Increment("webapp.put_user_self")
		log.Info("updating user info with auth")
		t := client.NewTiming()
		controller.UpdateAuthorizedUser(c, auth.GetCurrUsername())
		t.Send("put_user_self.response_time")
	})

	//post a new question
	authorized.POST("question", func(c *gin.Context) {
		client.Increment("webapp.post_question")
		log.Info("creating a new question with auth")
		t := client.NewTiming()
		controller.CreateQuestionAuth(c, auth.GetCurrentUserID())
		t.Send("post_question.response_time")
	})
	//update a question
	authorized.PUT("question/:question_id", func(c *gin.Context) {
		client.Increment("webapp.put_question")
		log.Info("updating a question with auth")
		t := client.NewTiming()
		controller.UpdateQuestionAuth(c, auth.GetCurrentUserID())
		t.Send("put_question.response_time")
	})
	//delete a question
	authorized.DELETE("question/:question_id", func(c *gin.Context) {
		client.Increment("webapp.delete_question")
		log.Info("deleting a question with auth")
		t := client.NewTiming()
		controller.DeleteQuestionAuth(c, auth.GetCurrentUserID())
		t.Send("delete_question.response_time")
	})

	//answer a question
	authorized.POST("question/:question_id/answer", func(c *gin.Context) {
		client.Increment("webapp.post_answer")
		log.Info("creating an answer for a question with auth")
		t := client.NewTiming()
		controller.CreateAnswer(c, auth.GetCurrentUserID())
		t.Send("post_answer.response_time")
	})
	//update a question's answer
	authorized.PUT("question/:question_id/answer/:answer_id", func(c *gin.Context) {
		client.Increment("webapp.put_answer")
		log.Info("updating an answer of a question with auth")
		t := client.NewTiming()
		controller.UpdateAnswer(c, auth.GetCurrentUserID())
		t.Send("put_answer.response_time")
	})
	//delete a question's answer
	authorized.DELETE("question/:question_id/answer/:answer_id", func(c *gin.Context) {
		client.Increment("webapp.delete_answer")
		log.Info("deleting an answer of a question with auth")
		t := client.NewTiming()
		controller.DeleteAnswer(c, auth.GetCurrentUserID())
		t.Send("delete_answer.response_time")
	})

	//upload a file to a question
	authorized.POST("question/:question_id/file", func(c *gin.Context) {
		client.Increment("webapp.post_question_file")
		log.Info("creating an file for a question with auth")
		t := client.NewTiming()
		controller.CreateQuestionFileAuth(c, auth.GetCurrentUserID())
		t.Send("post_question_file.response_time")
	})

	//delete a file to a question
	authorized.DELETE("question/:question_id/file/:file_id", func(c *gin.Context) {
		client.Increment("webapp.delete_question_file")
		log.Info("deleting an file of a question with auth")
		t := client.NewTiming()
		controller.DeleteQuestionFileAuth(c, auth.GetCurrentUserID())
		t.Send("delete_question_file.response_time")
	})

	//upload a file to an answer
	authorized.POST("question/:question_id/answer/:answer_id/file", func(c *gin.Context) {
		client.Increment("webapp.post_answer_file")
		log.Info("creating an file for an answer with auth")
		t := client.NewTiming()
		controller.CreateAnswerFileAuth(c, auth.GetCurrentUserID())
		t.Send("post_answer_file.response_time")
	})

	//delete a file to an answer
	authorized.DELETE("question/:question_id/answer/:answer_id/file/:file_id", func(c *gin.Context) {
		client.Increment("webapp.delete_answer_file")
		log.Info("deleting an file of an answer with auth")
		t := client.NewTiming()
		controller.DeleteAnswerFileAuth(c, auth.GetCurrentUserID())
		t.Send("delete_answer_file.response_time")
	})

	return r
}
