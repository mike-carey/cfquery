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

func StackErrors(errs []error) error {
	s := ""
	if len(errs) > 0 {
		s = "s"
	}

	err := fmt.Sprintf("%d error%s occured:\n", len(errs), s)
	for _, e := range errs {
		err += fmt.Sprintf("%q\n", e)
	}

	return errors.New(err)
}
