package config

import (
	"cloudcomputing/webapp/tool"
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

/*//local
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "MysqlPwd123",
		DBName:   "user_story",
	}
	return &dbConfig
}*/

//aws
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     tool.GetHostname(),//"localhost",
		Port:     3306,
		User:     tool.GetEnvVar("DB_USERNAME"),//"csye6225fall2020","root",
		Password: tool.GetEnvVar("DB_PASSWORD"),//"MysqlPwd123",
		DBName:   tool.GetEnvVar("DB_NAME"),//"csye6225",//"user_story",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
