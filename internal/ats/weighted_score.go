package ats

func WeightedScore(matched, missing []string) int {
	matchedScore := 0
	totalScore := 0

	for _, keyword := range matched {
		weight := getWeight(keyword)
		matchedScore += weight
		totalScore += weight
	}

	for _, keyword := range missing {
		totalScore += getWeight(keyword)
	}

	if totalScore == 0 {
		return 0
	}

	return (matchedScore * 100) / totalScore
}

func getWeight(keyword string) int {
	weight, exists := KeywordWeights[keyword]

	if exists {
		return weight
	}

	return 1
}