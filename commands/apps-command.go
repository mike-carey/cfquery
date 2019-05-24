package commands

import (
	"github.com/mike-carey/cfquery/query"
	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

type AppsCommand struct {
	BaseCommand

	apps []cfclient.App
}

func NewAppsCommand(inquisitor *query.Inquistor) *AppsCommand {
	return &AppsCommand{
		Inquistor: inquisitor
	}
}

func (c *AppsCommand) Execute() error {
	return nil
}

func (c *AppsCommand) TargetOptions() []string {
	return []string{}
}

func (c *AppsCommand) GroupByOptions() []string {
	return []string{
		"stack",
		"space",
		"org",
	}
}

func (c *AppsCommand) SortByOptions() []string {
	return []string{}
}

func (c *AppsCommand) SortBy(name string) {
	panic("Sort By Not supported")
}

func (c *AppsCommand) GroupBy(name ...string) {

}

func (c *AppsCommand) Target(name string) error {
	c.apps, err := inquisitor.GetAppService().GetAll()
	if err != nil {
		return err
	}
}
