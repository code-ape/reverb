package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestReverb(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Reverb Suite")
}
