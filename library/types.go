package library

// KeywordMatch stores info about matched/missing keywords
type KeywordMatch struct {
	Matched   []string
	Missing   []string
	Frequency map[string]int
}

// CategoryMatch stores percentage match per category
type CategoryMatch struct {
	Technical float64
	Soft      float64
}

// AnalysisResult combines keyword and category info
type AnalysisResult struct {
	MatchedKeywords   []string
	MissingKeywords   []string
	KeywordFrequency  map[string]int
	CategoryScore     map[string]float64
	MatchScore        float64
}