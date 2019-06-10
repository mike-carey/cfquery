package query

import (
	"errors"
	"reflect"
	"sync"
	"fmt"

	"github.com/cheekybits/genny/generic"

	"github.com/mike-carey/cfquery/cf"
	"github.com/mike-carey/cfquery/logger"
	"github.com/mike-carey/cfquery/util"
)

type CFObject generic.Type

type CFObjectService struct {
	Client  cf.CFClient
	storage CFObjectMap
	filled  bool
	mutex   *sync.Mutex
	serviceName string
	key string
}

func NewCFObjectService(client cf.CFClient) *CFObjectService {
	return &CFObjectService{
		Client:  client,
		storage: make(map[string]CFObject, 0),
		filled:  false,
		mutex:   &sync.Mutex{},
	}
}

func (s *CFObjectService) ServiceName() string {
	if s.serviceName == "" {
		name := fmt.Sprintf("%T", s)

		_name := strings.Split(name, ".")
		name = _name[len(_name)-1]

		s.serviceName = fmt.Sprintf("%s", name)
	}

	return s.serviceName
}

func (s *CFObjectService) Key() string {
	if s.key == "" {
		key := s.ServiceName()
		s.key = key[:len(key)-len("Service")]
		logger.Info(s.key)
	}

	return s.key
}

func (s *CFObjectService) lock() {
	s.mutex.Lock()
	logger.Infof("Locked %v", reflect.TypeOf(s))
}

func (s *CFObjectService) unlock() {
	s.mutex.Unlock()
	logger.Infof("Unlocked %v", reflect.TypeOf(s))
}

func (s *CFObjectService) GetStorage() (CFObjectMap, error) {
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

func (s *CFObjectService) GetManyByGuid(guids ...string) (CFObjectMap, error) {
	pool := make(CFObjectMap, len(guids))

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
		return nil, util.StackErrors(errs)
	}

	return pool, nil
}

func (s *CFObjectService) GetAll() (CFObjects, error) {
	s.lock()

	if s.filled {
		logger.Infof("Reusing storage")
		siSlice := make(CFObjects, 0, len(s.storage))
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

	go func(s *CFObjectService, sis CFObjects) {
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

func (i *Inquistor) NewCFObjectService() *CFObjectService {
	return NewCFObjectService(i.CFClient)
}

func (i *Inquistor) GetCFObjectService() *CFObjectService {
	class := &CFObjectService{}
	service := i.GetService(class.Key())

	return service.(*CFObjectService)
}
