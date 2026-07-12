package casing_test

import (
	"fmt"

	"github.com/toaweme/sintax"
	casing "github.com/toaweme/sintax/functions/text/case"
)

func render(tpl string, vars map[string]any) string {
	out, err := sintax.New(casing.Modifiers()).Render(tpl, vars)
	if err != nil {
		return fmt.Sprintf("error: %v", err)
	}
	return fmt.Sprintf("%v", out)
}

// ExampleToLower converts the value to lowercase.
func ExampleToLower() {
	fmt.Println(render(`{{ name | lower }}`, map[string]any{
		"name": "Ada LOVELACE",
	}))
	// Output: ada lovelace
}

// ExampleToLower_unicode lowercases a multi-byte string, keeping the accented
// letters intact.
func ExampleToLower_unicode() {
	fmt.Println(render(`{{ city | lower }}`, map[string]any{
		"city": "MÜNCHEN",
	}))
	// Output: münchen
}

// ExampleToLower_unchanged leaves a string that is already lowercase untouched.
func ExampleToLower_unchanged() {
	fmt.Println(render(`{{ tag | lower }}`, map[string]any{
		"tag": "already-lower",
	}))
	// Output: already-lower
}

// ExampleToUpper converts the value to uppercase.
func ExampleToUpper() {
	fmt.Println(render(`{{ code | upper }}`, map[string]any{
		"code": "us-east-1",
	}))
	// Output: US-EAST-1
}

// ExampleToUpper_literal uppercases a string literal with no variables.
func ExampleToUpper_literal() {
	fmt.Println(render(`{{ "hello" | upper }}`, nil))
	// Output: HELLO
}

// ExampleToUpper_unicode uppercases a multi-byte string, keeping the accent.
func ExampleToUpper_unicode() {
	fmt.Println(render(`{{ word | upper }}`, map[string]any{
		"word": "café",
	}))
	// Output: CAFÉ
}

// ExampleTitle title-cases a hyphen-separated slug and uppercases known
// acronyms such as API.
func ExampleTitle() {
	fmt.Println(render(`{{ slug | title }}`, map[string]any{
		"slug": "user-api-gateway",
	}))
	// Output: User API Gateway
}

// ExampleModelTitle formats a raw model identifier into a human-readable title.
func ExampleModelTitle() {
	fmt.Println(render(`{{ model | title_model }}`, map[string]any{
		"model": "openai/gpt-4o-mini",
	}))
	// Output: OpenAI: GPT 4o Mini
}

// ExampleTitle_extraAcronyms uppercases extra acronyms passed as parameters
// alongside the built-in list.
func ExampleTitle_extraAcronyms() {
	fmt.Println(render(`{{ heading | title:'seo','cta' }}`, map[string]any{
		"heading": "seo-and-cta-tips",
	}))
	// Output: SEO And CTA Tips
}

// ExampleTitle_acronyms uppercases multiple built-in acronyms in one slug.
func ExampleTitle_acronyms() {
	fmt.Println(render(`{{ heading | title }}`, map[string]any{
		"heading": "json-api-schema",
	}))
	// Output: JSON API Schema
}

// ExampleTitle_numbers keeps numeric segments as their own words.
func ExampleTitle_numbers() {
	fmt.Println(render(`{{ slug | title }}`, map[string]any{
		"slug": "hello-world-2023",
	}))
	// Output: Hello World 2023
}

// ExampleModelTitle_sizeSuffix normalizes a parameter-size suffix and the
// provider name.
func ExampleModelTitle_sizeSuffix() {
	fmt.Println(render(`{{ model | title_model }}`, map[string]any{
		"model": "meta-llama/llama-3.1-8b-instruct",
	}))
	// Output: Meta Llama: Llama 3.1 8B Instruct
}

// ExampleModelTitle_dateSuffix moves a trailing date into a parenthetical.
func ExampleModelTitle_dateSuffix() {
	fmt.Println(render(`{{ model | title_model }}`, map[string]any{
		"model": "anthropic/claude-3.5-haiku-20241022",
	}))
	// Output: Anthropic: Claude 3.5 Haiku (2024-10-22)
}

// ExampleModelTitle_ollama formats an Ollama tag with a quantization suffix.
func ExampleModelTitle_ollama() {
	fmt.Println(render(`{{ model | title_model }}`, map[string]any{
		"model": "olmo-3:7b-instruct-q4_K_M",
	}))
	// Output: Olmo 3 7B Instruct Q4_K_M
}

// ExampleSlug converts free text into a URL-friendly slug.
func ExampleSlug() {
	fmt.Println(render(`{{ title | slug }}`, map[string]any{
		"title": "Hello, World! Version 4.5",
	}))
	// Output: hello-world-version-4.5
}

// ExampleSlug_versionDots keeps dots between digits so version numbers survive.
func ExampleSlug_versionDots() {
	fmt.Println(render(`{{ title | slug }}`, map[string]any{
		"title": "Version 1.2.3 Release",
	}))
	// Output: version-1.2.3-release
}

// ExampleSlug_unicode drops non-ASCII letters rather than transliterating them.
func ExampleSlug_unicode() {
	fmt.Println(render(`{{ title | slug }}`, map[string]any{
		"title": "Café Münchën",
	}))
	// Output: caf-m-nch-n
}

// ExampleSlug_collapsesSpaces collapses runs of spaces and punctuation into
// single hyphens.
func ExampleSlug_collapsesSpaces() {
	fmt.Println(render(`{{ title | slug }}`, map[string]any{
		"title": "  Draft --- Notes & Ideas  ",
	}))
	// Output: draft-notes-ideas
}
