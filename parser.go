package sintax

import (
	"fmt"
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

func (p *StringParser) ParseVariable(s string) ([]Token, error) {
	tokens, err := p.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("failed to parse variable: %w", err)
	}

	return tokens, nil
}

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

		// create the appropriate token
		tokenType := p.detectTokenType(contents)
		tokens = append(tokens, p.createToken(tokenType, contents))

		// move `i` beyond the closer
		i = closerIndex + len(p.closer)
	}

	return tokens, nil
}

// regex check if it's a valid variable name
// or quoted string
// or number
// or boolean (true/false or 1/0 or yes/no)
func (p *StringParser) isVariable(s string) bool {
	// implement
	reg := regexp.MustCompile(`^([a-zA-Z_\.][a-zA-Z0-9_\.]*)$`)
	return reg.MatchString(s)
}

func (p *StringParser) detectTokenType(s string) TokenType {
	s = strings.TrimSpace(s)

	if strings.HasPrefix(s, "if") {
		return IfToken
	} else if strings.HasPrefix(s, "/if") {
		return IfEndToken
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
		return BaseToken{VariableToken, strings.TrimSpace(value), strings.TrimSpace(value), nil}
	case FilteredVariableToken:
		token := BaseToken{FilteredVariableToken, strings.TrimSpace(value), strings.TrimSpace(splitAndGetFirst(value)), nil}
		_, funcs := getVarAndFunctions(token)
		for _, f := range funcs {
			for _, p := range f.Args {
				if p.Var {
					if token.ParamVars == nil {
						token.ParamVars = make([]string, 0)
					}
					token.ParamVars = append(token.ParamVars, p.Value.(string))
				}
			}
		}
		return token
	case IfToken:
		return BaseToken{IfToken, trimPrefix(value, "if"), "", nil}
	case ElseToken:
		return BaseToken{ElseToken, "", "", nil}
	case IfEndToken:
		return BaseToken{IfEndToken, "", "", nil}
	case ShorthandIfToken:
		return BaseToken{ShorthandIfToken, value, "", nil}
	default:
		return BaseToken{TextToken, value, value, nil}
	}
}
