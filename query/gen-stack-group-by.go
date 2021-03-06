// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package query

import (
	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/mike-carey/cfquery/logger"
	"github.com/pkg/errors"
)

func StackGroupBy(items Stacks, getKey func(cfclient.Stack) (string, error)) (StackGroup, error) {
	pool := make(StackGroup, 0)

	for _, item := range items {
		key, err := getKey(item)
		if err != nil {
			logger.Errorf("Could not get key from item: %v", item)
			return nil, errors.Wrap(err, "Could not get key for item")
		}

		if _, ok := pool[key]; !ok {
			pool[key] = make(Stacks, 0)
		}

		logger.Infof("Adding item to %s entry", key)
		pool[key] = append(pool[key], item)
	}

	logger.Infof("Returning %d groups in slice", len(pool))
	return pool, nil
}

func StackGroupMapBy(items StackMap, getKey func(string, cfclient.Stack) (string, error)) (MappedStackMap, error) {
	pool := make(MappedStackMap, 0)

	for origKey, item := range items {
		key, err := getKey(origKey, item)
		if err != nil {
			logger.Errorf("Could not get key from item: %v", item)
			return nil, errors.Wrap(err, "Could not get key for item")
		}

		if _, ok := pool[key]; !ok {
			pool[key] = make(StackMap, 0)
		}

		logger.Infof("Adding %s item to %s entry", origKey, key)
		pool[key][origKey] = item
	}

	logger.Infof("Returning %d groups in map", len(pool))
	return pool, nil
}

func StackGroupMappedSliceBy(items StackGroup, getKey func(string, Stacks) (string, error)) (MappedStackGroup, error) {
	pool := make(MappedStackGroup, 0)

	for origKey, item := range items {
		key, err := getKey(origKey, item)
		if err != nil {
			logger.Errorf("Could not get key from item: %v", item)
			return nil, errors.Wrap(err, "Could not get key for item")
		}

		if _, ok := pool[key]; !ok {
			pool[key] = make(StackGroup, 0)
		}

		logger.Infof("Adding %s item to %s entry", origKey, key)
		pool[key][origKey] = item
	}

	logger.Infof("Returning %d groups in map", len(pool))
	return pool, nil
}
