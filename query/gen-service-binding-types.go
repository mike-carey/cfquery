// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package query

import "github.com/cloudfoundry-community/go-cfclient"

type ServiceBindings []cfclient.ServiceBinding
type ServiceBindingMap map[string]cfclient.ServiceBinding
type ServiceBindingGroup map[string]ServiceBindings
type MappedServiceBindingMap map[string]ServiceBindingMap
type MappedServiceBindingGroup map[string]ServiceBindingGroup

// func (i *ServiceBindings) ToJson() {
// 	enc := json.NewEncoder(os.Stdin)
// 	enc.SetIndent("", "    ")
// 	if err := enc.Encode(&a); err != nil {
// 		panic(err)
// 	}
