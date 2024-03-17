package date

var DefaultFormat = "Y-m-d H:i:s"

var Format = map[string]string{
	// day of the month, 2 digits with leading zeros: 01 to 31
	"d": "02",
	// a textual representation of a day, three letters: mon through sun
	"D": "Mon",
	// day of the month without leading zeros: 1 to 31
	"j": "2",
	// a full textual representation of the day of the week: sunday through saturday
	"l": "Monday",
	// iso-8601 numeric representation of the day of the week: 1 (for monday) through 7 (for sunday) - custom handling needed
	"N": "",
	// english ordinal suffix for the day of the month, 2 characters: st, nd, rd, or th - custom handling needed
	"S": "",
	// numeric representation of the day of the week: 0 (for sunday) through 6 (for saturday) - custom handling needed
	"w": "",
	// the day of the year (starting from 0): 0 through 365 - custom handling needed
	"z": "",
	// iso-8601 week number of year, weeks starting on monday: example: 42 (the 42nd week in the year)
	"W": "02",
	// a full textual representation of a month, such as january or march: january through december
	"F": "January",
	// numeric representation of a month, with leading zeros: 01 through 12
	"m": "01",
	// a short textual representation of a month, three letters: jan through dec
	"M": "Jan",
	// numeric representation of a month, without leading zeros: 1 through 12
	"n": "1",
	// number of days in the given month: 28 through 31 - custom handling needed
	"t": "",
	// whether it's a leap year: 1 if it is a leap year, 0 otherwise - custom handling needed
	"L": "",
	// iso-8601 week-numbering year. this has the same value as y, except that if the iso week number (w) belongs to the previous or next year, that year is used instead.
	"o": "2006",
	// an expanded full numeric representation of a year, at least 4 digits, with - for years bce, and + for years ce - custom handling needed
	"X": "",
	// an expanded full numeric representation if required, or a standard full numeral representation if possible - custom handling needed
	"x": "",
	// a full numeric representation of a year, at least 4 digits
	"Y": "2006",
	// a two-digit representation of a year
	"y": "06",
	// lowercase ante meridiem and post meridiem: am or pm
	"a": "pm",
	// uppercase ante meridiem and post meridiem: am or pm
	"A": "PM",
	// swatch internet time: 000 through 999 - custom handling needed
	"B": "",
	// 12-hour format of an hour without leading zeros: 1 through 12
	"g": "3",
	// 24-hour format of an hour without leading zeros: 0 through 23
	"G": "15",
	// 12-hour format of an hour with leading zeros: 01 through 12
	"h": "03",
	// 24-hour format of an hour with leading zeros: 00 through 23
	"H": "15",
	// minutes with leading zeros: 00 to 59
	"i": "04",
	// seconds with leading zeros: 00 through 59
	"s": "05",
	// microseconds - note that date() will always generate 000000 since it takes an int parameter, whereas datetime::format() does support microseconds if datetime was created with microseconds.
	"u": "000000",
	// milliseconds - same note applies as for u.
	"v": "000",
	// timezone identifier: examples: utc, gmt, atlantic/azores
	"e": "MST",
	// whether or not the date is in daylight saving time: 1 if daylight saving time, 0 otherwise - custom handling needed
	"I": "",
	// difference to greenwich time (gmt) without colon between hours and minutes: example: +0200
	"O": "-0700",
	// difference to greenwich time (gmt) with colon between hours and minutes: example: +02:00
	"P": "-07:00",
	// the same as p, but returns z instead of +00:00 (available as of php 8.0.0): examples: z or +02:00 - custom handling needed
	"p": "",
	// timezone abbreviation, if known; otherwise the gmt offset: examples: est, mdt
	"T": "MST",
	// timezone offset in seconds. the offset for timezones west of utc is always negative, and for those east of utc is always positive - custom handling needed
	"Z": "",
	// iso 8601 date: 2004-02-12t15:19:21+00:00
	"c": "2006-01-02T15:04:05-07:00",
	// rfc 2822 formatted date: example: thu, 21 dec 2000 16:01:07 +0200
	"r": "Mon, 02 Jan 2006 15:04:05 -0700",
	// seconds since the unix epoch (january 1 1970 00:00:00 gmt) - custom handling needed
	"U": "",
}
