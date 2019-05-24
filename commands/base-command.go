package commands

type Command interface {
	Execute() error
	TargetOptions() []string
	GroupByOptions() []string
	SortByOptions() []string

	SortBy(name string)
	GroupBy(name ...string)
	Target(name string)
}

type BaseCommand struct {
	Inquistor *Inquistor
}
