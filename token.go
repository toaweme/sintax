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
	ForToken
	ForEndToken
)

type Token interface {
	Type() TokenType
	Raw() string
	Name() string
	Params() []string
	WithDefault() bool
	LoopExpr() string
}

type BaseToken struct {
	TokenType  TokenType
	RawValue   string
	Var        string
	ParamVars  []string
	HasDefault bool
	// LoopExprValue holds the iteration expression for ForToken (e.g. "groups",
	// "items | filter:'a','b'"). For ForToken, Var holds the loop variable name
	// (e.g. "tx") and LoopExprValue holds the right-hand-side expression.
	LoopExprValue string
}

func (bt BaseToken) Type() TokenType   { return bt.TokenType }
func (bt BaseToken) Raw() string       { return bt.RawValue }
func (bt BaseToken) Name() string      { return bt.Var }
func (bt BaseToken) Params() []string  { return bt.ParamVars }
func (bt BaseToken) WithDefault() bool { return bt.HasDefault }
func (bt BaseToken) LoopExpr() string  { return bt.LoopExprValue }
