package query_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/cfquery/query"

	"github.com/cloudfoundry-community/go-cfclient"

	fakes "github.com/mike-carey/cfquery/cf/fakes"
)

var _ = Describe("App Queries", func() {

	var (
		fakeClient *fakes.FakeCFClient
		inquistor *Inquistor
		apps Apps
		spaces SpaceMap
	)

	BeforeEach(func() {
		fakeClient = new(fakes.FakeCFClient)
		inquistor = NewInquistor(fakeClient)

		apps = Apps{
			cfclient.App{
				Guid:      "app1",
				Name:      "app-1",
				SpaceGuid: "space1",
			},
			cfclient.App{
				Guid:      "app2",
				Name:      "app-2",
				SpaceGuid: "space1",
			},
			cfclient.App{
				Guid:      "app3",
				Name:      "app-3",
				SpaceGuid: "space2",
			},
			cfclient.App{
				Guid:      "app4",
				Name:      "app-4",
				SpaceGuid: "space2",
			},
			cfclient.App{
				Guid:      "app5",
				Name:      "app-5",
				SpaceGuid: "space3",
			},
			cfclient.App{
				Guid:      "app6",
				Name:      "app-6",
				SpaceGuid: "space4",
			},
		}

		spaces = map[string]cfclient.Space{
			"space1": cfclient.Space{
				Guid: "space1",
				Name: "space-1",
				OrganizationGuid: "org1",
			},
			"space2": cfclient.Space{
				Guid: "space2",
				Name: "space-2",
				OrganizationGuid: "org1",
			},
			"space3": cfclient.Space{
				Guid: "space3",
				Name: "space-3",
				OrganizationGuid: "org2",
			},
			"space4": cfclient.Space{
				Guid: "space4",
				Name: "space-4",
				OrganizationGuid: "org2",
			},
		}
	})

	It("Should group by Space then Org", func() {
		expect := map[string]map[string][]cfclient.App {
			"org1": map[string][]cfclient.App{
				"space1": []cfclient.App {
					apps[0],
					apps[1],
				},
				"space2": []cfclient.App {
					apps[2],
					apps[3],
				},
			},
			"org2": map[string][]cfclient.App{
				"space3": []cfclient.App {
					apps[4],
				},
				"space4": []cfclient.App {
					apps[5],
				},
			},
		}

		fakeClient.GetSpaceByGuidStub = func(guid string) (cfclient.Space, error) {
			return spaces[guid], nil
		}

		gas, err := apps.GroupBySpace(inquistor)
		Expect(err).To(BeNil())

		gags, err := gas.GroupByOrg(inquistor)
		Expect(err).To(BeNil())

		RecursiveMapCompare(expect, gags)
	})

	It("Should group by Org", func() {
		expect := map[string][]cfclient.App {
			"org1": []cfclient.App {
				apps[0],
				apps[1],
				apps[2],
				apps[3],
			},
			"org2": []cfclient.App {
				apps[4],
				apps[5],
			},
		}

		fakeClient.GetSpaceByGuidStub = func(guid string) (cfclient.Space, error) {
			return spaces[guid], nil
		}

		gas, err := apps.GroupByOrg(inquistor)
		Expect(err).To(BeNil())

		RecursiveMapCompare(expect, gas)
	})

})
