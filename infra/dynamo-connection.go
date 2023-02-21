package infra

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func DynamoConnection(uri string, region string, accessKey string, secretAccess string, sessionToken string) *dynamodb.DynamoDB {
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
