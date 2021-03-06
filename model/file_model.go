package model

import (
	"cloudcomputing/webapp/config"
	"cloudcomputing/webapp/entity"
	"errors"
	"fmt"
	guuid "github.com/google/uuid"
	"gopkg.in/alexcesaro/statsd.v2"
	"time"
)

//GetAllFiles Fetch all File data
func GetAllFiles(file *[]entity.File) (err error) {
	if err = config.DB.Find(&file).Error; err != nil {
		return err
	}
	return nil
}

//CreateFile ... Insert New data
func CreateFileAuth(file *entity.File, client *statsd.Client) (err error) {
	t := client.NewTiming()
	file.CreatedDate = time.Now()
	if err = config.DB.Create(&file).Error; err != nil {
		return err
	}
	t.Send("create_file.query_time")
	return nil
}

//CreateFile ... Insert New data
func CreateFile(file *entity.File, questionOrAnswerID string) (err error) {
	file.ID = guuid.New().String()
	file.CreatedDate = time.Now()
	file.S3ObjectName = fmt.Sprintf("%s/%s/%s", questionOrAnswerID, file.ID, file.FileName)
	if err = config.DB.Create(&file).Error; err != nil {
		return err
	}
	return nil
}

//GetFileByID ... Fetch only one file by Id
func GetFileByID(file *entity.File, id string, client *statsd.Client) (err error) {
	t := client.NewTiming()
	if err = config.DB.Where("id = ?", id).First(&file).Error; err != nil {
		return err
	}
	t.Send("get_file_by_id.query_time")
	return nil
}

//DeleteFile ... Delete file
func DeleteFile(file *entity.File, id string, client *statsd.Client) (err error) {
	t := client.NewTiming()
	if config.DB.Where("id = ?", id).First(&file); file.ID == "" {
		return errors.New("the file doesn't exist!!!")
	}
	config.DB.Where("id = ?", id).Delete(&file)
	t.Send("delete_file.query_time")
	return nil
}
