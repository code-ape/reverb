package parser2_test

import (
	. "github.com/code-ape/reverb/parser2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"strings"
)

var _ = Describe("Scanner", func() {

	Context("NewScanner", func() {

		It("takes an io.Reader and returns a *Scanner", func() {
			r := strings.NewReader("test string")
			s := NewScanner(r)
			Expect(s).Should(BeAssignableToTypeOf(&Scanner{}))
		})

	})

	Context("Next() and Char()", func() {
		var s *Scanner

		BeforeEach(func() {
			s = NewScanner(strings.NewReader("test string\nline 2"))
		})

		It("returns correct Character", func() {
			s.Next()
			Expect(s.Char()).Should(Equal(&Character{"t", 1,1,1}))
			s.Next()
			Expect(s.Char()).Should(Equal(&Character{"e", 2,1,2}))
		})

		It("registers line breaks correctly", func() {
			c := &Character{}
			for i:=0; i<13; i++ {
				s.Next()
				c = s.Char()
			}
			Expect(c).Should(Equal(&Character{"l", 13, 2, 1}))
		})

		It("returns false after end of string", func() {
			count := 0
			expected := len("text string\nline 2")
			for s.Next() {
				count ++
			}
			Expect(count).Should(Equal(expected))
		})

	})

})
