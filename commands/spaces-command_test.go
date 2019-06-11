package commands_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/cfquery/commands"

	"github.com/mike-carey/cfquery/query"

	fakes "github.com/mike-carey/cfquery/cf/fakes"
	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

var _ = Describe("SpacesCommand", func() {

	var (
		fakeClient *fakes.FakeCFClient
		inquisitor *query.Inquisitor

		spaces []cfclient.Space
	)

	BeforeEach(func() {
		fakeClient = new(fakes.FakeCFClient)
		inquisitor = query.NewInquisitor(fakeClient)

		spaces = []cfclient.Space{
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
		}

		fakeClient.ListSpacesReturns(spaces, nil)
	})

	It("Should GroupBy Org", func() {
		c := &SpacesCommand{}
		x := query.SpaceGroup{
			"org1": query.Spaces{
				spaces[0],
				spaces[1],
			},
		}

		i, e := c.Run(&Options{
			GroupBy: "org",
		}, inquisitor)

		Expect(e).To(BeNil())
		Expect(x).To(Equal(i))
	})

})
