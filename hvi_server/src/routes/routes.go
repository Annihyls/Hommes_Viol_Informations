package routes

import "hvi_server/configserver"

var serverConf *config.ServerConf = nil

func Init(serverconf *config.ServerConf) {
	serverConf = serverconf
	staticFilesInit()
	setHomePage()
	setLoginPage()
}
