package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapByPrimitiveA(doc primitive.A) []any {
	var a []any

	for _, v := range doc {

		switch typed := v.(type) {
		case primitive.D:
			newM := MakeMongoMap(typed)
			a = append(a, newM)
			continue

		case primitive.A:
			newA := MapByPrimitiveA(typed)
			a = append(a, newA)
			continue

		default:
			a = append(a, typed)
			continue
		}
	}

	return a
}

func MakeMongoMap(doc interface{}) primitive.M {
	obj := make(map[string]interface{})

	for _, prop := range doc.(primitive.D) {
		obj[prop.Key] = prop.Value

		switch valueTyped := prop.Value.(type) {
		case primitive.A:
			newObj := MapByPrimitiveA(valueTyped)
			obj[prop.Key] = newObj
			continue
		case primitive.D:
			newObj := MakeMongoMap(valueTyped)
			obj[prop.Key] = newObj
			continue
		default:
			obj[prop.Key] = prop.Value
			continue
		}
	}

	return obj
}
