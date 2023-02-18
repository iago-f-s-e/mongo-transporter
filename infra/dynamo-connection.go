package infra

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func DynamoConnection(uri string, region string, disableSSL bool) *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		DisableSSL:                aws.Bool(disableSSL),
		DisableEndpointHostPrefix: aws.Bool(true),
		Region:                    aws.String(region),
		Endpoint:                  aws.String(uri),
	})

	if err != nil {
		log.Fatal(err)
	}

	return dynamodb.New(sess)
}
