package parser2_test

import (
	. "github.com/code-ape/reverb/parser2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Character", func() {

	Context("NewCharacter", func() {

		It("Takes a string and three integers and returns a pointer to a Character", func() {
			char := NewCharacter("a", 1,1,1)
			Expect(char).Should(BeAssignableToTypeOf(&Character{}))
		})

	})

	Context("Variables of Character", func() {
		var c *Character
		BeforeEach(func() {
			c = NewCharacter("a",2,3,4)
		})

		It("has correct TextIndex", func() {
			Expect(c.TextIndex).Should(Equal(2))
		})

		It("has correct LineIndex", func() {
			Expect(c.LineIndex).Should(Equal(3))
		})

		It("has correct ColumnIndex", func() {
			Expect(c.ColumnIndex).Should(Equal(4))
		})

	})

})
