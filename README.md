# Quick Guide

Templating engine built for workflows, document generation, and data transformations.

```
{{ response | from:'json' | key:'orders' | filter:'status','paid' | pluck:'total' | sum | decimal:2 }}
```

---

## Features

- **Pipe syntax**: chain 51 built-in modifiers to transform any value in a single expression
- **Dependency resolution**: variables can reference each other; sintax resolves them in the correct order automatically
- **Cycle detection**: circular variable references are caught at resolution time, not at runtime
- **Nested data**: maps, slices, structs, and pointers are all resolved and rendered recursively
- **Conditionals**: `{{ if x }} … {{ else }} … {{ endif }}` blocks
- **Loops**: `{{ for v in items }} … {{ endfor }}` over slices and maps, with auto-bound index/key helpers
- **Typed errors**: parse failures, missing variables, and circular dependencies are surfaced as distinct, branchable error categories
- **Extensible**: register your own modifiers alongside the built-ins

---

## Template syntax

| Pattern | Example |
|---|---|
| Variable | `{{ name }}` |
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
String literals use single or double quotes; unquoted tokens resolve as variables, numbers, or booleans.

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
| `<v>_first`, `<v>_last` | both: booleans for the first/last iteration |
| `<v>_key` | map iterations only, when no explicit key name was bound |

Map iteration order is sorted by key. Loops nest freely and parent variables remain visible inside the body.

```
{{ for item in cart }}{{ item.name }} × {{ item.qty }}{{ if item_last }}.{{ else }}, {{ endif }}{{ endfor }}
```

---

## Special variables

These names resolve to runtime-generated values without being declared.

| Name | Returns |
|---|---|
| `{{ uuid }}` | UUID v4 string (alias for `uuidv4`) |
| `{{ uuidv1 }}` | UUID v1 string (time + MAC) |
| `{{ now }}` | timestamp, pipe into `format` |
| `{{ [] }}` | empty array, useful with `default` |
| `{{ {} }}` | empty object, useful with `default` |

---

## Items

Each item ships with structured docs under [`./_data/docs/sintax`](./_data/docs/sintax). The tables below are a quick reference.

### Text

Trim, case-shift, slugify, split, and reshape strings.

| Item | Description | Example |
|----------|-------------|---------|
| [`concat`](./_data/docs/sintax/text/concat.mdx) | Concat appends one or more strings to the value. | `{{ greeting \| concat:'!' }}` |
| [`join`](./_data/docs/sintax/text/join.mdx) | Join combines an array of strings into a single string with a separator. | `{{ tags \| join:',' }}` |
| [`lines`](./_data/docs/sintax/text/lines.mdx) | Lines splits a string or byte slice into an array of lines. | `{{ note \| lines }}` |
| [`lower`](./_data/docs/sintax/text/lower.mdx) | ToLower converts a string to lowercase. | `{{ email \| lower }}` |
| [`replace`](./_data/docs/sintax/text/replace.mdx) | Replace replaces all occurrences of a substring within the string value. | `{{ greeting \| replace:'world','everyone' }}` |
| [`replace_pattern`](./_data/docs/sintax/text/replace_pattern.mdx) | ReplacePattern replaces all regex matches within the string value. | `{{ text \| replace_pattern:'\s+',' ' }}` |
| [`reverse`](./_data/docs/sintax/text/reverse.mdx) | Reverse reverses the characters in a string. | `{{ name \| reverse }}` |
| [`sexy`](./_data/docs/sintax/text/sexy.mdx) | Sexy returns a bear ASCII art. | `{{ anything \| sexy }}` |
| [`shorten`](./_data/docs/sintax/text/shorten.mdx) | Shorten truncates a string to the given maximum character length. | `{{ description \| shorten:30 }}` |
| [`slug`](./_data/docs/sintax/text/slug.mdx) | Slug converts a string to a URL-friendly slug. | `{{ title \| slug }}` |
| [`split`](./_data/docs/sintax/text/split.mdx) | Split splits a string into an array using a separator. | `{{ csv_line \| split:',' }}` |
| [`title`](./_data/docs/sintax/text/title.mdx) | Title converts a hyphen-separated slug into a title-cased string. | `{{ slug \| title }}` |
| [`title_model`](./_data/docs/sintax/text/title_model.mdx) | ModelTitle formats an AI model identifier into a human-readable title. | `{{ model_id \| title_model }}` |
| [`trim`](./_data/docs/sintax/text/trim.mdx) | Trim removes leading and trailing whitespace, or the given character set. | `{{ name \| trim }}` |
| [`trim-prefix`](./_data/docs/sintax/text/trim-prefix.mdx) | TrimPrefix removes a leading prefix string or leading whitespace from the value. | `{{ path \| trim-prefix:'/' }}` |
| [`trim-suffix`](./_data/docs/sintax/text/trim-suffix.mdx) | TrimSuffix removes a trailing suffix string or trailing whitespace from the value. | `{{ url \| trim-suffix:'/' }}` |
| [`upper`](./_data/docs/sintax/text/upper.mdx) | ToUpper converts a string to uppercase. | `{{ name \| upper }}` |

