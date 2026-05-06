package tests

import (
	"testing"

	"smart-cv-checker/internal/text"
)

func TestNormalize(t *testing.T) {
	input := "Go Developer!!!\nDocker"
	expected := "go developer docker"

	result := text.Normalize(input)

	if result != expected {
		t.Errorf("Normalize failed: expected '%s', got '%s'", expected, result)
	}
}

func TestTokenize(t *testing.T) {
	input := "go developer docker"

	tokens := text.Tokenize(input)

	if len(tokens) != 3 {
		t.Errorf("Tokenize failed: expected 3 tokens, got %d", len(tokens))
	}
}

func TestFrequency(t *testing.T) {
	tokens := []string{"go", "go", "docker"}

	freq := text.Frequency(tokens)

	if freq["go"] != 2 {
		t.Errorf("Frequency failed for 'go'")
	}
	if freq["docker"] != 1 {
		t.Errorf("Frequency failed for 'docker'")
	}
}