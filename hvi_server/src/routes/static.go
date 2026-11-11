package routes

import (
	config "hvi_server/serverconfig"
	"net/http"
)

func staticFilesInit(serverconfig config.ServerConf) {
	fileServer := http.FileServer(http.Dir(serverconfig.Get(config.StaticDirPath)))
	http.Handle("/css/", fileServer)
}
