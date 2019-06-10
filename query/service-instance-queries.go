package query

import (
	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

func (g ServiceInstances) GroupBySpace(_ *Inquistor) (ServiceInstanceGroup, error) {
	return ServiceInstanceGroupBy(g, func(si cfclient.ServiceInstance) (string, error) {
		return si.SpaceGuid, nil
	})
}

func (g ServiceInstances) GroupByOrg(i *Inquistor) (ServiceInstanceGroup, error) {
	return ServiceInstanceGroupBy(g, func(si cfclient.ServiceInstance) (string, error) {
		s, e := i.GetSpaceService().GetByGuid(si.SpaceGuid)
		if e != nil {
			return "", e
		}

		return s.OrganizationGuid, nil
	})
}

// GroupBySpaceAndOrg ...
func (g ServiceInstances) GroupBySpaceAndOrg(i *Inquistor) (MappedServiceInstanceGroup, error) {
	ag, err := g.GroupBySpace(i)
	if err != nil {
		return nil, err
	}

	return ag.GroupByOrg(i)
}

func (g ServiceInstanceGroup) GroupByOrg(i *Inquistor) (MappedServiceInstanceGroup, error) {
	return ServiceInstanceGroupMappedSliceBy(g, func(_ string, apps ServiceInstances) (string, error) {
		s, e := i.GetSpaceService().GetByGuid(apps[0].SpaceGuid)
		if e != nil {
			return "", e
		}

		return s.OrganizationGuid, nil
	})
}
