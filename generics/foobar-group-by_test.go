package generics_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/cfquery/generics"
)

var _ = Describe("Foobar", func() {

	var (
		foos   []Foo
		bars   []Bar
		fooMap map[string]Foo
	)

	BeforeEach(func() {
		fooMap = make(map[string]Foo, 0)

		foo1, bar1 := NewFooBarPair("one")
		foos = append(foos, foo1)
		bars = append(bars, bar1)
		fooMap["one"] = foo1

		foo2 := NewFoo("two", &bar1)
		foos = append(foos, foo2)
		fooMap["two"] = foo2

		foo3 := NewFoo("three", &bar1)
		foos = append(foos, foo3)
		fooMap["three"] = foo3

		foo4, bar4 := NewFooBarPair("four")
		foos = append(foos, foo4)
		bars = append(bars, bar4)
		fooMap["four"] = foo4
	})

	It("Should group slice by bar name", func() {
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

	It("Should group map by bar name", func() {
		expect := map[string]map[string]Foo{
			"one": map[string]Foo{
				"one":   foos[0],
				"two":   foos[1],
				"three": foos[2],
			},
			"four": map[string]Foo{
				"four": foos[3],
			},
		}

		actual, err := FooGroupMapBy(fooMap, func(foo Foo) (string, error) {
			return foo.Bar.Name, nil
		})

		Expect(err).To(BeNil())
		Expect(actual).To(Equal(expect))
	})
})
