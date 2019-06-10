package query

import (
	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

func (g Apps) GroupBySpace(_ *Inquistor) (AppGroup, error) {
	return AppGroupBy(g, func(app cfclient.App) (string, error) {
		return app.SpaceGuid, nil
	})
}

func (g Apps) GroupByOrg(i *Inquistor) (AppGroup, error) {
	return AppGroupBy(g, func(app cfclient.App) (string, error) {
		s, e := i.GetSpaceService().GetByGuid(app.SpaceGuid)
		if e != nil {
			return "", e
		}

		return s.OrganizationGuid, nil
	})
}

// GroupBySpaceAndOrg ...
func (g Apps) GroupBySpaceAndOrg(i *Inquistor) (MappedAppGroup, error) {
	ag, err := g.GroupBySpace(i)
	if err != nil {
		return nil, err
	}

	return ag.GroupByOrg(i)
}

func (g AppGroup) GroupByOrg(i *Inquistor) (MappedAppGroup, error) {
	return AppGroupMappedSliceBy(g, func(_ string, apps Apps) (string, error) {
		s, e := i.GetSpaceService().GetByGuid(apps[0].SpaceGuid)
		if e != nil {
			return "", e
		}

		return s.OrganizationGuid, nil
	})
}
