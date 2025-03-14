package sintax

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/toaweme/date"
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
	"trim":    trim,
	"shorten": shorten,
	"length":  length,
}

func toFormat(val any, params []any) (any, error) {
	switch timeValue := val.(type) {
	case string:
		return timeValue, nil
	case time.Time:
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
	if varValue == "" {
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

var trim = func(s any, _ []any) (any, error) {
	switch v := s.(type) {
	case string:
		return strings.TrimSpace(v), nil
		// default:
		// 	return "", fmt.Errorf("trim requires a text argument")
	}
	return s, nil
}

var shorten = func(s any, args []any) (any, error) {
	str, ok := s.(string)
	if !ok {
		return "", fmt.Errorf("shorten requires a text argument")
	}

	if len(args) != 1 {
		return "", fmt.Errorf("shorten requires 1 argument")
	}
	length, err := strconv.Atoi(fmt.Sprint(args[0]))
	if err != nil {
		return "", fmt.Errorf("shorten requires a numeric argument")
	}

	if len(str) > length {
		return str[:length], nil
	}

	return str, nil
}

var length = func(s any, _ []any) (any, error) {
	str, ok := s.(string)
	if !ok {
		return "", fmt.Errorf("length requires a text argument")
	}

	return strconv.Itoa(len(str)), nil
}
