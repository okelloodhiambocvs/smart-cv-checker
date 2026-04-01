package library

import "strings" // <--- ADD THIS

// GenerateSuggestions returns simple improvement suggestions
func GenerateSuggestions(result AnalysisResult) []string {
	suggestions := []string{}

	if len(result.MissingKeywords) > 0 {
		suggestions = append(suggestions, "Add missing technical/soft skills: "+strings.Join(result.MissingKeywords, ", "))
	}

	for kw, freq := range result.KeywordFrequency {
		if freq < 2 {
			suggestions = append(suggestions, "Increase repetition of keyword: "+kw)
		}
	}

	for cat, score := range result.CategoryScore {
		if score < 50 {
			suggestions = append(suggestions, "Improve "+cat+" skill coverage")
		}
	}

	if result.MatchScore < 50 {
		suggestions = append(suggestions, "Consider revising CV to better match job description")
	}

	return suggestions
}