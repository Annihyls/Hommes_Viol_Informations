package routes

import (
	"html/template"
	"hvi_server/constants"
	"net/http"
	"fmt"
)

func staticFilesInit() {
	pathStatic := cstenv.Getenv(cstenv.PATH_STATIC)
	fileServer := http.FileServer(http.Dir(pathStatic))
	http.Handle("/css/", fileServer)
}

func RoutesInit() {
	staticFilesInit()
	homePage := fmt.Sprintf("%s/home.html", cstenv.Getenv(cstenv.PATH_TEMPLATES))
	tmpl := template.Must(template.ParseFiles(homePage))
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
