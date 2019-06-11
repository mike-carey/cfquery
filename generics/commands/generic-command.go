package commands

import (
	"fmt"
	"reflect"

	"github.com/mike-carey/cfquery/query"
	"github.com/iancoleman/strcase"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type ItemsCommand struct {
	CommandDefaults
}

func (c *ItemsCommand) Execute([]string) error {
	return workerFactory.Go(c)
}

func (c *ItemsCommand) Run(o *Options, i *query.Inquisitor) (interface{}, error) {
	service := i.GetItemService()

	a, err := service.GetAll()
	if err != nil {
		return nil, err
	}

	val := reflect.ValueOf(a)

	if gb := o.GroupBy; gb != "" {
		fn := val.MethodByName(fmt.Sprintf("GroupBy%s", strcase.ToCamel(gb)))

		if fn.Kind() != reflect.Func {
			panic(fmt.Sprintf("Missing method GroupBy%s", strcase.ToCamel(gb)))
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
