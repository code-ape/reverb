package parser2

type Token struct {
	Text        string
	TextIndex   int
	LineIndex   int
	ColumnIndex int	
	Type        string
}

func NewToken(text string, ti, li, ci int, tp string) *Token {
	return &Token{text, ti, li, ci, tp}
}