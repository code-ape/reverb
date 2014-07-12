package parser

type Parser struct {
	CursorEnv     string
	Match         bool
	LineNum       int
	CharNum       int
	LastLineLen   int
	line          string
	block         string
	Indicator     string
	Content       *Block
	CurrentBlock  *Block
	parsers       []func(string)
}



func NewParser() *Parser {
	p := Parser{}
	p.LineNum = 1
	p.CharNum = 0
	p.CurrentBlock = &Block{}
  p.CurrentBlock.Kind = "WHITESPACE"
  p.Content = p.CurrentBlock
  p.CursorEnv = p.CurrentBlock.Kind
	parsers := []func(string){p.IncrementCursor}
	p.RegisterParsers(parsers)
	return &p
}

func (p *Parser) RegisterParsers(parsers []func(string)) {
	p.parsers = append(p.parsers, parsers...)
}

func (p *Parser) Parse(s string) {
	for _, a := range s {
		p.Match = false
		for _, pars := range p.parsers {
			pars(string(a))
			if p.Match == true {
				break
			}
		}
	}

}

func (p *Parser) IncrementCursor(s string) {
	// increment line number and cursor position
	switch s {
	case "\n":
		p.LineNum++
		p.LastLineLen = p.CharNum
		p.CharNum = 0
	default:
		p.CharNum++
	}
}

func (p *Parser) NewBlock(kind string) *Block {
	i := &Block{}
	i.Kind = kind
	i.StartLineNum = p.LineNum
	i.StartCharNum = p.CharNum
	return i
}

func (p *Parser) AddBlock(kind string) {
	i := p.NewBlock(kind)
	i.ParentBlock = p.CurrentBlock
	p.CurrentBlock.ChildBlocks = append(p.CurrentBlock.ChildBlocks, i)
	p.CurrentBlock = i
	p.CursorEnv = kind
}

func (p *Parser) EndBlock() {
	p.CurrentBlock.EndLineNum = p.LineNum
	p.CurrentBlock.EndCharNum = p.CharNum
	p.CurrentBlock = p.CurrentBlock.ParentBlock
	p.CursorEnv = p.CurrentBlock.Kind
}

func (p *Parser) AddChar(s string) {
	p.CurrentBlock.Text += p.Indicator + s
	p.Indicator = ""
	p.Match = true
}

func (p *Parser) NewWhiteSpace() {
	p.AddBlock("WHITESPACE")
}

func (p *Parser) EndWhiteSpace() {
	p.EndBlock()
}

func (p *Parser) NewSingleComment() {
	p.AddBlock("SINGLE COMMENT")
	p.CurrentBlock.StartCharNum += -1
	p.AddChar("/")
}

func (p *Parser) EndSingleComment() {
	b := p.CurrentBlock
	p.EndBlock()
	b.EndCharNum = p.LastLineLen
	b.EndLineNum += -1
}

func (p *Parser) NewMultiComment() {
	p.AddBlock("MULTI COMMENT")
	p.CurrentBlock.StartCharNum += -1
	p.AddChar("*")
}

func (p *Parser) EndMultiComment() {
	p.AddChar("/")
	p.EndBlock()
}

func (p *Parser) NewCode() {
	p.AddBlock("CODE")
}

func (p *Parser) EndCode() {
	p.AddChar(";")
	p.EndBlock()
}

func (p *Parser) NewBody() {
  endChar := p.CharNum - 1
  p.CurrentBlock.EndLineNum = p.LineNum
  if endChar == 0 {
    p.CurrentBlock.EndLineNum += -1
    p.CurrentBlock.EndCharNum = p.LastLineLen
  } else {
    p.CurrentBlock.EndCharNum = endChar
  }
	p.NewWhiteSpace()
}

func (p *Parser) EndBody() {
	p.EndWhiteSpace()
}

func (p *Parser) NewChar() {
	p.AddBlock("CHAR")
	p.AddChar("'")
}

func (p *Parser) EndChar() {
	p.AddChar("'")
	p.EndBlock()
}

func (p *Parser) NewString() {
	p.AddBlock("STRING")
	p.AddChar("\"")
}

func (p *Parser) EndString() {
	p.AddChar("\"")
	p.EndBlock()
}
