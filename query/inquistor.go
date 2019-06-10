package query

import (
	"fmt"
	"reflect"

	"github.com/mike-carey/cfquery/cf"

	"github.com/iancoleman/strcase"
)

type Service interface {
	ServiceName() string
}

type Inquistor struct {
	CFClient cf.CFClient
	Services map[string]Service
}

func NewInquistor(cfClient cf.CFClient) *Inquistor {
	return &Inquistor{
		CFClient: cfClient,
		Services: make(map[string]Service, 0),
	}
}

func (i *Inquistor) GetService(name string) Service {
	if service, ok := i.Services[name]; ok {
		return service
	}

	className := strcase.ToCamel(name)
	serviceName := fmt.Sprintf("%sService", className)
	funcName := fmt.Sprintf("New%s", serviceName)

	fnZeroValue := reflect.ValueOf(nil)

	fn := reflect.ValueOf(i).MethodByName(funcName)
	if fn == fnZeroValue {
		panic(fmt.Sprintf("Unknown service %s", name))
	}

	serviceValue := fn.Call([]reflect.Value{})[0]

	if serviceValue == fnZeroValue {
		panic(fmt.Sprintf("Error creating service %s", serviceName))
	}

	service := serviceValue.Interface().(Service)

	i.Services[name] = service

	return service
}
