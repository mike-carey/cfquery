package query

import (
	"errors"
	"fmt"
	"sync"

	"github.com/mike-carey/cfquery/cf"
	"github.com/mike-carey/cfquery/logger"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

type Inquistor struct {
	CFClient               cf.CFClient
	AppService             *AppService
	ServiceBindingService  *ServiceBindingService
	ServiceInstanceService *ServiceInstanceService
	SpaceService           *SpaceService
	OrgService             *OrgService
}

func NewInquistor(cfClient cf.CFClient) *Inquistor {
	return &Inquistor{
		CFClient: cfClient,
	}
}

func StackErrors(errs []error) error {
	s := ""
	if len(errs) > 0 {
		s = "s"
	}

	err := fmt.Sprintf("%d error%s occured:\n", len(errs), s)
	for _, e := range errs {
		err += fmt.Sprintf("%q\n", e)
	}

	return errors.New(err)
}

func (i *Inquistor) GetServiceInstanceToAppRatio() (map[string][]cfclient.App, error) {
	var serviceBindings []cfclient.ServiceBinding
	// var serviceInstances []cfclient.ServiceInstance
	logger.Info("Getting service instance to app ratio")

	var wg sync.WaitGroup

	errs := make([]error, 0)

	errCh := make(chan error, 0)
	sbsCh := make(chan []cfclient.ServiceBinding)
	sisCh := make(chan []cfclient.ServiceInstance)

	wg.Add(2)

	go func() {
		defer wg.Done()

		logger.Info("Getting all service bindings")
		sbs, err := i.GetServiceBindingService().GetAll()
		if err != nil {
			errCh <- err
			return
		}

		logger.Infof("Received %d service-bindings", len(sbs))
		sbsCh <- sbs
	}()
	go func() {
		defer wg.Done()

		logger.Info("Getting all service instances")
		sis, err := i.GetServiceInstanceService().GetAll()
		if err != nil {
			errCh <- err
			return
		}

		logger.Infof("Received %d service-instances", len(sis))
		sisCh <- sis
	}()

	for j := 0; j < 2; j++ {
		select {
		case sbs := <-sbsCh:
			logger.Info("Populating serviceBindings array")
			serviceBindings = sbs
		case _ = <-sisCh:
			logger.Info("Populating serviceInstances array")
			// serviceInstances = sis
		case err := <-errCh:
			logger.Info("Populating error array")
			errs = append(errs, err)
		}
	}

	logger.Info("Waiting for service*s")
	wg.Wait()
	logger.Info("Done waiting for service*s")

	if len(errs) > 0 {
		return nil, StackErrors(errs)
	}

	type Result struct {
		ServiceInstance *cfclient.ServiceInstance
		App             *cfclient.App
	}

	resCh := make(chan Result, 0)

	wg.Add(len(serviceBindings))

	for _, sb := range serviceBindings {
		go func(sb cfclient.ServiceBinding) {
			logger.Infof("Grabbing app by guid: %s", sb.AppGuid)
			app, err := i.GetAppService().GetByGuid(sb.AppGuid)
			if err != nil {
				errCh <- err
				return
			}

			logger.Infof("Grabbing service instance by guid: %s", sb.ServiceInstanceGuid)
			si, err := i.GetServiceInstanceService().GetByGuid(sb.ServiceInstanceGuid)
			if err != nil {
				errCh <- err
				return
			}

			resCh <- Result{
				ServiceInstance: si,
				App:             app,
			}
			wg.Done()
		}(sb)
	}

	ret := make(map[string][]cfclient.App, 0)

	for _, _ = range serviceBindings {
		select {
		case res := <-resCh:
			if _, ok := ret[res.ServiceInstance.Guid]; !ok {
				ret[res.ServiceInstance.Guid] = make([]cfclient.App, 0)
			}

			ret[res.ServiceInstance.Guid] = append(ret[res.ServiceInstance.Guid], *res.App)
		case err := <-errCh:
			errs = append(errs, err)
		}
	}

	logger.Info("Waiting for mapping")
	wg.Wait()
	logger.Info("Done waiting for mapping")

	if len(errs) > 0 {
		return nil, StackErrors(errs)
	}

	return ret, nil
}
