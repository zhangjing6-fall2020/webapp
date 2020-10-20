package tool

import "os"

//Database username, password, hostname, and S3 bucket name get from ec2 instance
func GetBucketName() string {
	return os.Getenv("BUCKET_NAME")
}

func GetDBUserName() string {
	return os.Getenv("DB_USERNAME")
}

func GetDBPassword() string {
	return os.Getenv("DB_PASSWORD")
}

func GetHostname() string {
	return os.Getenv("HOSTNAME")
}