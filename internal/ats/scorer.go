package ats

func Score(matched, missing []string) int {
	total := len(matched) + len(missing)
	if total == 0 {
		return 0
	}

	score := (len(matched) * 100) / total
	return score
}