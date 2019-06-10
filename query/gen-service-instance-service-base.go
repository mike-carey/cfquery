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

type ServiceInstanceService struct {
	Client      cf.CFClient
	storage     ServiceInstanceMap
	filled      bool
	mutex       *sync.Mutex
	serviceName string
	key         string
}

func NewServiceInstanceService(client cf.CFClient) *ServiceInstanceService {
	return &ServiceInstanceService{
		Client:  client,
		storage: make(map[string]cfclient.ServiceInstance, 0),
		filled:  false,
		mutex:   &sync.Mutex{},
	}
}

func (s *ServiceInstanceService) ServiceName() string {
	if s.serviceName == "" {
		name := fmt.Sprintf("%T", s)

		_name := strings.Split(name, ".")
		name = _name[len(_name)-1]

		s.serviceName = fmt.Sprintf("%s", name)
	}

	return s.serviceName
}

func (s *ServiceInstanceService) Key() string {
	if s.key == "" {
		key := s.ServiceName()
		s.key = key[:len(key)-len("Service")]
		logger.Info(s.key)
	}

	return s.key
}

func (s *ServiceInstanceService) lock() {
	s.mutex.Lock()
	logger.Infof("Locked %v", reflect.TypeOf(s))
}

func (s *ServiceInstanceService) unlock() {
	s.mutex.Unlock()
	logger.Infof("Unlocked %v", reflect.TypeOf(s))
}

func (s *ServiceInstanceService) GetStorage() (ServiceInstanceMap, error) {
	_, err := s.GetAll()
	if err != nil {
		return nil, err
	}

	return s.storage, nil
}

func (s *ServiceInstanceService) GetByGuid(guid string) (*cfclient.ServiceInstance, error) {
	s.lock()

	defer s.unlock()

	if s.filled {
		if val, ok := s.storage[guid]; ok {
			return &val, nil
		}
	}

	logger.Infof("Did not find %s in storage", guid)
	item, err := s.Client.GetServiceInstanceByGuid(guid)
	if err != nil {
		return nil, err
	}

	// Save off in storage
	s.storage[guid] = item

	return &item, nil
}

func (s *ServiceInstanceService) GetManyByGuid(guids ...string) (ServiceInstanceMap, error) {
	pool := make(ServiceInstanceMap, len(guids))

	type Result struct {
		Guid   string
		Object *cfclient.ServiceInstance
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

func (s *ServiceInstanceService) GetAll() (ServiceInstances, error) {
	s.lock()

	if s.filled {
		logger.Infof("Reusing storage")
		siSlice := make(ServiceInstances, 0, len(s.storage))
		for _, si := range s.storage {
			siSlice = append(siSlice, si)
		}

		s.unlock()

		return siSlice, nil
	}

	logger.Infof("Calling out to CFClient")
	sis, err := s.Client.ListServiceInstances()
	if err != nil {
		return nil, err
	}

	go func(s *ServiceInstanceService, sis ServiceInstances) {
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

func (i *Inquistor) NewServiceInstanceService() *ServiceInstanceService {
	return NewServiceInstanceService(i.CFClient)
}

func (i *Inquistor) GetServiceInstanceService() *ServiceInstanceService {
	class := &ServiceInstanceService{}
	service := i.GetService(class.Key())

	return service.(*ServiceInstanceService)
}