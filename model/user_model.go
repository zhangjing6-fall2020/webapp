package model

import (
	"cloudcomputing/webapp/config"
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/tool"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	guuid "github.com/google/uuid"
	"time"
)

//GetAllUsers Fetch all user data
func GetAllUsers(user *[]entity.User) (err error) {
	if err = config.DB.Find(&user).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *entity.User) (err error) {
	user.ID = guuid.New().String()
	if !tool.CheckUsername(user.Username) {
		return errors.New("create user error: username is not email address!")
	}

	if err = tool.CheckPwd(8, 32, 2, user.Password); err != nil {
		return err
	}
	user.Password = tool.BcryptAndSalt(user.Password)
	user.AccountCreated = time.Now()
	user.AccountUpdated = time.Now()
	if err = config.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *entity.User, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByUsername(user *entity.User, username string) (err error) {
	if err = config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateUserWithSamePwd(user *entity.User, id string) (err error) {
	if !tool.CheckUsername(user.Username) {
		return errors.New("update user error: username is not email address!")
	}

	user.AccountUpdated = time.Now()
	config.DB.Save(&user)
	return nil
}

func UpdateUserWithDiffPwd(user *entity.User, id string) (err error) {
	if !tool.CheckUsername(user.Username) {
		return errors.New("update user error: username is not email address!")
	}

	if err = tool.CheckPwd(8, 32, 2, user.Password); err != nil {
		return err
	}
	user.Password = tool.BcryptAndSalt(user.Password)
	user.AccountUpdated = time.Now()
	config.DB.Save(&user)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *entity.User, id string) (err error) {
	if config.DB.Where("id = ?", id).First(&user); user.ID == "" {
		return errors.New("the user doesn't exist!!!")
	}
	config.DB.Where("id = ?", id).Delete(&user)
	return nil
}
