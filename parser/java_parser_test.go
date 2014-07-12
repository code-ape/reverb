package parser_test

import (
	. "github.com/code-ape/reverb/parser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
  "strings"
)

var _ = Describe("JavaParser", func() {

  Describe("NewParser", func() {

    It("returns a parser object", func() {
      p := NewJavaParser()
      Expect(p).Should(BeAssignableToTypeOf(&Parser{}))
    })

  })

})


var _ = Describe("JavaParser Integration", func() {

  var p *Parser

  BeforeEach(func() {
    p = NewJavaParser()
    p.Parse(java_code)
  })

  Context("after parsing", func() {

    It("has whitespace as head block", func() {
      Expect(p.Content.Kind).Should(Equal("WHITESPACE"))
    })

    Context("in head block", func() {
      Context("package declaration", func() {
        var pd *Block // pd -> package declaration
        BeforeEach(func() {
          pd = p.Content.ChildBlocks[0]
        })
        It("has correct Kind", func() {
          Expect(pd.Kind).Should(Equal("CODE"))
        })
        It("has correct start line number", func() {
          Expect(pd.StartLineNum).Should(Equal(1))
        })
        It("has correct end line number", func() {
          Expect(pd.EndLineNum).Should(Equal(1))
        })
        It("has correct start char number", func() {
          Expect(pd.StartCharNum).Should(Equal(1))
          })
        It("has correct end char number", func() {
          Expect(pd.EndCharNum).Should(Equal(25))
        })
        It("has correct text", func() {
          Expect(pd.Text).Should(Equal("package com.scribe.oauth;"))
        })
        It("has correct parent block", func() {
          Expect(pd.ParentBlock).Should(Equal(p.Content))
        })
        It("has corrent child blocks (none)", func() {
          Expect(pd.ChildBlocks).Should(BeZero())
        })        
      })
      Context("multiline comment", func() {
        var mc *Block // mc -> multiline comment
        BeforeEach(func() {
          mc = p.Content.ChildBlocks[1]
        })
        It("has correct Kind", func() {
          Expect(mc.Kind).Should(Equal("MULTI COMMENT"))
        })
        It("has correct start line number", func() {
          Expect(mc.StartLineNum).Should(Equal(2))
        })
        It("has correct end line number", func() {
          Expect(mc.EndLineNum).Should(Equal(4))
        })
        It("has correct start char number", func() {
          Expect(mc.StartCharNum).Should(Equal(1))
          })
        It("has correct end char number", func() {
          Expect(mc.EndCharNum).Should(Equal(10))
        })
        It("has correct text", func() {
          Expect(mc.Text).Should(Equal(multicomment_1))
        })
        It("has correct parent block", func() {
          Expect(mc.ParentBlock).Should(Equal(p.Content))
        })
        It("has corrent child blocks (none)", func() {
          Expect(mc.ChildBlocks).Should(BeZero())
        })
      })
      Context("regular import", func() {
        var ri *Block // ri -> regular import
        BeforeEach(func() {
          ri = p.Content.ChildBlocks[2]
        })
        It("has correct Kind", func() {
          Expect(ri.Kind).Should(Equal("CODE"))
        })
        It("has correct start line number", func() {
          Expect(ri.StartLineNum).Should(Equal(6))
        })
        It("has correct end line number", func() {
          Expect(ri.EndLineNum).Should(Equal(6))
        })
        It("has correct start char number", func() {
          Expect(ri.StartCharNum).Should(Equal(1))
          })
        It("has correct end char number", func() {
          Expect(ri.EndCharNum).Should(Equal(19))
        })
        It("has correct text", func() {
          Expect(ri.Text).Should(Equal("import com.foo.bar;"))
        })
        It("has correct parent block", func() {
          Expect(ri.ParentBlock).Should(Equal(p.Content))
        })
        It("has corrent child blocks (none)", func() {
          Expect(ri.ChildBlocks).Should(BeZero())
        })
      })
      Context("static star import", func() {
        var ssi *Block // ssi -> static star import
        BeforeEach(func() {
          ssi = p.Content.ChildBlocks[3]
        })
        It("has correct Kind", func() {
          Expect(ssi.Kind).Should(Equal("CODE"))
        })
        It("has correct start line number", func() {
          Expect(ssi.StartLineNum).Should(Equal(7))
        })
        It("has correct end line number", func() {
          Expect(ssi.EndLineNum).Should(Equal(7))
        })
        It("has correct start char number", func() {
          Expect(ssi.StartCharNum).Should(Equal(1))
          })
        It("has correct end char number", func() {
          Expect(ssi.EndCharNum).Should(Equal(28))
        })
        It("has correct text", func() {
          Expect(ssi.Text).Should(Equal("import static com.foo.baz.*;"))
        })
        It("has correct parent block", func() {
          Expect(ssi.ParentBlock).Should(Equal(p.Content))
        })
        It("has corrent child blocks (none)", func() {
          Expect(ssi.ChildBlocks).Should(BeZero())
        })
      })
      Context("single line comment after code", func() {
        var slc *Block // slc -> single line comment
        BeforeEach(func() {
          slc = p.Content.ChildBlocks[4]
        })
        It("has correct Kind", func() {
          Expect(slc.Kind).Should(Equal("SINGLE COMMENT"))
        })
        It("has correct start line number", func() {
          Expect(slc.StartLineNum).Should(Equal(7))
        })
        It("has correct end line number", func() {
          Expect(slc.EndLineNum).Should(Equal(7))
        })
        It("has correct start char number", func() {
          Expect(slc.StartCharNum).Should(Equal(30))
          })
        It("has correct end char number", func() {
          Expect(slc.EndCharNum).Should(Equal(53))
        })
        It("has correct text", func() {
          Expect(slc.Text).Should(Equal("//comment after code end"))
        })
        It("has correct parent block", func() {
          Expect(slc.ParentBlock).Should(Equal(p.Content))
        })
        It("has corrent child blocks (none)", func() {
          Expect(slc.ChildBlocks).Should(BeZero())
        })
      })
      Context("single line comment on own line", func() {
        var slc *Block // slc -> single line comment
        BeforeEach(func() {
          slc = p.Content.ChildBlocks[5]
        })
        It("has correct Kind", func() {
          Expect(slc.Kind).Should(Equal("SINGLE COMMENT"))
        })
        It("has correct start line number", func() {
          Expect(slc.StartLineNum).Should(Equal(8))
        })
        It("has correct end line number", func() {
          Expect(slc.EndLineNum).Should(Equal(8))
        })
        It("has correct start char number", func() {
          Expect(slc.StartCharNum).Should(Equal(1))
          })
        It("has correct end char number", func() {
          Expect(slc.EndCharNum).Should(Equal(22))
        })
        It("has correct text", func() {
          Expect(slc.Text).Should(Equal("// single line comment"))
        })
        It("has correct parent block", func() {
          Expect(slc.ParentBlock).Should(Equal(p.Content))
        })
        It("has corrent child blocks (none)", func() {
          Expect(slc.ChildBlocks).Should(BeZero())
        })
      })
      Context("method", func() {
        var m *Block // m -> method
        BeforeEach(func() {
          m = p.Content.ChildBlocks[6]
        })
        It("has correct Kind", func() {
          Expect(m.Kind).Should(Equal("CODE"))
        })
        It("has correct start line number", func() {
          Expect(m.StartLineNum).Should(Equal(10))
        })
        It("has correct end line number", func() {
          Expect(m.EndLineNum).Should(Equal(11))
        })
        It("has correct start char number", func() {
          Expect(m.StartCharNum).Should(Equal(1))
          })
        It("has correct end char number", func() {
          Expect(m.EndCharNum).Should(Equal(34))
        })
        It("has correct text", func() {
          Expect(m.Text).Should(Equal("public int testMethod(int var1, double var2, string var3,\n                         foo var4)\n"))
        })
        It("has correct parent block", func() {
          Expect(m.ParentBlock).Should(Equal(p.Content))
        })
        It("has child blocks (not none)", func() {
          Expect(m.ChildBlocks).ShouldNot(BeZero())
        })
        It("has 1 child block", func() {
          Expect(len(m.ChildBlocks)).Should(Equal(1))
        })
        
      })

    })
  })


})

// TEST2(&m)
func TEST2(b **Block) {
  It("TEST", func() {
    t := *b
    Expect(t).Should(BeAssignableToTypeOf(&Block{}))
    Expect(len(t.ChildBlocks)).Should(Equal(1))
  })
}



var multicomment_1 = strings.Join([]string{
      "/* multi line comment",
      "* another line ; \\",
      "ending **/",
    }, "\n")

var method_1 = strings.Join([]string{
      "public int testMethod(int var1, double var2, string var3,",
      "                         foo var4)",
      "{",
      "   return var1 * (var2- var3)/var1 + 3;",
      "}", 
    }, "\n")

var java_slice = []string{
      "package com.scribe.oauth;",
      multicomment_1,
      "",
      "import com.foo.bar;",
      "import static com.foo.baz.*; //comment after code end",
      "// single line comment",
      "",
      method_1,
    }

var java_code = strings.Join(java_slice, "\n")