package formatter_test

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/ghodss/yaml"
)

func EqualJson(expected interface{}) *equalEncodingMatcher {
	return &equalEncodingMatcher{
		format: "json",
		check: AreEqualJSON,
		expected: expected,
	}
}

func EqualYaml(expected interface{}) *equalEncodingMatcher {
	return &equalEncodingMatcher{
		format: "yaml",
		check: AreEqualYAML,
		expected: expected,
	}
}

type equalEncodingMatcher struct {
	format string
	check func(s1 string, s2 string) (bool, error)
	expected interface{}
}

func (matcher *equalEncodingMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nto be equivalently encoded as %s to \n\t%#v", actual, matcher.format, matcher.expected)
}

func (matcher *equalEncodingMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nnot to be equivalently encoded as %s to \n\t%#v", actual, matcher.format, matcher.expected)
}

func (matcher *equalEncodingMatcher) Match(actual interface{}) (success bool, err error) {
	e, ok := matcher.expected.(string)
	if !ok {
		return false, fmt.Errorf("Expected provided is not string :: '%#v'", matcher.expected)
	}

	a, ok := actual.(string)
	if !ok {
		return false, fmt.Errorf("Actual provided is not string :: '%#v'", matcher.expected)
	}

	return matcher.check(e, a)
}

func AreEqualJSON(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string into JSON :: %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string into JSON :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}

func AreEqualYAML(s1, s2 string) (bool, error) {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = yaml.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string into YAML :: %s", err.Error())
	}
	err = yaml.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string into YAML :: %s", err.Error())
	}

	return reflect.DeepEqual(o1, o2), nil
}