### Collections

Sort, filter, find, and reshape arrays and maps.

| Item | Description | Example |
|----------|-------------|---------|
| [`filter`](./_data/docs/sintax/collections/filter.mdx) | Filter returns a subset of a slice where a nested field matches a value. | `{{ items \| filter:'status','active' }}` |
| [`find`](./_data/docs/sintax/collections/find.mdx) | Find returns the first element in a slice or map where a field equals the given value. | `{{ users \| find:'id',42 }}` |
| [`first`](./_data/docs/sintax/collections/first.mdx) | First returns the first character of a string or the first element of a slice. | `{{ items \| first }}` |
| [`flatten`](./_data/docs/sintax/collections/flatten.mdx) | Flatten flattens a slice of slices by one level. | `{{ groups \| pluck:'items' \| flatten }}` |
| [`has`](./_data/docs/sintax/collections/has.mdx) | Has returns true if the slice or map contains the given value. | `{{ tags \| has:'featured' }}` |
| [`is`](./_data/docs/sintax/collections/is.mdx) | Is returns true if the value equals any of the given parameters. | `{{ status \| is:'active' }}` |
| [`key`](./_data/docs/sintax/collections/key.mdx) | Key extracts a value from a map or slice by key path or index. | `{{ user \| key:'name' }}` |
| [`last`](./_data/docs/sintax/collections/last.mdx) | Last returns the last character of a string or the last element of a slice. | `{{ items \| last }}` |
| [`map`](./_data/docs/sintax/collections/map.mdx) | Map converts a slice of maps into a map keyed by the given field's string value. | `{{ users \| map:'id' }}` |
| [`merge`](./_data/docs/sintax/collections/merge.mdx) | Merge converts a slice of maps into a map keyed by the given field's string value. | `{{ users \| merge:'id' }}` |
| [`pluck`](./_data/docs/sintax/collections/pluck.mdx) | Pluck extracts a single field from each element of a slice of maps and returns a slice of values. | `{{ users \| pluck:'id' }}` |
| [`sort`](./_data/docs/sintax/collections/sort.mdx) | Sort sorts a slice in ascending or descending order. | `{{ names \| sort }}` |
| [`sum`](./_data/docs/sintax/collections/sum.mdx) | Sum returns the numeric sum of the elements of a slice. | `{{ amounts \| sum }}` |
| [`wrap`](./_data/docs/sintax/collections/wrap.mdx) | Wrap wraps the value in a map under the given key. | `{{ name \| wrap:'user' }}` |

### Boolean

Compare values for use inside if/else blocks and conditional expressions.

| Item | Description | Example |
|----------|-------------|---------|
| [`eq`](./_data/docs/sintax/boolean/eq.mdx) | Eq returns true if the value equals the given parameter. | `{{ status \| eq:'active' }}` |
| [`gt`](./_data/docs/sintax/boolean/gt.mdx) | Gt returns true if the numeric value is greater than the threshold. | `{{ items_in_cart \| gt:0 }}` |
| [`gte`](./_data/docs/sintax/boolean/gte.mdx) | Gte returns true if the numeric value is greater than or equal to the threshold. | `{{ qty \| gte:1 }}` |
| [`not`](./_data/docs/sintax/boolean/not.mdx) | Not inverts the truthiness of the value. | `{{ is_active \| not }}` |

