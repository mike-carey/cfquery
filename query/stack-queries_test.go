package query_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/cfquery/query"

	fakes "github.com/mike-carey/cfquery/cf/fakes"
)

var _ = Describe("StackQueries", func() {

	var (
		fakeClient *fakes.FakeCFClient
		inquisitor Inquisitor
	)

	BeforeEach(func () {
		fakeClient = new(fakes.FakeCFClient)
		inquisitor = NewInquisitor(fakeClient)
	})

	Describe("GetStackByName", func() {
		stack := cfclient.Stack{
			Guid: "stack1",
			Name: "stack-1",
		}

		fakeClient.ListStacksByQueryReturns(stack, nil)

		s, err := inquisitor.GetStackByName("stack-1")

		Expect(err).NotTo(BeNil())
		Expect(s).To(Equal(stack))
	})

})
