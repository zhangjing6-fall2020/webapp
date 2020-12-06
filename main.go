package main

import (
	"cloudcomputing/webapp/config"
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/route"
	"cloudcomputing/webapp/tool"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alexcesaro/statsd.v2"
	"os"
)

var err error

func main() {
	log.Info("webapp starts...")

	//set up logrus
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	filename := "/var/log/webapp.log"
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Error(err)
	}
	log.SetOutput(f)

	//log.SetLevel(log.WarnLevel)

	//set up statsd
	client, err := statsd.New() // Connect to the UDP port 8125 by default.
	if err != nil {
		// If nothing is listening on the target port, an error is returned and
		// the returned client does nothing but is still usable. So we can
		// just log the error and go on.
		log.Error(err)
	}
	defer client.Close()

	//set up db
	cfg := mysql.Config{
		Addr:   tool.GetHostname(),            //"localhost",
		User:   tool.GetEnvVar("DB_USERNAME"), //"csye6225fall2020","root",
		Passwd: tool.GetEnvVar("DB_PASSWORD"), //"MysqlPwd123",
		DBName: tool.GetEnvVar("DB_NAME"),     //"csye6225",//"user_story",
		//TLSConfig: "custom",
	}
	config.DB, err = gorm.Open("mysql", cfg.FormatDSN())
	//config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
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
	r := route.SetupRouter(client)

	//running
	log.Info("webapp is running...")
	r.Run()
	log.Info("webapp ends...")
}
