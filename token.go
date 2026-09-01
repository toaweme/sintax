package sintax

import "strings"

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
	LoopExpr() string
}

// BaseToken is the concrete Token implementation shared by every token kind.
type BaseToken struct {
	TokenType TokenType
	RawValue  string
	Var       string
	ParamVars []string
	// LoopExprValue holds the iteration expression for ForToken (e.g. "groups",
	// "items | filter:'a','b'"). For ForToken, Var holds the loop variable name
	// (e.g. "tx") and LoopExprValue holds the right-hand-side expression.
	LoopExprValue string
	// SourceValue is the template this token was parsed from, and OffsetValue is
	// the byte index of its opening delimiter within it. Both are zero on a token
	// the engine built rather than parsed, such as the expression inside an if or
	// a for tag. The source travels on the token because the template modifier
	// re-enters the engine on a different one, so a position resolved against the
	// outer document would name the wrong template.
	SourceValue string
	OffsetValue int
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

// LoopExpr returns the iteration expression for a ForToken.
func (bt BaseToken) LoopExpr() string { return bt.LoopExprValue }

// position reads where this token was written, and reports an unknown position
// for a token that carries no source.
func (bt BaseToken) position() Position {
	if bt.SourceValue == "" {
		return Position{}
	}

	return positionAt(bt.SourceValue, bt.OffsetValue)
}

// positionAt turns a byte offset into a line, a column and the line's own text.
// Both counts start at 1, and an offset outside source has no position.
func positionAt(source string, offset int) Position {
	if offset < 0 || offset > len(source) {
		return Position{}
	}

	start := strings.LastIndexByte(source[:offset], '\n') + 1
	end := strings.IndexByte(source[start:], '\n')
	if end < 0 {
		end = len(source)
	} else {
		end += start
	}

	return Position{
		Line:   strings.Count(source[:start], "\n") + 1,
		Column: offset - start + 1,
		Text:   strings.TrimSuffix(source[start:end], "\r"),
	}
}
