package main

import (
	"reflect"
	"smart-cv-checker/library"
	"testing"
)

func TestParseText(t *testing.T) {
	text := "Go Developer, API & Docker!"
	expected := []string{"go", "developer", "api", "docker"}
	result := library.ParseText(text)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestAnalyzeCV(t *testing.T) {
	cv := library.ParseText("go api backend docker")
	job := library.ParseText("go api docker kubernetes")
	result := library.AnalyzeCV(cv, job)

	if result.MatchScore != 75 {
		t.Errorf("Expected 75, got %.0f", result.MatchScore)
	}
	if len(result.MissingKeywords) != 1 || result.MissingKeywords[0] != "kubernetes" {
		t.Errorf("Expected missing: [kubernetes], got %v", result.MissingKeywords)
	}
}