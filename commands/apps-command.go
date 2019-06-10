package commands

import (
	"fmt"
	"reflect"

	"github.com/mike-carey/cfquery/query"
	"github.com/iancoleman/strcase"
)

type AppsCommand struct {
	CommandDefaults
}

func (c *AppsCommand) Execute([]string) error {
	w, e := workerFactory.NewWorker(c)
	if e != nil {
		return e
	}

	return w.Work()
}

func (c *AppsCommand) Run(o *Options, i *query.Inquistor) (interface{}, error) {
	service := i.GetAppService()

	apps, err := service.GetAll()
	if err != nil {
		return nil, err
	}

	val := reflect.ValueOf(apps)

	if gb := o.GroupBy; gb != "" {
		fn := val.MethodByName(fmt.Sprintf("GroupBy%s", strcase.ToCamel(gb)))

		if fn == reflect.Zero(reflect.TypeOf(fn)) {
			panic(fmt.Sprintf("Missing method GrouBy%s", strcase.ToCamel(gb)))
		}

		res := fn.Call([]reflect.Value{
			reflect.ValueOf(i),
		})

		if e := res[1].Interface(); e != nil {
			return nil, e.(error)
		}

		val = res[0]
	}

	return val.Interface(), nil
}

func (c *AppsCommand) GroupByOptions() []string {
	return []string{
		// "stack",
		"space",
		"org",
		"space-and-org",
	}
}
