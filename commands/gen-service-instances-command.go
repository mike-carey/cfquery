// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package commands

import (
	"fmt"
	"reflect"

	"github.com/iancoleman/strcase"
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