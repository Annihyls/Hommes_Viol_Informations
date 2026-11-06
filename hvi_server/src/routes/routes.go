package routes

import "hvi_server/serverconfig"

var serverConf *config.ServerConf = nil

func Init(serverconf *config.ServerConf) {
	staticFilesInit(serverconf)
	setHomePage(serverconf)
	setLoginPage(serverconf)
}
