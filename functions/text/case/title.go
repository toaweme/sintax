package casing

import (
	"strings"
	"unicode"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameModelTitle is the template name for the ModelTitle modifier.
const ModifierNameModelTitle functions.ModifierName = "title_model"

// ModifierNameTitle is the template name for the Title modifier.
const ModifierNameTitle functions.ModifierName = "title"

// ModelTitle formats an AI model identifier into a human-readable title. It
// handles provider prefixes, model sizes, quantization suffixes, and date
// suffixes.
func ModelTitle(s string) (string, error) {
	return FormatModelTitle(s), nil
}

// Title converts a hyphen-separated slug into a title-cased string. Known
// acronyms (AI, API, GPT, etc.) are uppercased automatically, and any extra
// acronyms passed as params are merged with that built-in list.
func Title(s string, acronymsExtra ...string) (string, error) {
	all := append([]string{}, acronyms...)
	all = append(all, acronymsExtra...)
	return slugToTitle(s, all), nil
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
