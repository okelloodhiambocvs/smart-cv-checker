package library

// AnalyzeCV performs keyword matching, frequency counting, category analysis
func AnalyzeCV(cvWords, jobWords []string) AnalysisResult {
	cvFreq := make(map[string]int)
	for _, w := range cvWords {
		cvFreq[w]++
	}

	matched := []string{}
	missing := []string{}
	for _, kw := range jobWords {
		if cvFreq[kw] > 0 {
			matched = append(matched, kw)
		} else {
			missing = append(missing, kw)
		}
	}

	matchScore := 0.0
	if len(jobWords) > 0 {
		matchScore = float64(len(matched)) / float64(len(jobWords)) * 100
	}

	// Simple category analysis
	categories := map[string][]string{
		"Technical": {"go", "api", "docker", "kubernetes", "backend", "frontend", "database", "microservices"},
		"Soft":      {"communication", "teamwork", "leadership", "problem-solving", "collaboration"},
	}

	categoryScore := make(map[string]float64)
	for cat, kwList := range categories {
		count := 0
		for _, kw := range kwList {
			if contains(cvWords, kw) && contains(jobWords, kw) {
				count++
			}
		}
		if len(kwList) > 0 {
			categoryScore[cat] = float64(count) / float64(len(kwList)) * 100
		} else {
			categoryScore[cat] = 0
		}
	}

	return AnalysisResult{
		MatchedKeywords:  matched,
		MissingKeywords:  missing,
		KeywordFrequency: cvFreq,
		CategoryScore:    categoryScore,
		MatchScore:       matchScore,
	}
}

// Helper function
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}