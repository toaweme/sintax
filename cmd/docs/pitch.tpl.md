Templating engine built for workflows, document generation, and data transformations.

```
{{ response | from:'json' | key:'orders' | filter:'status','paid' | pluck:'total' | sum | decimal:2 }}
```

---

## Features

- **Pipe syntax** — chain [[ .TotalModifiers ]] built-in modifiers to transform any value in a single expression
- **Dependency resolution** — variables can reference each other; sintax resolves them in the correct order automatically
- **Cycle detection** — circular variable references are caught at resolution time, not at runtime
- **Nested data** — maps, slices, structs, and pointers are all resolved and rendered recursively
- **Conditionals** — `{{ if x }} … {{ else }} … {{ endif }}` blocks
- **Loops** — `{{ for v in items }} … {{ endfor }}` over slices and maps, with auto-bound index/key helpers
- **Typed errors** — parse failures, missing variables, and circular dependencies are surfaced as distinct, branchable error categories
- **Extensible** — register your own modifiers alongside the built-ins

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

Block control tags (`if`/`else`/`endif`/`for`/`endfor`) that sit alone on their own line are auto-trimmed —
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

Inside the loop body the following helpers are available — `<v>` is whatever name you bound the value to:

| Binding | Set on |
|---|---|
| `<v>_index` | both slice and map iterations (0-based) |
| `<v>_first`, `<v>_last` | both — booleans for the first/last iteration |
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
| `{{ now }}` | timestamp — pipe into `format` |
| `{{ [] }}` | empty array — useful with `default` |
| `{{ {} }}` | empty object — useful with `default` |
