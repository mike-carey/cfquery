package util

import (
	"errors"

	"github.com/mike-carey/cfquery/logger"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

func ItemGroupBy(items []Item, getKey func(Item) (string, error)) (map[string][]Item, error) {
	pool := make(map[string][]Item, 0)

	for _, item := range items {
		key, err := getKey(item)
		if err != nil {
			logger.Errorf("Could not get key from item: %v", item)
			return nil, errors.Wrap(err, "Could not get key for item")
		}

		if _, ok := pool[key]; !ok {
			pool[key] = make([]Item, 0)
		}

		logger.Infof("Adding item to %s entry", key)
		pool[key] = append(pool[key], item)
	}

	logger.Infof("Returning %d two groups", len(pool))
	return pool, nil
}
