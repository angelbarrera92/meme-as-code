package meme_as_code_test

import (
	. "meme_as_code/src"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	Context("When loading a config file", func() {
		It("returns error if the file does not exists", func() {
			config, err := GetConfigFromFile("NoExistFile.yaml")
			Expect(config).To(BeNil())
			Expect(err).To(HaveOccurred())
		})
		It("reads the file if it exists and is correct", func() {
			config, err := GetConfigFromFile("test-artifacts/config.yaml")

			Expect(err).To(BeNil())

			expected := &Config{
				OutputDir: "",
				Username:  "",
				Password:  "",
				Overrive:  false,
				Memes: []Meme{
					{
						Filename:   "one-just-dont-simply-01.jpg",
						TemplateId: "61579",
						Captions: []string{
							"One just dont simply",
							"automate memes"},
					},
					{
						Filename:   "one-just-dont-simply-02.jpg",
						TemplateId: "61579",
						Captions: []string{
							"One just dont simply",
							"automate memes in yaml"},
					},
				},
			}

			Expect(config).NotTo(BeNil())
			Expect(config).To(Equal(expected))

		})
	})
})
