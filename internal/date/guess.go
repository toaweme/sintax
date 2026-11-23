package date

import (
	"errors"
	"strings"
	"time"
)

var (
	// ErrUnknownLayout reports that a value is written in none of the layouts
	// Guess knows, so the caller has to say which layout it is written in.
	ErrUnknownLayout = errors.New("no built-in date layout matches")

	// ErrAmbiguousDate reports that a numeric value reads as a real date with
	// the day and the month either way round, and the two readings disagree.
	ErrAmbiguousDate = errors.New("day and month could be either way round")
)

// guessLayouts are the Go reference layouts Guess tries, in order, and that
// order is the whole contract. It runs from the machine forms nothing else can
// look like, through the written-out forms a month name pins down, to the
// numeric forms only one reading fits.
//
// They are Go layouts rather than the PHP-style layout language the rest of
// this package speaks, because that language has no escape and its T means a
// zone abbreviation, so it cannot write the literal T that separates the date
// and the time in ISO 8601.
var guessLayouts = []string{
	time.RFC3339Nano,
	time.RFC3339,
	"2006-01-02T15:04:05",
	"2006-01-02T15:04",
	"2006-01-02 15:04:05Z07:00",
	"2006-01-02 15:04:05",
	"2006-01-02 15:04",
	"2006-01-02",
	"Monday, January 2, 2006 15:04",
	"Monday, January 2, 2006",
	"January 2, 2006 15:04:05",
	"January 2, 2006 15:04",
	"January 2, 2006",
	"2 January 2006 15:04",
	"2 January 2006",
	"Jan 2, 2006 15:04",
	"Jan 2, 2006",
	"2 Jan 2006 15:04",
	"2 Jan 2006",
	"2006/01/02 15:04:05",
	"2006/01/02 15:04",
	"2006/01/02",
}

// rfc5322Layouts are the forms a mail date header takes once guessRFC5322 has
// dropped its comments, folded its whitespace and turned an obsolete zone name
// into an offset. The weekday and the seconds are optional, and the unpadded
// day reads a zero-padded one too.
var rfc5322Layouts = []string{
	"Mon, 2 Jan 2006 15:04:05 -0700",
	"2 Jan 2006 15:04:05 -0700",
	"Mon, 2 Jan 2006 15:04 -0700",
	"2 Jan 2006 15:04 -0700",
	"Monday, 2 Jan 2006 15:04:05 -0700",
	"Monday, 2 Jan 2006 15:04 -0700",
}

// obsoleteZones maps the zone names RFC 5322 still lets a date header carry to
// their offsets. UTC is not one of them, but servers write it anyway. Any other
// name is refused, because Go reads an unknown abbreviation as a zero offset
// and the time would be silently wrong.
var obsoleteZones = map[string]string{
	"UT":  "+0000",
	"UTC": "+0000",
	"GMT": "+0000",
	"Z":   "+0000",
	"EST": "-0500",
	"EDT": "-0400",
	"CST": "-0600",
	"CDT": "-0500",
	"MST": "-0700",
	"MDT": "-0600",
	"PST": "-0800",
	"PDT": "-0700",
}

// guessRFC5322 reads value as a mail date header, reporting false when it is
// not one. The zone is part of the form, so the result never takes a location
// from the caller.
func guessRFC5322(value string) (time.Time, bool) {
	stripped, ok := stripComments(value)
	if !ok {
		return time.Time{}, false
	}
	fields := strings.Fields(stripped)
	if len(fields) < 4 {
		return time.Time{}, false
	}
	zone := fields[len(fields)-1]
	if offset, known := obsoleteZones[strings.ToUpper(zone)]; known {
		fields[len(fields)-1] = offset
	}
	normalised := strings.Join(fields, " ")
	for _, layout := range rfc5322Layouts {
		if parsed, err := time.Parse(layout, normalised); err == nil {
			return parsed, true
		}
	}
	return time.Time{}, false
}

// stripComments drops every parenthesised comment from value, the way RFC 5322
// lets one trail a date header as in +0100 (CET). It reports false for a
// comment left open.
func stripComments(value string) (string, bool) {
	var out strings.Builder
	depth := 0
	for _, r := range value {
		switch {
		case r == '(':
			depth++
		case r == ')' && depth > 0:
			depth--
			if depth == 0 {
				out.WriteByte(' ')
			}
		case depth == 0:
			out.WriteRune(r)
		}
	}
	return out.String(), depth == 0
}

// ambiguousLayouts are the numeric forms whose first two fields could be the
// day and the month either way round. Each entry pairs the day-first reading
// with the month-first one, and both are unpadded so a single-digit field and
// a zero-padded one land on the same entry.
var ambiguousLayouts = [][2]string{
	{"2/1/2006 15:04:05", "1/2/2006 15:04:05"},
	{"2/1/2006 15:04", "1/2/2006 15:04"},
	{"2/1/2006", "1/2/2006"},
	{"2.1.2006 15:04", "1.2.2006 15:04"},
	{"2.1.2006", "1.2.2006"},
	{"2-1-2006 15:04", "1-2-2006 15:04"},
	{"2-1-2006", "1-2-2006"},
}

// Guess reads a date out of a value that does not say which layout it is
// written in, using loc for a value carrying no zone of its own.
//
// The layouts in guessLayouts are tried in their listed order and the first
// that parses the value whole wins. A mail date header in any form RFC 5322
// allows comes next, with an optional weekday, a one or two digit day,
// optional seconds, a numeric or obsolete named zone and comments. The numeric day-and-month forms come last
// and are handled as pairs, so 25/12/2026 reads as December 25 because only
// the day-first reading gives a real date, while 03/04/2026 reads two ways and
// is refused with ErrAmbiguousDate rather than guessed. A wrong-but-plausible
// date written silently is worse than an error naming the value, and the
// layout argument is right there to settle it.
func Guess(value string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		loc = time.UTC
	}
	for _, layout := range guessLayouts {
		if parsed, err := time.ParseInLocation(layout, value, loc); err == nil {
			return parsed, nil
		}
	}
	if parsed, ok := guessRFC5322(value); ok {
		return parsed, nil
	}
	for _, pair := range ambiguousLayouts {
		dayFirst, dayErr := time.ParseInLocation(pair[0], value, loc)
		monthFirst, monthErr := time.ParseInLocation(pair[1], value, loc)
		switch {
		case dayErr == nil && monthErr == nil:
			if dayFirst.Equal(monthFirst) {
				return dayFirst, nil
			}
			return time.Time{}, ErrAmbiguousDate
		case dayErr == nil:
			return dayFirst, nil
		case monthErr == nil:
			return monthFirst, nil
		}
	}
	return time.Time{}, ErrUnknownLayout
}
