/*
Package config implement loading
*/
package config

import (
	"fmt"
	"hvi_server/environnement"
	"hvi_server/utils"
)

type ServerConf struct {
	content map[string]string
}

func (c *ServerConf) Get(key string) string {
	return c.content[key]
}

func (c *ServerConf) GetTemplateDirPath() string {
	return c.content["PathToTemplateFiles"]
}

func (c *ServerConf) GetStaticDirPath() string {
	return c.content["PathToStaticFiles"]
}

func Load(cf ConfFile, envVarMap *env.VarMap) (*ServerConf, error) {
	confPath := envVarMap.Get(env.HVIServerConfiguration)
	confFile, err := cf.Open(confPath)
	if err != nil {
		return nil, fmt.Errorf("unable to open configuration file: %w", err)
	}
	defer cf.Close(confFile)
	configuration := &ServerConf{
		content: make(map[string]string),
	}
	for cf.CanReadLine(confFile) {
		line, err := cf.ReadLine(confFile)
		if err != nil {
			return nil, fmt.Errorf("unable to parse file: %w", err)
		}
		trimedLine := utils.StringTrimSpace(line)
		if trimedLine == "" || utils.StringHasPrefix(trimedLine, "#") {
			continue
		}
		key, val, ok := utils.StringCut(trimedLine, "=")
		if !ok {
			return nil, fmt.Errorf("invalid line: %q", trimedLine)
		}
		configuration.content[utils.StringTrimSpace(key)] = utils.StringTrimSpace(val)
	}
	return configuration, nil
}
