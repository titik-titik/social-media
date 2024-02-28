package tool

import "encoding/json"

func DeepCopy[T any](data T) T {
	jsonByte, marshallErr := json.Marshal(data)
	if marshallErr != nil {
		panic(marshallErr)
	}

	var result T
	unmarshallErr := json.Unmarshal(jsonByte, &result)
	if unmarshallErr != nil {
		panic(unmarshallErr)
	}

	return result
}
