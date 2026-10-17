package routes

import (
	"html/template"
	"net/http"
)

func staticFilesInit() {
	// Serveur pour les fichiers statiques (CSS, JS, images)
	fileServer := http.FileServer(http.Dir("../static"))
	http.Handle("/css/", fileServer) // Serve le CSS depuis /css/
}

func RoutesInit() {
	staticFilesInit()
	tmpl := template.Must(template.ParseFiles("../templates/home.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
