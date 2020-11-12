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

	p0 := 0
	pub.POST("user", func(c *gin.Context) {
		p0++
		t := client.NewTiming()
		controller.CreateUser(c, client)
		t.Send("post_user.response_time")
		client.Count("webapp.post_user", p0)
	})

	j := 0
	//get all the users
	pub.GET("users", func(c *gin.Context) {
		j++
		t := client.NewTiming()
		controller.GetUsers(c, client)
		t.Send("get_users.response_time")
		client.Count("webapp.get_users", j)
	})

	//pub.GET("/user/:email_address", controller.GetUserByEmail)
	//pub.GET("user/:id", controller.GetUserByID)

	p2 := 0
	pub.GET("user/:id", func(c *gin.Context) {
		p2++
		t := client.NewTiming()
		controller.GetUserByID(c, client)
		t.Send("get_user.response_time")
		client.Count("webapp.get_user", p2)
	})

	p3 := 0
	//delete the user given UserID
	pub.DELETE("user/:id", func(c *gin.Context) {
		p3++
		t := client.NewTiming()
		controller.DeleteUser(c, client)
		t.Send("delete_user.response_time")
		client.Count("webapp.delete_user", p3)
	})

	p4 := 0
	//get all the questions
	pub.GET("questions", func(c *gin.Context) {
		p4++
		t := client.NewTiming()
		controller.GetQuestions(c, client)
		t.Send("get_questions.response_time")
		client.Count("webapp.get_questions", p4)
	})

	p5 := 0
	//get a question according to given id
	pub.GET("question/:question_id", func(c *gin.Context) {
		p5++
		t := client.NewTiming()
		controller.GetQuestionByID(c, client)
		t.Send("get_question.response_time")
		client.Count("webapp.get_question", p5)
	})

	p6 := 0
	//get a question's answer
	pub.GET("question/:question_id/answer/:answer_id", func(c *gin.Context) {
		p6++
		t := client.NewTiming()
		controller.GetAnswerByID(c, client)
		t.Send("get_answer.response_time")
		client.Count("webapp.get_answer", p6)
	})

	a0 := 0
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

	a1 := 0
	//edit the information of the authorized user
	authorized.PUT("user/self", func(c *gin.Context) {
		a1++
		log.Info("updating user info with auth")
		t := client.NewTiming()
		controller.UpdateAuthorizedUser(c, auth.GetCurrUsername(), client)
		t.Send("put_user_self.response_time")
		client.Count("webapp.put_user_self", a1)
	})

	a2 := 0
	//post a new question
	authorized.POST("question", func(c *gin.Context) {
		a2++
		log.Info("creating a new question with auth")
		t := client.NewTiming()
		controller.CreateQuestionAuth(c, auth.GetCurrentUserID(client), client)
		t.Send("post_question.response_time")
		client.Count("webapp.post_question", a2)
	})

	a3 := 0
	//update a question
	authorized.PUT("question/:question_id", func(c *gin.Context) {
		a3++
		log.Info("updating a question with auth")
		t := client.NewTiming()
		controller.UpdateQuestionAuth(c, auth.GetCurrentUserID(client), client)
		t.Send("put_question.response_time")
		client.Count("webapp.put_question", a3)
	})

	a4 := 0
	//delete a question
	authorized.DELETE("question/:question_id", func(c *gin.Context) {
		a4++
		log.Info("deleting a question with auth")
		t := client.NewTiming()
		controller.DeleteQuestionAuth(c, auth.GetCurrentUserID(client), client)
		t.Send("delete_question.response_time")
		client.Count("webapp.delete_question", a4)
	})

	a5 := 0
	//answer a question
	authorized.POST("question/:question_id/answer", func(c *gin.Context) {
		a5++
		log.Info("creating an answer for a question with auth")
		t := client.NewTiming()
		controller.CreateAnswer(c, auth.GetCurrentUserID(client), client)
		t.Send("post_answer.response_time")
		client.Count("webapp.post_answer", a5)
	})

	a6 := 0
	//update a question's answer
	authorized.PUT("question/:question_id/answer/:answer_id", func(c *gin.Context) {
		a6++
		log.Info("updating an answer of a question with auth")
		t := client.NewTiming()
		controller.UpdateAnswer(c, auth.GetCurrentUserID(client), client)
		t.Send("put_answer.response_time")
		client.Count("webapp.put_answer", a6)
	})

	a7 := 0
	//delete a question's answer
	authorized.DELETE("question/:question_id/answer/:answer_id", func(c *gin.Context) {
		a7++
		log.Info("deleting an answer of a question with auth")
		t := client.NewTiming()
		controller.DeleteAnswer(c, auth.GetCurrentUserID(client), client)
		t.Send("delete_answer.response_time")
		client.Count("webapp.delete_answer", a7)
	})

	a8 := 0
	//upload a file to a question
	authorized.POST("question/:question_id/file", func(c *gin.Context) {
		a8++
		log.Info("creating an file for a question with auth")
		t := client.NewTiming()
		controller.CreateQuestionFileAuth(c, auth.GetCurrentUserID(client), client)
		t.Send("post_question_file.response_time")
		client.Count("webapp.post_question_file", a8)
	})

	a9 := 0
	//delete a file to a question
	authorized.DELETE("question/:question_id/file/:file_id", func(c *gin.Context) {
		a9++
		log.Info("deleting an file of a question with auth")
		t := client.NewTiming()
		controller.DeleteQuestionFileAuth(c, auth.GetCurrentUserID(client), client)
		t.Send("delete_question_file.response_time")
		client.Count("webapp.delete_question_file", a9)
	})

	a10 := 0
	//upload a file to an answer
	authorized.POST("question/:question_id/answer/:answer_id/file", func(c *gin.Context) {
		a10++
		log.Info("creating an file for an answer with auth")
		t := client.NewTiming()
		controller.CreateAnswerFileAuth(c, auth.GetCurrentUserID(client), client)
		t.Send("post_answer_file.response_time")
		client.Count("webapp.post_answer_file", a10)
	})

	a11 := 0
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
