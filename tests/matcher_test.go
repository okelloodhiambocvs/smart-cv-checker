package tests

import (
	"testing"

	"smart-cv-checker/internal/ats"
)

func TestMatchKeywords(t *testing.T) {
	cv := []string{"go", "docker"}
	job := []string{"go", "kubernetes"}

	matched, missing := ats.MatchKeywords(cv, job)

	if len(matched) != 1 || matched[0] != "go" {
		t.Errorf("Expected 'go' as matched")
	}

	if len(missing) != 1 || missing[0] != "kubernetes" {
		t.Errorf("Expected 'kubernetes' as missing")
	}
}

func TestMatchEmpty(t *testing.T) {
	cv := []string{}
	job := []string{"go"}

	matched, missing := ats.MatchKeywords(cv, job)

	if len(matched) != 0 {
		t.Errorf("Expected no matches")
	}

	if len(missing) != 1 {
		t.Errorf("Expected 1 missing keyword")
	}
}