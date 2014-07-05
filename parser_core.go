package main

type Parser struct {
  cursor_env string
  match bool
  line_num int
  char_num int
  last_line_len int
  line string
  block string
  indicator string
  blocks []*Item
  current_block *Item
  parsers []func(string)
}
 

type Item struct {
  kind string
  reference string
  Text string

  StartLineNum int
  StartCharNum int
  EndLineNum int
  EndCharNum int

  ParentItem *Item
  ChildItems []*Item

}

func NewParser() *Parser {
  p := Parser{}
  p.line_num = 1
  p.char_num = 0
  p.current_block = &Item{kind: "CORE"}
  p.blocks = []*Item{p.current_block}
  p.NewWhiteSpace()
  parsers := []func(string){p.IncrementCursor}
  p.RegisterParsers(parsers)
  return &p
}

func (p *Parser) RegisterParsers(parsers []func(string)) {
  p.parsers = append(p.parsers, parsers...)
}


func (p *Parser) Parse(s string) {
  for _,a := range s {
    p.match = false
    for _,pars := range p.parsers {
      pars(string(a))
      if p.match==true { break }
    }
  }
  
}

func (p *Parser) IncrementCursor(s string) {
  // increment line number and cursor position
  switch s {
  case "\n":
    p.line_num++
    p.last_line_len = p.char_num
    p.char_num = 0
  default:
    p.char_num++
  }
}

func (p *Parser) NewBlock(kind string) {
  i := &Item{}
  i.kind = kind
  i.ParentItem = p.current_block
  p.current_block.ChildItems = append(p.current_block.ChildItems, i)
  p.current_block = i
  p.cursor_env = kind
  i.StartLineNum = p.line_num
  i.StartCharNum = p.char_num
}

func (p *Parser) EndBlock() {
  p.current_block.EndLineNum = p.line_num
  p.current_block.EndCharNum = p.char_num
  p.current_block = p.current_block.ParentItem
  p.cursor_env = p.current_block.kind
}


func (p *Parser) AddChar(s string) {
  p.current_block.Text += p.indicator + s
  p.indicator = ""
  p.match = true
}

func (p *Parser) NewWhiteSpace() {
  p.NewBlock("WHITESPACE")
}


func (p *Parser) EndWhiteSpace() {
  p.EndBlock()
}


func (p *Parser) NewSingleComment() {
  p.NewBlock("SINGLE COMMENT")
  p.current_block.StartCharNum += -1
  p.AddChar("/")
}


func (p *Parser) EndSingleComment() {
  b := p.current_block
  p.EndBlock()
  b.EndCharNum = p.last_line_len
  b.EndLineNum += -1
}


func (p *Parser) NewMultiComment() {
  p.NewBlock("MULTI COMMENT")
  p.current_block.StartCharNum += -1
  p.AddChar("*")
}


func (p *Parser) EndMultiComment() {
  p.AddChar("/")
  p.EndBlock()
}


func (p *Parser) NewCode() {
  p.NewBlock("CODE")
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
  p.NewBlock("CHAR")
  p.AddChar("'")
}


func (p *Parser) EndChar() {
  p.AddChar("'")
  p.EndBlock()
}


func (p *Parser) NewString() {
  p.NewBlock("STRING")
  p.AddChar("\"")
}


func (p *Parser) EndString() {
  p.AddChar("\"")
  p.EndBlock()
}