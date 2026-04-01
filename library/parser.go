package library

import (
	"strings"
	"unicode"
)

// ParseText normalizes text: lowercase, remove punctuation, split words
func ParseText(text string) []string {
	text = strings.ToLower(text)
	words := strings.FieldsFunc(text, func(r rune) bool {
		return unicode.IsPunct(r) || unicode.IsSpace(r)
	})

	// Remove common stopwords
	stopwords := map[string]bool{
		"and": true, "the": true, "with": true,
		"a": true, "an": true, "to": true,
		"for": true, "in": true, "on": true,
	}

	var result []string
	for _, w := range words {
		if _, ok := stopwords[w]; !ok {
			result = append(result, w)
		}
	}
	return result
}