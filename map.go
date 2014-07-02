package main

import (
  "fmt"
  "os"
  "io"
)

func MapDependencies(target_file string) {
  // open input file
  fi, err := os.Open(target_file)
  if err != nil { panic(err) }
  // close fi on exit and check for its returned error
  defer func() {
    if err := fi.Close(); err != nil {
        panic(err)
    }
  }()

  // make a buffer to keep chunks that are read
  Parse, PrintImports := Parser()
  buf := make([]byte, 1024)
  for {
    // read a chunk
    n, err := fi.Read(buf)
    if err != nil && err != io.EOF { panic(err) }
    if n == 0 { break }

    fmt.Println(string(buf[:n]))
    Parse(string(buf[:n]))

    // write a chunk
    //if _, err := fo.Write(buf[:n]); err != nil {
    //  panic(err)
    //}
  }

  PrintImports()


}

func Parser() (func(string), func()) {
  status := "CODE"
  substatus := ""
  blocks := make([]string, 1)
  //imports := make([]string, 1)

  AddChar := func(c string) {
    blocks[len(blocks)-1] = blocks[len(blocks)-1] + string(c)
  }

  NewBlock := func(init_string string) {
    blocks = append(blocks, init_string)
  }

  Parse := func(code_chunk string) {
    for _,a := range code_chunk {
      b := string(a)
      switch status {
      case "CODE":
        switch substatus {
        case "/":
          switch b {
          case "/":
            status = "SINGLE COMMENT"
            substatus = ""
            NewBlock("//")

          case "*":
            status = "MULTI COMMENT"
            substatus = ""
            NewBlock("/*")
          }

        default:
          switch b {
          case "/":
            substatus = "/"

          case "'":
            status = "CHAR"
            substatus = ""
            NewBlock("'")

          case "\"":
            status = "STRING"
            substatus = ""
            NewBlock("\"")
          default:
            substatus = ""
            AddChar(b)
            
          }
        }

      case "SINGLE COMMENT":
        switch b {
        case "\n":
          status = "CODE"
          substatus = ""
          NewBlock("")
        default:
          AddChar(b)
        }

      case "MULTI COMMENT":
        switch substatus {
        case "*":
          switch b {
          case "/":
            status = "CODE"
            substatus = ""
            AddChar("*/")
            NewBlock("")
          default:
            substatus = ""
            AddChar("*" + b)
          }
        default:
          switch b {
          case "*":
            substatus = "*"
          default:
            AddChar(b)
          }
        }

      case "STRING":
        switch substatus {
        case "\\":
          AddChar(b)
        default:
          switch b {
          case "\\":
            substatus = "\\"
          case "\"":
            status = "CODE"
            substatus = ""
            AddChar("\"")
            NewBlock("")
          default:
            substatus = ""
            AddChar(b)
          }
        }

      case "CHAR":
        switch substatus {
        case "\\":
          AddChar(b)
        default:
          switch b {
          case "\\":
            substatus = "\\"
          case "'":
            status = "CODE"
            substatus = ""
            NewBlock("'")
          }
        } 

      }
    }
  }
  PrintImports := func() {
    for _,b := range blocks {
      fmt.Println("{{{",b, "}}}")
    }
  }
  return Parse, PrintImports
}


