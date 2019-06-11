package query_test

import (
	"reflect"
	"io"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mike-carey/cfquery/cf/fakes"
	. "github.com/mike-carey/cfquery/query"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

func getItemServiceName() string {
	return fmt.Sprintf("%v", reflect.TypeOf(ItemService{}))
}

func getItemName() string {
	return fmt.Sprintf("%v", reflect.TypeOf(Item{}))
}

func newItem(guid string) Item {
	return Item{
		Guid: guid,
	}
}

func ItemShouldImplementService(service Service) {
	io.WriteString(GinkgoWriter, fmt.Sprintf("If this did not compile, it indicates that %v does not implement Service", reflect.TypeOf(&Item{})))
}

var _ = Describe(getItemServiceName()+"Base", func() {

	var (
		fakeClient *fakes.FakeCFClient
		service    *ItemService
	)

	BeforeEach(func() {
		fakeClient = new(fakes.FakeCFClient)
		service = NewItemService(fakeClient)
	})

	It("Should Implement Service", func() {
		ItemShouldImplementService(service)
	})

	It("Should Get All "+getItemName()+"s", func() {
		By("Calling the CFClient")
		expect := Items{
			newItem("one"),
			newItem("two"),
			newItem("three"),
		}

		fakeClient.ListItemsReturns(expect, nil)

		actual, err := service.GetAllItems()

		Expect(err).To(BeNil())
		Expect(actual).Should(ConsistOf(expect))

		Expect(fakeClient.ListItemsCallCount()).To(Equal(1))

		By("Using the storage")

		actual2, err2 := service.GetAllItems()

		Expect(err2).To(BeNil())
		Expect(actual2).Should(ConsistOf(expect))

		// It should still equal 1
		Expect(fakeClient.ListItemsCallCount()).To(Equal(1))
	})

})
