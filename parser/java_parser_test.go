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
      Context("multiline comment", func() {
        var mc *Item
        BeforeEach(func() {
          mc = p.Content.ChildItems[1]
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
        It("has correct parent item", func() {
          Expect(mc.ParentItem).Should(Equal(p.Content))
        })
      })
    
    })

  })



})


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