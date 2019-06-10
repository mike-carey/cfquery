package query_test

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestQuery(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Query Suite")
}

func getReflectionInfo(subject interface{}) (reflect.Value, reflect.Type, reflect.Kind) {
	v := reflect.ValueOf(subject)
	t := v.Type()
	k := t.Kind()

	return v, t, k
}

func writef(message string, args ...interface{}) {
	GinkgoWriter.Write([]byte(fmt.Sprintf(message + "\n", args...)))
}

func RecursiveMapCompare(expect interface{}, actual interface{}) {
	expectValue, expectType, expectKind := getReflectionInfo(expect)
	actualValue, actualType, actualKind := getReflectionInfo(actual)

	writef("Asserting that expected and actual kinds match: %v ?= %v", expectKind, actualKind)
	Expect(expectKind).To(Equal(actualKind))


	switch expectKind {
	case reflect.Slice:
		writef("Found a slice, assuming this is the end")
		Expect(expect).Should(ConsistOf(actual))
	case reflect.Map:
		writef("Found a map")

		_true := reflect.ValueOf(true)
		_false := reflect.ValueOf(false)

		a := reflect.MakeMap(reflect.MapOf(actualType.Key(), _true.Type()))

		for _, key := range actualValue.MapKeys() {
			a.SetMapIndex(key, _false)
		}

		zero := reflect.Zero(expectType.Elem())
		writef("Checking against each entry")
		for _, key := range expectValue.MapKeys() {
			a.SetMapIndex(key, _true)

			writef("Recursively checking that values match at '%v'", key)
			expectEntry := expectValue.MapIndex(key)
			actualEntry := actualValue.MapIndex(key)

			Expect(actualEntry).NotTo(Equal(zero), fmt.Sprintf("Missing entry in actual map at '%v'\n", key))

			RecursiveMapCompare(expectEntry.Interface(), actualEntry.Interface())
		}

		prettyA := reflect.MakeSlice(reflect.SliceOf(actualType.Key()), 0, 0)
		for _, key := range a.MapKeys() {
			if !a.MapIndex(key).Bool() {
				prettyA.Set(key)
			}
		}

		Expect(prettyA.Len()).To(Equal(0), fmt.Sprintf("Found extra keys in actual map: '%q'", prettyA))
	default:
		writef("It's not a slice or a map, offloading to Gomega")
		Expect(expect).To(Equal(actual))
	}
}
