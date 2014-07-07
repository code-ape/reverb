package parser

type Parser struct {
	CursorEnv     string
	Match         bool
	LineNum       int
	CharNum       int
	LastLineLen   int
	line          string
	block         string
	indicator     string
	Content       *Item
	current_block *Item
	parsers       []func(string)
}

type Item struct {
	Kind      string
	reference string
	Text      string

	StartLineNum int
	StartCharNum int
	EndLineNum   int
	EndCharNum   int

	ParentItem *Item
	ChildItems []*Item
}

func NewParser() *Parser {
	p := Parser{}
	p.LineNum = 1
	p.CharNum = 0
	p.current_block = &Item{Kind: "WHITESPACE"}
	p.Content = p.current_block
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

func (p *Parser) NewBlock(kind string) *Item {
	i := &Item{}
	i.Kind = kind
	i.StartLineNum = p.LineNum
	i.StartCharNum = p.CharNum
	return i
}

func (p *Parser) AddBlock(kind string) {
	i := p.NewBlock(kind)
	i.ParentItem = p.current_block
	p.current_block.ChildItems = append(p.current_block.ChildItems, i)
	p.current_block = i
	p.CursorEnv = kind
}

func (p *Parser) EndBlock() {
	p.current_block.EndLineNum = p.LineNum
	p.current_block.EndCharNum = p.CharNum
	p.current_block = p.current_block.ParentItem
	p.CursorEnv = p.current_block.Kind
}

func (p *Parser) AddChar(s string) {
	p.current_block.Text += p.indicator + s
	p.indicator = ""
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
	p.current_block.StartCharNum += -1
	p.AddChar("/")
}

func (p *Parser) EndSingleComment() {
	b := p.current_block
	p.EndBlock()
	b.EndCharNum = p.LastLineLen
	b.EndLineNum += -1
}

func (p *Parser) NewMultiComment() {
	p.AddBlock("MULTI COMMENT")
	p.current_block.StartCharNum += -1
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
