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
	t := statsDClient.NewTiming()
	if err = config.DB.Find(&user).Error; err != nil {
		return err
	}
	t.Send("get_all_users.query_time")
	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *entity.User) (err error) {
	t := statsDClient.NewTiming()
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
	t.Send("create_user.query_time")
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *entity.User, id string) (err error) {
	t := statsDClient.NewTiming()
	if err = config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}
	t.Send("get_user_by_id.query_time")
	return nil
}

func GetUserByUsername(user *entity.User, username string) (err error) {
	t := statsDClient.NewTiming()
	if err = config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return err
	}
	t.Send("get_user_by_username.query_time")
	return nil
}

//UpdateUser ... Update user
func UpdateUserWithSamePwd(user *entity.User, id string) (err error) {
	t := statsDClient.NewTiming()
	if !tool.CheckUsername(user.Username) {
		return errors.New("update user error: username is not email address!")
	}

	user.AccountUpdated = time.Now()
	config.DB.Save(&user)
	t.Send("update_user_with_same_pwd.query_time")
	return nil
}

func UpdateUserWithDiffPwd(user *entity.User, id string) (err error) {
	t := statsDClient.NewTiming()
	if !tool.CheckUsername(user.Username) {
		return errors.New("update user error: username is not email address!")
	}

	if err = tool.CheckPwd(8, 32, 2, user.Password); err != nil {
		return err
	}
	user.Password = tool.BcryptAndSalt(user.Password)
	user.AccountUpdated = time.Now()
	config.DB.Save(&user)
	t.Send("update_user_with_diff_pwd.query_time")
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *entity.User, id string) (err error) {
	t := statsDClient.NewTiming()
	if config.DB.Where("id = ?", id).First(&user); user.ID == "" {
		return errors.New("the user doesn't exist!!!")
	}
	config.DB.Where("id = ?", id).Delete(&user)
	t.Send("delete_user.query_time")
	return nil
}
