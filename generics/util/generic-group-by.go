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

	logger.Infof("Returning %d groups in slice", len(pool))
	return pool, nil
}

func ItemGroupMapBy(items map[string]Item, getKey func(Item) (string, error)) (map[string]map[string]Item, error) {
	pool := make(map[string]map[string]Item, 0)

	for origKey, item := range items {
		key, err := getKey(item)
		if err != nil {
			logger.Errorf("Could not get key from item: %v", item)
			return nil, errors.Wrap(err, "Could not get key for item")
		}

		if _, ok := pool[key]; !ok {
			pool[key] = make(map[string]Item, 0)
		}

		logger.Infof("Adding %s item to %s entry", origKey, key)
		pool[key][origKey] = item
	}

	logger.Infof("Returning %d groups in map", len(pool))
	return pool, nil
}
