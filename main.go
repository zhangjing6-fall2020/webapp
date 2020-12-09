package main

import (
	"cloudcomputing/webapp/config"
	"cloudcomputing/webapp/entity"
	"cloudcomputing/webapp/route"
	"cloudcomputing/webapp/tool"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alexcesaro/statsd.v2"
	"os"
	"time"
)

var err error

/*type Result struct {
	id              int
	user            string
	host            string
	connection_type string
}*/

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
	rootCertPool := x509.NewCertPool()
	//download from https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html
	pem := `-----BEGIN CERTIFICATE-----
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
	if ok := rootCertPool.AppendCertsFromPEM([]byte(pem)); !ok {
		log.Fatal("Failed to append PEM.")
	}

	mysql.RegisterTLSConfig("custom", &tls.Config{
		RootCAs: rootCertPool,
	})

	cfg := mysql.Config{
		Addr:                 fmt.Sprintf("%s:%d", tool.GetHostname(), 3306), //"localhost",
		User:                 tool.GetEnvVar("DB_USERNAME"),                  //"csye6225fall2020","root",
		Passwd:               tool.GetEnvVar("DB_PASSWORD"),                  //"MysqlPwd123","Znt9yjNTp5NR"
		DBName:               tool.GetEnvVar("DB_NAME"),                      //"csye6225",//"user_story",
		Net:                  "tcp",
		AllowNativePasswords: true,
		TLSConfig:            "custom",
		Loc:                  time.Local,
		ParseTime:            true,
	}
	config.DB, err = gorm.Open("mysql", cfg.FormatDSN())
	//config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
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

	/*not working
	sql := `SELECT id, user, host, connection_type
       FROM performance_schema.threads pst
       INNER JOIN information_schema.processlist isp
       ON pst.processlist_id = isp.id;`
	var result Result
	config.DB.Raw(sql).Scan(&result)

	log.Infof("ssl connection:\n%v", result)
	fmt.Printf("ssl connection:\n%v", result)*/

	log.Info("waiting for request...")
	r := route.SetupRouter(client)

	//running
	log.Info("webapp is running...")
	r.Run()
	log.Info("webapp ends...")
}
