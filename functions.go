package sintax

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"

	"github.com/contentforward/date"
)

type GlobalModifier func(val any, params []any) (any, error)

var BuiltinFunctions = map[string]GlobalModifier{
	"format":  toFormat,
	"default": toDefault,
	"json":    toJSON,
	"yaml":    toYAML,
	"sexy":    toSexy,
	"lines":   toLines,
	"join":    toJoin,
}

func toFormat(val any, params []any) (any, error) {
	switch timeValue := val.(type) {
	case string:
		log.Trace().Msgf("formatting string: %s", timeValue)
		return timeValue, nil
	case time.Time:
		log.Trace().Msgf("formatting time: %s", timeValue)
		d := date.NewFormatter(timeValue, date.DefaultMapping)
		format := date.DefaultFormat
		if len(params) > 0 {
			format = params[0].(string)
		}

		goDateFormat, err := d.Render(format)
		if err != nil {
			return nil, fmt.Errorf("failed to apply format filter '%s': %w", params[0], err)
		}
		return timeValue.Format(goDateFormat), nil
	}

	return nil, fmt.Errorf("format function expected string or time.Time, got %T", val)
}

func toJoin(val any, params []any) (any, error) {
	switch v := val.(type) {
	case []string:
		return strings.Join(v, "\n"), nil
	}

	return nil, fmt.Errorf("join function expected array of strings, got %T", val)
}

func toSexy(val any, params []any) (any, error) {
	sexyBear := `
	 ʕ•ᴥ•ʔ
	/\o-o/\
	 | ᴥ |
	 \_|_/
	`
	return sexyBear, nil
}

func isParam(params []any, index int, name string) bool {
	if len(params) <= index {
		return false
	}

	return params[index] == name
}

func toDefault(varValue any, params []any) (any, error) {
	if varValue == nil {
		return params[0], nil
	}

	return varValue, nil
}

func toJSON(val any, params []any) (any, error) {
	if len(params) > 0 && isParam(params, 0, "pretty") {
		jsonBytes, err := json.MarshalIndent(val, "", "  ")
		if err != nil {
			return "", fmt.Errorf("failed to apply json filter: %w", err)
		}

		return string(jsonBytes), nil
	}
	jsonBytes, err := json.Marshal(val)
	if err != nil {
		return "", fmt.Errorf("failed to apply json filter: %w", err)
	}

	return string(jsonBytes), nil
}

func toYAML(val any, params []any) (any, error) {
	yamlBytes, err := yaml.Marshal(val)
	if err != nil {
		return "", fmt.Errorf("failed to apply yaml filter: %w", err)
	}

	return string(yamlBytes), nil
}

func toLines(val any, params []any) (any, error) {
	if val == nil {
		return nil, nil
	}

	switch v := val.(type) {
	case string:
		return strings.Split(v, "\n"), nil
	case []byte:
		return strings.Split(string(v), "\n"), nil
	}

	return nil, fmt.Errorf("lines function expected string, got %T", val)
}
