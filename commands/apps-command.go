package commands

func (c *AppsCommand) GroupByOptions() []string {
	return []string{
		// "stack",
		"space",
		"org",
		"space-and-org",
	}
}
