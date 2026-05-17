package web

import (
	"html/template"
	"net/http"

	"ajirascan/internal/ats"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		tmpl := template.Must(
			template.ParseFiles("templates/index.html"),
		)

		tmpl.Execute(w, nil)
		return
	}

	cv := r.FormValue("cv")
	job := r.FormValue("job")

	result := ats.Analyze(cv, job)

	tmpl := template.Must(
		template.ParseFiles("templates/index.html"),
	)

	tmpl.Execute(w, result)
}