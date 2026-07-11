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

// ExampleToUpper converts the value to uppercase.
func ExampleToUpper() {
	fmt.Println(render(`{{ code | upper }}`, map[string]any{
		"code": "us-east-1",
	}))
	// Output: US-EAST-1
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

// ExampleSlug converts free text into a URL-friendly slug.
func ExampleSlug() {
	fmt.Println(render(`{{ title | slug }}`, map[string]any{
		"title": "Hello, World! Version 4.5",
	}))
	// Output: hello-world-version-4.5
}
