package controller

import (
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	log.Info("getting all users")
	var users []entity.User
	err := model.GetAllUsers(&users)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, user := range users {
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

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	log.Info("creating user")
	var user entity.User
	c.BindJSON(&user)

	var checkUser entity.User
	if err := model.GetUserByUsername(&checkUser, *user.Username); err == nil {
		log.Error("the email has been registered")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "the email has been registered!",
		})
		return
	}

	err := model.CreateUser(&user)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"id":              user.ID,
			"first_name":      user.FirstName,
			"last_name":       user.LastName,
			"username":        user.Username,
			"account_created": user.AccountCreated,
			"account_updated": user.AccountUpdated,
		})
	}
	log.Info("user created")
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	log.Info("getting user by id")
	id := c.Params.ByName("id")
	var user entity.User
	err := model.GetUserByID(&user, id)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
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

//GetUserByUsername ... Get the user by username
//Used in authorized get method endpoint: "/user/self"
func GetUserByUsername(c *gin.Context, username string) {
	log.Info("getting user by username")
	var user entity.User
	err := model.GetUserByUsername(&user, username)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
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

//UpdateAuthorizedUser ... update the user
//Used in authorized put method endpoint: "/user/self"
func UpdateAuthorizedUser(c *gin.Context, username string) {
	log.Info("updating user")
	var user entity.User
	err := model.GetUserByUsername(&user, username)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
	}
	//get the authorized user information
	oriID := user.ID
	var oriUsername string = *user.Username
	oriCreatedTime := user.AccountCreated
	oriPwd := user.Password
	c.BindJSON(&user)

	//if the user wants to change id, username or create time, it's not allowed
	if user.ID != oriID || *user.Username != oriUsername || user.AccountCreated != oriCreatedTime {
		log.Error("id, username and AccountCreated are readonly")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id, username and AccountCreated are readonly!!!",
		})
		return
	} else if oriPwd == user.Password {
		//if the user isn't updating password, don't need to check the password
		err = model.UpdateUserWithSamePwd(&user, user.ID)
		if err != nil {
			log.Error(err)
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
		//if the user is updating password, do check the password
		err = model.UpdateUserWithDiffPwd(&user, user.ID)
		if err != nil {
			log.Error(err)
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
	log.Info("user updated")
}

//UpdateUser ... Update the user information
func UpdateUserWithSamePwd(c *gin.Context) {
	var user entity.User
	id := c.Params.ByName("id")
	err := model.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&user)
	err = model.UpdateUserWithSamePwd(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
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

func UpdateUserWithDiffPwd(c *gin.Context) {
	var user entity.User
	id := c.Params.ByName("id")
	err := model.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&user)
	err = model.UpdateUserWithDiffPwd(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
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

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	log.Info("deleting user")
	var user entity.User
	id := c.Params.ByName("id")
	err := model.DeleteUser(&user, id)

	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
	log.Info("user deleted")
}
