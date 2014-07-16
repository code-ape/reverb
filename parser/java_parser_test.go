package parser_test

import (
	. "github.com/code-ape/reverb/parser"
  . "github.com/code-ape/reverb/factory"

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

  var p *Parser = &Parser{}
  p.Content = &Block{}

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
        f := MakeTestFactory(&p)
        m := f.NewMap()
        f.BeforeEach(func() {
          m["childnum"] = 0
          m["kind"] = "CODE"
          m["startline"] = 1
          m["endline"] = 1
          m["startchar"] = 1
          m["endchar"] = 25
          m["text"] = "package com.scribe.oauth;"
          m["parent"] = p.Content
          m["children"] = BeZero()
        })
        f.RunAll(m)
      })

      Context("multiline comment", func() {
        f := MakeTestFactory(&p)
        m := f.NewMap()
        f.BeforeEach(func() {
          m["childnum"] = 1
          m["kind"] = "MULTI COMMENT"
          m["startline"] = 2
          m["endline"] = 4
          m["startchar"] = 1
          m["endchar"] = 10
          m["text"] = multicomment_1
          m["parent"] = p.Content
          m["children"] = BeZero()
        })
        f.RunAll(m)
      })

      Context("regular import", func() {
        f := MakeTestFactory(&p)
        m := f.NewMap()
        f.BeforeEach(func() {
          m["childnum"] = 2
          m["kind"] = "CODE"
          m["startline"] = 6
          m["endline"] = 6
          m["startchar"] = 1
          m["endchar"] = 19
          m["text"] = "import com.foo.bar;"
          m["parent"] = p.Content
          m["children"] = BeZero()
        })
        f.RunAll(m)
      })

      Context("static star import", func() {
        f := MakeTestFactory(&p)
        m := f.NewMap()
        f.BeforeEach(func() {
          m["childnum"] = 3
          m["kind"] = "CODE"
          m["startline"] = 7
          m["endline"] = 7
          m["startchar"] = 1
          m["endchar"] = 28
          m["text"] = "import static com.foo.baz.*;"
          m["parent"] = p.Content
          m["children"] = BeZero()
        })
        f.RunAll(m)
      })

      Context("single line comment after code", func() {
        f := MakeTestFactory(&p)
        m := f.NewMap()
        f.BeforeEach(func() {
          m["childnum"] = 4
          m["kind"] = "SINGLE COMMENT"
          m["startline"] = 7
          m["endline"] = 7
          m["startchar"] = 30
          m["endchar"] = 53
          m["text"] = "//comment after code end"
          m["parent"] = p.Content
          m["children"] = BeZero()
        })
        f.RunAll(m)
      })

      Context("single line comment on own line", func() {
        f := MakeTestFactory(&p)
        m := f.NewMap()
        f.BeforeEach(func() {
          m["childnum"] = 5
          m["kind"] = "SINGLE COMMENT"
          m["startline"] = 8
          m["endline"] = 8
          m["startchar"] = 1
          m["endchar"] = 22
          m["text"] = "// single line comment"
          m["parent"] = p.Content
          m["children"] = BeZero()
        })
        f.RunAll(m)
      })

      Context("method", func() {
        f := MakeTestFactory(&p)
        m := f.NewMap()
        f.BeforeEach(func() {
          m["childnum"] = 6
          m["kind"] = "CODE"
          m["startline"] = 10
          m["endline"] = 14
          m["startchar"] = 1
          m["endchar"] = 1
          m["text"] = "public int testMethod(int var1, double var2, string var3,\n                         foo var4)\n"
          m["parent"] = p.Content
          m["children"] = HaveLen(1)
        })
        f.RunAll(m)
      })

      Context("whitespace in method", func() {
        f := MakeMethodTestFactory(&p, 6)
        m := f.NewMap()
        f.BeforeEach(func() {
          m["childnum"] = 0
          m["kind"] = "WHITESPACE"
          m["startline"] = 12
          m["endline"] = 14
          m["startchar"] = 1
          m["endchar"] = 1
          m["text"] = ""
          m["parent"] = p.Content.ChildBlocks[6]
          m["children"] = HaveLen(1)
        })
        f.RunAll(m)
      })

      Context("class", func() {
        f := MakeTestFactory(&p)
        m := f.NewMap()
        f.BeforeEach(func() {
          m["childnum"] = 7
          m["kind"] = "CODE"
          m["startline"] = 15
          m["endline"] = 21
          m["startchar"] = 1
          m["endchar"] = 1
          m["text"] = "public class OAuth10aServiceImpl implements OAuthService "
          m["parent"] = p.Content
          m["children"] = HaveLen(2)
        })
        f.RunAll(m)
      })

      Context("char declaration", func() {
        f := MakeTestFactory(&p)
        m := f.NewMap()
        f.BeforeEach(func() {
          m["childnum"] = 8
          m["kind"] = "CODE"
          m["startline"] = 22
          m["endline"] = 22
          m["startchar"] = 1
          m["endchar"] = 14
          m["text"] = "char ch = ;"
          m["parent"] = p.Content
          m["children"] = HaveLen(1)
        })
        f.RunAll(m)
      })

      Context("char", func() {
        f := MakeMethodTestFactory(&p, 8)
        m := f.NewMap()
        f.BeforeEach(func() {
          m["childnum"] = 0
          m["kind"] = "CHAR"
          m["startline"] = 22
          m["endline"] = 22
          m["startchar"] = 11
          m["endchar"] = 13
          m["text"] = "'a'"
          m["parent"] = p.Content.ChildBlocks[8]
          m["children"] = BeZero()
        })
        f.RunAll(m)
      })

    })
  })


})


