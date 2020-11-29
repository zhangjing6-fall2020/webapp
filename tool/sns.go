package tool

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	log "github.com/sirupsen/logrus"
)

var sns_session *session.Session
var sns_client *sns.SNS

func initSNSSession() *session.Session {
	if sns_session == nil {
		log.Info("initialize SNS session")

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
			log.Info("loaded SNS session")
			sns_session = newSess
		}
	}

	log.Info("initialized SNS session")
	return sns_session
}

func initSNSClient() *sns.SNS {
	if sns_client == nil {
		sns_session = initSNSSession()
		// Create SNS client
		sns_client = sns.New(sns_session)
	}
	log.Info("initialized SNS client")
	return sns_client
}

func PublishMessageOnSNS(message string) (*sns.PublishOutput, error) {
	log.Info("start to publish message in sns")
	fmt.Println("start to publish message in sns")
	input := &sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String("arn:aws:sns:us-east-1:907204364947:topic"),
	}

	result, err := initSNSClient().Publish(input)
	if err != nil {
		log.Errorf("Failed to publish message: %v\n", err)
		fmt.Printf("Failed to publish message: %v\n", err)
		return nil, err
	}
	log.Info("completed to publish message in sns")
	fmt.Println("completed to publish message in sns")
	return result, nil
}