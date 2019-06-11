package commands_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/cfquery/commands"

	fakes "github.com/mike-carey/cfquery/commands/fakes"

	"github.com/mike-carey/cfquery/query"
)

var _ = Describe("Worker", func() {

	var (
		fakeCommand *fakes.FakeCommand
		options *Options
		inquisitor *query.Inquisitor
		worker *Worker
	)

	BeforeEach(func() {
		fakeCommand = new(fakes.FakeCommand)
		options = &Options{}
		inquisitor = &query.Inquisitor{}
		worker = &Worker{
			Command: fakeCommand,
			Options: options,
			Inquisitor: inquisitor,
		}
	})

	Describe("Option validation", func() {
		var err error

		It("Should throw error on bad Target", func() {
			options.Target = "nothere"

			By("Empty TargetOptions")
			fakeCommand.TargetOptionsReturns([]string{})
			err = worker.Validate()
			Expect(err).NotTo(BeNil())

			By("Not in TargetOptions")
			fakeCommand.TargetOptionsReturns([]string{"foo"})
			err = worker.Validate()
			Expect(err).NotTo(BeNil())
		})

		It("Should throw error on bad SortBy", func() {
			options.SortBy = "nothere"

			By("Empty SortByOptions")
			fakeCommand.SortByOptionsReturns([]string{})
			err = worker.Validate()
			Expect(err).NotTo(BeNil())

			By("Not in SortByOptions")
			fakeCommand.SortByOptionsReturns([]string{"foo"})
			err = worker.Validate()
			Expect(err).NotTo(BeNil())
		})

		It("Should throw error on bad GroupBy", func() {
			options.GroupBy = "nothere"

			By("Empty GroupByOptions")
			fakeCommand.GroupByOptionsReturns([]string{})
			err = worker.Validate()
			Expect(err).NotTo(BeNil())

			By("Not in GroupByOptions")
			fakeCommand.GroupByOptionsReturns([]string{"foo"})
			err = worker.Validate()
			Expect(err).NotTo(BeNil())
		})
	})

	// It("Should execute command Run method", func() {
	// 	worker.Work()
	// 	Expect(fakeCommand.RunCallCount()).To(Equal(1))
	// })

})
