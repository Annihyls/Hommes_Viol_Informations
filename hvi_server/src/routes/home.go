package routes

import (
	"html/template"
	"hvi_server/serverconfig"
	"net/http"
	"path/filepath"
)

func setHomePage(serverconfig *config.ServerConf) {
	tmplPath := filepath.Join(serverconfig.GetTemplateDirPath(), "home.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
