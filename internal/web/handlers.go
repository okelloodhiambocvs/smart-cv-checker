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

	file, _, err := r.FormFile("cvfile")

	if err == nil {

		defer file.Close()

		buf := make([]byte, 1024*1024)

		n, _ := file.Read(buf)

		cv = string(buf[:n])
	}

	job := r.FormValue("job")

	result := ats.Analyze(cv, job)

	tmpl := template.Must(
		template.ParseFiles("templates/index.html"),
	)

	tmpl.Execute(w, result)
}