package main

import (
	"flag"
	"fmt"
	"os"
	"smart-cv-checker/internal/ats"
)

func read(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func main() {
	cv := flag.String("cv", "", "")
	job := flag.String("job", "", "")
	flag.Parse()

	if *cv == "" || *job == "" {
		fmt.Println("usage: -cv file -job file")
		return
	}

	res := ats.Analyze(read(*cv), read(*job))

	fmt.Println("MATCH SCORE:", res.Score)
	fmt.Println("MATCHED:", res.Matched)
	fmt.Println("MISSING:", res.Missing)
}