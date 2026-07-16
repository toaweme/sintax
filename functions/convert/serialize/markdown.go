package serialize

import (
	"errors"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameMarkdown is the template name for the Markdown modifier.
const ModifierNameMarkdown functions.ModifierName = "markdown"

// Markdown converts an HTML string to Markdown, so a fetched page or a rich-text
// field can be reduced to plain prose a later step can store or render.
//
// It ships as a stub that returns an error until you inject a converter. A full
// HTML-to-Markdown converter is a heavy dependency, and the core stays free of
// third-party dependencies rather than forcing one on every consumer. Injecting
// one is a map entry.
//
//	func toMarkdown(html string) (string, error) {
//		return htmltomarkdown.ConvertString(html)
//	}
//
//	mods := serialize.Modifiers()
//	mods[string(serialize.ModifierNameMarkdown)] = functions.Wrap(toMarkdown)
//	out, err := sintax.New(mods).Render(tpl, vars)
func Markdown(html string) (string, error) {
	return "", errors.New("markdown function needs to be injected")
}
