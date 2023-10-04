package handler

import (
	"html/template"
	"net/http"
)

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/message.html"))
	tmpl.Execute(w, nil)
}
