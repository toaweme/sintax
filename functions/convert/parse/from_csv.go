package parse

import (
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameFromCSV is the template name for the FromCSV modifier.
const ModifierNameFromCSV functions.ModifierName = "from_csv"

// FromCSV parses a CSV string into a list of rows, so a serialized table (a
// report export, a spreadsheet dump) becomes data that later template steps can
// filter and pluck.
//
// The first record is the header row, and each remaining row becomes one map
// keyed by header. Cells stay strings, so coerce them downstream with other
// modifiers if you need numbers. Fully blank lines are skipped, and a row with
// fewer cells than the header pads the missing columns with an empty string.
func FromCSV(value string) ([]map[string]any, error) {
	reader := csv.NewReader(strings.NewReader(value))
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
