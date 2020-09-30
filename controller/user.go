package controller

import (
	"cloudcomputing/webapp/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var users []model.User
	err := model.GetAllUsers(&users)
	if err != nil {
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
	var user model.User
	c.BindJSON(&user)

	var checkUser model.User
	if err := model.GetUserByUsername(&checkUser, *user.Username); err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "the email has been registered!",
		})
		return
	}

	err := model.CreateUser(&user)
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

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user model.User
	err := model.GetUserByID(&user, id)
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

//GetUserByID ... Get the user by id
func GetUserByUsername(c *gin.Context) {
	username := c.Params.ByName("username")
	var user model.User
	err := model.GetUserByUsername(&user, username)
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

//UpdateUser ... Update the user information
func UpdateUserWithSamePwd(c *gin.Context) {
	var user model.User
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
	var user model.User
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
	var user model.User
	id := c.Params.ByName("id")
	err := model.DeleteUser(&user, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
