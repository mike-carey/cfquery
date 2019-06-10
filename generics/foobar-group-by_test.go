package generics_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/cfquery/generics"
)

var _ = Describe("Foobar", func() {

	var (
		foos   Foos
		bars   []Bar
		fooMap FooMap
		fooGroup FooGroup
	)

	BeforeEach(func() {
		fooGroup = make(FooGroup, 0)
		fooMap = make(FooMap, 0)

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

		fooGroup["one"] = []Foo{
			foo1,
			foo2,
			foo3,
		}

		foo4, bar4 := NewFooBarPair("four")
		foos = append(foos, foo4)
		bars = append(bars, bar4)
		fooMap["four"] = foo4

		fooGroup["four"] = []Foo{
			foo4,
		}

		foo5, bar5 := NewFooBarPair("five")
		foos = append(foos, foo5)
		bars = append(bars, bar5)
		fooMap["five"] = foo5

		fooGroup["five"] = []Foo{
			foo5,
		}
	})

	It("Should group slice by bar name", func() {
		expect := FooGroup{
			"one": Foos{
				foos[0],
				foos[1],
				foos[2],
			},
			"four": Foos{
				foos[3],
			},
			"five": Foos{
				foos[4],
			},
		}

		actual, err := FooGroupBy(foos, func(foo Foo) (string, error) {
			return foo.Bar.Name, nil
		})

		Expect(err).To(BeNil())
		Expect(actual).To(Equal(expect))
	})

	It("Should group map by bar name", func() {
		expect := MappedFooMap{
			"one": FooMap{
				"one":   foos[0],
				"two":   foos[1],
				"three": foos[2],
			},
			"four": FooMap{
				"four": foos[3],
			},
			"five": FooMap{
				"five": foos[4],
			},
		}

		actual, err := FooGroupMapBy(fooMap, func(_ string, foo Foo) (string, error) {
			return foo.Bar.Name, nil
		})

		Expect(err).To(BeNil())
		Expect(actual).To(Equal(expect))
	})

	It("Should group mapped slice by first character", func() {
		expect := MappedFooGroup{
			"o": FooGroup{
				"one": Foos{
					foos[0],
					foos[1],
					foos[2],
				},
			},
			"f": FooGroup{
				"four": Foos{
					foos[3],
				},
				"five": Foos{
					foos[4],
				},
			},
		}

		actual, err := FooGroupMappedSliceBy(fooGroup, func(key string, _ Foos) (string, error) {
			return string([]rune(key)[0]), nil
		})

		Expect(err).To(BeNil())
		Expect(actual).To(Equal(expect))
	})

})
