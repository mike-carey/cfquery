/**
 *
 */
package config_test

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/spf13/afero"

	. "github.com/mike-carey/cfquery/config"
)

var WHITESPACE = []string{" ", "\n", "\t"}

func StripWhitespace(str string) string {
	for _, char := range WHITESPACE {
		str = strings.Replace(str, char, "", -1)
	}

	return str
}

var _ = Describe("Config", func() {

	const configFile = "config.json"
	fs := afero.NewOsFs()

	Describe("Loading", func() {
		Context("Valid JSON", func() {
			data := []byte("{\"foundation1\":{\"api_url\":\"https://api.domain.com\",\"user\":\"admin\",\"password\":\"password\"}}")

			BeforeEach(func() {
				afero.WriteFile(fs, configFile, data, 0644)
			})

			AfterEach(func() {
				fs.Remove(configFile)
			})

			It("Should parse out values", func() {
				config, err := LoadConfig(configFile)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(config["foundation1"]).NotTo(BeNil())

				foundation := config["foundation1"]

				Expect(foundation.ApiAddress).To(Equal("https://api.domain.com"))
				Expect(foundation.Username).To(Equal("admin"))
				Expect(foundation.Password).To(Equal("password"))
			})
		})

		Context("Invalid JSON", func() {
			data := []byte("{'foundation1': {'api_url': 'https://api.domain.com', 'user': 'admin', 'password': 'password'} }")

			BeforeEach(func() {
				afero.WriteFile(fs, configFile, data, 0644)
			})

			AfterEach(func() {
				fs.Remove(configFile)
			})

			It("Should return an error", func() {
				_, err := LoadConfig(configFile)
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})
