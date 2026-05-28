_The sections below are for embedding sintax inside a Go program — instantiating the engine, registering
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

    // Variables can reference each other — sintax resolves them in order
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

`sintax.New(overrides)` takes an optional `map[string]sintax.GlobalModifier` — pass `nil` for the built-ins.
`Render` resolves variables first, then renders. Use `RenderSafe` if values inside the variable map should
**not** be re-parsed for `{{ ... }}` tokens (handy when those values may originate from untrusted input).

---

## Public API

The engine speaks through a small, stable set of interfaces. Everything else (`StringParser`,
`StringRenderer`, `BaseToken`) is an implementation detail you can replace via `NewWith`.
[[ range .Interfaces ]]
### `[[ .Name ]]`
[[ if .Comment ]]
[[ trim .Comment ]]
[[ end ]]
```go
type [[ trim .Snippet ]]
```
[[ end ]]
---

## Error handling

Every failure mode has a sentinel error you can match with `errors.Is`. Wrap or compare — never string-match.

```go
resolved, err := s.ResolveVariables(vars)
switch {
case errors.Is(err, sintax.ErrCircularDependency):
    // variable A references B which references A
case errors.Is(err, sintax.ErrVariableNotFound):
    // referenced variable missing — use | default:'...' to avoid this
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

Pass a map of overrides to `sintax.New`. Overrides also replace built-ins of the same name — useful for
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

**YAML serialization** — requires `gopkg.in/yaml.v3`

```go
overrides := map[string]sintax.GlobalModifier{
    "yaml": func(value any, params []any) (any, error) {
        b, err := yaml.Marshal(value)
        return string(b), err
    },
}
```

**HTML → Markdown** — requires `github.com/JohannesKaufmann/html-to-markdown/v2`

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
