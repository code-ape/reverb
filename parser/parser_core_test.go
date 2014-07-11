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

    It("sets CursorEnv to CurrentBlock Kind", func() {
      Expect(p.CursorEnv).Should(Equal(p.CurrentBlock.Kind))
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
    var b *Block

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
      len_child_i := len(p.Content.ChildBlocks)
      p.AddBlock(TestKind)
      len_child_f := len(p.Content.ChildBlocks)
      Expect(len_child_f).Should(Equal(len_child_i + 1))
    })

    It("new block is child of previous block", func() {
      b_i := p.CurrentBlock
      p.AddBlock(TestKind)
      Expect(b_i.ChildBlocks).Should(ContainElement(p.CurrentBlock))
    })

    It("previous block is parent of new block", func() {
      b_i := p.CurrentBlock
      p.AddBlock(TestKind)
      Expect(p.CurrentBlock.ParentBlock).Should(Equal(b_i))
    })

  })

  Describe("EndBlock", func() {
    
    Context("after AddBlock has been called", func() {
      BeforeEach(func() {
        p = NewParser()
        p.AddBlock(TestKind)
      })

      It("sets correct end line and char numbers", func() {
        b := p.CurrentBlock
        p.EndBlock()
        Expect(b.EndLineNum).Should(Equal(p.LineNum))
        Expect(b.EndCharNum).Should(Equal(p.CharNum))
      })

      It("sets current block to parent block", func() {
        par := p.CurrentBlock.ParentBlock
        p.EndBlock()
        Expect(p.CurrentBlock).Should(Equal(par))
      })

      It("sets cursor to parent block kind", func() {
        par := p.CurrentBlock.ParentBlock
        p.EndBlock()
        Expect(p.CursorEnv).Should(Equal(par.Kind))
      })

    })

  })

  Describe("AddChar", func() {
    It("adds indicator then given char to text of currentblock", func() {
      p.CurrentBlock.Text = "a"
      p.Indicator = "b"
      p.AddChar("c")
      Expect(p.CurrentBlock.Text).Should(Equal("abc"))
      })
  })

})
