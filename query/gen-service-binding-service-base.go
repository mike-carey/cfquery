// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package query

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/mike-carey/cfquery/cf"
	"github.com/mike-carey/cfquery/logger"
	"github.com/mike-carey/cfquery/util"
)

type ServiceBindingService struct {
	Client      cf.CFClient
	storage     ServiceBindingMap
	filled      bool
	mutex       *sync.Mutex
	serviceName string
	key         string
}

func NewServiceBindingService(client cf.CFClient) *ServiceBindingService {
	return &ServiceBindingService{
		Client:  client,
		storage: make(map[string]cfclient.ServiceBinding, 0),
		filled:  false,
		mutex:   &sync.Mutex{},
	}
}

func (s *ServiceBindingService) ServiceName() string {
	if s.serviceName == "" {
		name := fmt.Sprintf("%T", s)

		_name := strings.Split(name, ".")
		name = _name[len(_name)-1]

		s.serviceName = fmt.Sprintf("%s", name)
	}

	return s.serviceName
}

func (s *ServiceBindingService) Key() string {
	if s.key == "" {
		key := s.ServiceName()
		s.key = key[:len(key)-len("Service")]
		logger.Info(s.key)
	}

	return s.key
}

func (s *ServiceBindingService) lock() {
	s.mutex.Lock()
	logger.Infof("Locked %v", reflect.TypeOf(s))
}

func (s *ServiceBindingService) unlock() {
	s.mutex.Unlock()
	logger.Infof("Unlocked %v", reflect.TypeOf(s))
}

func (s *ServiceBindingService) GetStorage() (ServiceBindingMap, error) {
	_, err := s.GetAllServiceBindings()
	if err != nil {
		return nil, err
	}

	return s.storage, nil
}

func (s *ServiceBindingService) GetServiceBindingByGuid(guid string) (*cfclient.ServiceBinding, error) {
	s.lock()

	defer s.unlock()

	if s.filled {
		if val, ok := s.storage[guid]; ok {
			return &val, nil
		}
	}

	logger.Infof("Did not find %s in storage", guid)
	item, err := s.Client.GetServiceBindingByGuid(guid)
	if err != nil {
		return nil, err
	}

	// Save off in storage
	s.storage[guid] = item

	return &item, nil
}

func (s *ServiceBindingService) GetManyServiceBindingsByGuid(guids ...string) (ServiceBindingMap, error) {
	pool := make(ServiceBindingMap, len(guids))

	type Result struct {
		Guid   string
		Object *cfclient.ServiceBinding
		Error  error
	}

	resCh := make(chan Result, len(guids))

	for _, guid := range guids {
		go func(guid string) {
			obj, err := s.GetServiceBindingByGuid(guid)
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

func (s *ServiceBindingService) GetAllServiceBindings() (ServiceBindings, error) {
	s.lock()

	if s.filled {
		logger.Infof("Reusing storage")
		siSlice := make(ServiceBindings, 0, len(s.storage))
		for _, si := range s.storage {
			siSlice = append(siSlice, si)
		}

		s.unlock()

		return siSlice, nil
	}

	logger.Infof("Calling out to CFClient")
	sis, err := s.Client.ListServiceBindings()
	if err != nil {
		return nil, err
	}

	go func(s *ServiceBindingService, sis ServiceBindings) {
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

func (i *Inquisitor) NewServiceBindingService() *ServiceBindingService {
	return NewServiceBindingService(i.CFClient)
}

func (i *Inquisitor) GetServiceBindingService() *ServiceBindingService {
	class := &ServiceBindingService{}
	service := i.GetService(class.Key())

	return service.(*ServiceBindingService)
}
