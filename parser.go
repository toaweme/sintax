package sintax

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

type StringParser struct{}

func NewStringParser() *StringParser {
	return &StringParser{}
}

var _ Parser = (*StringParser)(nil)

// start of a token
const opener = '{'
const closer = '}'

func (p *StringParser) Parse(template string) ([]Token, error) {
	if !strings.ContainsRune(template, opener) && !strings.ContainsRune(template, closer) {
		return []Token{BaseToken{TokenType: TextToken, RawValue: template}}, nil
	}
	
	var sb = &strings.Builder{}
	var tokens []Token
	startIndex := -1
	
	i := 0
	totalRunes := len(template)
	for i < totalRunes {
		char := template[i]
		
		if char == opener && peek(template, i) == opener {
			if sb.Len() > 0 {
				tokens = append(tokens, BaseToken{TokenType: TextToken, RawValue: sb.String()})
				sb.Reset()
			}
			
			afterOpeningIndex := p.skipWhitespace(template, i+2)
			
			i = afterOpeningIndex
			startIndex = i
			sb.Reset()
			continue
		} else if char == closer && peek(template, i) == closer {
			contents := template[startIndex:i]
			token := p.createToken(p.detectTokenType(contents), contents)
			tokens = append(tokens, token)
			i = i + 2
			startIndex = -1
			continue
		} else if startIndex == -1 {
			sb.WriteByte(char)
		}
		i++
	}
	
	// if there's any text left in the buffer, add it as a text token
	if sb.Len() > 0 {
		tokens = append(tokens, BaseToken{TokenType: TextToken, RawValue: sb.String()})
		sb.Reset()
	}
	
	return tokens, nil
}

func (p *StringParser) ParseVariable(s string) ([]Token, error) {
	tokens, err := p.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("failed to parse variable: %w", err)
	}
	
	return tokens, nil
}

func (p *StringParser) skipWhitespace(s string, i int) int {
	for i < len(s) && unicode.IsSpace(rune(s[i])) {
		i++
	}
	
	return i
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

func (p *StringParser) createToken(tokenType TokenType, value string) Token {
	switch tokenType {
	case VariableToken:
		return BaseToken{VariableToken, trimPrefix(value, "")}
	case FilteredVariableToken:
		return BaseToken{FilteredVariableToken, trimPrefix(value, "")}
	case IfToken:
		return BaseToken{IfToken, trimPrefix(value, "if")}
	case ElseToken:
		return BaseToken{ElseToken, ""}
	case IfEndToken:
		return BaseToken{IfEndToken, ""}
	case ShorthandIfToken:
		return BaseToken{ShorthandIfToken, trimPrefix(value, "")}
	default:
		return BaseToken{TextToken, value}
	}
}

func peek(s string, i int) byte {
	if i+1 < len(s) {
		return s[i+1]
	}
	return 0
}
