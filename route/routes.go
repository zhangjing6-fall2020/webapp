package route

import (
	"cloudcomputing/webapp/controller"
	"cloudcomputing/webapp/model"
	"cloudcomputing/webapp/tool"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var currUsername string

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	pub := r.Group("/v1")
	{
		pub.POST("user", controller.CreateUser)
		pub.GET("user", controller.GetUsers)
		//pub.GET("/user/:email_address", controller.GetUserByEmail)
		//pub.GET("user/:id", controller.GetUserByID)
		//pub.PUT("user/:id", controller.UpdateUser)
		pub.DELETE("user/:id", controller.DeleteUser)
	}

	authorized := r.Group("/v1", basicAuth())
	authorized.GET("/user/self", func(c *gin.Context) {
		var user model.User
		err := model.GetUserByUsername(&user, currUsername)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"id":              user.ID,
				"first_name":      user.FirstName,
				"last_name":       user.LastName,
				"username":        user.Username,
				"account_created": user.AccountCreated,
				"account_updated": user.AccountUpdated,
			})
		}
	})
	//authorized.GET("/user/:id", controller.GetUserByID)
	//authorized.GET("/user/:email_address", controller.GetUserByEmail)
	//authorized.PUT("/user/:id", controller.UpdateUser)
	authorized.PUT("/user/self", func(c *gin.Context) {
		var user model.User
		err := model.GetUserByUsername(&user, currUsername)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		}
		oriID := user.ID
		var oriUsername string = *user.Username
		oriCreatedTime := user.AccountCreated
		oriPwd := user.Password
		c.BindJSON(&user)

		if user.ID != oriID || *user.Username != oriUsername || user.AccountCreated != oriCreatedTime {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "id, username and AccountCreated are readonly!!!",
			})
			return
		} else if oriPwd == user.Password {
			err = model.UpdateUserWithSamePwd(&user, user.ID)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"id":              user.ID,
					"first_name":      user.FirstName,
					"last_name":       user.LastName,
					"username":        user.Username,
					"account_created": user.AccountCreated,
					"account_updated": user.AccountUpdated,
				})
			}
		} else {
			err = model.UpdateUserWithDiffPwd(&user, user.ID)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"id":              user.ID,
					"first_name":      user.FirstName,
					"last_name":       user.LastName,
					"username":        user.Username,
					"account_created": user.AccountCreated,
					"account_updated": user.AccountUpdated,
				})
			}
		}
	})

	return r
}

func basicAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			respondWithError(401, "Unauthorized", c)
			return
		}
		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !authenticateUser(pair[0], pair[1]) {
			respondWithError(401, "Unauthorized", c)
			return
		}

		c.Next()
	}
}

func authenticateUser(username, password string) bool {
	currUsername = username
	var user model.User
	err := model.GetUserByUsername(&user, username)
	if err != nil {
		fmt.Println("search user by email error: ", err)
		return false
	}
	if !tool.VerifyPasswd(user.Password, password) {
		fmt.Println("password is not right!!!")
		return false
	}
	return true
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}
