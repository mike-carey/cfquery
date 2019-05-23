package util

import (
	"reflect"
)

func MapKeys(subject interface{}) []string {
	mapKeys := reflect.ValueOf(subject).MapKeys()
	keys := make([]string, len(mapKeys))

	for i, key := range mapKeys {
		keys[i] = key.String()
		i++
	}

	return keys
}
