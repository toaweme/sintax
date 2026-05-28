package docs

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	codereader "github.com/awee-ai/go-code-reader"
	"github.com/awee-ai/go-code-reader/readers/golang"
	"github.com/awee-ai/go-code-reader/utils"
	"github.com/toaweme/cli"
	"github.com/toaweme/log"
)

type DocsConfig struct {
	OutputDir  string `arg:"output_dir" short:"o" env:"OUTPUT_DIR" default:"_data/docs/sintax" help:"Directory to write per-modifier MDX files"`
	ProjectDir string `arg:"project_dir" short:"p" env:"PROJECT_DIR" default:".." help:"Root of the sintax module (contains functions/)"`
}

type DocsCommand struct {
	cli.BaseCommand[DocsConfig]
}

func NewDocsCommand() *DocsCommand {
	return &DocsCommand{}
}

func (c DocsCommand) Help() string {
	return "Generate MDX documentation for all built-in modifiers"
}

// ModifierParam documents one parameter of a modifier.
type ModifierParam struct {
	Index    int
	Variadic bool
	Type     string
	Required bool
}

// Example documents one usage scenario for a modifier.
//
// A short example may consist of just a Template (the legacy single-line form).
// A rich example pairs a Title, optional Input lines (variable assignments),
// the Template, and the rendered Output.
type Example struct {
	Title    string
	Input    []string
	Template string
	Output   []string
}

// Interface holds a public interface declaration to display in the README.
type Interface struct {
	Name    string
	Comment string
	Snippet string
}

// rootInterfaceOrder pins the display order of public interfaces in the
// README — the most user-facing types come first, internal building blocks
// last. Interfaces present in the source but missing from this list are
// appended in source order so a new interface never silently disappears.
var rootInterfaceOrder = []string{
	"Sintax",
	"Parser",
	"Renderer",
	"Token",
}

// rootInterfacePackage names the Go package whose interfaces appear in the
// README's "Interfaces" section.
const rootInterfacePackage = "sintax"

// extractRootInterfaces returns the interfaces declared in the root sintax
// package, ordered by rootInterfaceOrder with any unlisted interfaces
// appended in source order.
func extractRootInterfaces(content *codereader.Content) []Interface {
	if content == nil {
		return nil
	}
	byName := make(map[string]Interface, len(content.Interfaces))
	var fallback []Interface
	for _, iface := range content.Interfaces {
		if iface.Position.Package != rootInterfacePackage {
			continue
		}
		out := Interface{
			Name:    iface.Name,
			Comment: strings.TrimSpace(iface.Comment),
			Snippet: strings.TrimSpace(iface.Snippet),
		}
		byName[iface.Name] = out
		fallback = append(fallback, out)
	}

	if len(byName) == 0 {
		return nil
	}

	out := make([]Interface, 0, len(byName))
	seen := make(map[string]bool, len(rootInterfaceOrder))
	for _, name := range rootInterfaceOrder {
		if iface, ok := byName[name]; ok {
			out = append(out, iface)
			seen[name] = true
		}
	}
	for _, iface := range fallback {
		if !seen[iface.Name] {
			out = append(out, iface)
		}
	}
	return out
}

// ModifierDoc holds the extracted documentation for one modifier.
type ModifierDoc struct {
	Name        string
	Category    string
	Description string
	Accepts     []string
	Params      []ModifierParam
	Returns     string
	Examples    []Example
}

func (c DocsCommand) Run(opts cli.GlobalOptions, unknown cli.Unknowns) error {
	projectDir := c.Inputs.ProjectDir
	if !filepath.IsAbs(projectDir) {
		projectDir = filepath.Join(opts.Cwd, projectDir)
	}

	if err := GenerateWithDocgen(context.Background(), projectDir); err != nil {
		return err
	}

	// README generation uses sintax-specific templates and interface extraction
	// that aren't yet part of the docgen SDK. Keep using the old renderer for now.
	content, err := c.readProjectContent(projectDir)
	if err != nil {
		return fmt.Errorf("failed to read project for README: %w", err)
	}

	mods := c.extractModifiers(content)
	ifaces := extractRootInterfaces(content)

	pitchTpl, err := os.ReadFile(filepath.Join(projectDir, "cmd", "docs", "pitch.tpl.md"))
	if err != nil {
		return fmt.Errorf("failed to read pitch template: %w", err)
	}
	devTpl, err := os.ReadFile(filepath.Join(projectDir, "cmd", "docs", "dev.tpl.md"))
	if err != nil {
		return fmt.Errorf("failed to read dev template: %w", err)
	}

	renderer := NewModifierRenderer()
	readmeContent, err := renderer.RenderREADME(mods, ifaces, string(pitchTpl), string(devTpl))
	if err != nil {
		return fmt.Errorf("failed to render README: %w", err)
	}
	readmePath := filepath.Join(projectDir, "README.md")
	if err := os.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
		return fmt.Errorf("failed to write README: %w", err)
	}
	log.Info("docs", "write", "README.md", "to", readmePath)

	return nil
}

