// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package query_test

import (
	. "github.com/mike-carey/cfquery/query"
	. "github.com/onsi/gomega"
	. "github.com/onsi/ginkgo"
	"fmt"
	"io"
	"reflect"

	"github.com/cloudfoundry-community/go-cfclient"
	"github.com/mike-carey/cfquery/cf/fakes"
)

func getServiceInstanceServiceName() string {
	return fmt.Sprintf("%v", reflect.TypeOf(ServiceInstanceService{}))
}

func getServiceInstanceName() string {
	return fmt.Sprintf("%v", reflect.TypeOf(cfclient.ServiceInstance{}))
}

func newServiceInstance(guid string) cfclient.ServiceInstance {
	return cfclient.ServiceInstance{
		Guid: guid,
	}
}

func ServiceInstanceShouldImplementService(service Service) {
	io.WriteString(GinkgoWriter, fmt.Sprintf("If this did not compile, it indicates that %v does not implement Service", reflect.TypeOf(&cfclient.ServiceInstance{})))
}

var _ = Describe(getServiceInstanceServiceName()+"Base", func() {

	var (
		fakeClient *fakes.FakeCFClient
		service    *ServiceInstanceService
	)

	BeforeEach(func() {
		fakeClient = new(fakes.FakeCFClient)
		service = NewServiceInstanceService(fakeClient)
	})

	It("Should Implement Service", func() {
		ServiceInstanceShouldImplementService(service)
	})

	It("Should Get All "+getServiceInstanceName()+"s", func() {
		By("Calling the CFClient")
		expect := ServiceInstances{
			newServiceInstance("one"),
			newServiceInstance("two"),
			newServiceInstance("three"),
		}

		fakeClient.ListServiceInstancesReturns(expect, nil)

		actual, err := service.GetAllServiceInstances()

		Expect(err).To(BeNil())
		Expect(actual).Should(ConsistOf(expect))

		Expect(fakeClient.ListServiceInstancesCallCount()).To(Equal(1))

		By("Using the storage")

		actual2, err2 := service.GetAllServiceInstances()

		Expect(err2).To(BeNil())
		Expect(actual2).Should(ConsistOf(expect))

		// It should still equal 1
		Expect(fakeClient.ListServiceInstancesCallCount()).To(Equal(1))
	})

})
