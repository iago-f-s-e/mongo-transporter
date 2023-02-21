package utils

import (
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MakeDynamoExpressionToUpdate(fields interface{}) string {
	expressions := []string{}

	if typed, ok := fields.(primitive.M); ok {
		for key := range typed {
			expression := fmt.Sprintf("%s = :%s", key, key)

			expressions = append(expressions, expression)
		}

		return "SET " + strings.Join(expressions, ", ")
	}

	return ""
}

func MakeDynamoExpressionToRemove(fields interface{}) string {
	expressions := []string{}

	if typed, ok := fields.(primitive.M); ok {
		for key := range typed {
			expressions = append(expressions, key)
		}

		return "REMOVE " + strings.Join(expressions, ", ")
	}

	return ""
}

func MakeDynamoExpressionValues(toUpdate interface{}) primitive.M {
	values := make(map[string]interface{})

	if typed, ok := toUpdate.(primitive.M); ok {
		for key, value := range typed {
			values[":"+key] = value
		}

		return values
	}

	return values
}

func MakeDynamoExpressionValuesToRemove(toUpdate interface{}) primitive.M {
	values := make(map[string]interface{})

	if typed, ok := toUpdate.(primitive.M); ok {
		for key := range typed {
			values[":"+key] = nil
		}

		return values
	}

	return values
}
