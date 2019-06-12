package query_test

import (
	// "reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/cfquery/query"

	fakes "github.com/mike-carey/cfquery/cf/fakes"
)

var _ = Describe("Inquisitor", func() {

	var (
		fakeClient *fakes.FakeCFClient
		inquisitor Inquisitor
	)

	BeforeEach(func() {
		fakeClient = new(fakes.FakeCFClient)
	})

	It("Should be new-able and implement Inquisitor", func() {
		inquisitor = NewInquisitor(fakeClient)

		Expect(inquisitor).NotTo(BeNil())
	})

	// It("Should know how to Get the proper service", func() {
	// 	service := inquisitor.GetService("app")
	//
	// 	Expect(reflect.TypeOf(service)).To(Equal(reflect.TypeOf(&AppService{})))
	// })

})
