package tests

import (
	"testing"
	"smart-cv-checker/internal/text"
)

func TestNormalize(t *testing.T) {
	out := text.Normalize("Go Developer!!!\nDocker")
	if out != "go developer docker" {
		t.Errorf("unexpected normalize output: %s", out)
	}
}