package utils

import (
	"errors"
)

func ConvertInterfaceToInt64(value interface{}) (int64, error) {
	switch value := value.(type) {
	case float64:
		return int64(value), nil

	case float32:
		return int64(value), nil

	case int64:
		return value, nil

	case int32:
		return int64(value), nil

	case int16:
		return int64(value), nil

	case int8:
		return int64(value), nil

	case int:
		return int64(value), nil

	default:
		return -1, errors.New("invalid type for conversion")
	}
}
