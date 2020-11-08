package auth

import (
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/model"
	"cloudcomputing/webapp/tool"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strings"
)

var currUsername string

//verify the basic auth with username and password
func BasicAuth() gin.HandlerFunc {
	log.Info("auth...")
	return func(c *gin.Context) {
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			respondWithError(401, "Unauthorized", c)
			log.Error("unauthorized request type")
			return
		}
		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !authenticateUser(pair[0], pair[1]) {
			respondWithError(401, "Unauthorized", c)
			log.Error("unauthorized user")
			return
		}

		c.Next()
	}
}

//check the username and password are the same as in the database to verify the user
func authenticateUser(username, password string) bool {
	currUsername = username
	var user entity.User
	err := model.GetUserByUsername(&user, username)
	if err != nil {
		fmt.Println("search user by email error: ", err)
		log.Errorf("email doesn't exist: %v", err)
		return false
	}
	if !tool.VerifyPasswd(user.Password, password) {
		fmt.Println("password is not right!!!")
		log.Error("password is not right!")
		return false
	}
	return true
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}

func GetCurrUsername()  string {
	return currUsername
}

func GetCurrentUserID()  string {
	var user entity.User
	err := model.GetUserByUsername(&user, currUsername)
	if err != nil {
		log.Error("failed to get current user")
		return err.Error()
	} else {
		return user.ID
	}
}
