package sintax

import "strings"

// blockEnd locates the token closing the block opened at tokens[start], counting
// nested blocks of the same kind so an inner block never closes an outer one. It
// reports the index of the block's own top-level else, or -1. ok is false when
// the tag opens no block or when nothing closes it.
func blockEnd(tokens []Token, start, end int) (elseIdx, endIdx int, ok bool) {
	if start < 0 || start >= len(tokens) {
		return -1, -1, false
	}

	opener := tokens[start].Type()
	closer, isBlock := blockCloser(opener)
	if !isBlock {
		return -1, -1, false
	}

	elseIdx = -1
	depth := 0
	for j := start + 1; j < end && j < len(tokens); j++ {
		switch t := tokens[j].Type(); {
		case t == opener:
			depth++
		case t == closer:
			if depth == 0 {
				return elseIdx, j, true
			}
			depth--
		case t == ElseToken && opener == IfToken && depth == 0 && elseIdx == -1:
			elseIdx = j
		}
	}

	return -1, -1, false
}

func blockCloser(opener TokenType) (TokenType, bool) {
	switch opener {
	case IfToken:
		return IfEndToken, true
	case ForToken:
		return ForEndToken, true
	default:
		return UndefinedToken, false
	}
}

// loopSpec is the names one for block binds over its body, derived from the
// element name, so a loop over "tx" reads its index as "tx_index".
type loopSpec struct {
	// Element is the name bound to the current item.
	Element string
	// Position is the name the template chose in the paired form "for k, v in xs",
	// and is empty in the single-name form.
	Position string
	Index    string
	First    string
	Last     string
	// MapKey holds the map key where the template chose no Position of its own.
	MapKey string
}

func parseLoopSpec(tok Token) loopSpec {
	spec := strings.TrimSpace(tok.Name())

	var loop loopSpec
	loop.Element = spec
	if position, element, paired := strings.Cut(spec, ","); paired {
		loop.Position = strings.TrimSpace(position)
		loop.Element = strings.TrimSpace(element)
	}

	if loop.Element == "" {
		return loop
	}

	loop.Index = loop.Element + "_index"
	loop.First = loop.Element + "_first"
	loop.Last = loop.Element + "_last"
	loop.MapKey = loop.Element + "_key"

	return loop
}

// bindings is the set of names the loop puts in reach of its body. MapKey is
// bound only over a map, which is not known until the iterable is read, so it
// counts as bound wherever the template named no position of its own.
func (s loopSpec) bindings() map[string]bool {
	if s.Element == "" {
		if s.Position == "" {
			return nil
		}

		return map[string]bool{s.Position: true}
	}

	bound := map[string]bool{
		s.Element: true,
		s.Index:   true,
		s.First:   true,
		s.Last:    true,
	}
	if s.Position != "" {
		bound[s.Position] = true
	} else {
		bound[s.MapKey] = true
	}

	return bound
}
