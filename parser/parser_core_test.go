package parser_test

import (
	. "github.com/code-ape/reverb/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParserCore", func() {



  Describe("NewParser", func(){

    var (
      p *Parser
      null_p *Parser
    )

    BeforeEach(func() {
      p = NewParser()
      null_p = &Parser{}
    })

    It("returns parser pointer", func() {
      Expect(p).Should(BeAssignableToTypeOf(null_p))
    })
  })

  Describe("IncrementCursor", func() {

    It("increments char number", func() {
      p := NewParser()
      line_i, char_i := p.LineNum, p.CharNum
      p.IncrementCursor("s")
      line_f, char_f := p.LineNum, p.CharNum
      Expect(line_f).Should(Equal( line_i ))
      Expect(char_f).Should(Equal( char_i + 1 ))
    })
  
    It("increments line number", func() {
      p := NewParser()
      line_i, char_i := p.LineNum, p.CharNum
      p.IncrementCursor("\n")
      line_f, char_f := p.LineNum, p.CharNum
      Expect(line_f).Should(Equal( line_i + 1) )
      Expect(char_f).Should(Equal( char_i ))
    })
  })

})
