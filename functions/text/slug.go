package text

import (
	"regexp"
	"strings"

	"github.com/toaweme/sintax/functions"
)

// ModifierNameSlug is the template name for the Slug modifier.
const ModifierNameSlug functions.ModifierName = "slug"

// Slug converts a string to a URL-friendly slug.
// Lowercases, replaces spaces with hyphens, and removes non-alphanumeric characters.
//
// value: string
// returns: string
//
// example: turn a blog post title into a URL slug
// in:  title = "Welcome to the Coffee Club!"
// tpl: {{ title | slug }}
// out: welcome-to-the-coffee-club
//
// example: slugify a product name
// in:  name = "Premium Tea & Honey Set"
// tpl: {{ name | slug }}
// out: premium-tea-honey-set
func Slug(value any, params []any) (any, error) {
	slug, err := functions.ValueString(value)
	if err != nil {
		return nil, err
	}

	return titleToSlug(slug), nil
}

func titleToSlug(title string) string {
	// Convert to lowercase
	slug := strings.ToLower(title)

	// Replace chars that aren't alphanumeric, spaces, hyphens, or dots with spaces
	reg := regexp.MustCompile(`[^a-z0-9\s\-\.]`)
	slug = reg.ReplaceAllString(slug, " ")

	// Replace multiple spaces with a single space
	regMultiSpace := regexp.MustCompile(`\s+`)
	slug = regMultiSpace.ReplaceAllString(slug, " ")

	// trim spaces from beginning and end
	slug = strings.TrimSpace(slug)

	// change non-numeric dots to spaces (keep dots between numbers)
	regNonDigitDots := regexp.MustCompile(`(\D)\.+(\D)`)
	slug = regNonDigitDots.ReplaceAllString(slug, "$1 $2")

	regDotAfterNonDigit := regexp.MustCompile(`(\D)\.+(\d)`)
	slug = regDotAfterNonDigit.ReplaceAllString(slug, "$1 $2")

	regDotBeforeNonDigit := regexp.MustCompile(`(\d)\.+(\D)`)
	slug = regDotBeforeNonDigit.ReplaceAllString(slug, "$1 $2")

	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	// Replace consecutive hyphens with a single hyphen
	regMultiHyphen := regexp.MustCompile(`-+`)
	slug = regMultiHyphen.ReplaceAllString(slug, "-")

	return slug
}
