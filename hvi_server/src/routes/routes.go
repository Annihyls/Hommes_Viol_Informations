package routes

import "hvi_server/serverconfig"

func Init(serverconf config.ServerConf) {
	staticFilesInit(serverconf)
	setHomePage(serverconf)
	setLoginPage(serverconf)
}
