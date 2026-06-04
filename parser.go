package sintax

import (
	"regexp"
	"strings"
)

type StringParser struct {
	opener string
	closer string
}

func NewStringParser() *StringParser {
	return &StringParser{
		opener: "{{",
		closer: "}}",
	}
}

var _ Parser = (*StringParser)(nil)

func (p *StringParser) Parse(template string) ([]Token, error) {
	var tokens []Token

	i := 0
	for {
		// find the next occurrence of `opener`
		openerIndex := strings.Index(template[i:], p.opener)
		if openerIndex == -1 {
			// no more `opener` in the substring; treat the rest as text
			if i < len(template) {
				tokens = append(tokens, BaseToken{
					TokenType: TextToken,
					RawValue:  template[i:],
				})
			}
			break
		}

		// adjust to absolute index from the slice-based Index
		openerIndex += i

		// everything before opener is text
		if openerIndex > i {
			tokens = append(tokens, BaseToken{
				TokenType: TextToken,
				RawValue:  template[i:openerIndex],
			})
		}

		// find the next occurrence of `closer`, after the opener
		startOfInner := openerIndex + len(p.opener)
		closerIndex := strings.Index(template[startOfInner:], p.closer)
		if closerIndex == -1 {
			// no matching closer found: might treat the rest of the string as text
			tokens = append(tokens, BaseToken{
				TokenType: TextToken,
				RawValue:  template[openerIndex:], // everything from opener
			})
			break
		}

		// adjust to absolute index
		closerIndex += startOfInner

		// extract the substring (contents) between opener and closer
		contents := template[startOfInner:closerIndex]

		// detect Jinja-style trim markers: {{- expr -}}
		trimLeft := strings.HasPrefix(contents, "-")
		trimRight := strings.HasSuffix(contents, "-")
		if trimLeft {
			contents = contents[1:]
		}
		if trimRight && len(contents) > 0 {
			contents = contents[:len(contents)-1]
		}

		// {{- }}: strip trailing whitespace from previous text token (incl. newlines)
		if trimLeft {
			stripPrevTextRight(tokens, true)
		}

		// create the appropriate token
		tokenType := p.detectTokenType(contents)
		tokens = append(tokens, p.createToken(tokenType, contents))

		// move `i` beyond the closer
		i = closerIndex + len(p.closer)

		// {{ -}}: strip leading whitespace from following text token (incl. newlines).
		// we do this by advancing `i` past any whitespace + optional newlines.
		if trimRight {
			for i < len(template) {
				c := template[i]
				if c == ' ' || c == '\t' || c == '\r' || c == '\n' {
					i++
					continue
				}
				break
			}
		}
	}

	// post-pass: auto-trim whitespace around control tags that sit alone on a line.
	tokens = autoTrimBlockLines(tokens)

	return tokens, nil
}

// stripPrevTextRight strips trailing whitespace from the last token if it is
// a TextToken. when `includeNewlines` is true, newlines are also stripped.
func stripPrevTextRight(tokens []Token, includeNewlines bool) {
	if len(tokens) == 0 {
		return
	}
	last := tokens[len(tokens)-1]
	if last.Type() != TextToken {
		return
	}
	bt, ok := last.(BaseToken)
	if !ok {
		return
	}
	cut := func(r rune) bool {
		if r == ' ' || r == '\t' {
			return true
		}
		if includeNewlines && (r == '\n' || r == '\r') {
			return true
		}
		return false
	}
	bt.RawValue = strings.TrimRightFunc(bt.RawValue, cut)
	bt.Var = bt.RawValue
	tokens[len(tokens)-1] = bt
}

// autoTrimBlockLines removes the surrounding whitespace + newline for control
// tags (if/else/endif/for/endfor) that sit alone on their own line. specifically:
//   - if the preceding text token's tail (after the last \n) is all whitespace,
//     strip that trailing whitespace; and
//   - if the following text token starts with optional whitespace then \n,
//     strip up to and including that newline.
//
// the same rule applies at the start/end of the template (treating the
// template boundary as a virtual newline).
func autoTrimBlockLines(tokens []Token) []Token {
	if len(tokens) == 0 {
		return tokens
	}
	out := make([]Token, len(tokens))
	copy(out, tokens)

	isBlock := func(t Token) bool {
		switch t.Type() {
		case IfToken, ElseToken, IfEndToken, ForToken, ForEndToken:
			return true
		}
		return false
	}

	for i, tok := range out {
		if !isBlock(tok) {
			continue
		}

		// previous text tail check: at start-of-template (i==0), or the prev text
		// contains a '\n' followed only by whitespace until the end.
		prevOK := i == 0
		var prevBt BaseToken
		if i > 0 {
			prev := out[i-1]
			if prev.Type() == TextToken {
				if bt, ok := prev.(BaseToken); ok {
					tail := bt.RawValue
					nl := strings.LastIndexByte(tail, '\n')
					if nl >= 0 {
						rest := tail[nl+1:]
						allWS := true
						for _, c := range rest {
							if c != ' ' && c != '\t' && c != '\r' {
								allWS = false
								break
							}
						}
						if allWS {
							prevOK = true
							prevBt = bt
						}
					}
				}
			}
		}

		// next text head check
		nextOK := i == len(out)-1
		var nextBt BaseToken
		var nextIdx int
		var nextStripUntil int
		if i < len(out)-1 {
			next := out[i+1]
			if next.Type() == TextToken {
				if bt, ok := next.(BaseToken); ok {
					head := bt.RawValue
					j := 0
					for j < len(head) && (head[j] == ' ' || head[j] == '\t' || head[j] == '\r') {
						j++
					}
					if j < len(head) && head[j] == '\n' {
						nextOK = true
						nextBt = bt
						nextIdx = i + 1
						nextStripUntil = j + 1
					} else if j == len(head) {
						// trailing whitespace, no newline → end of template
						nextOK = true
						nextBt = bt
						nextIdx = i + 1
						nextStripUntil = j
					}
				}
			}
		}

		if !prevOK || !nextOK {
			continue
		}

		// apply: strip trailing whitespace from prev (back to last newline kept)
		if i > 0 && out[i-1].Type() == TextToken {
			tail := prevBt.RawValue
			nl := strings.LastIndexByte(tail, '\n')
			prevBt.RawValue = tail[:nl+1]
			prevBt.Var = prevBt.RawValue
			out[i-1] = prevBt
		}

		// strip leading whitespace + newline from next
		if i < len(out)-1 && out[i+1].Type() == TextToken {
			nextBt.RawValue = nextBt.RawValue[nextStripUntil:]
			nextBt.Var = nextBt.RawValue
			out[nextIdx] = nextBt
		}
	}

	return out
}

