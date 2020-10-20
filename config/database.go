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

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     tool.GetHostname(),//"localhost",
		Port:     3306,
		User:     tool.GetDBUserName(),//"root",
		Password: tool.GetDBPassword(),//"MysqlPwd123",
		DBName:   "user_story",
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
