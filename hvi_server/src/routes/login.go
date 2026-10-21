package routes

import (
	"html/template"
	"path/filepath"
	"net/http"
)

func setLoginPage() {
	var url string = "/login"
	var tmplPath string = filepath.Join(getTemplateDirPath(), "login.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	http.HandleFunc(url, func(w http.ResponseWriter, req *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
