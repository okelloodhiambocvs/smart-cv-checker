package tests

import (
	"testing"
	"smart-cv-checker/internal/ats"
)

func TestEngine(t *testing.T) {
	res := ats.Analyze("go docker", "go kubernetes")

	if res.Score != 50 {
		t.Errorf("expected 50 got %d", res.Score)
	}
}