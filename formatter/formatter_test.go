package formatter_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/mike-carey/cfquery/formatter"
)

var _ = Describe("Formatter", func() {

	var formatter *Formatter

	It("Should encode into json", func() {
		formatter = NewFormatter(JSON)

		v := map[string]string{
			"foo": "bar",
			"bar": "baz",
		}

		buff, err := formatter.Format(v)

		Expect(err).To(BeNil())

		s := buff.String()

		expect := "{\"bar\":\"baz\",\"foo\":\"bar\"}"

		Expect(expect).To(EqualJson(s))
	})

	It("Should encode into yaml", func() {
		formatter = NewFormatter(YAML)

		v := map[string]string{
			"foo": "bar",
			"bar": "baz",
		}

		buff, err := formatter.Format(v)

		Expect(err).To(BeNil())

		s := buff.String()
		GinkgoWriter.Write([]byte(s))

		expect := `---
bar: baz
foo: bar`

		Expect(expect).To(EqualYaml(s))
	})

})
