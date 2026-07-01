// Package convert provides template modifiers for converting between data formats.
package convert

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameFrom is the template name for the From modifier.
const ModifierNameFrom functions.ModifierName = "from"

// From parses the string value as the given format and returns the parsed result.
//
// value: string
// param:0: string ("json" | "csv")
// returns: map for json, []map[string]any for csv
//
// example: parse a JSON string into a map
// in:  body = "{\"name\": \"Alice\", \"role\": \"admin\"}"
// tpl: {{ body | from:'json' }}
// out: {"name": "Alice", "role": "admin"}
//
// example: parse a CSV string into a list of rows
// in:  body = "name,age\nAlice,30\nBob,25"
// tpl: {{ body | from:'csv' }}
// out: [{"name": "Alice", "age": "30"}, {"name": "Bob", "age": "25"}]
func From(value any, params []any) (any, error) {
	if len(params) == 0 {
		return nil, fmt.Errorf("from function requires a format parameter")
	}

	switch {
	case functions.IsParam(params, 0, "json"):
		return fromJSON(value)
	case functions.IsParam(params, 0, "csv"):
		return fromCSV(value)
	}

	return nil, fmt.Errorf("unsupported format in from function")
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
