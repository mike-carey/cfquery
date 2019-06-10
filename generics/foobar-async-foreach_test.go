package generics_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/cfquery/generics"
)

var _ = Describe("FoobarAsyncForeach", func() {
	It("Should cycle through every Foo and return proper Bars", func() {
		foos := make(Foos, 0)
		bars := make([]Bar, 0)

		for _, i := range []string{"one", "two", "three"} {
			foo, bar := NewFooBarPair(i)
			foos = append(foos, foo)
			bars = append(bars, bar)
		}

		_bars, errs := ForEachFooToBar(foos, func(foo Foo) (Bar, error) {
			return *foo.Bar, nil
		})

		Expect(errs).To(BeEmpty())
		Expect(_bars).Should(ConsistOf(bars))
	})

	It("Should cycle through every Foo and return errors", func() {
		foos := make(Foos, 0)
		bars := make([]Bar, 0)

		for _, i := range []string{"one", "two", "three"} {
			foo, bar := NewFooBarPair(i)
			foos = append(foos, foo)
			if foo.Name != "one" {
				bars = append(bars, bar)
			}
		}

		err := errors.New("Name is one")

		_bars, errs := ForEachFooToBar(foos, func(foo Foo) (Bar, error) {
			if foo.Name == "one" {
				return Bar{}, err
			}
			return *foo.Bar, nil
		})

		Expect(errs).Should(ConsistOf(err))
		Expect(_bars).Should(ConsistOf(bars))
	})
})
