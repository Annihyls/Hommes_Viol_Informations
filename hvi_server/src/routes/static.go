package routes

import (
	"hvi_server/serverconfig"
	"net/http"
)

func staticFilesInit(serverconfig *config.ServerConf) {
	fileServer := http.FileServer(http.Dir(serverconfig.GetStaticDirPath()))
	http.Handle("/css/", fileServer)
}
