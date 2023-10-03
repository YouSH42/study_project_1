package handler

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, _ *http.Request) {

	// HTML 템플릿 렌더링
	tmpl := template.Must(template.ParseFiles("templates/counter.html"))
	tmpl.Execute(w, nil)
}
