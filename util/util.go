package util

import (
	"reflect"

	"github.com/iancoleman/strcase"
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

func TranslateKeyToClassName(key string) string {
	return strcase.ToCamel(key)
}
