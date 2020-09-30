package main

import (
	"cloudcomputing/webapp/config"
	"cloudcomputing/webapp/model"
	"cloudcomputing/webapp/route"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close()

	config.DB.AutoMigrate(&model.User{})

	/*user := model.User{ID: "6", FirstName: "jin", LastName: "z", Password: "123",
		EmailAddress: "123@gmail.com", AccountCreated: time.Now(), AccountUpdated: time.Now()}
	if err = config.DB.Create(&user).Error; err != nil {
		fmt.Println("Create User error: ", err)
	}*/

	r := route.SetupRouter()

	//running
	r.Run()
}
