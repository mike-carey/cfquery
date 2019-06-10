package commands

import (
	"github.com/mike-carey/cfquery/config"
	"github.com/mike-carey/cfquery/query"
)

type Commands struct {
	Apps AppsCommand `command:"apps"`
	Spaces SpacesCommand `command:"spaces"`
	ServiceInstances ServiceInstancesCommand `command:"service-instances"`
}

type GlobalOptions struct {
	Verbose bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	Config string `short:"c" long:"config" description:"The config file to load" default:"cfquery.json"`
	Format string `short:"f" long:"format" description:"The format of printing the objects" default:"table"`
}

type Options struct {
	*GlobalOptions
	Foundations config.Foundations

	GroupBy string `short:"g" long:"group-by" description:"Groups entries on this property"`
	SortBy string `short:"s" long:"sort-by" description:"Sorts the entries by this property"`
	Target string `short:"t" long:"target" description:"Changes what the main object to print is"`
}

type Command interface {
	TargetOptions() []string
	GroupByOptions() []string
	SortByOptions() []string

	Execute([]string) error

	Run(*Options, *query.Inquistor) (interface{}, error)
}

type CommandDefaults struct {}

// TargetOptions ...
func (c *CommandDefaults) TargetOptions() []string {
	return []string{}
}

// GroupByOptions ...
func (c *CommandDefaults) GroupByOptions() []string {
	return []string{}
}

// SortByOptions ...
func (c *CommandDefaults) SortByOptions() []string {
	return []string{}
}
