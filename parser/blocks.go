package parser

type Block struct {
	Item
	ParentBlock *Block
	ChildBlocks []*Block
}

func (b *Block) GetChildren() []*Block {
	if len(b.ChildBlocks) == 1 {
		if b.ChildBlocks[0].Kind == "WHITESPACE" {
			return b.ChildBlocks[0].ChildBlocks
		}
	}
	return b.ChildBlocks
}

type Item struct {
	Kind string
	Text string

	StartLineNum int
	StartCharNum int
	EndLineNum   int
	EndCharNum   int
}

type Package struct {
	Block
	Static      bool
	PackageName string
}

type Variable struct {
	Name        string
	Type        string
	Declaration *VariableDeclaration
	// Scope ???
}

type VariableDeclaration struct {
	Block
	Variable   *Variable
	Assignemnt *ValueAssignment
}

type ValueAssignment struct {
	Block
	Target       *Variable
	Dependencies *Variable
}

type MethodDeclaration struct {
	Block
	Modifier       string // public, private
	ReturnType     string
	Name           string
	Arguments      *[]Variable
	MethodVariable *Variable
}

type MethodCall struct {
	Declaration *MethodDeclaration
	Arguments   *[]Variable
}
