package parser2

import (
	"fmt"
	"bufio"
	"io"
	"os"
)

type Scanner struct {
	scanner     *bufio.Scanner
	End         bool
	TextIndex   int
	LineIndex   int
	ColumnIndex int
	CurrentChar *Character
}


func NewScanner(r io.Reader) *Scanner {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)
	s := &Scanner{}
	s.scanner = scanner
	s.End = false
	s.TextIndex = 0
	s.LineIndex = 1
	s.ColumnIndex = 0
	s.CurrentChar = nil
	return s
}


func (s *Scanner) Next() bool {

	if more := s.scanner.Scan(); more == false {
		s.End = true
		if err := s.scanner.Err(); err != nil {
    	fmt.Fprintln(os.Stderr, "reading input:", err)
    }
    return false
	}

	c := s.scanner.Text()
	
	s.TextIndex ++

	if s.CurrentChar == nil {
		s.ColumnIndex ++
	} else {
		if s.CurrentChar.Char == "\n" {
			s.LineIndex ++
			s.ColumnIndex = 1
		} else {
			s.ColumnIndex ++
		}
	}

	char := NewCharacter(c, s.TextIndex, s.LineIndex, s.ColumnIndex)
	s.CurrentChar = char

	return true
}

func (s *Scanner) Char() *Character {
	return s.CurrentChar
}