func MakeTestFactory(pp **Parser) *Factory {

  f := NewTestFactory()

  var b *Block // b -> package declaration

  f.JustBeforeEach(func() {
    p := *pp
    b = p.Content.ChildBlocks[f.Int("childnum")]
  })
  f.It("kind", "has correct Kind", func() {
    Expect(b.Kind).Should(Equal(f.String("kind")))
  })
  f.It("startline", "has correct start line number", func() {
    Expect(b.StartLineNum).Should(Equal(f.Int("startline")))
  })
  f.It("endline", "has correct end line number", func() {
    Expect(b.EndLineNum).Should(Equal(f.Int("endline")))
  })
  f.It("startchar", "has correct start char number", func() {
    Expect(b.StartCharNum).Should(Equal(f.Int("startchar")))
  })
  f.It("endchar", "has correct end char number", func() {
    Expect(b.EndCharNum).Should(Equal(f.Int("endchar")))
  })
  f.It("text", "has correct text", func() {
    Expect(b.Text).Should(Equal(f.String("text")))
  })
  f.It("parent", "has correct parent block", func() {
    Expect(b.ParentBlock).Should(Equal( f.Val("parent").(*Block) ))
  })
  f.It("children", "has corrent child blocks (none)", func() {
    Expect(b.GetChildren()).Should(f.Matcher("children"))
  })

  return f
}


func MakeMethodTestFactory(pp **Parser, method_index int) *Factory {
  f := NewTestFactory()

  var b *Block // b -> package declaration

  f.JustBeforeEach(func() {
    p := *pp
    b = p.Content.ChildBlocks[method_index].ChildBlocks[f.Int("childnum")]
  })
  f.It("kind", "has correct Kind", func() {
    Expect(b.Kind).Should(Equal(f.String("kind")))
  })
  f.It("startline", "has correct start line number", func() {
    Expect(b.StartLineNum).Should(Equal(f.Int("startline")))
  })
  f.It("endline", "has correct end line number", func() {
    Expect(b.EndLineNum).Should(Equal(f.Int("endline")))
  })
  f.It("startchar", "has correct start char number", func() {
    Expect(b.StartCharNum).Should(Equal(f.Int("startchar")))
  })
  f.It("endchar", "has correct end char number", func() {
    Expect(b.EndCharNum).Should(Equal(f.Int("endchar")))
  })
  f.It("text", "has correct text", func() {
    Expect(b.Text).Should(Equal(f.String("text")))
  })
  f.It("parent", "has correct parent block", func() {
    Expect(b.ParentBlock).Should(Equal( f.Val("parent").(*Block) ))
  })
  f.It("children", "has corrent child blocks (none)", func() {
    Expect(b.GetChildren()).Should(f.Matcher("children"))
  })

  return f
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

var class_1 = strings.Join([]string{
      "public class OAuth10aServiceImpl implements OAuthService {",
      "   private static final String VERSION = \"1.0\";",
      "   public Token getRequestToken(int timeout, TimeUnit unit)",
      "   {",
      "   return getRequestToken(new TimeoutTuner(timeout, unit));",
      "   }",
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
      class_1,
      "char ch = 'a';",
    }

var java_code = strings.Join(java_slice, "\n")