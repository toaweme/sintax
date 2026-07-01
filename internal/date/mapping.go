package date

// DefaultFormat is the default date format used when no format string is given.
var DefaultFormat = "Y-m-d H:i:s"

// DefaultMapping maps PHP-style date format characters to Go reference time
// layout tokens. Only characters with a direct, correct Go equivalent are
// included: year, month, day, hour, minute, second, and timezone offset.
//
// Characters not present here are copied through to the output unchanged.
var DefaultMapping = map[string]string{
	// year
	"Y": "2006",
	"y": "06",

	// month
	"m": "01",
	"n": "1",
	"M": "Jan",
	"F": "January",

	// day
	"d": "02",
	"j": "2",
	"D": "Mon",
	"l": "Monday",

	// hour
	// G has no unpadded 24-hour token in Go's reference layout, so it renders
	// zero-padded like H.
	"H": "15",
	"G": "15",
	"h": "03",
	"g": "3",

	// minute, second
	"i": "04",
	"s": "05",

	// meridiem, paired with h/g
	"a": "pm",
	"A": "PM",

	// timezone offset
	"T": "MST",
	"O": "-0700",
	"P": "-07:00",
}
