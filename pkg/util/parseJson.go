package util

import (
	"encoding/json"
)

func ParseJSON(jsonData []byte, target interface{}) error {
	return json.Unmarshal(jsonData, target)
}

func MarshalSturct(i interface{}) ([]byte, error) {
	result, err := json.MarshalIndent(i, "", "    ")
	if err != nil {
		return nil, err
	}
	return result, nil
}
