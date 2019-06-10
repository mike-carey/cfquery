package query_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/cfquery/query"

	"github.com/cloudfoundry-community/go-cfclient"

	fakes "github.com/mike-carey/cfquery/cf/fakes"
)

var _ = Describe("Service Instance Queries", func() {

	var (
		fakeClient *fakes.FakeCFClient
		inquistor *Inquistor
		serviceInstances ServiceInstances
		spaces SpaceMap
	)

	BeforeEach(func() {
		fakeClient = new(fakes.FakeCFClient)
		inquistor = NewInquistor(fakeClient)

		serviceInstances = ServiceInstances{
			cfclient.ServiceInstance{
				Guid:      "service-instance1",
				Name:      "service-instance-1",
				SpaceGuid: "space1",
			},
			cfclient.ServiceInstance{
				Guid:      "service-instance2",
				Name:      "service-instance-2",
				SpaceGuid: "space1",
			},
			cfclient.ServiceInstance{
				Guid:      "service-instance3",
				Name:      "service-instance-3",
				SpaceGuid: "space2",
			},
			cfclient.ServiceInstance{
				Guid:      "service-instance4",
				Name:      "service-instance-4",
				SpaceGuid: "space2",
			},
			cfclient.ServiceInstance{
				Guid:      "service-instance5",
				Name:      "service-instance-5",
				SpaceGuid: "space3",
			},
			cfclient.ServiceInstance{
				Guid:      "service-instance6",
				Name:      "service-instance-6",
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
		expect := map[string]map[string][]cfclient.ServiceInstance {
			"org1": map[string][]cfclient.ServiceInstance{
				"space1": []cfclient.ServiceInstance {
					serviceInstances[0],
					serviceInstances[1],
				},
				"space2": []cfclient.ServiceInstance {
					serviceInstances[2],
					serviceInstances[3],
				},
			},
			"org2": map[string][]cfclient.ServiceInstance{
				"space3": []cfclient.ServiceInstance {
					serviceInstances[4],
				},
				"space4": []cfclient.ServiceInstance {
					serviceInstances[5],
				},
			},
		}

		fakeClient.GetSpaceByGuidStub = func(guid string) (cfclient.Space, error) {
			return spaces[guid], nil
		}

		gas, err := serviceInstances.GroupBySpace(inquistor)
		Expect(err).To(BeNil())

		gags, err := gas.GroupByOrg(inquistor)
		Expect(err).To(BeNil())

		RecursiveMapCompare(expect, gags)
	})

	It("Should group by Org", func() {
		expect := map[string][]cfclient.ServiceInstance {
			"org1": []cfclient.ServiceInstance {
				serviceInstances[0],
				serviceInstances[1],
				serviceInstances[2],
				serviceInstances[3],
			},
			"org2": []cfclient.ServiceInstance {
				serviceInstances[4],
				serviceInstances[5],
			},
		}

		fakeClient.GetSpaceByGuidStub = func(guid string) (cfclient.Space, error) {
			return spaces[guid], nil
		}

		gas, err := serviceInstances.GroupByOrg(inquistor)
		Expect(err).To(BeNil())

		RecursiveMapCompare(expect, gas)
	})

})
