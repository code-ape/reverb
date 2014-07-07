package parser_test

import (
	. "github.com/code-ape/reverb/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const TestKind = "TEST_KIND"

var _ = Describe("ParserCore", func() {

  var p *Parser

	Describe("NewParser", func() {

		var null_p *Parser

		BeforeEach(func() {
			p = NewParser()
			null_p = &Parser{}
		})

		It("returns parser pointer", func() {
			Expect(p).Should(BeAssignableToTypeOf(null_p))
		})

	})


	Describe("IncrementCursor", func() {

    BeforeEach(func(){
      p = NewParser()
    })

		It("increments char number", func() {
			line_i, char_i := p.LineNum, p.CharNum
			p.IncrementCursor("s")
			Expect(p.LineNum).Should(Equal(line_i))
			Expect(p.CharNum).Should(Equal(char_i + 1))
		})

		It("increments line number", func() {
			line_i, char_i := p.LineNum, p.CharNum
			p.IncrementCursor("\n")
			Expect(p.LineNum).Should(Equal(line_i + 1))
			Expect(p.CharNum).Should(Equal(char_i))
		})

	})

  Describe("NewBlock", func() {
    var b *Item

    BeforeEach(func() {
      p = NewParser()
      b = p.NewBlock(TestKind)
    })

    It("is has the kind it was inited with", func() {
      Expect(b.Kind).Should(Equal(TestKind))
    })

    It("inits with the correct line and char number", func() {
      Expect(b.StartLineNum).Should(Equal(p.LineNum))
      Expect(b.StartCharNum).Should(Equal(p.CharNum))
    })

  })

  Describe("AddBlock", func() {

    BeforeEach(func() {
      p = NewParser()

    })

    It("adds a block under the current one", func() {
      len_child_i := len(p.Content.ChildItems)
      p.AddBlock(TestKind)
      len_child_f := len(p.Content.ChildItems)
      Expect(len_child_f).Should(Equal(len_child_i + 1))
      })

  })

})
