package routes

func getTemplateDirPath() string {
	return serverConf.Get("PathToTemplateFiles")
}

func getStaticDirPath() string {
	return serverConf.Get("PathToStaticFiles")
}
