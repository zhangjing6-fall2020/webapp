package main

import (
	"cloudcomputing/webapp/config"
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/route"
	"fmt"
	"github.com/jinzhu/gorm"
	"sort"
)

var err error

func main() {
	test()
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close()

	config.DB.AutoMigrate(&entity.User{})
	config.DB.AutoMigrate(&entity.Question{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	config.DB.AutoMigrate(&entity.Answer{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").AddForeignKey("question_id", "questions(id)", "RESTRICT", "RESTRICT")
	config.DB.AutoMigrate(&entity.Category{})
	config.DB.AutoMigrate(&entity.QuestionCategory{}).AddForeignKey("question_id", "questions(id)", "RESTRICT", "RESTRICT").AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")

	r := route.SetupRouter()

	//running
	r.Run()
}

func test()  {
	categories := entity.Categories{
		{
			ID: "1",
			Category: "a",
		},
		{
			ID: "2",
			Category: "b",
		},
		{
			ID: "3",
			Category: "c",
		},
	}

	for _, c := range categories{
		fmt.Println(c.ID+":"+c.Category)
	}
	sort.Sort(categories)
	for _, c := range categories{
		fmt.Println(c.ID+":"+c.Category)
	}
}