func (c DocsCommand) readProjectContent(cwd string) (*codereader.Content, error) {
	module, err := utils.ReadModulePath(cwd)
	if err != nil {
		return nil, fmt.Errorf("failed to read module path: %w", err)
	}

	// Walk the whole module so we collect root-package interfaces (Sintax,
	// Parser, Renderer, Token) alongside the per-package modifier funcs under
	// functions/. extractModifiers already ignores the root package by name.
	reader := golang.NewReader(golang.Config{
		Dir:                 cwd,
		Module:              module,
		Structs:             true,
		Interfaces:          true,
		Constants:           true,
		Vars:                true,
		Functions:           true,
		ExpandStdlibStructs: false,
	})
	content, err := reader.Walk()
	if err != nil {
		return nil, fmt.Errorf("failed to read project: %w", err)
	}

	return content, nil
}

func (c DocsCommand) filterOutTests(content *codereader.Content) codereader.ContentByFile {
	groups := content.GroupContentByFile()
	for file := range groups {
		if strings.HasSuffix(file, "_test.go") {
			delete(groups, file)
		}
	}
	return groups
}

func (c DocsCommand) extractModifiers(content *codereader.Content) []ModifierDoc {
	var mods []ModifierDoc

	groups := c.filterOutTests(content)
	for _, cont := range groups {
		category := categoryFromContent(cont)
		if category == "" || category == "functions" {
			continue
		}

		// build funcName → templateName from ModifierName constants
		nameMap := buildNameMap(cont.Constants)
		if len(nameMap) == 0 {
			continue
		}

		// match vars (var Trim = func(...))
		for _, varGroup := range cont.Vars {
			for _, v := range varGroup {
				templateName, ok := nameMap[v.Name]
				if !ok || v.Comment == "" {
					continue
				}
				mod := parseDocComment(v.Comment, templateName, category)
				mods = append(mods, mod)
			}
		}

		// match functions (func Filter(...))
		for _, fn := range cont.Functions {
			templateName, ok := nameMap[fn.Name]
			if !ok || fn.Comment == "" {
				continue
			}
			mod := parseDocComment(fn.Comment, templateName, category)
			mods = append(mods, mod)
		}
	}

	return mods
}

// buildNameMap extracts ModifierName constants and returns funcName → templateName.
// e.g., const ModifierNameTrim = "trim" → "Trim" → "trim"
func buildNameMap(constantGroups []codereader.Constants) map[string]string {
	nameMap := make(map[string]string)
	for _, group := range constantGroups {
		for _, c := range group {
			if !strings.HasPrefix(c.Name, "ModifierName") {
				continue
			}
			funcName := strings.TrimPrefix(c.Name, "ModifierName")
			// c.Value is the raw AST literal including Go quotes: `"trim"` → strip them
			templateName := strings.Trim(c.Value, `"`)
			nameMap[funcName] = templateName
		}
	}
	return nameMap
}

// categoryFromContent returns the Go package name from any element in the file content.
func categoryFromContent(cont *codereader.Content) string {
	for _, group := range cont.Constants {
		for _, c := range group {
			if c.Position.Package != "" {
				return c.Position.Package
			}
		}
	}
	for _, group := range cont.Vars {
		for _, v := range group {
			if v.Position.Package != "" {
				return v.Position.Package
			}
		}
	}
	for _, fn := range cont.Functions {
		if fn.Position.Package != "" {
			return fn.Position.Package
		}
	}
	return ""
}

