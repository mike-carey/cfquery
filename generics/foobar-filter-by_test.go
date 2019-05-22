package generics_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/cfquery/generics"
)

var _ = Describe("FoobarFilterBy", func() {

	var (
		foos []Foo
		bars []Bar
	)

	BeforeEach(func() {
		foo1, bar1 := NewFooBarPair("one")
		foos = append(foos, foo1)
		bars = append(bars, bar1)

		foo2 := NewFoo("two", &bar1)
		foos = append(foos, foo2)

		foo3 := NewFoo("three", &bar1)
		foos = append(foos, foo3)

		foo4, bar4 := NewFooBarPair("four")
		foos = append(foos, foo4)
		bars = append(bars, bar4)
	})

	It("Should filter by bar name", func() {
		expect := []Foo{
			foos[0],
			foos[3],
		}

		actual, err := FooFilterBy(foos, func(foo Foo) (bool, error) {
			return foo.Name == foo.Bar.Name, nil
		})

		Expect(err).To(BeNil())
		Expect(actual).To(Equal(expect))
	})

})