### Convert

Move between Go values, JSON, YAML, and other serialized formats.

| Item | Description | Example |
|----------|-------------|---------|
| [`from`](./_data/docs/sintax/convert/from.mdx) | From parses the string value as the given format and returns the parsed result. | `{{ body \| from:'json' }}` |
| [`json`](./_data/docs/sintax/convert/json.mdx) | JSON serializes the value to a JSON string. | `{{ user \| json }}` |
| [`markdown`](./_data/docs/sintax/convert/markdown.mdx) | Markdown converts an HTML string to Markdown. | `{{ html_content \| markdown }}` |
| [`yaml`](./_data/docs/sintax/convert/yaml.mdx) | YAML serializes or parses a value as YAML. | `{{ config \| yaml }}` |

### Utilities

Defaults, lengths, line numbers, and date formatting.

| Item | Description | Example |
|----------|-------------|---------|
| [`decimal`](./_data/docs/sintax/utils/decimal.mdx) | Decimal formats a number with a fixed number of decimal places. | `{{ amount \| decimal:2 }}` |
| [`default`](./_data/docs/sintax/utils/default.mdx) | Default returns the fallback value if the input is nil or an empty string. | `{{ name \| default:'anonymous' }}` |
| [`format`](./_data/docs/sintax/utils/format.mdx) | Format formats a time.Time value using a date format string. | `{{ created_at \| format:'YYYY-MM-DD' }}` |
| [`length`](./_data/docs/sintax/utils/length.mdx) | Length returns the number of characters in a string, bytes in a byte slice, or elements in a slice/array/map. | `{{ name \| length }}` |
| [`line-numbers`](./_data/docs/sintax/utils/line-numbers.mdx) | LineNumbers prepends each line of the string with its zero-based line number. | `{{ note \| line-numbers }}` |

### File System

Pull pieces out of file paths - directory, name, and extension.

| Item | Description | Example |
|----------|-------------|---------|
| [`dirname`](./_data/docs/sintax/fs/dirname.mdx) | Dirname returns the directory portion of a file path. | `{{ file_path \| dirname }}` |
| [`ext`](./_data/docs/sintax/fs/ext.mdx) | FilenameExt returns the file extension without the leading dot. | `{{ file_path \| ext }}` |
| [`ext-dot`](./_data/docs/sintax/fs/ext-dot.mdx) | FilenameExtDot returns the file extension including the leading dot. | `{{ file_path \| ext-dot }}` |
| [`ext-prepend`](./_data/docs/sintax/fs/ext-prepend.mdx) | FilenamePrependExt inserts an additional extension before the existing file extension. | `{{ file_path \| ext-prepend:'min' }}` |
| [`ext-trim`](./_data/docs/sintax/fs/ext-trim.mdx) | FilenameTrimExt returns the file path without its extension. | `{{ file_path \| ext-trim }}` |
| [`file`](./_data/docs/sintax/fs/file.mdx) | File reads a file's contents from an allowlisted directory. | `{{ "greeting.tpl" \| file }}` |
| [`filename`](./_data/docs/sintax/fs/filename.mdx) | Filename returns the base file name from a path, including the extension. | `{{ file_path \| filename }}` |

`file` reads from disk, so it only works against an allowlist of directories you
pass in: `BuiltinFunctions(overrides, safeDirs)`. Paths are resolved against
those dirs and any `..` escape is rejected. With no `safeDirs` (`nil`), `file`
always errors.

### Money

Convert numbers between currency units like dollars and cents.

| Item | Description | Example |
|----------|-------------|---------|
| [`currency`](./_data/docs/sintax/money/currency.mdx) | Currency converts a numeric value between currency units by applying a unit multiplier ratio. | `{{ price \| currency:1,100 }}` |

---

