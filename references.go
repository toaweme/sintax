package sintax

import "strings"

// Reference is one name a template reads from outside itself.
type Reference struct {
	// Name is the reference as written, dots included.
	Name string
	// Root is the first segment of Name, which is the name a variable set
	// answers with the value the rest of the path walks into.
	Root string
	// Optional marks a position that reads a miss as false rather than failing,
	// which is what an if condition does.
	Optional bool
	// Default marks a pipeline that answers a miss itself.
	Default bool
}

// References reports every name tokens read from outside themselves, in written
// order. A loop's own bindings are left out, since the loop supplies them, and
// the collection it reads is reported in the scope around the block.
//
// A block tag nothing closes opens no block, so its body is read as the tokens
// it is.
func References(tokens []Token) ([]Reference, error) {
	var s refScan
	if err := s.walk(tokens); err != nil {
		return nil, err
	}

	return s.out, nil
}

type refScan struct {
	out   []Reference
	bound []map[string]bool
}

func (s *refScan) walk(tokens []Token) error {
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		switch token.Type() {
		case TextToken, ElseToken, IfEndToken, ForEndToken, UndefinedToken:

		case VariableToken, FilteredVariableToken:
			s.record(token, false)

		case IfToken:
			_, end, ok := blockEnd(tokens, i, len(tokens))
			if !ok {
				continue
			}
			if err := s.expr(token.Raw(), true); err != nil {
				return err
			}
			if err := s.walk(tokens[i+1 : end]); err != nil {
				return err
			}
			i = end

		case ForToken:
			_, end, ok := blockEnd(tokens, i, len(tokens))
			if !ok {
				continue
			}
			if err := s.expr(token.LoopExpr(), false); err != nil {
				return err
			}
			s.bound = append(s.bound, parseLoopSpec(token).bindings())
			err := s.walk(tokens[i+1 : end])
			s.bound = s.bound[:len(s.bound)-1]
			if err != nil {
				return err
			}
			i = end
		}
	}

	return nil
}

// expr reads the single expression a control tag carries. The renderer
// evaluates the same text through parseExpr, so an expression it would refuse
// is refused here rather than reported as a name.
func (s *refScan) expr(expr string, optional bool) error {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return nil
	}

	token, err := parseExpr(expr)
	if err != nil {
		return err
	}
	s.record(token, optional)

	return nil
}

func (s *refScan) record(token Token, optional bool) {
	if name := token.Name(); name != "" && !s.isBound(name) {
		s.out = append(s.out, Reference{Name: name, Root: rootName(name), Optional: optional, Default: hasDefault(token)})
	}

	// a modifier argument carries no default of its own
	for _, param := range token.Params() {
		if !s.isBound(param) {
			s.out = append(s.out, Reference{Name: param, Root: rootName(param), Optional: optional})
		}
	}
}

// isBound reads a dotted reference by its first segment, so a path into the
// current element is answered by the element.
func (s *refScan) isBound(name string) bool {
	root := rootName(name)
	for _, frame := range s.bound {
		if frame[root] {
			return true
		}
	}

	return false
}

func rootName(name string) string {
	if cut := strings.IndexByte(name, '.'); cut >= 0 {
		return name[:cut]
	}

	return name
}

func hasDefault(token Token) bool {
	if token.Type() != FilteredVariableToken {
		return false
	}

	_, funcs := varAndFuncs(token)
	for _, fn := range funcs {
		if fn.Name == "default" {
			return true
		}
	}

	return false
}
