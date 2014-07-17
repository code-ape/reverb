package token

type Tokenizer struct {
	Kind     string
	Language string

	Matchers []Matcher
}