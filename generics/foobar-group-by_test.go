package generics_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/cfquery/generics"
)

var _ = Describe("Foobar", func() {

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

	It("Should sort by bar name", func() {
		expect := map[string][]Foo{
			"one": []Foo{
				foos[0],
				foos[1],
				foos[2],
			},
			"four": []Foo{
				foos[3],
			},
		}

		actual, err := FooGroupBy(foos, func(foo Foo) (string, error) {
			return foo.Bar.Name, nil
		})

		Expect(err).To(BeNil())
		Expect(actual).To(Equal(expect))
	})

})
