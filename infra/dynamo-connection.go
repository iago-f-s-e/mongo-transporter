package infra

import (
	"log"
	"mongo_transporter/constants"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func DynamoConnection(uri string, region string, accessKey string, secretAccess string, sessionToken string) *dynamodb.DynamoDB {
	if region == constants.ReceiverDefaultLocally {
		sess, err := session.NewSession(&aws.Config{
			DisableEndpointHostPrefix: aws.Bool(true),
			Region:                    aws.String(region),
			Endpoint:                  aws.String(uri),
		})

		if err != nil {
			log.Fatal(err)
		}

		return dynamodb.New(sess)
	}

	sess, err := session.NewSession(&aws.Config{
		DisableEndpointHostPrefix: aws.Bool(true),
		Region:                    aws.String(region),
		Credentials:               credentials.NewStaticCredentials(accessKey, secretAccess, sessionToken),
	})

	if err != nil {
		log.Fatal(err)
	}

	return dynamodb.New(sess)
}
