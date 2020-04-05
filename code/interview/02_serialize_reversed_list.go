package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func createMap(splits []string, value string) map[string]interface{} {
	result := make(map[string]interface{})
	if len(splits) == 0 {
		return nil
	}
	if len(splits) == 1 {
		result[splits[0]] = value
		return result
	}
	result[splits[0]] = createMap(splits[1:len(splits)], value)
	return result

}

func storeToMap(result map[string]interface{}, splits []string, value string) (map[string]interface{}, error) {
	if len(splits) == 0 {
		return nil, nil
	}
	if len(splits) == 1 {
		result[splits[0]] = value
		return result, nil
	}
	var err error
	if v, ok := result[splits[0]]; ok {
		vMap, ok := v.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("error: %v can not store to map", splits)
		}
		result[splits[0]], err = storeToMap(vMap, splits[1:len(splits)], value)
		if err != nil {
			return nil, err
		}
	} else {
		result[splits[0]] = createMap(splits[1:len(splits)], value)
	}
	return result, nil

}

func SerializeReversed(revList map[string]string) (jsonStr string, err error) {
	result := make(map[string]interface{})
	for k, v := range revList {
		splits := strings.Split(v, ".")
		_, err := storeToMap(result, splits, k)
		if err != nil {
			return "", err
		}
	}
	jsonByte, err := json.Marshal(result)
	return string(jsonByte), err
}
