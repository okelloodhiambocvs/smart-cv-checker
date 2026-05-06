package tests

import (
	"testing"
	"smart-cv-checker/internal/ats"
)

func TestMatcher(t *testing.T) {
	cv := []string{"go", "docker"}
	job := []string{"go", "kubernetes"}

	m, miss := ats.MatchKeywords(cv, job)

	if len(m) != 1 || len(miss) != 1 {
		t.Errorf("matcher incorrect")
	}
}