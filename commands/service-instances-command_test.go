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
		inquisitor query.Inquisitor

		serviceInstances []cfclient.ServiceInstance
		// apps []cfclient.App
		spaces []cfclient.Space
	)

	BeforeEach(func() {
		fakeClient = new(fakes.FakeCFClient)
		inquisitor = query.NewInquisitor(fakeClient)

		serviceInstances = []cfclient.ServiceInstance{
			cfclient.ServiceInstance{
				Name: "service-instance-1",
				SpaceGuid: "space1",
			},
			cfclient.ServiceInstance{
				Name: "service-instance-2",
				SpaceGuid: "space1",
			},
		}

		// apps = []cfclient.App{
		// 	cfclient.App{
		// 		Name: "app-1",
		// 		SpaceGuid: "space1",
		// 	},
		// 	cfclient.App{
		// 		Name: "app-2",
		// 		SpaceGuid: "space1",
		// 	},
		// }

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

		fakeClient.ListServiceInstancesReturns(serviceInstances, nil)
		// fakeClient.ListAppsReturns(apps, nil)
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
		c := &ServiceInstancesCommand{}
		x := query.ServiceInstanceGroup{
			"space1": query.ServiceInstances{
				serviceInstances[0],
				serviceInstances[1],
			},
		}

		i, e := c.Run(&Options{
			GroupBy: "space",
		}, inquisitor)

		Expect(e).To(BeNil())
		Expect(x).To(Equal(i))
	})

	It("Should GroupBy Org", func() {
		c := &ServiceInstancesCommand{}
		x := query.ServiceInstanceGroup{
			"org1": query.ServiceInstances{
				serviceInstances[0],
				serviceInstances[1],
			},
		}

		i, e := c.Run(&Options{
			GroupBy: "org",
		}, inquisitor)

		Expect(e).To(BeNil())
		Expect(x).To(Equal(i))
	})

	It("Should GroupBy SpaceAndOrg", func() {
		c := &ServiceInstancesCommand{}
		x := query.MappedServiceInstanceGroup{
			"org1": query.ServiceInstanceGroup{
				"space1": query.ServiceInstances{
					serviceInstances[0],
					serviceInstances[1],
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
