package parser2_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestParser2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parser2 Suite")
}
