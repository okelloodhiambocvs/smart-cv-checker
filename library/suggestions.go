package library

import "strings"

func GenerateSuggestions(missing []string, freq map[string]int, cvText string) []string {
	suggestions := []string{}

	if len(missing) > 0 {
		suggestions = append(suggestions, "Add missing keywords: "+strings.Join(missing, ", "))
	}

	for word, count := range freq {
		if count > 3 {
			suggestions = append(suggestions, "Avoid repeating '"+word+"' too many times")
		}
	}

	if strings.Contains(cvText, "responsibilities") || strings.Contains(cvText, "we are looking for") {
		suggestions = append(suggestions, "Avoid copying job description phrases")
	}

	actionWords := []string{"developed", "built", "designed", "implemented", "created"}
	foundAction := false
	for _, w := range actionWords {
		if strings.Contains(cvText, w) {
			foundAction = true
			break
		}
	}
	if !foundAction {
		suggestions = append(suggestions, "Use action verbs like 'developed', 'built', 'designed'")
	}

	return suggestions
}