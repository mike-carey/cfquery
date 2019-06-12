//

package query

// Inquisitor ...
type Inquisitor interface {
	GetService(name string) Service
	NewAppService() *AppService
	GetAppService() *AppService
	NewOrgService() *OrgService
	GetOrgService() *OrgService
	NewServiceBindingService() *ServiceBindingService
	GetServiceBindingService() *ServiceBindingService
	NewServiceInstanceService() *ServiceInstanceService
	GetServiceInstanceService() *ServiceInstanceService
	NewSpaceService() *SpaceService
	GetSpaceService() *SpaceService
}
