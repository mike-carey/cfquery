package query_test

import (
	"reflect"

	"github.com/mike-carey/cfquery/cf/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/cfquery/query"
)

type FooService struct {}

var _ = Describe("Inquistor", func() {

	var (
		fakeClient *fakes.FakeCFClient
		inquistor *Inquistor
	)

	BeforeEach(func() {
		fakeClient = new(fakes.FakeCFClient)
		inquistor = NewInquistor(fakeClient)
	})

	It("Should know how to Get the proper service", func() {
		service := inquistor.GetService("app")

		Expect(reflect.TypeOf(service)).To(Equal(reflect.TypeOf(&AppService{})))
	})

})