// parseDocComment parses a cleaned doc comment string into a ModifierDoc.
//
// Expected format (all sections optional except description):
//
//	First line(s) until a blank line = description.
//	value: string, array
//	param:0: string
//	param:0?: string
//	param:...: string
//	param:...?: string
//	returns: string
//
// Examples come in two shapes — a legacy single-line form:
//
//	example: {{ value | modifier }}
//
// or a titled multi-field block (each field below the title may repeat):
//
//	example: title for this case
//	in:  name = "  Alice  "
//	tpl: {{ name | trim }}
//	out: Alice
func parseDocComment(text, name, category string) ModifierDoc {
	mod := ModifierDoc{Name: name, Category: category}

	var descLines []string
	inDesc := true

	// current example being built — flushed when a new top-level keyword is seen
	var curEx *Example
	flush := func() {
		if curEx != nil {
			mod.Examples = append(mod.Examples, *curEx)
			curEx = nil
		}
	}

	for _, raw := range strings.Split(strings.TrimSpace(text), "\n") {
		line := strings.TrimSpace(raw)

		switch {
		case strings.HasPrefix(line, "value:"):
			inDesc = false
			flush()
			for _, t := range strings.Split(strings.TrimSpace(strings.TrimPrefix(line, "value:")), ",") {
				if s := strings.TrimSpace(t); s != "" {
					mod.Accepts = append(mod.Accepts, s)
				}
			}

		case strings.HasPrefix(line, "param:"):
			inDesc = false
			flush()
			if p, ok := parseParam(line); ok {
				mod.Params = append(mod.Params, p)
			}

		case strings.HasPrefix(line, "returns:"):
			inDesc = false
			flush()
			mod.Returns = strings.TrimSpace(strings.TrimPrefix(line, "returns:"))

		case strings.HasPrefix(line, "example:"):
			inDesc = false
			flush()
			rest := strings.TrimSpace(strings.TrimPrefix(line, "example:"))
			// legacy single-line form: `example: {{ ... }}` becomes a bare-template example
			if strings.HasPrefix(rest, "{{") {
				mod.Examples = append(mod.Examples, Example{Template: rest})
				break
			}
			curEx = &Example{Title: rest}

		case strings.HasPrefix(line, "in:") && curEx != nil:
			curEx.Input = append(curEx.Input, strings.TrimSpace(strings.TrimPrefix(line, "in:")))

		case strings.HasPrefix(line, "tpl:") && curEx != nil:
			curEx.Template = strings.TrimSpace(strings.TrimPrefix(line, "tpl:"))

		case strings.HasPrefix(line, "out:") && curEx != nil:
			curEx.Output = append(curEx.Output, strings.TrimSpace(strings.TrimPrefix(line, "out:")))

		default:
			if inDesc {
				if line == "" {
					if len(descLines) > 0 {
						inDesc = false
					}
				} else {
					descLines = append(descLines, line)
				}
			}
		}
	}
	flush()

	mod.Description = strings.Join(descLines, " ")
	return mod
}

// parseParam parses a line like:
//
//	param:0: string       (required positional)
//	param:0?: string      (optional positional)
//	param:...: string     (required variadic)
//	param:...?: string    (optional variadic)
func parseParam(line string) (ModifierParam, bool) {
	re := regexp.MustCompile(`^param:(\d+|\.\.\.)(\?)?:\s*(.+)$`)
	m := re.FindStringSubmatch(line)
	if m == nil {
		return ModifierParam{}, false
	}
	variadic := m[1] == "..."
	optional := m[2] == "?"
	idx := 0
	if !variadic {
		idx, _ = strconv.Atoi(m[1])
	}
	return ModifierParam{
		Index:    idx,
		Variadic: variadic,
		Type:     strings.TrimSpace(m[3]),
		Required: !optional,
	}, true
}

func groupByCategory(mods []ModifierDoc) map[string][]ModifierDoc {
	cats := make(map[string][]ModifierDoc)
	for _, mod := range mods {
		cats[mod.Category] = append(cats[mod.Category], mod)
	}
	for cat := range cats {
		sort.Slice(cats[cat], func(i, j int) bool { return cats[cat][i].Name < cats[cat][j].Name })
	}
	return cats
}

func sortedKeys(m map[string][]ModifierDoc) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

var _ cli.Command[DocsConfig] = (*DocsCommand)(nil)
