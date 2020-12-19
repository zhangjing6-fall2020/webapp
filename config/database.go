package config

import (
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/tool"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"time"
)

var DB *gorm.DB

//download from https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html
const pem = `-----BEGIN CERTIFICATE-----
MIIEBjCCAu6gAwIBAgIJAMc0ZzaSUK51MA0GCSqGSIb3DQEBCwUAMIGPMQswCQYD
VQQGEwJVUzEQMA4GA1UEBwwHU2VhdHRsZTETMBEGA1UECAwKV2FzaGluZ3RvbjEi
MCAGA1UECgwZQW1hem9uIFdlYiBTZXJ2aWNlcywgSW5jLjETMBEGA1UECwwKQW1h
em9uIFJEUzEgMB4GA1UEAwwXQW1hem9uIFJEUyBSb290IDIwMTkgQ0EwHhcNMTkw
ODIyMTcwODUwWhcNMjQwODIyMTcwODUwWjCBjzELMAkGA1UEBhMCVVMxEDAOBgNV
BAcMB1NlYXR0bGUxEzARBgNVBAgMCldhc2hpbmd0b24xIjAgBgNVBAoMGUFtYXpv
biBXZWIgU2VydmljZXMsIEluYy4xEzARBgNVBAsMCkFtYXpvbiBSRFMxIDAeBgNV
BAMMF0FtYXpvbiBSRFMgUm9vdCAyMDE5IENBMIIBIjANBgkqhkiG9w0BAQEFAAOC
AQ8AMIIBCgKCAQEArXnF/E6/Qh+ku3hQTSKPMhQQlCpoWvnIthzX6MK3p5a0eXKZ
oWIjYcNNG6UwJjp4fUXl6glp53Jobn+tWNX88dNH2n8DVbppSwScVE2LpuL+94vY
0EYE/XxN7svKea8YvlrqkUBKyxLxTjh+U/KrGOaHxz9v0l6ZNlDbuaZw3qIWdD/I
6aNbGeRUVtpM6P+bWIoxVl/caQylQS6CEYUk+CpVyJSkopwJlzXT07tMoDL5WgX9
O08KVgDNz9qP/IGtAcRduRcNioH3E9v981QO1zt/Gpb2f8NqAjUUCUZzOnij6mx9
McZ+9cWX88CRzR0vQODWuZscgI08NvM69Fn2SQIDAQABo2MwYTAOBgNVHQ8BAf8E
BAMCAQYwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUc19g2LzLA5j0Kxc0LjZa
pmD/vB8wHwYDVR0jBBgwFoAUc19g2LzLA5j0Kxc0LjZapmD/vB8wDQYJKoZIhvcN
AQELBQADggEBAHAG7WTmyjzPRIM85rVj+fWHsLIvqpw6DObIjMWokpliCeMINZFV
ynfgBKsf1ExwbvJNzYFXW6dihnguDG9VMPpi2up/ctQTN8tm9nDKOy08uNZoofMc
NUZxKCEkVKZv+IL4oHoeayt8egtv3ujJM6V14AstMQ6SwvwvA93EP/Ug2e4WAXHu
cbI1NAbUgVDqp+DRdfvZkgYKryjTWd/0+1fS8X1bBZVWzl7eirNVnHbSH2ZDpNuY
0SBd8dj5F6ld3t58ydZbrTHze7JJOd8ijySAp4/kiu9UfZWuTPABzDa/DSdz9Dk/
zPW4CXXvhLmE02TA9/HeCw3KEHIwicNuEfw=
-----END CERTIFICATE-----`

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func ConfigDB() error {
	rootCertPool := x509.NewCertPool()
	if ok := rootCertPool.AppendCertsFromPEM([]byte(pem)); !ok {
		return errors.New("failed to append PEM")
	}

	mysql.RegisterTLSConfig("custom", &tls.Config{
		RootCAs: rootCertPool,
	})

	cfg := mysql.Config{
		Addr:                 fmt.Sprintf("%s:%d", tool.GetHostname(), 3306), //"localhost",
		User:                 tool.GetEnvVar("DB_USERNAME"),                  //"csye6225fall2020","root",
		Passwd:               tool.GetEnvVar("DB_PASSWORD"),                  //"Znt9yjNTp5NR"
		DBName:               tool.GetEnvVar("DB_NAME"),                      //"csye6225", "user_story",
		Net:                  "tcp",
		AllowNativePasswords: true,
		TLSConfig:            "custom",
		Loc:                  time.Local,
		ParseTime:            true,
	}
	DB, err := gorm.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer DB.Close()

	DB.AutoMigrate(&entity.User{})
	DB.AutoMigrate(&entity.Question{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	DB.AutoMigrate(&entity.Answer{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").AddForeignKey("question_id", "questions(id)", "RESTRICT", "RESTRICT")
	DB.AutoMigrate(&entity.Category{})
	DB.AutoMigrate(&entity.QuestionCategory{}).AddForeignKey("question_id", "questions(id)", "RESTRICT", "RESTRICT").AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
	DB.AutoMigrate(&entity.File{})
	DB.AutoMigrate(&entity.AnswerFile{}).AddForeignKey("id", "files(id)", "RESTRICT", "RESTRICT").AddForeignKey("answer_id", "answers(id)", "RESTRICT", "RESTRICT")
	DB.AutoMigrate(&entity.QuestionFile{}).AddForeignKey("id", "files(id)", "RESTRICT", "RESTRICT").AddForeignKey("question_id", "questions(id)", "RESTRICT", "RESTRICT")
	log.Info("created tables in database")

	return nil
}
