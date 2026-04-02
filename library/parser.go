package library

import (
	"regexp"
	"strings"
)

// NormalizeText: lowercase, remove punctuation, normalize spaces
func NormalizeText(text string) string {
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, "\n", " ")
	re := regexp.MustCompile(`[^\w\s]`) // remove punctuation
	text = re.ReplaceAllString(text, "")
	return text
}

// ExtractKeywords: filter stopwords and short words
func ExtractKeywords(text string) []string {
	words := strings.Fields(text)

	stopWords := map[string]bool{
		"the": true, "and": true, "is": true, "in": true, "to": true,
		"of": true, "a": true, "with": true, "for": true,
		"on": true, "as": true, "at": true, "by": true,
	}

	keywords := []string{}
	seen := make(map[string]bool)

	for _, word := range words {
		if len(word) < 3 || stopWords[word] {
			continue
		}
		if !seen[word] {
			keywords = append(keywords, word)
			seen[word] = true
		}
	}

	return keywords
}

// CountWordFrequency: counts frequency of each word in text
func CountWordFrequency(text string) map[string]int {
	freq := make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		freq[word]++
	}

	return freq
}