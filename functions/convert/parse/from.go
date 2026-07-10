// Package parse provides modifiers that parse a serialized value into data.
package parse

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameFrom is the template name for the From modifier.
const ModifierNameFrom functions.ModifierName = "from"

// From parses the string value as the named format and returns the parsed data,
// so a serialized payload (an API response body, a config blob) can be turned
// back into a map or list of rows that later template steps can index into.
// The format argument is required and only "json" and "csv" are supported.
// An unknown format, a missing format, or a non-string value all return an error.
//
// value: string
// param:0: string, the format to parse, one of "json" or "csv"
// returns: map[string]any for json, []map[string]any for csv
//
// The "json" format expects a top-level JSON object and returns it as a map.
// Numbers decode to native int64 or float64 (a value with a decimal point or
// exponent becomes float64, otherwise int64) rather than json.Number, so
// downstream numeric modifiers see real numbers. A top-level JSON array or
// scalar is not an object and returns an error.
//
// example: parse a JSON object into a map
// in:  body = "{\"name\": \"Alice\", \"role\": \"admin\"}"
// tpl: {{ body | from:'json' }}
// out: map[name:Alice role:admin]
//
// example: numbers decode to native int64 and float64
// in:  body = "{\"count\": 3, \"ratio\": 1.5}"
// tpl: {{ body | from:'json' }}
// out: map[count:3 ratio:1.5]
//
// The "csv" format treats the first record as the header row and returns one
// map per remaining row, keyed by header. Cells stay strings, so coerce them
// downstream with other modifiers if you need numbers. Fully blank lines are
// skipped, and a row with fewer cells than the header pads the missing columns
// with an empty string.
//
// example: parse a CSV string into a list of rows
// in:  body = "name,age\nAlice,30\nBob,25"
// tpl: {{ body | from:'csv' }}
// out: [map[age:30 name:Alice] map[age:25 name:Bob]]
//
// example: a header-only CSV yields an empty list
// in:  body = "name,age\n"
// tpl: {{ body | from:'csv' }}
// out: []
func From(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, errors.New("from function requires a format parameter")
	}

	switch {
	case functions.IsParam(params, 0, "json"):
		return fromJSON(value)
	case functions.IsParam(params, 0, "csv"):
		return fromCSV(value)
	}

	return nil, errors.New("unsupported format in from function")
}

func fromJSON(value any) (any, error) {
	val, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("from function expected string for json, got %T", value)
	}

	dec := json.NewDecoder(strings.NewReader(val))
	dec.UseNumber()

	var raw map[string]any
	if err := dec.Decode(&raw); err != nil {
		return nil, fmt.Errorf("failed to convert JSON to map: %w", err)
	}

	return functions.ConvertNumbersJSON(raw), nil
}

// fromCSV parses a CSV string into a list of header-keyed maps. The first
// non-empty record is treated as the header. Values stay as strings — callers
// can coerce downstream via other modifiers.
func fromCSV(value any) (any, error) {
	val, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("from function expected string for csv, got %T", value)
	}

	reader := csv.NewReader(strings.NewReader(val))
	reader.FieldsPerRecord = -1
	reader.TrimLeadingSpace = true

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to parse CSV: %w", err)
	}

	if len(records) == 0 {
		return []map[string]any{}, nil
	}

	headers := records[0]
	rows := make([]map[string]any, 0, len(records)-1)
	for _, rec := range records[1:] {
		// skip fully empty lines
		empty := true
		for _, cell := range rec {
			if strings.TrimSpace(cell) != "" {
				empty = false
				break
			}
		}
		if empty {
			continue
		}

		row := make(map[string]any, len(headers))
		for i, h := range headers {
			if i < len(rec) {
				row[h] = rec[i]
			} else {
				row[h] = ""
			}
		}
		rows = append(rows, row)
	}

	return rows, nil
}
