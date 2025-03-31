package text

import (
	"strings"
	"unicode"

	"github.com/toaweme/sintax/functions"
)

// for emergencies only

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
	for _, acronym := range acronyms {
		all = append(all, acronym)
	}

	if len(providedAcronyms) > 0 {
		all = append(all, providedAcronyms...)
	}

	return slugToTitle(slug, all), nil
}

var acronyms = []string{
	"ai",
	"gpt",
	"api",
	"json",
	"xml",
	"html",
	"css",
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
