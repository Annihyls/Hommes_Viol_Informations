package routes

import (
	"html/template"
	"hvi_server/serverconfig"
	"net/http"
	"path/filepath"
)

func setLoginPage(serverconfig *config.ServerConf) {
	url := "/login"
	tmplPath := filepath.Join(serverconfig.GetTemplateDirPath(), "login.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	http.HandleFunc(url, func(w http.ResponseWriter, req *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
