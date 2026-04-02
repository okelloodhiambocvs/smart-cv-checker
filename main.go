package main

import (
	"flag"
	"fmt"
	"os"
	"smart-cv-checker/library"
)

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	return string(data)
}

func main() {
	cvFile := flag.String("cv", "", "Path to CV file")
	jobFile := flag.String("job", "", "Path to Job Description file")
	flag.Parse()

	var cvText, jobText string

	if *cvFile != "" && *jobFile != "" {
		cvText = readFile(*cvFile)
		jobText = readFile(*jobFile)
	} else if len(flag.Args()) == 2 {
		cvText = flag.Args()[0]
		jobText = flag.Args()[1]
	} else {
		fmt.Println("Usage: go run . -cv <cv_file> -job <job_file>")
		fmt.Println("   or: go run . \"<cv text>\" \"<job text>\"")
		return
	}

	result := library.AnalyzeCV(cvText, jobText)

	fmt.Println("====== ATS ANALYSIS RESULT ======")
	fmt.Println("Match Score:", result.MatchScore, "%")
	fmt.Println("Authenticity Score:", result.AuthenticityScore, "%")
	fmt.Println("\nMatched Keywords:", result.MatchedKeywords)
	fmt.Println("Missing Keywords:", result.MissingKeywords)
	fmt.Println("\nKeyword Frequency:", result.KeywordFrequency)
	fmt.Println("\nSuggestions:")
	for _, s := range result.Suggestions {
		fmt.Println("-", s)
	}
}