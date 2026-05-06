package tests

import (
	"testing"
	"smart-cv-checker/internal/ats"
)

func TestScore(t *testing.T) {
	score := ats.Score([]string{"go"}, []string{"docker"})

	if score != 50 {
		t.Errorf("expected 50 got %d", score)
	}
}