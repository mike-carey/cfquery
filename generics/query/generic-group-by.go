package query

import (
	"errors"

	"github.com/mike-carey/cfquery/logger"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

func ItemGroupBy(items Items, getKey func(Item) (string, error)) (ItemGroup, error) {
	pool := make(ItemGroup, 0)

	for _, item := range items {
		key, err := getKey(item)
		if err != nil {
			logger.Errorf("Could not get key from item: %v", item)
			return nil, errors.Wrap(err, "Could not get key for item")
		}

		if _, ok := pool[key]; !ok {
			pool[key] = make(Items, 0)
		}

		logger.Infof("Adding item to %s entry", key)
		pool[key] = append(pool[key], item)
	}

	logger.Infof("Returning %d groups in slice", len(pool))
	return pool, nil
}

func ItemGroupMapBy(items ItemMap, getKey func(string, Item) (string, error)) (MappedItemMap, error) {
	pool := make(MappedItemMap, 0)

	for origKey, item := range items {
		key, err := getKey(origKey, item)
		if err != nil {
			logger.Errorf("Could not get key from item: %v", item)
			return nil, errors.Wrap(err, "Could not get key for item")
		}

		if _, ok := pool[key]; !ok {
			pool[key] = make(ItemMap, 0)
		}

		logger.Infof("Adding %s item to %s entry", origKey, key)
		pool[key][origKey] = item
	}

	logger.Infof("Returning %d groups in map", len(pool))
	return pool, nil
}

func ItemGroupMappedSliceBy(items ItemGroup, getKey func(string, Items) (string, error)) (MappedItemGroup, error) {
	pool := make(MappedItemGroup, 0)

	for origKey, item := range items {
		key, err := getKey(origKey, item)
		if err != nil {
			logger.Errorf("Could not get key from item: %v", item)
			return nil, errors.Wrap(err, "Could not get key for item")
		}

		if _, ok := pool[key]; !ok {
			pool[key] = make(ItemGroup, 0)
		}

		logger.Infof("Adding %s item to %s entry", origKey, key)
		pool[key][origKey] = item
	}

	logger.Infof("Returning %d groups in map", len(pool))
	return pool, nil
}