// variableNameRe matches a bare variable name (letters, digits, underscore and
// dots). Compiled once at package load: detectTokenType runs it against every
// token, so compiling per call dominated parse allocations.
var variableNameRe = regexp.MustCompile(`^([a-zA-Z_.][a-zA-Z0-9_.]*)$`)

// isVariable reports whether s is a bare variable name (as opposed to a quoted
// string, number, boolean, or filtered expression).
func (p *StringParser) isVariable(s string) bool {
	return variableNameRe.MatchString(s)
}

func (p *StringParser) detectTokenType(s string) TokenType {
	s = strings.TrimSpace(s)

	if s == "endif" {
		return IfEndToken
	} else if s == "endfor" {
		return ForEndToken
	} else if strings.HasPrefix(s, "for ") {
		return ForToken
	} else if strings.HasPrefix(s, "if ") || s == "if" {
		return IfToken
	} else if strings.HasPrefix(s, "else") {
		return ElseToken
	} else if strings.Contains(s, " ? ") && strings.Contains(s, " : ") {
		return ShorthandIfToken
	} else if p.isVariable(s) {
		return VariableToken
	} else if strings.Contains(s, "|") {
		return FilteredVariableToken
	}

	return UndefinedToken
}

func trimPrefix(s string, prefix string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, prefix) {
		return strings.TrimSpace(s[len(prefix):])
	}
	return s
}

func splitAndGetFirst(s string) string {
	parts := strings.Split(s, "|")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

func (p *StringParser) createToken(tokenType TokenType, value string) Token {
	switch tokenType {
	case VariableToken:
		return BaseToken{TokenType: VariableToken, RawValue: strings.TrimSpace(value), Var: strings.TrimSpace(value)}
	case FilteredVariableToken:
		token := BaseToken{
			TokenType: FilteredVariableToken,
			RawValue:  strings.TrimSpace(value),
			Var:       strings.TrimSpace(splitAndGetFirst(value)),
		}

		hasDefault := false
		// extract the modifier variables
		_, funcs := getVarAndFunctions(token)
		for _, f := range funcs {
			if f.Name == "default" {
				hasDefault = true
			}
			for _, p := range f.Args {
				if p.Var {
					if token.ParamVars == nil {
						token.ParamVars = make([]string, 0)
					}
					token.ParamVars = append(token.ParamVars, p.Value.(string))
				}
			}
		}
		token.HasDefault = hasDefault
		return token
	case IfToken:
		return BaseToken{TokenType: IfToken, RawValue: trimPrefix(value, "if")}
	case ElseToken:
		return BaseToken{TokenType: ElseToken}
	case IfEndToken:
		return BaseToken{TokenType: IfEndToken}
	case ShorthandIfToken:
		return BaseToken{TokenType: ShorthandIfToken, RawValue: value}
	case ForToken:
		loopVar, expr := parseForExpr(value)
		return BaseToken{TokenType: ForToken, RawValue: strings.TrimSpace(value), Var: loopVar, LoopExprValue: expr}
	case ForEndToken:
		return BaseToken{TokenType: ForEndToken}
	default:
		return BaseToken{TokenType: TextToken, RawValue: value, Var: value}
	}
}

// parseForExpr extracts the loop variable specification and iteration expression
// from a "for X in Y" or "for K, V in Y" string. the leading "for" has already
// been recognised; we split on " in ". the returned spec is either "v" (single
// var) or "k,v" (paired) — the renderer splits this on comma.
func parseForExpr(s string) (string, string) {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "for")
	s = strings.TrimSpace(s)
	idx := strings.Index(s, " in ")
	if idx < 0 {
		return strings.TrimSpace(s), ""
	}
	lhs := strings.TrimSpace(s[:idx])
	expr := strings.TrimSpace(s[idx+len(" in "):])
	// normalise pair form: collapse internal whitespace so "k, v" → "k,v"
	if strings.Contains(lhs, ",") {
		parts := strings.Split(lhs, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		lhs = strings.Join(parts, ",")
	}
	return lhs, expr
}
