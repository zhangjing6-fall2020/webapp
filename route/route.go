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
		client.Count("webapp.post_user", p0)
		defer client.NewTiming().Send("post_user.response_time")
		controller.CreateUser(c)
	})

	j := 0
	//get all the users
	pub.GET("users", func(c *gin.Context) {
		j++
		client.Count("webapp.get_users", j)
		t := client.NewTiming()
		controller.GetUsers(c)
		t.Send("get_users.response_time")
	})

	//pub.GET("/user/:email_address", controller.GetUserByEmail)
	//pub.GET("user/:id", controller.GetUserByID)

	p2 := 0
	pub.GET("user/:id", func(c *gin.Context) {
		p2++
		client.Count("webapp.get_user", p2)
		t := client.NewTiming()
		controller.GetUserByID(c)
		t.Send("get_user.response_time")
	})

	p3 := 0
	//delete the user given UserID
	pub.DELETE("user/:id", func(c *gin.Context) {
		p3++
		client.Count("webapp.delete_user", p3)
		t := client.NewTiming()
		controller.DeleteUser(c)
		t.Send("delete_user.response_time")
	})

	p4 := 0
	//get all the questions
	pub.GET("questions", func(c *gin.Context) {
		p4++
		client.Count("webapp.get_questions", p4)
		t := client.NewTiming()
		controller.GetQuestions(c)
		t.Send("get_questions.response_time")
	})

	p5 := 0
	//get a question according to given id
	pub.GET("question/:question_id", func(c *gin.Context) {
		p5++
		client.Count("webapp.get_question", p5)
		t := client.NewTiming()
		controller.GetQuestionByID(c)
		t.Send("get_question.response_time")
	})

	p6 := 0
	//get a question's answer
	pub.GET("question/:question_id/answer/:answer_id", func(c *gin.Context) {
		p6++
		client.Count("webapp.get_answer", p6)
		t := client.NewTiming()
		controller.GetAnswerByID(c)
		t.Send("get_answer.response_time")
	})

	a0 := 0
	//authorized endpoints, basic auth is required with username and password
	authorized := r.Group("/v1", auth.BasicAuth())
	//get the information of the authorized user
	authorized.GET("userself", func(c *gin.Context) {
		a0++
		// Count a counter.
		client.Count("webapp.get_user_self", a0)
		log.Info("getting self user info")
		t := client.NewTiming()
		controller.GetUserByUsername(c, auth.GetCurrUsername())
		t.Send("get_user_self.response_time")
	})

	a1 := 0
	//edit the information of the authorized user
	authorized.PUT("user/self", func(c *gin.Context) {
		a1++
		client.Count("webapp.put_user_self", a1)
		log.Info("updating user info with auth")
		t := client.NewTiming()
		controller.UpdateAuthorizedUser(c, auth.GetCurrUsername())
		t.Send("put_user_self.response_time")
	})

	a2 := 0
	//post a new question
	authorized.POST("question", func(c *gin.Context) {
		a2++
		client.Count("webapp.post_question", a2)
		log.Info("creating a new question with auth")
		t := client.NewTiming()
		controller.CreateQuestionAuth(c, auth.GetCurrentUserID())
		t.Send("post_question.response_time")
	})

	a3 := 0
	//update a question
	authorized.PUT("question/:question_id", func(c *gin.Context) {
		a3++
		client.Count("webapp.put_question", a3)
		log.Info("updating a question with auth")
		t := client.NewTiming()
		controller.UpdateQuestionAuth(c, auth.GetCurrentUserID())
		t.Send("put_question.response_time")
	})

	a4 := 0
	//delete a question
	authorized.DELETE("question/:question_id", func(c *gin.Context) {
		a4++
		client.Count("webapp.delete_question", a4)
		log.Info("deleting a question with auth")
		t := client.NewTiming()
		controller.DeleteQuestionAuth(c, auth.GetCurrentUserID())
		t.Send("delete_question.response_time")
	})

	a5 := 0
	//answer a question
	authorized.POST("question/:question_id/answer", func(c *gin.Context) {
		a5++
		client.Count("webapp.post_answer", a5)
		log.Info("creating an answer for a question with auth")
		t := client.NewTiming()
		controller.CreateAnswer(c, auth.GetCurrentUserID())
		t.Send("post_answer.response_time")
	})

	a6 := 0
	//update a question's answer
	authorized.PUT("question/:question_id/answer/:answer_id", func(c *gin.Context) {
		a6++
		client.Count("webapp.put_answer", a6)
		log.Info("updating an answer of a question with auth")
		t := client.NewTiming()
		controller.UpdateAnswer(c, auth.GetCurrentUserID())
		t.Send("put_answer.response_time")
	})

	a7 := 0
	//delete a question's answer
	authorized.DELETE("question/:question_id/answer/:answer_id", func(c *gin.Context) {
		a7++
		client.Count("webapp.delete_answer", a7)
		log.Info("deleting an answer of a question with auth")
		t := client.NewTiming()
		controller.DeleteAnswer(c, auth.GetCurrentUserID())
		t.Send("delete_answer.response_time")
	})

	a8 := 0
	//upload a file to a question
	authorized.POST("question/:question_id/file", func(c *gin.Context) {
		a8++
		client.Count("webapp.post_question_file", a8)
		log.Info("creating an file for a question with auth")
		t := client.NewTiming()
		controller.CreateQuestionFileAuth(c, auth.GetCurrentUserID())
		t.Send("post_question_file.response_time")
	})

	a9 := 0
	//delete a file to a question
	authorized.DELETE("question/:question_id/file/:file_id", func(c *gin.Context) {
		a9++
		client.Count("webapp.delete_question_file", a9)
		log.Info("deleting an file of a question with auth")
		t := client.NewTiming()
		controller.DeleteQuestionFileAuth(c, auth.GetCurrentUserID())
		t.Send("delete_question_file.response_time")
	})

	a10 := 0
	//upload a file to an answer
	authorized.POST("question/:question_id/answer/:answer_id/file", func(c *gin.Context) {
		a10++
		client.Count("webapp.post_answer_file", a10)
		log.Info("creating an file for an answer with auth")
		t := client.NewTiming()
		controller.CreateAnswerFileAuth(c, auth.GetCurrentUserID())
		t.Send("post_answer_file.response_time")
	})

	a11 := 0
	//delete a file to an answer
	authorized.DELETE("question/:question_id/answer/:answer_id/file/:file_id", func(c *gin.Context) {
		a11++
		client.Count("webapp.delete_answer_file", a11)
		log.Info("deleting an file of an answer with auth")
		t := client.NewTiming()
		controller.DeleteAnswerFileAuth(c, auth.GetCurrentUserID())
		t.Send("delete_answer_file.response_time")
	})

	return r
}
