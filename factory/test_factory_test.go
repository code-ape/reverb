package factory_test

import (
	. "github.com/code-ape/reverb/factory"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TestFactory", func() {

	Context("", func() {

		factory := MakeTestFactory()

		m1 := factory.NewMap()
		factory.BeforeEach(func() {
			m1["exp"] = 1
			m1["val"] = 1
		})
		factory.RunAll(m1)

		m2 := factory.NewMap()
		factory.BeforeEach(func() {
			m2["exp"] = 2
			m2["val"] = 2
		})
		factory.RunAll(m2)

	})

})

func MakeTestFactory() *Factory {
	var a int

	f := NewTestFactory()

	f.JustBeforeEach(func() {
		a = f.Int("exp")
	})

	f.It("test 1", "exp = val", func() {
		Expect(a).Should(Equal(f.Int("val")))
	})

	return f
}
