package query

import (
	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

func (g Spaces) GroupByOrg(_ *Inquistor) (SpaceGroup, error) {
	return SpaceGroupBy(g, func(space cfclient.Space) (string, error) {
		return space.OrganizationGuid, nil
	})
}