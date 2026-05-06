package tests

import (
	"testing"
	"smart-cv-checker/internal/ats"
)

func TestEngine_BasicMatch(t *testing.T) {
	res := ats.Analyze(
		"Go developer with Docker experience",
		"Go Docker Kubernetes Backend",
	)

	if res.Score == 0 {
		t.Errorf("expected non-zero score")
	}

	if len(res.Matched) == 0 {
		t.Errorf("expected matched keywords")
	}
}