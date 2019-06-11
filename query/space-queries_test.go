package query_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/cfquery/query"

	"github.com/cloudfoundry-community/go-cfclient"

	fakes "github.com/mike-carey/cfquery/cf/fakes"
)

var _ = Describe("Space Queries", func() {

	var (
		fakeClient *fakes.FakeCFClient
		inquisitor *Inquisitor
		spaces Spaces
	)

	BeforeEach(func() {
		fakeClient = new(fakes.FakeCFClient)
		inquisitor = NewInquisitor(fakeClient)

		spaces = Spaces{
			cfclient.Space{
				Guid: "space1",
				Name: "space-1",
				OrganizationGuid: "org1",
			},
			cfclient.Space{
				Guid: "space2",
				Name: "space-2",
				OrganizationGuid: "org1",
			},
			cfclient.Space{
				Guid: "space3",
				Name: "space-3",
				OrganizationGuid: "org2",
			},
			cfclient.Space{
				Guid: "space4",
				Name: "space-4",
				OrganizationGuid: "org2",
			},
		}
	})

	It("Should group by Org", func() {
		expect := map[string][]cfclient.Space {
			"org1": []cfclient.Space {
				spaces[0],
				spaces[1],
			},
			"org2": []cfclient.Space {
				spaces[2],
				spaces[3],
			},
		}

		gas, err := spaces.GroupByOrg(inquisitor)
		Expect(err).To(BeNil())

		RecursiveMapCompare(expect, gas)
	})

})
