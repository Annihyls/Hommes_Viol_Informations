package routes

import (
        "hvi_server/configserver"
)

func getTemplateDirPath() string {
        return config.Current["PathToTemplateFiles"]
}

func getStaticDirPath() string {
        return config.Current["PathToStaticFiles"]
}
