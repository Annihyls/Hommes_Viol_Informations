package routes

import "net/http"

func staticFilesInit() {
	fileServer := http.FileServer(http.Dir(getStaticDirPath()))
	http.Handle("/css/", fileServer)
}
