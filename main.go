package main

import (
	"cloudcomputing/webapp/config"
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/monitor"
	"cloudcomputing/webapp/route"
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var err error

func main() {
	log.Info("webapp starts...")

	//set up logrus
	monitor.SetUpLog()

	//set up statsD
	monitor.SetUpStatsD()

	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
		log.Errorf("failed to connect to database: %v", err)
	}
	defer config.DB.Close()

	config.DB.AutoMigrate(&entity.User{})
	config.DB.AutoMigrate(&entity.Question{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	config.DB.AutoMigrate(&entity.Answer{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").AddForeignKey("question_id", "questions(id)", "RESTRICT", "RESTRICT")
	config.DB.AutoMigrate(&entity.Category{})
	config.DB.AutoMigrate(&entity.QuestionCategory{}).AddForeignKey("question_id", "questions(id)", "RESTRICT", "RESTRICT").AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
	config.DB.AutoMigrate(&entity.File{}) //.AddForeignKey("question_id", "questions(id)", "RESTRICT", "RESTRICT").AddForeignKey("answer_id", "categories(id)", "RESTRICT", "RESTRICT")
	config.DB.AutoMigrate(&entity.AnswerFile{}).AddForeignKey("id", "files(id)", "RESTRICT", "RESTRICT").AddForeignKey("answer_id", "answers(id)", "RESTRICT", "RESTRICT")
	config.DB.AutoMigrate(&entity.QuestionFile{}).AddForeignKey("id", "files(id)", "RESTRICT", "RESTRICT").AddForeignKey("question_id", "questions(id)", "RESTRICT", "RESTRICT")
	log.Info("created tables in database")

	log.Info("waiting for request...")
	r := route.SetupRouter()

	//running
	log.Info("webapp is running...")
	r.Run()
	log.Info("webapp ends...")
}
