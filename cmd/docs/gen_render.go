package docs

import (
	"fmt"
	"strings"
	"text/template"
)

// ModifierRenderer renders modifier documentation to MDX.
type ModifierRenderer struct{}

func NewModifierRenderer() *ModifierRenderer {
	return &ModifierRenderer{}
}

// sanitizeMDX escapes characters that MDX interprets as JSX expressions.
func sanitizeMDX(s string) string {
	s = strings.ReplaceAll(s, "{", `\{`)
	s = strings.ReplaceAll(s, "}", `\}`)
	return s
}

// sanitizeFrontmatter produces a safe single-line YAML string value.
func sanitizeFrontmatter(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, `"`, `'`)
	return s
}

// sanitizeTableCell escapes characters that break MDX table cells.
func sanitizeTableCell(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "|", `\|`)
	return sanitizeMDX(s)
}

func firstSentence(s string) string {
	if idx := strings.IndexByte(s, '.'); idx >= 0 {
		return s[:idx+1]
	}
	return s
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

// RenderModifier produces the MDX content for a single modifier, mirroring RenderComponent.
func (r *ModifierRenderer) RenderModifier(mod ModifierDoc) string {
	var b strings.Builder

	b.WriteString("---\n")
	fmt.Fprintf(&b, "title: \"%s\"\n", sanitizeFrontmatter(mod.Name))
	fmt.Fprintf(&b, "description: \"%s\"\n", sanitizeFrontmatter(firstSentence(mod.Description)))
	fmt.Fprintf(&b, "category: %s\n", mod.Category)
	b.WriteString("---\n\n")

	b.WriteString(sanitizeMDX(mod.Description) + "\n")

	if len(mod.Accepts) > 0 {
		b.WriteString("\n## Input\n\n")
		for _, a := range mod.Accepts {
			fmt.Fprintf(&b, "`%s` ", a)
		}
		b.WriteString("\n")
	}

	b.WriteString("\n## Parameters\n\n")
	if len(mod.Params) > 0 {
		b.WriteString("| # | Type | Required |\n")
		b.WriteString("|---|------|----------|\n")
		for _, p := range mod.Params {
			req := "no"
			if p.Required {
				req = "yes"
			}
			idx := fmt.Sprintf("%d", p.Index)
			if p.Variadic {
				idx = "..."
			}
			fmt.Fprintf(&b, "| %s | `%s` | %s |\n", idx, p.Type, req)
		}
	} else {
		b.WriteString("None\n")
	}

	b.WriteString("\n## Returns\n\n")
	if mod.Returns != "" {
		fmt.Fprintf(&b, "`%s`\n", mod.Returns)
	} else {
		b.WriteString("`any`\n")
	}

	if len(mod.Examples) > 0 {
		b.WriteString("\n## Examples\n")
		for i, ex := range mod.Examples {
			renderExample(&b, ex, i+1)
		}
	}

	return b.String()
}

// renderExample writes a single example block. Bare-template (legacy) examples
// render as a fenced template snippet; rich examples render with titled
// sub-sections for input, template, and output.
func renderExample(b *strings.Builder, ex Example, num int) {
	bare := ex.Title == "" && len(ex.Input) == 0 && len(ex.Output) == 0
	if bare {
		if num == 1 {
			b.WriteString("\n```\n")
			b.WriteString(ex.Template + "\n")
			b.WriteString("```\n")
			return
		}
		b.WriteString("\n```\n" + ex.Template + "\n```\n")
		return
	}

	title := ex.Title
	if title == "" {
		title = fmt.Sprintf("Example %d", num)
	}
	fmt.Fprintf(b, "\n### %s\n\n", capitalize(title))

	if len(ex.Input) > 0 {
		b.WriteString("**Input**\n\n```\n")
		for _, in := range ex.Input {
			b.WriteString(in + "\n")
		}
		b.WriteString("```\n\n")
	}

	if ex.Template != "" {
		b.WriteString("**Template**\n\n```\n")
		b.WriteString(ex.Template + "\n")
		b.WriteString("```\n")
	}

	if len(ex.Output) > 0 {
		b.WriteString("\n**Output**\n\n```\n")
		for _, out := range ex.Output {
			b.WriteString(out + "\n")
		}
		b.WriteString("```\n")
	}
}

// READMEGroup holds one category's worth of modifiers for README rendering.
type READMEGroup struct {
	Name        string
	Title       string
	Description string
	Modifiers   []ModifierDoc
}

// READMEData is passed to the README Go template.
type READMEData struct {
	TotalModifiers int
	Groups         []READMEGroup
	Interfaces     []Interface
}

// buildREADMEData turns parsed modifier docs and root interfaces into the data
// shape expected by the pitch and dev templates. Both templates share the same
// shape so the same data flows into either.
func buildREADMEData(mods []ModifierDoc, ifaces []Interface) READMEData {
	cats := groupByCategory(mods)
	catNames := orderedCategories(cats)

	groups := make([]READMEGroup, 0, len(catNames))
	for _, cat := range catNames {
		g := lookupGroup(cat)
		groups = append(groups, READMEGroup{
			Name:        cat,
			Title:       g.Title,
			Description: g.Description,
			Modifiers:   cats[cat],
		})
	}
	return READMEData{
		TotalModifiers: len(mods),
		Groups:         groups,
		Interfaces:     ifaces,
	}
}

// readmeFuncMap returns the template helpers shared by every README/index
// template: trimming, first-sentence extraction, and first-example pickup.
func readmeFuncMap() template.FuncMap {
	return template.FuncMap{
		"firstSentence": firstSentence,
		"firstExample": func(examples []Example) string {
			for _, ex := range examples {
				if ex.Template != "" {
					return ex.Template
				}
			}
			return ""
		},
		"trim": strings.TrimSpace,
	}
}

// executeTemplate parses and runs one of the docs templates ([[ ]] delimiters)
// against the supplied data, returning the rendered body.
func executeTemplate(name, content string, data any) (string, error) {
	tmpl, err := template.New(name).Delims("[[", "]]").Funcs(readmeFuncMap()).Parse(content)
	if err != nil {
		return "", fmt.Errorf("parse %s template: %w", name, err)
	}
	var b strings.Builder
	if err := tmpl.Execute(&b, data); err != nil {
		return "", fmt.Errorf("render %s template: %w", name, err)
	}
	return b.String(), nil
}

// renderModifierListing produces the grouped modifier reference table as
// Markdown. linkFor builds the per-modifier href so the same listing can be
// reused from the README (relative to repo root) and the docs index (relative
// to its own folder). Pipe characters inside the example column are escaped
// to `\|` so the GFM table parser doesn't split a row mid-cell — without that
// escape, the trailing `{{ ... }}` of the example leaks out of its inline
// code span and MDX/fumadocs tries to parse it as a JSX expression.
func (r *ModifierRenderer) renderModifierListing(data READMEData, linkFor func(category, name string) string) string {
	var b strings.Builder
	for _, group := range data.Groups {
		fmt.Fprintf(&b, "### %s\n\n", group.Title)
		if group.Description != "" {
			fmt.Fprintf(&b, "%s\n\n", group.Description)
		}
		b.WriteString("| Modifier | Description | Example |\n")
		b.WriteString("|----------|-------------|---------|\n")
		for _, mod := range group.Modifiers {
			example := ""
			for _, ex := range mod.Examples {
				if ex.Template != "" {
					example = ex.Template
					break
				}
			}
			fmt.Fprintf(&b, "| [`%s`](%s) | %s | `%s` |\n",
				mod.Name,
				linkFor(mod.Category, mod.Name),
				sanitizeTableCell(firstSentence(mod.Description)),
				strings.ReplaceAll(example, "|", `\|`),
			)
		}
		b.WriteString("\n")
	}
	return b.String()
}

// RenderREADME glues the pitch and dev templates together with a modifier
// listing in between, producing the repository README.
func (r *ModifierRenderer) RenderREADME(mods []ModifierDoc, ifaces []Interface, pitchTpl, devTpl string) (string, error) {
	data := buildREADMEData(mods, ifaces)

	pitch, err := executeTemplate("pitch", pitchTpl, data)
	if err != nil {
		return "", err
	}
	dev, err := executeTemplate("dev", devTpl, data)
	if err != nil {
		return "", err
	}
	listing := r.renderModifierListing(data, func(cat, name string) string {
		return fmt.Sprintf("./_data/docs/sintax/%s/%s.mdx", cat, name)
	})

	var b strings.Builder
	// b.WriteString("# sintax\n\n")
	b.WriteString("# Quick Guide\n\n")
	b.WriteString(strings.TrimSpace(pitch))
	b.WriteString("\n\n---\n\n## Modifiers\n\n")
	b.WriteString("Each modifier ships with structured docs (input, parameters, return type, and worked\n")
	b.WriteString("examples) under [`_data/docs/sintax`](./_data/docs/sintax). The tables below are a quick\n")
	b.WriteString("reference — click through for sample inputs and outputs.\n\n")
	b.WriteString(listing)
	b.WriteString("---\n\n")
	b.WriteString(strings.TrimSpace(dev))
	b.WriteString("\n")
	return b.String(), nil
}

// RenderCategoryIndex produces an index.mdx for a modifier category, mirroring RenderGroupIndex.
func (r *ModifierRenderer) RenderCategoryIndex(category string, mods []ModifierDoc) string {
	var b strings.Builder

	g := lookupGroup(category)
	title := g.Title
	description := g.Description
	if description == "" {
		description = title + " modifiers"
	}

	b.WriteString("---\n")
	fmt.Fprintf(&b, "title: \"%s\"\n", title)
	fmt.Fprintf(&b, "description: \"%s\"\n", description)
	b.WriteString("---\n\n")

	fmt.Fprintf(&b, "## %s\n\n", title)
	b.WriteString("| Modifier | Description |\n")
	b.WriteString("|----------|-------------|\n")
	for _, mod := range mods {
		fmt.Fprintf(&b, "| [`%s`](./%s) | %s |\n",
			mod.Name, mod.Name, sanitizeTableCell(firstSentence(mod.Description)))
	}
	b.WriteString("\n")

	return b.String()
}

// RenderIndex produces the top-level _data/docs/sintax/index.mdx for the docs
// site: YAML frontmatter, the pitch template body, then the modifier listing.
// The Public API and Go integration sections live exclusively in the README.
func (r *ModifierRenderer) RenderIndex(mods []ModifierDoc, pitchTpl string) (string, error) {
	data := buildREADMEData(mods, nil)

	pitch, err := executeTemplate("pitch", pitchTpl, data)
	if err != nil {
		return "", err
	}
	listing := r.renderModifierListing(data, func(cat, name string) string {
		return fmt.Sprintf("./%s/%s", cat, name)
	})

	var b strings.Builder
	b.WriteString("---\n")
	b.WriteString("title: \"Quick Guide\"\n")
	b.WriteString("icon: \"GraduationCap\"\n")
	b.WriteString("description: \"An overview of the functionality and all of the modifiers\"\n")
	b.WriteString("---\n\n")
	b.WriteString(strings.TrimSpace(pitch))
	b.WriteString("\n\n---\n\n## Modifiers\n\n")
	b.WriteString("Each modifier links to its own page with full inputs, parameters, and examples.\n\n")
	b.WriteString(listing)
	return b.String(), nil
}
