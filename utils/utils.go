package utils

import (
	"log"
	"reflect"
)

func ToMapInterface(input interface{}) map[string]interface{} {
	inputVal := reflect.ValueOf(input)
	inputType := reflect.TypeOf(input)

	if inputType.Kind() != reflect.Map {
		log.Fatal("Input is not a map")
	}

	output := make(map[string]interface{})
	for _, key := range inputVal.MapKeys() {
		output[key.String()] = inputVal.MapIndex(key).Interface()
	}
	return output
}
