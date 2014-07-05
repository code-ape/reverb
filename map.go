package main

import (
  "fmt"
  "os"
  "bufio"
  "github.com/code-ape/reverb/parser"
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
  //Parse, PrintImports := Parser()

  scanner := bufio.NewScanner(fi)
  scanner.Split(bufio.ScanRunes)

  p := parser.NewJavaParser()

  for scanner.Scan() {
    p.Parse(scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    fmt.Println("ERROR WITH SCANNER", scanner.Err())
  }

  p.PrintBlocks()

}