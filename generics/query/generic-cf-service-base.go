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

type Item generic.Type

type ItemService struct {
	Client  cf.CFClient
	storage ItemMap
	filled  bool
	mutex   *sync.Mutex
	serviceName string
	key string
}

func NewItemService(client cf.CFClient) *ItemService {
	return &ItemService{
		Client:  client,
		storage: make(map[string]Item, 0),
		filled:  false,
		mutex:   &sync.Mutex{},
	}
}

func (s *ItemService) ServiceName() string {
	if s.serviceName == "" {
		name := fmt.Sprintf("%T", s)

		_name := strings.Split(name, ".")
		name = _name[len(_name)-1]

		s.serviceName = fmt.Sprintf("%s", name)
	}

	return s.serviceName
}

func (s *ItemService) Key() string {
	if s.key == "" {
		key := s.ServiceName()
		s.key = key[:len(key)-len("Service")]
		logger.Info(s.key)
	}

	return s.key
}

func (s *ItemService) lock() {
	s.mutex.Lock()
	logger.Infof("Locked %v", reflect.TypeOf(s))
}

func (s *ItemService) unlock() {
	s.mutex.Unlock()
	logger.Infof("Unlocked %v", reflect.TypeOf(s))
}

func (s *ItemService) GetStorage() (ItemMap, error) {
	_, err := s.GetAllItems()
	if err != nil {
		return nil, err
	}

	return s.storage, nil
}

func (s *ItemService) GetItemByGuid(guid string) (*Item, error) {
	s.lock()

	defer s.unlock()

	if s.filled {
		if val, ok := s.storage[guid]; ok {
			return &val, nil
		}
	}

	logger.Infof("Did not find %s in storage", guid)
	item, err := s.Client.GetItemByGuid(guid)
	if err != nil {
		return nil, err
	}

	// Save off in storage
	s.storage[guid] = item

	return &item, nil
}

func (s *ItemService) GetManyItemsByGuid(guids ...string) (ItemMap, error) {
	pool := make(ItemMap, len(guids))

	type Result struct {
		Guid   string
		Object *Item
		Error  error
	}

	resCh := make(chan Result, len(guids))

	for _, guid := range guids {
		go func(guid string) {
			obj, err := s.GetItemByGuid(guid)
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

func (s *ItemService) GetAllItems() (Items, error) {
	s.lock()

	if s.filled {
		logger.Infof("Reusing storage")
		siSlice := make(Items, 0, len(s.storage))
		for _, si := range s.storage {
			siSlice = append(siSlice, si)
		}

		s.unlock()

		return siSlice, nil
	}

	logger.Infof("Calling out to CFClient")
	sis, err := s.Client.ListItems()
	if err != nil {
		return nil, err
	}

	go func(s *ItemService, sis Items) {
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

func (i *inquisitor) NewItemService() *ItemService {
	return NewItemService(i.CFClient)
}

func (i *inquisitor) GetItemService() *ItemService {
	class := &ItemService{}
	service := i.GetService(class.Key())

	return service.(*ItemService)
}
