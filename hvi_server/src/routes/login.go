package routes

import (
	"html/template"
	"hvi_server/serverconfig"
	"net/http"
	"path/filepath"
)

func setLoginPage(serverconfig *config.ServerConf) {
	var url string = "/login"
	var tmplPath string = filepath.Join(serverconfig.GetTemplateDirPath(), "login.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	http.HandleFunc(url, func(w http.ResponseWriter, req *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
