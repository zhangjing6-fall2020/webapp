package tool

import (
	"os"
	"strings"
)

//Database username, password, hostname, and S3 bucket name get from ec2 instance
//ec2
func GetBucketName() string {
	return os.Getenv("BUCKET_NAME")
}

//local
/*func GetBucketName() string {
	return "webapp.jing.zhang"
}*/

func GetDBUserName() string {
	return os.Getenv("DB_USERNAME")
}

func GetDBPassword() string {
	return os.Getenv("DB_PASSWORD")
}

func GetHostname() string {
	hostname := os.Getenv("DBHOSTNAME")
	return strings.TrimRight(hostname, ":3306")
}

func GetEnvVar(env string) string {
	return os.Getenv(env)
}
