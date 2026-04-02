package library

import "strings"

func AnalyzeCV(cvText, jobText string) AnalysisResult {
	cvText = NormalizeText(cvText)
	jobText = NormalizeText(jobText)

	jobKeywords := ExtractKeywords(jobText)
	freq := CountWordFrequency(cvText)

	matched := []string{}
	missing := []string{}

	// Keyword weights
	weights := map[string]int{}
	for _, k := range jobKeywords {
		if len(k) > 6 {
			weights[k] = 3
		} else if len(k) > 4 {
			weights[k] = 2
		} else {
			weights[k] = 1
		}
	}

	totalWeight := 0
	matchedWeight := 0

	for _, jk := range jobKeywords {
		totalWeight += weights[jk]
		if freq[jk] > 0 {
			matched = append(matched, jk)
			matchedWeight += weights[jk]
		} else {
			missing = append(missing, jk)
		}
	}

	// Base score
	score := 0
	if totalWeight > 0 {
		score = (matchedWeight * 100) / totalWeight
	}

	// Penalties
	penalty := 0
	for _, count := range freq {
		if count > 3 {
			penalty += 2
		}
	}
	if strings.Contains(cvText, "responsibilities") || strings.Contains(cvText, "we are looking for") || strings.Contains(cvText, "requirements") {
		penalty += 20
	}
	if cvText == jobText {
		penalty += 30
	}
	score -= penalty
	if score < 0 {
		score = 0
	}

	// Boost: action verbs
	actionWords := []string{"developed", "built", "designed", "implemented", "created"}
	boost := 0
	for _, w := range actionWords {
		if strings.Contains(cvText, w) {
			boost += 3
		}
	}
	score += boost
	if score > 100 {
		score = 100
	}

	authScore := 100 - penalty
	if authScore < 0 {
		authScore = 0
	}

	suggestions := GenerateSuggestions(missing, freq, cvText)

	return AnalysisResult{
		MatchScore:        score,
		AuthenticityScore: authScore,
		MatchedKeywords:   matched,
		MissingKeywords:   missing,
		KeywordFrequency:  freq,
		Suggestions:       suggestions,
	}
}