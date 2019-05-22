// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package generics

import (
	"github.com/mike-carey/cfquery/logger"
	"github.com/pkg/errors"
)

func FooGroupBy(items []Foo, getKey func(Foo) (string, error)) (map[string][]Foo, error) {
	pool := make(map[string][]Foo, 0)

	for _, item := range items {
		key, err := getKey(item)
		if err != nil {
			logger.Errorf("Could not get key from item: %v", item)
			return nil, errors.Wrap(err, "Could not get key for item")
		}

		if _, ok := pool[key]; !ok {
			pool[key] = make([]Foo, 0)
		}

		logger.Infof("Adding item to %s entry", key)
		pool[key] = append(pool[key], item)
	}

	logger.Infof("Returning %d two groups", len(pool))
	return pool, nil
}