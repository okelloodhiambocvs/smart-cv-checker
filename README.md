## smart-cv-checker
Smart-ATS is a beginner-friendly Go project that simulates an Applicant Tracking System. It analyzes a CV/resume against a Job Description, extracts keywords, calculates match scores, categorizes skills, and provides actionable suggestions for improving CVs.

smart-ats/
├─ main.go              # Entry point: orchestrates reading CV/Job and calls library
├─ go.mod               # Go module
├─ sample_cv.txt        # Optional sample CV input
├─ sample_job.txt       # Optional sample Job Description
├─ ats_test.go          # Unit tests
└─ library/             # Core library for ATS logic
   ├─ analyzer.go       # Core analysis: matching, scoring, category analysis
   ├─ parser.go         # Input parsing, normalization, keyword extraction
   ├─ suggestions.go    # Suggestion generation logic
   └─ types.go          # Shared types (e.g., KeywordMatch, CategoryMatch)
