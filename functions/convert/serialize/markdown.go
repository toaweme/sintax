package serialize

import (
	"errors"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameMarkdown is the template name for the Markdown modifier.
const ModifierNameMarkdown functions.ModifierName = "markdown"

// Markdown converts an HTML string to Markdown. The library ships only a stub,
// because a full HTML-to-Markdown converter pulls in heavy dependencies the core
// should not force on every consumer. Applications that want this modifier inject
// their own implementation by overriding the modifier entry. Until injected,
// calling it returns an error.
func Markdown(value any) (string, error) {
	return "", errors.New("markdown function needs to be injected")
}
