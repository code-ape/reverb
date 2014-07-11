package parser

import "fmt"

func NewJavaParser() *Parser {
	p := NewParser()
	parsers := []func(string){p.ParseWhiteSpace, p.ParseContent, p.ParseCode, p.ParseString,
		p.ParseChar, p.ParseSingleComment, p.ParseMultiComment}
	p.RegisterParsers(parsers)
	return p
}

func (p *Parser) ParseWhiteSpace(s string) {
	if p.CursorEnv != "WHITESPACE" {
		return
	}
	if s == "\n" || s == "\t" || s == "\r" || s == " " {
		p.Match = true
	} else {
		p.CursorEnv = "CONTENT"
	}
}

func (p *Parser) ParseContent(s string) {
	if p.CursorEnv != "CONTENT" && p.CursorEnv != "CODE" {
		return
	}
	switch p.Indicator {
	case "/":
		switch s {
		case "/":
			p.NewSingleComment()
		case "*":
			p.NewMultiComment()
		default:
			if p.CursorEnv == "CONTENT" {
				p.NewCode()
			}
		}
	default:
		switch s {
		case "/":
			p.Indicator = "/"
		default:
			if p.CursorEnv == "CONTENT" {
				p.NewCode()
			}
		}
	}
}

func (p *Parser) ParseCode(s string) {
	if p.CursorEnv != "CODE" {
		return
	}

	switch s {
	case "/":
		p.Indicator = "/"
	case "'":
		p.NewChar()
	case "\"":
		p.NewString()
	case "{":
		p.NewBody()
	case "}":
		p.EndBody()
	case ";":
		p.EndCode()
	default:
		p.AddChar(s)
	}
}

func (p *Parser) ParseSingleComment(s string) {
	if p.CursorEnv != "SINGLE COMMENT" {
		return
	}
	switch s {
	case "\n":
		p.EndSingleComment()
	default:
		p.AddChar(s)
	}
}

func (p *Parser) ParseMultiComment(s string) {
	if p.CursorEnv != "MULTI COMMENT" {
		return
	}

	switch p.Indicator {
	case "*":
		switch s {
		case "/":
			p.EndMultiComment()
    case "*":
      p.AddChar("")
      p.Indicator = "*"
		default:
			p.AddChar(s)
		}
	default:
		switch s {
		case "*":
			p.Indicator = "*"
		default:
			p.AddChar(s)
		}
	}
}

func (p *Parser) ParseString(s string) {
	if p.CursorEnv != "STRING" {
		return
	}
	switch p.Indicator {
	case "\\":
		p.AddChar(s)
	default:
		switch s {
		case "\\":
			p.Indicator = "\\"
		case "\"":
			p.EndString()
		default:
			p.AddChar(s)
		}
	}
}

func (p *Parser) ParseChar(s string) {
	if p.CursorEnv != "CHAR" {
		return
	}
	switch p.Indicator {
	case "\\":
		p.AddChar(s)
	default:
		switch s {
		case "\\":
			p.Indicator = "\\"
		case "'":
			p.EndChar()
		default:
			p.AddChar(s)
		}
	}
}

func (p *Parser) PrintBlocks() {
	PrintBlockLoop([]*Item{p.Content})
}

func PrintBlockLoop(blocks []*Item) {
	for _, b := range blocks {
		fmt.Println("Kind:", b.Kind)
		fmt.Println("Text:", b.Text)
		fmt.Println("CHILDREN:")
		PrintBlockLoop(b.ChildItems)
		fmt.Println("END")
	}
}
