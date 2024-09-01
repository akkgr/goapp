package views

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFS(Templates(), "home.html", "navbar.html", "base.html"))
	tmpl.Execute(w, nil)
}
