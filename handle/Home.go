package handle

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, _ *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)

}
