package docs

import "sort"

// Group describes one modifier category — a logical grouping that becomes a
// section heading in the rendered docs (top-level index, per-category index,
// and the README).
type Group struct {
	Title       string
	Description string
}

// groupOrder lists categories in the order they should appear in indexes and
// the README — the most-reached-for categories first, infrastructure-flavored
// ones last. Categories not listed here are still rendered (see
// orderedCategories) but appended alphabetically after the curated set.
var groupOrder = []string{
	"text",
	"collections",
	"boolean",
	"convert",
	"utils",
	"fs",
	"money",
}

// groups is the registry of category metadata, keyed by Go package name.
// Descriptions are written for public docs — short, plain, and lead with what
// the modifiers in the group help you do, not how they're implemented.
var groups = map[string]Group{
	"text": {
		Title:       "Text",
		Description: "Trim, case-shift, slugify, split, and reshape strings.",
	},
	"collections": {
		Title:       "Collections",
		Description: "Sort, filter, find, and reshape arrays and maps.",
	},
	"boolean": {
		Title:       "Boolean",
		Description: "Compare values for use inside if/else blocks and conditional expressions.",
	},
	"convert": {
		Title:       "Convert",
		Description: "Move between Go values, JSON, YAML, and other serialized formats.",
	},
	"utils": {
		Title:       "Utilities",
		Description: "Defaults, lengths, line numbers, and date formatting.",
	},
	"fs": {
		Title:       "File System",
		Description: "Pull pieces out of file paths — directory, name, and extension.",
	},
	"money": {
		Title:       "Money",
		Description: "Convert numbers between currency units like dollars and cents.",
	},
}

// lookupGroup returns the metadata for a category, falling back to a
// capitalized name and an empty description when the category is not
// registered. Use this everywhere a Title or Description is needed so that an
// unregistered category never produces a blank heading.
func lookupGroup(category string) Group {
	if g, ok := groups[category]; ok {
		return g
	}
	return Group{Title: capitalize(category)}
}

// orderedCategories returns category names in the curated groupOrder, then
// appends any categories present in `present` that weren't listed there
// (alphabetically). This guarantees that adding a new modifier package never
// silently drops it from the rendered output — the worst case is alphabetical
// placement until someone updates groupOrder.
func orderedCategories(present map[string][]ModifierDoc) []string {
	out := make([]string, 0, len(present))
	seen := make(map[string]bool, len(groupOrder))
	for _, c := range groupOrder {
		if _, ok := present[c]; ok {
			out = append(out, c)
			seen[c] = true
		}
	}
	var extras []string
	for c := range present {
		if !seen[c] {
			extras = append(extras, c)
		}
	}
	sort.Strings(extras)
	return append(out, extras...)
}
