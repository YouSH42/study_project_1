package handler

import (
	"html/template"
	"net/http"
)

func AdminHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/admin.html"))
	tmpl.Execute(w, nil)
}
