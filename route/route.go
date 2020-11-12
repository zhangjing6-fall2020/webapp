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

	var p0 int64
	pub.POST("user", func(c *gin.Context) {
		t := client.NewTiming()
		controller.CreateUser(c, client)
		t.Send("post_user.response_time")
		p0++
		client.Count("webapp.post_user", p0)
	})

	//var j int64
	//get all the users
	pub.GET("users", func(c *gin.Context) {
		//t := client.NewTiming()
		controller.GetUsers(c, client)
		//t.Send("get_users.response_time")
		//j++
		//client.Count("webapp.get_users", j)
	})

	//pub.GET("/user/:email_address", controller.GetUserByEmail)
	//pub.GET("user/:id", controller.GetUserByID)

	var p2 int64
	pub.GET("user/:id", func(c *gin.Context) {
		p2++
		t := client.NewTiming()
		controller.GetUserByID(c, client)
		t.Send("get_user.response_time")
		client.Count("webapp.get_user", p2)
	})

	var p3 int64
	//delete the user given UserID
	pub.DELETE("user/:id", func(c *gin.Context) {
		p3++
		t := client.NewTiming()
		controller.DeleteUser(c, client)
		t.Send("delete_user.response_time")
		client.Count("webapp.delete_user", p3)
	})

	var p4 int64
	//get all the questions
	pub.GET("questions", func(c *gin.Context) {
		p4++
		t := client.NewTiming()
		controller.GetQuestions(c, client)
		t.Send("get_questions.response_time")
		client.Count("webapp.get_questions", p4)
	})

	var p5 int64
	//get a question according to given id
	pub.GET("question/:question_id", func(c *gin.Context) {
		p5++
		t := client.NewTiming()
		controller.GetQuestionByID(c, client)
		t.Send("get_question.response_time")
		client.Count("webapp.get_question", p5)
	})

	var p6 int64
	//get a question's answer
	pub.GET("question/:question_id/answer/:answer_id", func(c *gin.Context) {
		p6++
		t := client.NewTiming()
		controller.GetAnswerByID(c, client)
		t.Send("get_answer.response_time")
		client.Count("webapp.get_answer", p6)
	})

	var a0 int64
	//authorized endpoints, basic auth is required with username and password
	authorized := r.Group("/v1", auth.BasicAuth(client))
	//get the information of the authorized user
	authorized.GET("userself", func(c *gin.Context) {
		a0++
		log.Info("getting self user info")
		t := client.NewTiming()
		controller.GetUserByUsername(c, auth.GetCurrUsername(), client)
		t.Send("get_user_self.response_time")
		client.Count("webapp.get_user_self", a0)
	})

	var a1 int64
	//edit the information of the authorized user
	authorized.PUT("user/self", func(c *gin.Context) {
		a1++
		log.Info("updating user info with auth")
		t := client.NewTiming()
		controller.UpdateAuthorizedUser(c, auth.GetCurrUsername(), client)
		t.Send("put_user_self.response_time")
		client.Count("webapp.put_user_self", a1)
	})

	var a2 int64
	//post a new question
	authorized.POST("question", func(c *gin.Context) {
		a2++
		log.Info("creating a new question with auth")
		t := client.NewTiming()
		controller.CreateQuestionAuth(c, auth.GetCurrentUserID(client), client)
		t.Send("post_question.response_time")
		client.Count("webapp.post_question", a2)
	})

	var a3 int64
	//update a question
	authorized.PUT("question/:question_id", func(c *gin.Context) {
		a3++
		log.Info("updating a question with auth")
		t := client.NewTiming()
		controller.UpdateQuestionAuth(c, auth.GetCurrentUserID(client), client)
		t.Send("put_question.response_time")
		client.Count("webapp.put_question", a3)
	})

	var a4 int64
	//delete a question
	authorized.DELETE("question/:question_id", func(c *gin.Context) {
		a4++
		log.Info("deleting a question with auth")
		t := client.NewTiming()
		controller.DeleteQuestionAuth(c, auth.GetCurrentUserID(client), client)
		t.Send("delete_question.response_time")
		client.Count("webapp.delete_question", a4)
	})

	var a5 int64
	//answer a question
	authorized.POST("question/:question_id/answer", func(c *gin.Context) {
		a5++
		log.Info("creating an answer for a question with auth")
		t := client.NewTiming()
		controller.CreateAnswer(c, auth.GetCurrentUserID(client), client)
		t.Send("post_answer.response_time")
		client.Count("webapp.post_answer", a5)
	})

	var a6 int64
	//update a question's answer
	authorized.PUT("question/:question_id/answer/:answer_id", func(c *gin.Context) {
		a6++
		log.Info("updating an answer of a question with auth")
		t := client.NewTiming()
		controller.UpdateAnswer(c, auth.GetCurrentUserID(client), client)
		t.Send("put_answer.response_time")
		client.Count("webapp.put_answer", a6)
	})

	var a7 int64
	//delete a question's answer
	authorized.DELETE("question/:question_id/answer/:answer_id", func(c *gin.Context) {
		a7++
		log.Info("deleting an answer of a question with auth")
		t := client.NewTiming()
		controller.DeleteAnswer(c, auth.GetCurrentUserID(client), client)
		t.Send("delete_answer.response_time")
		client.Count("webapp.delete_answer", a7)
	})

	var a8 int64
	//upload a file to a question
	authorized.POST("question/:question_id/file", func(c *gin.Context) {
		a8++
		log.Info("creating an file for a question with auth")
		t := client.NewTiming()
		controller.CreateQuestionFileAuth(c, auth.GetCurrentUserID(client), client)
		t.Send("post_question_file.response_time")
		client.Count("webapp.post_question_file", a8)
	})

	var a9 int64
	//delete a file to a question
	authorized.DELETE("question/:question_id/file/:file_id", func(c *gin.Context) {
		a9++
		log.Info("deleting an file of a question with auth")
		t := client.NewTiming()
		controller.DeleteQuestionFileAuth(c, auth.GetCurrentUserID(client), client)
		t.Send("delete_question_file.response_time")
		client.Count("webapp.delete_question_file", a9)
	})

	var a10 int64
	//upload a file to an answer
	authorized.POST("question/:question_id/answer/:answer_id/file", func(c *gin.Context) {
		a10++
		log.Info("creating an file for an answer with auth")
		t := client.NewTiming()
		controller.CreateAnswerFileAuth(c, auth.GetCurrentUserID(client), client)
		t.Send("post_answer_file.response_time")
		client.Count("webapp.post_answer_file", a10)
	})

	var a11 int64
	//delete a file to an answer
	authorized.DELETE("question/:question_id/answer/:answer_id/file/:file_id", func(c *gin.Context) {
		a11++
		log.Info("deleting an file of an answer with auth")
		t := client.NewTiming()
		controller.DeleteAnswerFileAuth(c, auth.GetCurrentUserID(client), client)
		t.Send("delete_answer_file.response_time")
		client.Count("webapp.delete_answer_file", a11)
	})

	return r
}
