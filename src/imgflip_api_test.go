package meme_as_code_test

import (
	. "meme_as_code/src"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ImgFlip", func() {
	Context("When parsing the result of the API", func() {
		It("returns error if the field success is not true", func() {
			apiString := `{"success":false,"error_message":"No texts specified. Remember, API request params are http parameters not JSON."} `
			apiRawResp := strings.NewReader(apiString)
			url, err := GetUrl(apiRawResp)
			Expect(url).To(Equal(""))
			Expect(err).To(HaveOccurred())
		})
		It("returns the correct url if success", func() {
			apiString := `{"success":true,"data":{"url":"https:\/\/i.imgflip.com\/5ep068.jpg","page_url":"https:\/\/imgflip.com\/i\/5ep068"}}`
			apiRawResp := strings.NewReader(apiString)
			url, err := GetUrl(apiRawResp)
			Expect(url).To(Equal("https://i.imgflip.com/5ep068.jpg"))
			Expect(err).To(BeNil())
		})
	})
})
