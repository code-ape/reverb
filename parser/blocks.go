package parser

type Block struct {
  Item
  ParentBlock *Block
  ChildBlocks []*Block
}

type Item struct {
  Kind      string
  Text      string

  StartLineNum int
  StartCharNum int
  EndLineNum   int
  EndCharNum   int
}


type Variable struct {
  Name string
  Type string
  Declaration *VariableDeclaration
  // Scope ???
}


type VariableDeclaration struct {
  Item
  Variable *Variable
  Assignemnt *ValueAssignment
}


type ValueAssignment struct {
  Item
  Target *Variable
  Dependencies *Variable
}


type MethodDeclaration struct {
  Item
  Modifier string // public, private
  ReturnType string
  Name string
  Arguments *[]Variable
  MethodVariable *Variable
}


type MethodCall struct {
  Declaration *MethodDeclaration
  Arguments *[]Variable
}