package library

type KeywordMatch struct {
	Keyword string
	Count   int
}

type AnalysisResult struct {
	MatchScore        int
	AuthenticityScore int
	MatchedKeywords   []string
	MissingKeywords   []string
	KeywordFrequency  map[string]int
	Suggestions       []string
}