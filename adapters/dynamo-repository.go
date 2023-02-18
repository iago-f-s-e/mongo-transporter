package adapters

import (
	"context"
	"log"
	"mongo_transporter/utils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r DynamoReceiver) SetupCollection(ctx context.Context) {
	r.dbClient.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: r.GetTableName(),
	})

	attributeDefinitions := []*dynamodb.AttributeDefinition{
		{
			AttributeName: aws.String("_id"),
			AttributeType: aws.String("S"),
		},
	}

	keySchema := []*dynamodb.KeySchemaElement{
		{
			AttributeName: aws.String("_id"),
			KeyType:       aws.String("HASH"),
		},
	}

	provisionedThroughput := &dynamodb.ProvisionedThroughput{
		ReadCapacityUnits:  aws.Int64(25),
		WriteCapacityUnits: aws.Int64(25),
	}

	_, err := r.dbClient.CreateTable(&dynamodb.CreateTableInput{
		ProvisionedThroughput: provisionedThroughput,
		AttributeDefinitions:  attributeDefinitions,
		KeySchema:             keySchema,
		TableName:             r.GetTableName(),
	})

	if err != nil {
		log.Fatal(err)
	}

}

func (r DynamoReceiver) InsertOnCollection(ctx context.Context, documents []interface{}) {
	batchSize := 25

	for i := 0; i < batchSize; i += batchSize {
		end := i + batchSize

		if end > len(documents) {
			end = len(documents)
		}

		batch := &dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]*dynamodb.WriteRequest{
				r.collectionName: {},
			},
		}

		for _, document := range documents[i:end] {
			utils.ConvertMongoToDynamo(&document, "M")

			item, err := dynamodbattribute.MarshalMap(document)

			if err != nil {
				log.Fatal(err)
			}

			batch.RequestItems[r.collectionName] = append(batch.RequestItems[r.collectionName], &dynamodb.WriteRequest{
				PutRequest: &dynamodb.PutRequest{
					Item: item,
				},
			})
		}

		_, err := r.dbClient.BatchWriteItem(batch)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func (r DynamoReceiver) ReflectWatchOnInsert(ctx context.Context, fullDocument primitive.M) {
	var document interface{} = fullDocument

	utils.ConvertMongoToDynamo(&document, "M")

	item, err := dynamodbattribute.MarshalMap(document)

	if err != nil {
		log.Fatal(err)
	}

	_, err = r.dbClient.PutItem(&dynamodb.PutItemInput{
		Item:      item,
		TableName: r.GetTableName(),
	})

	if err != nil {
		log.Fatal(err)
	}
}

func (r DynamoReceiver) ReflectWatchOnDelete(ctx context.Context, id primitive.ObjectID) {
	_, err := r.dbClient.DeleteItem(&dynamodb.DeleteItemInput{
		Key:       r.GetKey(id),
		TableName: r.GetTableName(),
	})

	if err != nil {
		log.Fatal(err)
	}
}

func (r DynamoReceiver) ReflectWatchOnUpdate(ctx context.Context, id primitive.ObjectID, updatedFields primitive.D, removedFields primitive.M) {

	if len(updatedFields) > 0 {
		var fields interface{} = utils.MakeMongoMap(updatedFields)
		expression := utils.MakeDynamoExpressionToUpdate(fields)

		var values interface{} = utils.MakeDynamoExpressionValues(fields)
		utils.ConvertMongoToDynamo(&values, "M")
		expressionValues, err := dynamodbattribute.MarshalMap(values)

		if err != nil {
			log.Fatal(err)
		}

		params := &dynamodb.UpdateItemInput{
			TableName:                 r.GetTableName(),
			Key:                       r.GetKey(id),
			ExpressionAttributeValues: expressionValues,
			UpdateExpression:          aws.String(expression),
		}

		_, err = r.dbClient.UpdateItem(params)

		if err != nil {
			log.Fatal(err)
		}
	}

	if len(removedFields) > 0 {
		expression := utils.MakeDynamoExpressionToRemove(removedFields)

		params := &dynamodb.UpdateItemInput{
			TableName:        r.GetTableName(),
			Key:              r.GetKey(id),
			UpdateExpression: aws.String(expression),
		}

		_, err := r.dbClient.UpdateItem(params)

		if err != nil {
			log.Fatal(err)
		}
	}
}
