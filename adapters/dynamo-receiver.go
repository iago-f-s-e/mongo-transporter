package adapters

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DynamoReceiver struct {
	dbUri          string
	dbName         string
	collectionName string
	dbClient       *dynamodb.DynamoDB
}

func NewDynamoReceiver(dbUri string, dbName string, dbCollectName string, dbClient *dynamodb.DynamoDB) DynamoReceiver {
	return DynamoReceiver{
		dbUri:          dbUri,
		dbName:         dbName,
		collectionName: dbCollectName,
		dbClient:       dbClient,
	}
}

func (r DynamoReceiver) GetCollectionName() string {
	return r.collectionName
}

func (r DynamoReceiver) GetKey(id primitive.ObjectID) map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"_id": {
			S: aws.String(id.Hex()),
		},
	}
}

func (r DynamoReceiver) GetTableName() *string {
	return aws.String(r.collectionName)
}
