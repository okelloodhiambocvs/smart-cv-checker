package ats

import "fmt"

func Suggestions(missing []string) []string {
	var suggestions []string

	for _, keyword := range missing {
		suggestions = append(
			suggestions,
			fmt.Sprintf("Consider adding experience or skills related to '%s'", keyword),
		)
	}

	return suggestions
}

func SectionSuggestions(missingSections []string) []string {
	var suggestions []string

	for _, section := range missingSections {
		suggestions = append(
			suggestions,
			fmt.Sprintf("Add a '%s' section to improve ATS completeness", section),
		)
	}

	return suggestions
}