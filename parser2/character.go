package parser2


type Character struct {
	Char        string
	TextIndex   int
	LineIndex   int
	ColumnIndex int
}


func NewCharacter(c string, ti, li, ci int) *Character{
	return &Character{Char: c, TextIndex: ti, LineIndex: li, ColumnIndex: ci}
}