_The sections below are for embedding sintax inside a Go program: instantiating the engine, registering
custom modifiers, surfacing typed errors, and integrating with workflow orchestrators. The template syntax
itself is documented above._

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
    s := sintax.New(nil)

    // Variables can reference each other, sintax resolves them in order
    vars := map[string]any{
        "env":     "production",
        "prefix":  "{{ env | upper }}",
        "db_name": "{{ prefix }}_database",
    }

    resolved, _ := s.ResolveVariables(vars)
    fmt.Println(resolved["db_name"]) // PRODUCTION_database

    // Render a one-off template
    out, _ := s.Render("Hello, {{ name | title }}!", map[string]any{"name": "alice-cooper"})
    fmt.Println(out) // Hello, Alice Cooper!
}
```

`sintax.New(overrides)` takes an optional `map[string]sintax.GlobalModifier`; pass `nil` for the built-ins.
`Render` resolves variables first, then renders. Use `RenderSafe` if values inside the variable map should
**not** be re-parsed for `{{ ... }}` tokens (handy when those values may originate from untrusted input).

---

## Public API

The engine speaks through a small, stable set of interfaces. Everything else (`StringParser`,
`StringRenderer`, `BaseToken`) is an implementation detail you can replace via `NewWith`.

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
resolved, err := s.ResolveVariables(vars)
switch {
case errors.Is(err, sintax.ErrCircularDependency):
    // variable A references B which references A
case errors.Is(err, sintax.ErrVariableNotFound):
    // referenced variable missing, use | default:'...' to avoid this
case errors.Is(err, sintax.ErrParseFailed):
    // template grammar error
case errors.Is(err, sintax.ErrRenderFailed):
    // a modifier returned an error or the value type was unexpected
case err != nil:
    // unclassified failure
}
```

Other errors worth knowing about: `sintax.ErrInvalidTokenType`, `sintax.ErrFunctionNotFound`,
`sintax.ErrFunctionApplyFailed`.

---

## Custom modifiers

Pass a map of overrides to `sintax.New`. Overrides also replace built-ins of the same name, useful for
sandboxing or instrumenting a modifier.

```go
overrides := map[string]sintax.GlobalModifier{
    "redact": func(value any, params []any) (any, error) {
        return "***", nil
    },
}
s := sintax.New(overrides)
// {{ secret | redact }} → ***
```

`GlobalModifier` is `func(value any, params []any) (any, error)`. The first positional argument from the
template flows in as `value`; everything after the modifier name (separated by `,`) shows up in `params`.

### Adding special variables

`uuid`, `now`, `[]`, and `{}` are registered globally. Add your own with `SetSpecialVariable`:

```go
sintax.SetSpecialVariable("today", func() any {
    return time.Now().Format("2006-01-02")
})
// {{ today }} → 2026-04-30
```

---

## Optional extensions

A handful of modifiers ship as stubs so the core stays dependency-light. Wire them up in your project:

**YAML serialization**: requires `gopkg.in/yaml.v3`

```go
overrides := map[string]sintax.GlobalModifier{
    "yaml": func(value any, params []any) (any, error) {
        b, err := yaml.Marshal(value)
        return string(b), err
    },
}
```

**HTML → Markdown**: requires `github.com/JohannesKaufmann/html-to-markdown/v2`

```go
overrides := map[string]sintax.GlobalModifier{
    "markdown": func(value any, params []any) (any, error) {
        html, _ := functions.ValueString(value)
        conv := converter.NewConverter(converter.WithPlugins(base.NewBasePlugin()))
        return conv.ConvertString(html)
    },
}
```

---

## Workflow engine integration

Sintax exposes the dependency graph so an orchestrator can run nodes in topological order, then evaluate
conditional branches between them:

```go
vars := map[string]any{
    "api_result": `{{ fetch_data | default:{} }}`,
    "items":      `{{ api_result | from:'json' | key:'items' | default:[] }}`,
    "count":      `{{ items | length }}`,
}

// Topological order: [fetch_data, api_result, items, count]
deps, _ := s.ExtractDependencies(vars)

// Evaluate a conditional branch
proceed, _ := s.ResolveCondition("{{ count | gt:0 }}", vars)
```

Use `ExtractDependenciesWithOptions` / `ResolveVariablesWithOptions` and a populated
`sintax.ResolveOptions{LiteralVars: ...}` to mark variables whose values must **not** be parsed as templates.

---

## License

MIT
