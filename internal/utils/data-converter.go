package utils

import (
	"encoding/json"
)

func Convert[T any](data interface{}) (T, error) {
	var result T
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(dataBytes, &result); err != nil {
		return result, err
	}

	return result, nil
}
