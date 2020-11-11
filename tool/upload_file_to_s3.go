package tool

import (
	"cloudcomputing/webapp/entity"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	log "github.com/sirupsen/logrus"
)

func exitErrorf(msg string, args ...interface{}) {
	log.Errorf(msg+"\n", args...)
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	//os.Exit(1)
}

var sess *session.Session
var svc *s3.S3
//var statsDClient *statsd.Client = monitor.SetUpStatsD()

//Create a session using the setup Region and credentials
//https://docs.aws.amazon.com/sdk-for-go/api/aws/session/
func initSession() *session.Session {
	if sess == nil {
		log.Trace("initialize s3 session")
		newSess, err := session.NewSessionWithOptions(session.Options{
			// Specify profile to load for the session's config
			Profile: "dev",

			// Provide SDK Config options, such as Region.
			Config: aws.Config{
				Region: aws.String("us-east-1"),
			},

			// Force enable Shared Config support
			SharedConfigState: session.SharedConfigEnable,
		})

		if err != nil {
			log.Error("can't load the aws session")
		} else {
			log.Trace("loaded s3 session")
			sess = newSess
		}
	}

	return sess
}

func initClient() *s3.S3 {
	if svc == nil {
		sess = initSession()
		// Create S3 service client
		svc = s3.New(sess)
	}

	return svc
}

func listBuckets() {
	result, err := initClient().ListBuckets(nil)
	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}

}

func listBucketItems(bucketName string) {
	resp, err := initClient().ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucketName)})
	if err != nil {
		exitErrorf("Unable to list items in bucket %q, %v", bucketName, err)
	}

	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}

}

func UploadFile(bucketName string, fileHeader *multipart.FileHeader, objectName string) error {
	//t := statsDClient.NewTiming()
	sess = initSession()
	uploader := s3manager.NewUploader(sess)

	file, err := fileHeader.Open()
	if err != nil {
		fmt.Printf("Unable to open file %v", err)
		return err
	}
	defer file.Close()

	/*file, err := os.Open(filename)*/

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectName),
		Body:   file,
	})
	if err != nil {
		// Print the error and exit.
		//exitErrorf("Unable to upload %q to %q, %v", objectName, bucketName, err)
		fmt.Printf("Unable to upload %q to %q, %v", objectName, bucketName, err)
		return err
	}

	fmt.Printf("Successfully uploaded %q to %q\n", objectName, bucketName)
	//t.Send("upload_file.call_s3_service_time")
	return nil
}

/*
Output:
{
  Body: buffer(0xc000188f80)
}
*/
func GetTorrentMetaData(bucketName, objectName string) {
	svc = initClient()
	input := &s3.GetObjectTorrentInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectName),
	}

	result, err := svc.GetObjectTorrent(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

/*
Output:
{
  Grants: [{
      Grantee: {
        DisplayName: "ibh7cw+dev",
        ID: "c9aa2e9801a08bff81f203f709ebe15f223510aceb094da962d18e4dee697738",
        Type: "CanonicalUser"
      },
      Permission: "FULL_CONTROL"
    }],
  Owner: {
    DisplayName: "ibh7cw+dev",
    ID: "c9aa2e9801a08bff81f203f709ebe15f223510aceb094da962d18e4dee697738"
  }
}
*/
func GetAclMetaData(bucketName, objectName string) {
	svc = initClient()
	input := &s3.GetObjectAclInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectName),
	}

	result, err := svc.GetObjectAcl(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				fmt.Println(s3.ErrCodeNoSuchKey, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

/*
Output:
{
  AcceptRanges: "bytes",
  Body: buffer(0xc0001882c0),
  ContentLength: 98782,
  ContentType: "binary/octet-stream",
  ETag: "\"8cbba5a8dd9fdac7adcd570da1164eb8\"",
  LastModified: 2020-10-18 14:08:25 +0000 UTC
}
*/
func GetObjectMetaData(bucketName, objectName string) entity.Metadata {
	//t := statsDClient.NewTiming()
	svc = initClient()
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectName),
	}

	result, err := svc.GetObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				log.Errorf("s3 error no such key: %v", aerr.Error())
				fmt.Println(s3.ErrCodeNoSuchKey, aerr.Error())
			default:
				log.Error(aerr.Error())
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Fatalf(err.Error())
		}
	}

	fmt.Println(result)
	//t.Send("get_object_metaData.call_s3_service_time")
	return entity.Metadata{
		AcceptRanges:  result.AcceptRanges,
		ContentLength: result.ContentLength,
		ContentType:   result.ContentType,
		ETag:          result.ETag,
		LastModified:  result.LastModified,
	}
}

func DeleteFile(bucketName, filename string) error {
	//t := statsDClient.NewTiming()
	svc = initClient()
	if _, err := svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucketName), Key: aws.String(filename)}); err != nil {
		fmt.Printf("Unable to delete object %q from bucket %q, %v", filename, bucketName, err)
		return err
	}

	err := svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
	})

	if err != nil {
		// Print the error and exit.
		fmt.Printf("Unable to delete %q to %q, %v", filename, bucketName, err)
		return err
	}

	fmt.Printf("Successfully deleted %q to %q\n", filename, bucketName)
	//t.Send("delete_file.call_s3_service_time")
	return nil
}
