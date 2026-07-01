package text

import (
	"strings"
	"unicode"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameModelTitle is the template name for the ModelTitle modifier.
const ModifierNameModelTitle functions.ModifierName = "title_model"

// ModifierNameTitle is the template name for the Title modifier.
const ModifierNameTitle functions.ModifierName = "title"

// ModelTitle formats an AI model identifier into a human-readable title.
// Handles provider prefixes, model sizes, quantization suffixes, and date suffixes.
//
// value: string
// returns: string
//
// example: format a hosted model id
// in:  model_id = "openai/gpt-4o"
// tpl: {{ model_id | title_model }}
// out: OpenAI: GPT 4o
//
// example: format an Ollama tag with quantization
// in:  model_id = "llama3.1:8b-instruct-q4_K_M"
// tpl: {{ model_id | title_model }}
// out: Llama3.1 8B Instruct Q4_K_M
func ModelTitle(value any, params []any) (any, error) {
	model, err := functions.ValueString(value)
	if err != nil {
		return nil, err
	}
	return FormatModelTitle(model), nil
}

// Title converts a hyphen-separated slug into a title-cased string.
// Known acronyms (AI, API, GPT, etc.) are uppercased automatically.
// Additional acronyms can be passed as parameters.
//
// value: string
// param:...?: string
// returns: string
//
// example: title-case a blog slug
// in:  slug = "welcome-to-the-club"
// tpl: {{ slug | title }}
// out: Welcome To The Club
//
// example: keep recognized acronyms uppercase
// in:  slug = "the-ai-revolution"
// tpl: {{ slug | title }}
// out: The AI Revolution
//
// example: pass extra acronyms to keep upper-case
// in:  post_slug = "seo-and-cta-tips"
// tpl: {{ post_slug | title:'seo','cta' }}
// out: SEO And CTA Tips
func Title(value any, params []any) (any, error) {
	// convert slug to title
	slug, err := functions.ValueString(value)
	if err != nil {
		return nil, err
	}

	providedAcronyms, err := functions.ParamStringList(params)
	if err != nil {
		return nil, err
	}

	var all []string
	all = append(all, acronyms...)

	if len(providedAcronyms) > 0 {
		all = append(all, providedAcronyms...)
	}

	return slugToTitle(slug, all), nil
}

var acronyms = []string{
	"ai",
	"gpt",
	"oss",
	"api",
	"json",
	"xml",
	"html",
	"css",
	"tars",
}

func foundAcronym(word string, acronyms []string) bool {
	for _, acronym := range acronyms {
		if strings.EqualFold(word, acronym) {
			return true
		}
	}
	return false
}

func slugToTitle(slug string, acronyms []string) string {
	words := strings.Split(slug, "-")

	for i, word := range words {
		if word == "" {
			continue
		}

		lowerWord := strings.ToLower(word)
		if foundAcronym(lowerWord, acronyms) {
			words[i] = strings.ToUpper(lowerWord)
			continue
		}

		runes := []rune(word)
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		words[i] = string(runes)
	}

	return strings.Join(words, " ")
}
