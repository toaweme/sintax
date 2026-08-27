package date

import (
	"fmt"
	"time"
)

// Parser reads time.Time values out of strings written in a PHP-style date
// format, translating each recognized character to its Go reference time
// layout token. It is the inverse of Formatter and shares its mapping, so a
// layout that renders a value also reads it back.
type Parser struct {
	mapping map[string]string
}

// NewParser returns a Parser that reads using the given character-to-Go-layout mapping.
func NewParser(mapping map[string]string) *Parser {
	return &Parser{mapping: mapping}
}

// Parse reads value according to format, in loc when the layout carries no zone
// of its own. Characters without a mapping entry are matched literally, so the
// separators in "F j, Y" have to appear in the value exactly as written.
//
// A value that does not match the layout is an error rather than a zero time,
// since a silent zero would travel downstream as a real date of January 1, year 1.
func (p *Parser) Parse(value string, format string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		loc = time.UTC
	}
	parsed, err := time.ParseInLocation(translate(p.mapping, format), value, loc)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date %q with format %q: %w", value, format, err)
	}
	return parsed, nil
}
