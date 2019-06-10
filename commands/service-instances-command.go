package commands

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
