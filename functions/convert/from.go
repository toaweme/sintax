package convert

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/toaweme/sintax/functions"
)

var FromDefinition = functions.ModifierDefinition{
	Func: From,
	AcceptedValue: []functions.Type{
		functions.TypeString,
	},
	AcceptedParams: []functions.Param{
		{
			Type:  functions.TypeString,
			Index: 0,
			Value: "json",
		},
	},
}

func From(value any, params []any) (any, error) {
	if len(params) > 0 && functions.IsParam(params, 0, "json") {
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

	return nil, fmt.Errorf("unsupported format in from function")
}
