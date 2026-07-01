package convert

import (
	"fmt"

	"github.com/toaweme/sintax/functions"
	// "github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	// "github.com/JohannesKaufmann/html-to-markdown/v2/plugin/base"
	// "github.com/JohannesKaufmann/html-to-markdown/v2/plugin/table"
)

// func Markdown2(value any, params []any) (any, error) {
// html, err := functions.ValueString(value)
// if err != nil {
// 	return nil, fmt.Errorf("markdown function expected a string, got %T: %w", value, err)
// }
//
// conv := converter.NewConverter(
// 	converter.WithPlugins(
// 		base.NewBasePlugin(),
// 		table.NewTablePlugin(),
// 	),
// )
//
// markdown, err := conv.ConvertString(html)
// if err != nil {
// 	log.Fatal(err)
// }
//
// return markdown, nil
// }

// ModifierNameMarkdown is the template name for the Markdown modifier.
const ModifierNameMarkdown functions.ModifierName = "markdown"

// Markdown converts an HTML string to Markdown.
// Must be injected by the consumer; returns an error by default.
//
// value: string
// returns: string
//
// example: convert a simple HTML snippet
// in:  html_content = "<h1>Welcome</h1><p>Thanks for joining.</p>"
// tpl: {{ html_content | markdown }}
// out: # Welcome
// out: (blank line)
// out: Thanks for joining.
func Markdown(value any, params []any) (any, error) {
	return nil, fmt.Errorf("markdown function needs to be injected")
}
