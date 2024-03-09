package sintax

type TokenType int

const (
	UndefinedToken TokenType = iota - 1
	TextToken
	VariableToken
	FilteredVariableToken
	IfToken
	ElseToken
	IfEndToken
	ShorthandIfToken
)

type Token interface {
	Type() TokenType
	Raw() string
}

type BaseToken struct {
	TokenType TokenType
	RawValue  string
}

func (bt BaseToken) Type() TokenType { return bt.TokenType }
func (bt BaseToken) Raw() string     { return bt.RawValue }
