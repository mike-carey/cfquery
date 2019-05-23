package query

import (
	"errors"
	"reflect"
	"sync"

	"github.com/cheekybits/genny/generic"

	"github.com/mike-carey/cfquery/cf"
	"github.com/mike-carey/cfquery/logger"
	"github.com/mike-carey/cfquery/util"
)

type CFObject generic.Type

type CFObjectService struct {
	Client  cf.CFClient
	storage map[string]CFObject
	filled  bool
	mutex   *sync.Mutex
}

func NewCFObjectService(client cf.CFClient) *CFObjectService {
	return &CFObjectService{
		Client:  client,
		storage: make(map[string]CFObject, 0),
		filled:  false,
		mutex:   &sync.Mutex{},
	}
}

func (s *CFObjectService) lock() {
	s.mutex.Lock()
	logger.Infof("Locked %v", reflect.TypeOf(s))
}

func (s *CFObjectService) unlock() {
	s.mutex.Unlock()
	logger.Infof("Unlocked %v", reflect.TypeOf(s))
}

func (s *CFObjectService) GetStorage() (map[string]CFObject, error) {
	_, err := s.GetAll()
	if err != nil {
		return nil, err
	}

	return s.storage, nil
}

func (s *CFObjectService) GetByGuid(guid string) (*CFObject, error) {
	s.lock()

	defer s.unlock()

	if s.filled {
		if val, ok := s.storage[guid]; ok {
			return &val, nil
		}
	}

	logger.Infof("Did not find %s in storage", guid)
	item, err := s.Client.GetCFObjectByGuid(guid)
	if err != nil {
		return nil, err
	}

	// Save off in storage
	s.storage[guid] = item

	return &item, nil
}

func (s *CFObjectService) GetManyByGuid(guids ...string) (map[string]CFObject, error) {
	pool := make(map[string]CFObject, len(guids))

	type Result struct {
		Guid   string
		Object *CFObject
		Error  error
	}

	resCh := make(chan Result, len(guids))

	for _, guid := range guids {
		go func(guid string) {
			obj, err := s.GetByGuid(guid)
			res := Result{
				Guid:   guid,
				Error:  err,
				Object: obj,
			}

			resCh <- res
		}(guid)
	}

	errs := make([]error, 0)

	for _, _ = range guids {
		select {
		case res := <-resCh:
			if res.Error != nil {
				errs = append(errs, res.Error)
			}

			pool[res.Guid] = *res.Object
		}
	}

	if len(errs) > 0 {
		return nil, StackErrors(errs)
	}

	return pool, nil
}

func (s *CFObjectService) GetAll() ([]CFObject, error) {
	s.lock()

	if s.filled {
		logger.Infof("Reusing storage")
		siSlice := make([]CFObject, 0, len(s.storage))
		for _, si := range s.storage {
			siSlice = append(siSlice, si)
		}

		s.unlock()

		return siSlice, nil
	}

	logger.Infof("Calling out to CFClient")
	sis, err := s.Client.ListCFObjects()
	if err != nil {
		return nil, err
	}

	go func(s *CFObjectService, sis []CFObject) {
		logger.Infof("Storing contents to storage")
		for _, si := range sis {
			s.storage[si.Guid] = si
		}

		logger.Infof("Done storing contents to storage")
		s.filled = true

		s.unlock()
	}(s, sis)

	logger.Infof("Returning list while populating happens")
	return sis, nil
}

func (i *Inquistor) GetCFObjectService() *CFObjectService {
	if i.CFObjectService == nil {
		i.CFObjectService = NewCFObjectService(i.CFClient)
	}

	return i.CFObjectService
}
