package query

import (
	"errors"

	"github.com/mike-carey/cfquery/logger"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

func ItemFilterBy(items Items, check func(Item) (bool, error)) (Items, error) {
	pool := make(Items, 0)

	for _, item := range items {
		ok, err := check(item)
		if err != nil {
			logger.Errorf("Could not check filter on: %v", item)
			return nil, errors.Wrap(err, "Could not check filter for item")
		}

		if ok {
			logger.Infof("Adding item to entry")
			pool = append(pool, item)
		}
	}

	logger.Infof("Returning filtered slice of size %d", len(pool))
	return pool, nil
}

func ItemFilterMapBy(items ItemMap, check func(Item) (bool, error)) (ItemMap, error) {
	pool := make(ItemMap, 0)

	for key, item := range items {
		ok, err := check(item)
		if err != nil {
			logger.Errorf("Could not check filter on: %v", item)
			return nil, errors.Wrap(err, "Could not check filter for item")
		}

		if ok {
			logger.Infof("Adding %s to entry", key)
			pool[key] = item
		}
	}

	logger.Infof("Returning filtered map of size %d", len(pool))
	return pool, nil
}
