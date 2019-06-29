package config

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

var _ = Describe("config", func() {

	Describe("Validate", func() {
		var settings *Settings

		BeforeEach(func() {
			settings = DefaultSettings()
			settings.Storage.Type = "postgres"
			settings.Storage.URI = "test_uri"
		})

		Context("when config is valid", func() {
			It("returns no error", func() {
				err := settings.Validate()
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("when port is missing", func() {
			It("returns an error", func() {
				settings.Server.Port = 0
				err := settings.Validate()
				Expect(err).Should(HaveOccurred())
			})
		})

		Context("when storage type is missing", func() {
			It("returns an error", func() {
				settings.Storage.Type = ""
				err := settings.Validate()
				Expect(err).Should(HaveOccurred())
			})
		})

		Context("when storage URI is missing", func() {
			It("returns an error", func() {
				settings.Storage.URI = ""
				err := settings.Validate()
				Expect(err).Should(HaveOccurred())
			})
		})

	})

})
