package sintax

// TokenType identifies the syntactic kind of a parsed Token.
type TokenType int

// The token kinds produced by Parser.Parse.
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

// Token is a single parsed unit of a template, such as a text run, a
// variable reference, or a control-flow marker.
type Token interface {
	Type() TokenType
	Raw() string
	Name() string
	Params() []string
	WithDefault() bool
	LoopExpr() string
}

// BaseToken is the concrete Token implementation shared by every token kind.
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
	// parsedVar and parsedFuncs cache the result of getVarAndFunctions for
	// FilteredVariableToken, computed once at parse time. renderVariable would
	// otherwise re-split and re-classify RawValue on every render, which
	// dominated allocations on modifier-heavy templates. nil parsedFuncs means
	// "not cached" and the renderer falls back to parsing on demand.
	parsedVar   string
	parsedFuncs []Func
}

// Type returns the token's kind.
func (bt BaseToken) Type() TokenType { return bt.TokenType }

// Raw returns the token's original, unparsed source text.
func (bt BaseToken) Raw() string { return bt.RawValue }

// Name returns the variable or loop variable name referenced by the token.
func (bt BaseToken) Name() string { return bt.Var }

// Params returns the token's raw parameter strings, if any.
func (bt BaseToken) Params() []string { return bt.ParamVars }

// WithDefault reports whether the token has a default-value fallback.
func (bt BaseToken) WithDefault() bool { return bt.HasDefault }

// LoopExpr returns the iteration expression for a ForToken.
func (bt BaseToken) LoopExpr() string { return bt.LoopExprValue }
