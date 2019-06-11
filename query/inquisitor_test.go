package query_test

import (
	"reflect"

	"github.com/mike-carey/cfquery/cf/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/cfquery/query"
)

type FooService struct {}

var _ = Describe("Inquisitor", func() {

	var (
		fakeClient *fakes.FakeCFClient
		inquisitor *Inquisitor
	)

	BeforeEach(func() {
		fakeClient = new(fakes.FakeCFClient)
		inquisitor = NewInquisitor(fakeClient)
	})

	It("Should know how to Get the proper service", func() {
		service := inquisitor.GetService("app")

		Expect(reflect.TypeOf(service)).To(Equal(reflect.TypeOf(&AppService{})))
	})

})
