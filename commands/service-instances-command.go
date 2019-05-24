package commands

import ()

type ServiceInstancesCommand struct {
	BaseCommand
}

func (c *ServiceInstancesCommand) Execute() error {
	return nil
}

func (c *ServiceInstancesCommand) TargetOptions() []string {
	return []string{}
}

func (c *ServiceInstancesCommand) GroupByOptions() []string {
	return []string{
		"stack",
		"space",
		"org"
	}
}

func (c *ServiceInstancesCommand) SortByOptions() []string {
	return []string{}
}

func (c *ServiceInstancesCommand) SortBy(string name) {

}

func (c *ServiceInstancesCommand) GroupBy(string ...name) {

}

func (c *ServiceInstancesCommand) Target(string name) {

}
