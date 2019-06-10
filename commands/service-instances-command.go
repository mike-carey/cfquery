package commands

import (
	"github.com/mike-carey/cfquery/query"
)

type ServiceInstancesCommand struct {
	CommandDefaults
}

func (c *ServiceInstancesCommand) Execute([]string) error {
	w, e := workerFactory.NewWorker(c)
	if e != nil {
		return e
	}

	return w.Work()
}

func (c *ServiceInstancesCommand) Run(o *Options, i *query.Inquistor) (interface{}, error) {
	service := i.GetServiceInstanceService()
	//
	// if o.Target == "apps" {
	// 	_, err := service.GetApps(i)
	// 	if err != nil {
	// 		return err
	// 	}
	// } else {
		sis, err := service.GetAll()
		if err != nil {
			return nil, err
		}
	// }

	return sis, nil
}

func (c *ServiceInstancesCommand) TargetOptions() []string {
	return []string{
		"apps",
	}
}

func (c *ServiceInstancesCommand) GroupByOptions() []string {
	return []string{
		"space",
		"org",
		"space-and-org",
	}
}
