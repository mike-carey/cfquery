package util

import (
	"errors"

	"github.com/mike-carey/cfquery/logger"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

func ItemFilterBy(items []Item, check func(Item) (bool, error)) ([]Item, error) {
	pool := make([]Item, 0)

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
