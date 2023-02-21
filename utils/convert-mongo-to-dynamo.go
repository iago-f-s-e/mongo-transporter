package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertMongoToDynamo(doc *interface{}, primitiveType string) {

	if primitiveType == "A" {
		for index, obj := range (*doc).(primitive.A) {

			switch objTyped := obj.(type) {
			case primitive.M:
				for key, prop := range objTyped {

					switch typed := prop.(type) {
					case primitive.ObjectID:
						obj.(primitive.M)[key] = typed.Hex()
						continue

					case primitive.DateTime:
						obj.(primitive.M)[key] = typed.Time().Format("2006-01-02T15:04:05.000Z")
						continue

					case primitive.M:
						converted := obj.(primitive.M)[key]

						ConvertMongoToDynamo(&converted, "M")

						obj.(primitive.M)[key] = converted
						continue

					case primitive.A:
						var converted interface{} = typed

						ConvertMongoToDynamo(&converted, "A")

						obj.(primitive.M)[key] = converted
						continue

					default:
						continue
					}

				}

				(*doc).(primitive.A)[index] = obj
				continue

			default:

				(*doc).(primitive.A)[index] = objTyped

				continue
			}

		}
	} else {
		for key, value := range (*doc).(primitive.M) {

			switch typed := value.(type) {
			case primitive.ObjectID:
				(*doc).(primitive.M)[key] = typed.Hex()
				continue

			case primitive.DateTime:
				(*doc).(primitive.M)[key] = typed.Time().Format("2006-01-02T15:04:05.000Z")
				continue

			case primitive.A:
				var converted interface{} = typed

				ConvertMongoToDynamo(&converted, "A")

				(*doc).(primitive.M)[key] = converted
				continue

			case primitive.M:
				var converted interface{} = typed

				ConvertMongoToDynamo(&converted, "M")

				(*doc).(primitive.M)[key] = converted
				continue

			case []primitive.M:
				var arrConverted []interface{}

				for _, v := range typed {
					var converted interface{} = v

					ConvertMongoToDynamo(&converted, "M")

					arrConverted = append(arrConverted, converted)
				}

				(*doc).(primitive.M)[key] = arrConverted
				continue

			default:
				continue
			}
		}
	}
}
