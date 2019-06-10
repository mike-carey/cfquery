package util_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/cfquery/util"
)

type Foo struct {
	Name string
}

var _ = Describe("Util", func() {

	Describe("MapKeys", func() {
		description := "Should pull out all keys from a(n) %s map"

		It(fmt.Sprintf(description, "empty"), func() {
			By("interface map")

			interfaceDict := make(map[string]interface{}, 0)
			interfaceKeys := MapKeys(interfaceDict)

			Expect(interfaceKeys).To(BeEmpty())

			By("string map")

			stringDict := make(map[string]string, 0)
			stringKeys := MapKeys(stringDict)

			Expect(stringKeys).To(BeEmpty())

			By("int map")

			intDict := make(map[string]int, 0)
			intKeys := MapKeys(intDict)

			Expect(intKeys).To(BeEmpty())

			By("slice map")

			sliceDict := make(map[string][]interface{}, 0)
			sliceKeys := MapKeys(sliceDict)

			Expect(sliceKeys).To(BeEmpty())

			By("map map")

			mapDict := make(map[string]map[string]interface{}, 0)
			mapKeys := MapKeys(mapDict)

			Expect(mapKeys).To(BeEmpty())

			By("foo map")

			fooDict := make(map[string]Foo, 0)
			fooKeys := MapKeys(fooDict)

			Expect(fooKeys).To(BeEmpty())
		})

		It(fmt.Sprintf(description, "non-empty"), func() {
			expectedSlice := []string{"one", "two", "three",}

			By("interface map")

			interfaceDict := map[string]interface{} {
				"one": 1,
				"two": "TWO",
				"three": byte(76),
			}

			interfaceKeys := MapKeys(interfaceDict)

			Expect(interfaceKeys).Should(ConsistOf(expectedSlice))

			By("string map")

			stringDict := map[string]string {
				"one": "ONE",
				"two": "TWO",
				"three": "SEVENTY-SIX",
			}
			stringKeys := MapKeys(stringDict)

			Expect(stringKeys).Should(ConsistOf(expectedSlice))

			By("int map")

			intDict := map[string]int{
				"one": 1,
				"two": 2,
				"three": 76,
			}
			intKeys := MapKeys(intDict)

			Expect(intKeys).Should(ConsistOf(expectedSlice))

			By("slice map")

			sliceDict := map[string][]interface{} {
				"one": []interface{} {
					1,
					"TWO",
				},
				"two": []interface{} {
					2,
					"SEVENTY-SIX",
				},
				"three": []interface{} {
					76,
					"ONE",
				},
			}
			sliceKeys := MapKeys(sliceDict)

			Expect(sliceKeys).Should(ConsistOf(expectedSlice))

			By("map map")

			mapDict := map[string]map[string]interface{} {
				"one": map[string]interface{} {
					"two": 1,
					"three": "TWO",
				},
				"two": map[string]interface{} {
					"three": 2,
					"one": "SEVENTY-SIX",
				},
				"three": map[string]interface{} {
					"one": 76,
					"two": "ONE",
				},
			}
			mapKeys := MapKeys(mapDict)

			Expect(mapKeys).Should(ConsistOf(expectedSlice))

			By("foo map")

			fooDict := map[string]Foo {
				"one": Foo{Name: "one",},
				"two": Foo{Name: "two",},
				"three": Foo{Name: "three",},
			}
			fooKeys := MapKeys(fooDict)

			Expect(fooKeys).Should(ConsistOf(expectedSlice))
		})

	})

	Describe("StringSliceContains", func() {
		pool := []string {
			"foo",
			"bar",
			"baz",
		}

		It("Should contain the string: foo", func() {
			doesContain := StringSliceContains(pool, "foo")
			Expect(doesContain).To(BeTrue())
		})

		It("Should not contain the string: nothere", func() {
			doesNotContain := StringSliceContains(pool, "nothere")
			Expect(doesNotContain).To(BeFalse())
		})
	})

})
