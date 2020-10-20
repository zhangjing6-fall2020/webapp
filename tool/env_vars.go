package tool

import "os"

//Database username, password, hostname, and S3 bucket name get from ec2 instance
func GetBucketName() string {
	return os.Getenv("Bucket_Name")
}

func GetDBUserName() string {
	return os.Getenv("db_username")
}

func GetDBPassword() string {
	return os.Getenv("db_password")
}

func GetHostname() string {
	return os.Getenv("hostname")
}