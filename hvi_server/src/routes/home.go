package routes

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func setHomePage() {
	var tmplPath string = filepath.Join(getTemplateDirPath(), "home.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
