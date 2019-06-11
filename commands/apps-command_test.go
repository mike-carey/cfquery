package commands_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/cfquery/commands"

	"github.com/mike-carey/cfquery/query"

	fakes "github.com/mike-carey/cfquery/cf/fakes"
	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

var _ = Describe("AppsCommand", func() {

	var (
		fakeClient *fakes.FakeCFClient
		inquisitor *query.Inquisitor

		apps []cfclient.App
		spaces []cfclient.Space
	)

	BeforeEach(func() {
		fakeClient = new(fakes.FakeCFClient)
		inquisitor = query.NewInquisitor(fakeClient)

		apps = []cfclient.App{
			cfclient.App{
				Name: "app-1",
				SpaceGuid: "space1",
			},
			cfclient.App{
				Name: "app-2",
				SpaceGuid: "space1",
			},
		}

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

		fakeClient.ListAppsReturns(apps, nil)
		fakeClient.GetSpaceByGuidStub = func (guid string) (cfclient.Space, error) {
			for _, space := range spaces {
				if space.Guid == guid {
					return space, nil
				}
			}

			return cfclient.Space{}, fmt.Errorf("Space not found: %s", guid)
		}
	})

	It("Should GroupBy Space", func() {
		c := &AppsCommand{}
		x := query.AppGroup{
			"space1": query.Apps{
				apps[0],
				apps[1],
			},
		}

		i, e := c.Run(&Options{
			GroupBy: "space",
		}, inquisitor)

		Expect(e).To(BeNil())
		Expect(x).To(Equal(i))
	})

	It("Should GroupBy Org", func() {
		c := &AppsCommand{}
		x := query.AppGroup{
			"org1": query.Apps{
				apps[0],
				apps[1],
			},
		}

		i, e := c.Run(&Options{
			GroupBy: "org",
		}, inquisitor)

		Expect(e).To(BeNil())
		Expect(x).To(Equal(i))
	})

	It("Should GroupBy SpaceAndOrg", func() {
		c := &AppsCommand{}
		x := query.MappedAppGroup{
			"org1": query.AppGroup{
				"space1": query.Apps{
					apps[0],
					apps[1],
				},
			},
		}

		i, e := c.Run(&Options{
			GroupBy: "space-and-org",
		}, inquisitor)

		Expect(e).To(BeNil())
		Expect(x).To(Equal(i))
	})

})
