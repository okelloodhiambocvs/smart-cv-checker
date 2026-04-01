package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"smart-cv-checker/library"
)

func main() {
	var cvText, jobText string

	// If exactly 2 args, treat as raw text
	if len(os.Args) == 3 {
		cvText = os.Args[1]
		jobText = os.Args[2]
	} else if len(os.Args) == 5 && os.Args[1] == "-cv" && os.Args[3] == "-job" {
		// Read files
		cvFile := os.Args[2]
		jobFile := os.Args[4]

		cvBytes, err := ioutil.ReadFile(cvFile)
		if err != nil {
			fmt.Println("Error reading CV file:", err)
			return
		}
		jobBytes, err := ioutil.ReadFile(jobFile)
		if err != nil {
			fmt.Println("Error reading Job file:", err)
			return
		}
		cvText = string(cvBytes)
		jobText = string(jobBytes)
	} else {
		fmt.Println("Usage:")
		fmt.Println("  go run . [CV_TEXT] [JOB_DESCRIPTION_TEXT]")
		fmt.Println("  or")
		fmt.Println("  go run . -cv [CV_FILE] -job [JOB_FILE]")
		return
	}

	// Step 1: Parse and normalize text
	cvWords := library.ParseText(cvText)
	jobWords := library.ParseText(jobText)

	// Step 2: Analyze keywords and categories
	result := library.AnalyzeCV(cvWords, jobWords)

	// Step 3: Generate suggestions
	suggestions := library.GenerateSuggestions(result)

	// Step 4: Display result
	fmt.Printf("Match Score: %.0f%%\n\n", result.MatchScore)
	fmt.Println("Matched Keywords:", result.MatchedKeywords)
	fmt.Println("Missing Keywords:", result.MissingKeywords)
	fmt.Println("Keyword Frequency:", result.KeywordFrequency)
	fmt.Println("Category Analysis:")
	for cat, score := range result.CategoryScore {
		fmt.Printf("- %s: %.0f%%\n", cat, score)
	}
	fmt.Println("\nSuggestions:")
	for _, s := range suggestions {
		fmt.Println("-", s)
	}
}