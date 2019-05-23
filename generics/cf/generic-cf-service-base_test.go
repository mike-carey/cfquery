package query_test

import (
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mike-carey/cfquery/cf/fakes"
	. "github.com/mike-carey/cfquery/query"

	"github.com/cheekybits/genny/generic"
)

type CFObject generic.Type

func getCFObjectServiceName() string {
	return fmt.Sprintf("%v", reflect.TypeOf(CFObjectService{}))
}

func getCFObjectName() string {
	return fmt.Sprintf("%v", reflect.TypeOf(CFObject{}))
}

func newCFObject(guid string) CFObject {
	return CFObject{
		Guid: guid,
	}
}

var _ = Describe(getCFObjectServiceName()+"Base", func() {

	var (
		fakeClient *fakes.FakeCFClient
		service    *CFObjectService
	)

	BeforeEach(func() {
		fakeClient = new(fakes.FakeCFClient)
		service = NewCFObjectService(fakeClient)
	})

	It("Should Get All "+getCFObjectName()+"s", func() {
		By("Calling the CFClient")
		expect := []CFObject{
			newCFObject("one"),
			newCFObject("two"),
			newCFObject("three"),
		}

		fakeClient.ListCFObjectsReturns(expect, nil)

		actual, err := service.GetAll()

		Expect(err).To(BeNil())
		Expect(actual).Should(ConsistOf(expect))

		Expect(fakeClient.ListCFObjectsCallCount()).To(Equal(1))

		By("Using the storage")

		actual2, err2 := service.GetAll()

		Expect(err2).To(BeNil())
		Expect(actual2).Should(ConsistOf(expect))

		// It should still equal 1
		Expect(fakeClient.ListCFObjectsCallCount()).To(Equal(1))
	})

})
