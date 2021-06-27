package meme_as_code_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMemeAsCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MemeAsCode Suite")
}
