# sintax

[![Quality](https://github.com/toaweme/sintax/actions/workflows/quality.yml/badge.svg)](https://github.com/toaweme/sintax/actions/workflows/quality.yml)
<a href="https://code.toawe.me/toaweme/sintax/health">
    <picture>
        <source media="(prefers-color-scheme: dark)" srcset="https://code.toawe.me/toaweme/sintax/badge-dark.svg">
        <source media="(prefers-color-scheme: light)" srcset="https://code.toawe.me/toaweme/sintax/badge.svg">
        <img alt="sintax health" src="https://code.toawe.me/toaweme/sintax/badge.svg">
    </picture>
</a>
[![Go Reference](https://img.shields.io/badge/Docs-pkg.go.dev-blue)](https://pkg.go.dev/github.com/toaweme/sintax)
[![GitHub Tag](https://img.shields.io/github/v/tag/toaweme/sintax?label=Tag&color=green)](https://github.com/toaweme/sintax/releases)
[![Dependencies](https://img.shields.io/badge/Dependencies-0-brightgreen)](go.mod)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue)](/LICENSE)

## Templating engine

Zero dependency templating engine built for workflows, document generation, and data transformations.

```
{{ response | from:'json' | key:'orders' | filter:'status','paid' | pluck:'total' | sum | decimal:2 }}
```

---

## Features

- **Pipe syntax**: chain built-in modifiers to transform any value in a single expression
- **Nested data**: maps, slices, structs, and pointers are all resolved and rendered recursively
- **Conditionals**: `{{ if x }} … {{ else }} … {{ endif }}` blocks
- **Loops**: `{{ for v in items }} … {{ endfor }}` over slices and maps, with auto-bound index/key helpers
- **Nested templates**: the `template` modifier re-enters the engine to render a loaded string (e.g. a
  file's contents) as its own template, guarded against runaway recursion
- **Typed errors**: missing variables, unknown functions, and malformed tokens are surfaced as distinct errors
- **Extensible**: register your own modifiers alongside the built-ins, or override a built-in by name
- **Zero dependencies**: the core engine only imports the Go standard library

---

## Install

```sh
go get github.com/toaweme/sintax
```

---

## Quick start

```go
package main

import (
    "fmt"
    "github.com/toaweme/sintax"
)

func main() {
    s := sintax.New(sintax.BuiltinFunctions(nil, nil))

    out, _ := s.Render("Hello, {{ name | title }}!", map[string]any{"name": "alice-cooper"})
    fmt.Println(out) // Hello, Alice Cooper!
}
```

`sintax.New(funcs)` takes the exact set of modifiers the engine can call - it does not merge in any
built-ins on its own. `sintax.BuiltinFunctions(overrides, safeDirs)` builds that set: pass `nil` for both
to get every built-in with no overrides, or a `map[string]sintax.GlobalModifier` to add or replace
modifiers by name, and a list of directories the `file` modifier is allowed to read from.

---

## Template syntax

| Pattern | Example |
|---|---|
| Variable | `{{ name }}` |
| Nested field access | `{{ user \| key:'name' }}` |
| Modifier chain | `{{ text \| trim \| upper }}` |
| Modifier with args | `{{ items \| join:',' }}` |
| Variable as argument | `{{ text \| trim-prefix:prefix_var }}` |
| Fallback to literal | `{{ value \| default:'n/a' }}` |
| Fallback to empty array | `{{ items \| default:[] }}` |
| Fallback to empty object | `{{ user \| default:{} }}` |
| If / else | `{{ if active }}yes{{ else }}no{{ endif }}` |
| Loop over a slice | `{{ for x in items }}- {{ x }}\n{{ endfor }}` |
| Loop with index | `{{ for i, x in items }}{{ i }}:{{ x }} {{ endfor }}` |
| Loop over a map | `{{ for k, v in headers }}{{ k }}={{ v }} {{ endfor }}` |

**Modifier syntax:** the name and the first argument are separated by `:`, additional arguments by `,`.
String literals use single or double quotes; `[]` and `{}` are empty-collection literals; other unquoted
tokens resolve as variables, numbers, or booleans.

**Variable names are literal keys, not paths.** A variable is looked up by its exact name in the vars map -
there is no `obj.field` dot-notation. `{{ user.name }}` looks for a variable literally named `user.name`; it does
**not** descend into a `user` map. To read a nested field, pipe the value through the `key` modifier:
`{{ user | key:'name' }}`. `key` also accepts a dotted path to reach deeper, e.g. `{{ order | key:'meta.total' }}`.

**Block tags use `endif` and `endfor`** to close.

### Whitespace control

A leading or trailing `-` inside a tag eats whitespace on that side, the same way Jinja and Go templates do:

| Pattern | Effect |
|---|---|
| `{{- expr }}` | strip trailing whitespace (including newlines) from the text **before** this tag |
| `{{ expr -}}` | strip leading whitespace (including newlines) from the text **after** this tag |
| `{{- expr -}}` | both at once |

Block control tags (`if`/`else`/`endif`/`for`/`endfor`) that sit alone on their own line are auto-trimmed:
the surrounding indentation and the line's newline are removed automatically, so you don't have to write `-` on
every block tag just to keep your output clean. Use the explicit `{{-` / `-}}` form when a tag shares a line
with text or you want extra whitespace eaten.

---

## Loops

`{{ for v in xs }}` iterates over slices, arrays, and maps. Two binding forms are supported:

```
{{ for v in xs }} … {{ endfor }}        # bind value
{{ for i, v in xs }} … {{ endfor }}     # bind index/key + value
```

Inside the loop body the following helpers are available, where `<v>` is whatever name you bound the value to:

| Binding | Set on |
|---|---|
| `<v>_index` | both slice and map iterations (0-based) |
| `<v>_first`, `<v>_last` | both booleans for the first/last iteration |
| `<v>_key` | map iterations only, when no explicit key name was bound |

Map iteration order is sorted by key. Loops nest freely and parent variables remain visible inside the body.

```
{{ for item in cart }}{{ item | key:'name' }} × {{ item | key:'qty' }}{{ if item_last }}.{{ else }}, {{ endif }}{{ endfor }}
```

---

## Modifiers

### Text

Trim, case-shift, slugify, split, and reshape strings.

| Item | Description | Example |
| --- | --- | --- |
| `concat` | Concat appends one or more strings to the value. | `{{ greeting \| concat:'!' }}` |
| `join` | Join combines an array of strings into a single string with a separator. | `{{ tags \| join:',' }}` |
| `lines` | Lines splits a string or byte slice into an array of lines. | `{{ note \| lines }}` |
| `lower` | ToLower converts a string to lowercase. | `{{ email \| lower }}` |
| `replace` | Replace replaces all occurrences of a substring within the string value. | `{{ greeting \| replace:'world','everyone' }}` |
| `replace_pattern` | ReplacePattern replaces all regex matches within the string value. | `{{ text \| replace_pattern:'\s+',' ' }}` |
| `reverse` | Reverse reverses the characters in a string. | `{{ name \| reverse }}` |
| `shorten` | Shorten truncates a string to the given maximum character length. | `{{ description \| shorten:30 }}` |
| `slug` | Slug converts a string to a URL-friendly slug. | `{{ title \| slug }}` |
| `split` | Split splits a string into an array using a separator. | `{{ csv_line \| split:',' }}` |
| `template` | Template renders its incoming string value as a nested sintax template, so a value loaded from a file (or any string variable) can itself contain `{{ ... }}` markup. | `{{ "partial.tpl" \| file \| template }}` |
| `title` | Title converts a hyphen-separated slug into a title-cased string. | `{{ slug \| title }}` |
| `title_model` | ModelTitle formats an AI model identifier into a human-readable title. | `{{ model_id \| title_model }}` |
| `trim` | Trim removes leading and trailing whitespace, or the given character set. | `{{ name \| trim }}` |
| `trim-prefix` | TrimPrefix removes a leading prefix string or leading whitespace from the value. | `{{ path \| trim-prefix:'/' }}` |
| `trim-suffix` | TrimSuffix removes a trailing suffix string or trailing whitespace from the value. | `{{ url \| trim-suffix:'/' }}` |
| `upper` | ToUpper converts a string to uppercase. | `{{ name \| upper }}` |

### Collections

Sort, filter, find, and reshape arrays and maps.

| Item | Description | Example |
| --- | --- | --- |
| `filter` | Filter returns a subset of a slice where a nested field matches a value. | `{{ items \| filter:'status','active' }}` |
| `find` | Find returns the first element in a slice or map where a field equals the given value. | `{{ users \| find:'id',42 }}` |
| `first` | First returns the first character of a string or the first element of a slice. | `{{ items \| first }}` |
| `flatten` | Flatten flattens a slice of slices by one level. | `{{ groups \| pluck:'items' \| flatten }}` |
| `has` | Has returns true if the slice or map contains the given value. | `{{ tags \| has:'featured' }}` |
| `is` | Is returns true if the value equals any of the given parameters. | `{{ status \| is:'active' }}` |
| `key` | Key extracts a value from a map or slice by key path or index. | `{{ user \| key:'name' }}` |
| `last` | Last returns the last character of a string or the last element of a slice. | `{{ items \| last }}` |
| `map` | Map converts a slice of maps into a map keyed by the given field's string value. | `{{ users \| map:'id' }}` |
| `merge` | Merge converts a slice of maps into a map keyed by the given field's string value. | `{{ users \| merge:'id' }}` |
| `pluck` | Pluck extracts a single field from each element of a slice of maps and returns a slice of values. | `{{ users \| pluck:'id' }}` |
| `sort` | Sort sorts a slice in ascending or descending order. | `{{ names \| sort }}` |
| `sum` | Sum returns the numeric sum of the elements of a slice. | `{{ amounts \| sum }}` |
| `wrap` | Wrap wraps the value in a map under the given key. | `{{ name \| wrap:'user' }}` |

### Boolean

Compare values for use inside if/else blocks and conditional expressions.

| Item | Description | Example |
| --- | --- | --- |
| `eq` | Eq returns true if the value equals the given parameter. | `{{ status \| eq:'active' }}` |
| `gt` | Gt returns true if the numeric value is greater than the threshold. | `{{ items_in_cart \| gt:0 }}` |
| `gte` | Gte returns true if the numeric value is greater than or equal to the threshold. | `{{ qty \| gte:1 }}` |
| `not` | Not inverts the truthiness of the value. | `{{ is_active \| not }}` |

### Convert

Move between Go values, JSON, YAML, and other serialized formats.

| Item | Description | Example |
| --- | --- | --- |
| `from` | From parses the string value as the given format and returns the parsed result. | `{{ body \| from:'json' }}` |
| `json` | JSON serializes the value to a JSON string. | `{{ user \| json }}` |
| `markdown` | Markdown converts an HTML string to Markdown. | `{{ html_content \| markdown }}` |
| `yaml` | YAML serializes or parses a value as YAML. | `{{ config \| yaml }}` |

### Utilities

Defaults, lengths, line numbers, and date formatting.

| Item | Description | Example |
| --- | --- | --- |
| `decimal` | Decimal formats a number with a fixed number of decimal places. | `{{ amount \| decimal:2 }}` |
| `default` | Default returns the fallback value if the input is nil or an empty string. | `{{ name \| default:'anonymous' }}` |
| `format` | Format formats a time.Time value using a date format string. | `{{ created_at \| format:'YYYY-MM-DD' }}` |
| `length` | Length returns the number of characters in a string, bytes in a byte slice, or elements in a slice/array/map. | `{{ name \| length }}` |
| `line-numbers` | LineNumbers prepends each line of the string with its zero-based line number. | `{{ note \| line-numbers }}` |

### File System

Pull pieces out of file paths - directory, name, and extension - and read files from an allowlist of directories.

| Item | Description | Example |
| --- | --- | --- |
| `dirname` | Dirname returns the directory portion of a file path. | `{{ file_path \| dirname }}` |
| `ext` | FilenameExt returns the file extension without the leading dot. | `{{ file_path \| ext }}` |
| `ext-dot` | FilenameExtDot returns the file extension including the leading dot. | `{{ file_path \| ext-dot }}` |
| `ext-prepend` | FilenamePrependExt inserts an additional extension before the existing file extension. | `{{ file_path \| ext-prepend:'min' }}` |
| `ext-trim` | FilenameTrimExt returns the file path without its extension. | `{{ file_path \| ext-trim }}` |
| `file` | File reads a file's contents as a string. The path is resolved against the `safeDirs` passed to `BuiltinFunctions`, and `..` traversal outside them is rejected. | `{{ "greeting.tpl" \| file }}` |
| `filename` | Filename returns the base file name from a path, including the extension. | `{{ file_path \| filename }}` |

### Money

Convert numbers between currency units like dollars and cents.

| Item | Description | Example |
| --- | --- | --- |
| `currency` | Currency converts a numeric value between currency units by applying a unit multiplier ratio. | `{{ price \| currency:1,100 }}` |

---

_The sections below are for embedding sintax inside a Go program: instantiating the engine, registering
custom modifiers, and surfacing typed errors. The template syntax itself is documented above._

## Public API

The engine speaks through a small, stable set of interfaces. Everything else (`StringParser`,
`StringRenderer`, `BaseToken`) is an implementation detail.

### `Sintax`

```go
type Sintax interface {
	Render(template string, vars map[string]any) (any, error)
}
```

### `Parser`

```go
type Parser interface {
	Parse(template string) ([]Token, error)
}
```

### `Renderer`

```go
type Renderer interface {
	Render(tokens []Token, vars map[string]any) (any, error)
}
```

### `Token`

```go
type Token interface {
	Type() TokenType
	Raw() string
	Name() string
	Params() []string
	WithDefault() bool
	LoopExpr() string
}
```

---

## Error handling

Every failure mode has a sentinel error you can match with `errors.Is`. Wrap or compare, never string-match.

```go
_, err := s.Render(template, vars)
switch {
case errors.Is(err, sintax.ErrVariableNotFound):
    // referenced variable missing, use | default:'...' to avoid this
case errors.Is(err, sintax.ErrFunctionNotFound):
    // a modifier name in the pipe chain isn't registered
case errors.Is(err, sintax.ErrFunctionApplyFailed):
    // a modifier returned an error or the value type was unexpected
case errors.Is(err, sintax.ErrInvalidTokenType):
    // the parser produced a token the renderer doesn't know how to handle
case errors.Is(err, sintax.ErrMaxDepthExceeded):
    // the `template` modifier recursed past the nesting limit
case err != nil:
    // unclassified failure, e.g. a malformed if/for block
}
```

---

## Custom modifiers

Pass a map of overrides to `sintax.BuiltinFunctions`. Overrides also replace built-ins of the same name,
useful for sandboxing or instrumenting a modifier.

```go
overrides := map[string]sintax.GlobalModifier{
    "redact": func(value any, params []any) (any, error) {
        return "***", nil
    },
}
s := sintax.New(sintax.BuiltinFunctions(overrides, nil))
// {{ secret | redact }} → ***
```

`GlobalModifier` is `func(value any, params []any) (any, error)`. The first positional argument from the
template flows in as `value`; everything after the modifier name (separated by `,`) shows up in `params`.

---

## Optional extensions

`yaml` and `markdown` ship as stubs that return an error until you inject a real implementation, so the
core stays dependency-light. Wire in whatever library you prefer; the examples below just show one option each.

**YAML serialization**: the example uses `gopkg.in/yaml.v3`, but any YAML library works

```go
overrides := map[string]sintax.GlobalModifier{
    "yaml": func(value any, params []any) (any, error) {
        b, err := yaml.Marshal(value)
        return string(b), err
    },
}
s := sintax.New(sintax.BuiltinFunctions(overrides, nil))
```

**HTML to Markdown**: the example uses `github.com/JohannesKaufmann/html-to-markdown/v2`, but any converter works

```go
overrides := map[string]sintax.GlobalModifier{
    "markdown": func(value any, params []any) (any, error) {
        html, _ := functions.ValueString(value)
        conv := converter.NewConverter(converter.WithPlugins(base.NewBasePlugin()))
        return conv.ConvertString(html)
    },
}
s := sintax.New(sintax.BuiltinFunctions(overrides, nil))
```

---

## Contributing

`sintax` uses an issue-first workflow. Open an issue describing the change and wait for a maintainer to approve the approach (the `approved` label) before you open a pull request. PRs that don't reference an approved issue are flagged by a bot and usually closed, so the issue step saves you wasted work.

Every commit must be signed off for the [Developer Certificate of Origin](https://developercertificate.org/) with `git commit -s`. A CI check enforces this on every commit in a pull request.

Full flow in [CONTRIBUTING.md](CONTRIBUTING.md), ground rules in [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).

## Hosted code and health reports

Reports for this repo are hosted by our <a href="https://code.toawe.me">code viewer</a>, which also serves the badges and cards above.

<p align="center">
  <a href="https://code.toawe.me/toaweme/sintax/health"><picture><source media="(prefers-color-scheme: dark)" srcset="https://code.toawe.me/toaweme/sintax/card.svg"><source media="(prefers-color-scheme: light)" srcset="https://code.toawe.me/toaweme/sintax/card-light.svg"><img alt="sintax health" src="https://code.toawe.me/toaweme/sintax/card-light.svg" width="48%"></picture></a>
  <a href="https://code.toawe.me/toaweme/sintax/code"><picture><source media="(prefers-color-scheme: dark)" srcset="https://code.toawe.me/toaweme/sintax/code-card.svg"><source media="(prefers-color-scheme: light)" srcset="https://code.toawe.me/toaweme/sintax/code-card-light.svg"><img alt="sintax code" src="https://code.toawe.me/toaweme/sintax/code-card-light.svg" width="48%"></picture></a>
</p>

---

Made with ❤️ in Lithuania 🇱🇹.
</content>
